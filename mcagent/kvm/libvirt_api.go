package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"fmt"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"strconv"
	"strings"
)

var libvirtConn *libvirt.Connect

func GetQemuConnect() (*libvirt.Connect, error) {
	if libvirtConn != nil {
		return libvirtConn, nil
	}

	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("GetQemuConnect error:", err)
		return conn, err
	}

	if conn != nil {
		libvirtConn = conn
	}

	return conn, err
}

func ConvertVmStatus(status libvirt.DomainState) string {
	var res string
	switch status {
	case libvirt.DOMAIN_NOSTATE:
		res = "no status"
	case libvirt.DOMAIN_RUNNING:
		res = "running"
	case libvirt.DOMAIN_SHUTDOWN:
		res = "shutdown"
	case libvirt.DOMAIN_BLOCKED:
		res = "blocked"
	case libvirt.DOMAIN_PAUSED:
		res = "paused"
	case libvirt.DOMAIN_CRASHED:
		res = "crashed"
	case libvirt.DOMAIN_PMSUSPENDED:
		res = "pm suspended"
	case libvirt.DOMAIN_SHUTOFF:
		res = "shutoff"
	default:
		res = "unknown"
	}
	return res
}

func ConvertImageFile2MgoVM(vm *mcmodel.McVm, file string) bool {
	// /opt/vm_instances/windows10-40G-0.qcow2
	vm.FullPath = file
	arr := strings.Split(file, "/")
	if len(arr) < 3 || !strings.Contains(arr[1], "opt") {
		fmt.Printf("The directory(%s) isn't valid --> skip (we didn't create this vm)\n",
			vm.FullPath)
		return false
	}
	name := arr[3]
	fmt.Printf("%s\n", arr)
	fmt.Printf("%s\n", arr[3])
	vm.Filename = name[:strings.LastIndexAny(name,".")]
	list := strings.Split(vm.Filename, "-")
	if list[0] == "windows10" {
		vm.OS = "win10"
	} else if list[0] == "ubuntu18" {
		vm.OS= "ubuntu18"
	} else if list[0] == "ubuntu16" {
		vm.OS = "ubuntu16"
	} else {
		return false
	}
	vm.VmIndex, _ = strconv.Atoi(list[2])
	fmt.Printf("ConvertImageFile2MgoVM: vmIndex %d, %s\n", vm.VmIndex, list[2])
	vm.Image = fmt.Sprintf("%s-%s", list[0], list[1])
	vm.Hdd, _ = strconv.Atoi(list[1][:strings.LastIndexAny(list[1],"G")])
	return true
}

func DumpMcVirtInfo() {
	vmList, netList, imgList := GetMcVirtInfo()
	mcmodel.DumpVmList(vmList)
	mcmodel.DumpNetworkList(netList)
	mcmodel.DumpImageList(imgList)
}

func GetVmByLibvirt() (vmList []mcmodel.McVm){
	// Get SnapVms Domains
	doms, err := GetDomainListAll()
	if err != nil {
		fmt.Println("GetVmByLibvirt error:", err)
		return vmList
	}

	if err == nil {
		for _, dom := range doms {
			// vm name
			var vm mcmodel.McVm
			name, _ := dom.GetName()

			vm.Name = name
			//****************************************************************
			xmlstr, _ := dom.GetXMLDesc(0)
			domcfg := &libvirtxml.Domain{}
			err = domcfg.Unmarshal(xmlstr)
			if domcfg.VCPU != nil {
				vcpu := domcfg.VCPU.Value
				vm.Cpu = int(vcpu)
			}
			if domcfg.Memory != nil {
				vm.Ram = int(domcfg.Memory.Value / 1024)
			}
			if domcfg.Devices != nil {
				var res = false
				devices := domcfg.Devices
				for _, disk := range devices.Disks {
					if disk.Source != nil && disk.Source.File != nil {
						res = ConvertImageFile2MgoVM(&vm, disk.Source.File.File)
					}
				}
				if res == false {
					// hcmp dit not create this vm!
					continue
				}
				interfaces := domcfg.Devices.Interfaces
				for _, intf := range interfaces {
					if intf.Source != nil && intf.Source.Network != nil {
						vm.Network = intf.Source.Network.Network
					}
					if intf.MAC != nil {
						vm.Mac = intf.MAC.Address
					}
					//for _, ip := range intf.IP {
					//	fmt.Printf("%s\n", ip.Address)
					//}
				}
				graps := domcfg.Devices.Graphics
				for _, grap := range graps {
					if grap.VNC != nil {
						vm.VncPort = fmt.Sprintf("%d", grap.VNC.Port)
					}
				}
			}
			//****************************************************************
			// os type
			//ostype, _ := dom.GetOSType()

			// status
			status, _, _ := dom.GetState()
			vm.CurrentStatus = ConvertVmStatus(status)

			// ip address
			domifs, _ := dom.ListAllInterfaceAddresses(0)
			for _, intf := range domifs {
				//fmt.Printf("   intf: %s, %s", intf.Name, intf.Hwaddr)
				for _, ip := range intf.Addrs {
					//fmt.Printf(", %s", ip.Addr)
					vm.IpAddr = ip.Addr
				}
			}
			cfg := config.GetMcGlobalConfig()
			vm.RemoteAddr = fmt.Sprintf("%s:%d",
				cfg.ServerIp,
				cfg.DnatBasePortNum + vm.VmIndex)
			vm.PublicRemoteAddr= fmt.Sprintf("%s:%d",
				cfg.ServerPublicIp,
				cfg.DnatBasePortNum + vm.VmIndex)
			//config.AllocateVmIndex(uint(vm.VmIndex))
			//fmt.Printf("\n")
			vm.IsCreated = true
			vmList = append(vmList, vm)
		}
	}

	return vmList
}

func GetNetworkByLibvirt() (netList []mcmodel.McNetworks){

	// Get Networks
	nets, err := GetAllNetwork()
	if err != nil {
		fmt.Println("GetNetworkByLibvirt error:", err)
		return netList
	}

	for _, net := range nets {
		var network mcmodel.McNetworks
		name, _ := net.GetName()
		network.Name = name
		bridge, _ := net.GetBridgeName()
		network.Bridge = bridge

		xmlstr, _ := net.GetXMLDesc(0)
		netcfg := &libvirtxml.Network{}
		err = netcfg.Unmarshal(xmlstr)
		mode := netcfg.Forward.Mode
		network.Mode = mode
		fmt.Println("Network:", mode)
		if netcfg.MAC != nil {
			mac := netcfg.MAC.Address
			network.Mac = mac
		}
		network.Uuid = netcfg.UUID
		for _, Ip := range netcfg.IPs {
			netIp := Ip.Address
			netNetmask := Ip.Netmask
			network.Ip = netIp
			network.Netmask = netNetmask
			network.Prefix = Ip.Prefix
			if Ip.DHCP != nil {
				netDhcp := Ip.DHCP
				for _, dhcprange := range netDhcp.Ranges {
					//fmt.Printf("   range %d: %s, %s\n", j, dhcprange.Start, dhcprange.End)
					network.DhcpStart = dhcprange.Start
					network.DhcpEnd = dhcprange.End
				}
				//for _, host := range netDhcp.Hosts {
				//	fmt.Printf("   host: %s, %s, %s, %s", host.ID, host.MAC, host.Name, host.IP)
				//}
			}
		}
		dhcps, _ := net.GetDHCPLeases()
		for _, dhcp := range dhcps {
			//fmt.Printf("   dhcp %d: %s, %s, %s, %s\n",
			//	i, dhcp.Iface, dhcp.Mac, dhcp.IPaddr, dhcp.Hostname)
			var host mcmodel.McNetHost
			host.Mac = dhcp.Mac
			host.Ip = dhcp.IPaddr
			host.Hostname = dhcp.Hostname
			network.Host = append(network.Host, host)
		}
		//network.Dump()
		netList = append(netList, network)
	}

	return netList
}

func GetMcServerInfo() mcmodel.McServerMsg{
	var server mcmodel.McServerMsg
	vmList := GetVmByLibvirt()
	netList := GetNetworkByLibvirt()
	imgList := GetImages()

	server.SerialNumber = config.GetSerialNumber()
	server.Ip = config.GetMcGlobalConfig().ServerIp
	server.Port = config.GetMcGlobalConfig().ServerPort
	server.Mac = config.GetMcGlobalConfig().ServerMac
	server.PublicIp = config.GetMcGlobalConfig().ServerPublicIp
	server.Vms = &vmList
	server.Networks = &netList
	server.Images = &imgList
	return server
}

func GetMcVirtInfo() (vmList []mcmodel.McVm, netList []mcmodel.McNetworks, imgList []mcmodel.McImages) {
	vmList = GetVmByLibvirt()
	netList = GetNetworkByLibvirt()
	imgList = GetImages()
	return  vmList, netList, imgList
}

func GetMcVirtInfoDebug() (vmList []mcmodel.McVm, netList []mcmodel.McNetworks, imgList []mcmodel.McImages) {
	// Get SnapVms Domains
	doms, err := GetDomainListAll()
	if err == nil {
		fmt.Println("--------------------------------")
		fmt.Println("VMs")
		for index, dom := range doms {
			// vm name
			var vm mcmodel.McVm
			name, _ := dom.GetName()
			fmt.Printf("%d. %s\n", index, name)

			vm.Name = name
			//****************************************************************
			xmlstr, _ := dom.GetXMLDesc(0)
			domcfg := &libvirtxml.Domain{}
			err = domcfg.Unmarshal(xmlstr)
			if domcfg.VCPU != nil {
				vcpu := domcfg.VCPU.Value
				fmt.Printf("   cpu: %d, %s\n", vcpu,
					domcfg.VCPU.Placement)
				vm.Cpu = int(vcpu)
			}
			if domcfg.Memory != nil {
				fmt.Printf("   memory: %d\n", domcfg.Memory.Value/1024)
				vm.Ram = int(domcfg.Memory.Value / 1024)
			}
			if domcfg.Devices != nil {
				devices := domcfg.Devices
				for _, disk := range devices.Disks {
					if disk.Source != nil && disk.Source.File != nil {
						fmt.Printf("   device: source file %s\n", disk.Source.File.File)
						ConvertImageFile2MgoVM(&vm, disk.Source.File.File)
					}
				}
				interfaces := domcfg.Devices.Interfaces
				for _, intf := range interfaces {
					fmt.Printf("   network: ")
					if intf.Source != nil {
						fmt.Printf("%s, %s",
							intf.Source.Network.Network,
							intf.Source.Network.Bridge)
						vm.Network = intf.Source.Network.Network
					}
					if intf.MAC != nil {
						fmt.Printf(", %s", intf.MAC.Address)
						vm.Mac = intf.MAC.Address
					}
					for _, ip := range intf.IP {
						fmt.Printf("%s", ip.Address)
					}
					fmt.Printf("\n")
				}
			}
			fmt.Printf("   Filename: %s\n", vm.Filename)
			fmt.Printf("   Image: %s\n", vm.Image)
			fmt.Printf("   OS: %s\n", vm.OS)
			fmt.Printf("   hdd: %d\n", vm.Hdd)
			//****************************************************************
			// os type
			ostype, _ := dom.GetOSType()
			fmt.Printf("   os: %s\n", ostype)

			// status
			status, _, _ := dom.GetState()
			fmt.Printf("   status: %s\n", ConvertVmStatus(status))
			vm.CurrentStatus = ConvertVmStatus(status)

			// ip address
			domifs, _ := dom.ListAllInterfaceAddresses(0)
			for i, intf := range domifs {
				fmt.Printf("   intf%d: %s, %s", i, intf.Name, intf.Hwaddr)
				for _, ip := range intf.Addrs {
					fmt.Printf(", %s", ip.Addr)
					vm.IpAddr = ip.Addr
				}
			}
			fmt.Printf("\n")
			vm.IsCreated = true
			vm.Dump()
			vmList = append(vmList, vm)
		}
	}

	// Get Networks
	nets, err := GetAllNetwork()
	if err == nil {
		fmt.Println("--------------------------------")
		fmt.Println("Networks")
		for index, net := range nets {
			var network mcmodel.McNetworks
			name, _ := net.GetName()
			fmt.Printf("%d. %s\n", index, name)
			network.Name = name
			bridge, _ := net.GetBridgeName()
			fmt.Printf("   bridge: %s\n", bridge)
			network.Bridge = bridge

			xmlstr, _ := net.GetXMLDesc(0)
			//fmt.Println(index, ": ", xmlstr)
			netcfg := &libvirtxml.Network{}
			err = netcfg.Unmarshal(xmlstr)
			mode := netcfg.Forward.Mode
			network.Mode = mode
			if netcfg.MAC != nil {
				mac := netcfg.MAC.Address
				fmt.Printf("   mac: %s\n", mac)
				network.Mac = mac
			}
			//fmt.Printf("   mode: %s\n", mode)
			network.Uuid = netcfg.UUID
			for j, Ip := range netcfg.IPs {
				netIp := Ip.Address
				netNetmask := Ip.Netmask
				network.Ip = netIp
				network.Netmask = netNetmask
				network.Prefix = Ip.Prefix
				fmt.Printf("   ip %d: %s/%s\n", j, netIp, netNetmask)
				if Ip.DHCP != nil {
					netDhcp := Ip.DHCP
					for _, dhcprange := range netDhcp.Ranges {
						fmt.Printf("   range %d: %s, %s\n", j, dhcprange.Start, dhcprange.End)
						network.DhcpStart = dhcprange.Start
						network.DhcpEnd = dhcprange.End
					}
					for _, host := range netDhcp.Hosts {
						fmt.Printf("   host: %s, %s, %s, %s", host.ID, host.MAC, host.Name, host.IP)
					}
				}
			}
			dhcps, _ := net.GetDHCPLeases()
			for i, dhcp := range dhcps {
				fmt.Printf("   dhcp %d: %s, %s, %s, %s\n",
					i, dhcp.Iface, dhcp.Mac, dhcp.IPaddr, dhcp.Hostname)
				var host mcmodel.McNetHost
				host.Mac = dhcp.Mac
				host.Ip = dhcp.IPaddr
				host.Hostname = dhcp.Hostname
				network.Host = append(network.Host, host)
			}
			network.Dump()
			netList = append(netList, network)
		}
	}

	// Get Images
	imgList = GetImages()
	if err == nil {
		fmt.Println("--------------------------------")
		fmt.Println("Images")
		for index, img := range imgList {
			fmt.Printf("%d. %s, %s\n", index, img.Variant, img.Name)
			img.Dump()
		}
	}
	return  vmList, netList, imgList
}