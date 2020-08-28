package kvm

import (
	"encoding/xml"
	"fmt"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func CreateXmlNetworkTestApi() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
	defer conn.Close()

	netcfg := &libvirtxml.Network{}
	// name
	netcfg.Name = "net11"

	// forward
	forward := &libvirtxml.NetworkForward{}
	forward.Mode = "nat"
	forwardNat := &libvirtxml.NetworkForwardNAT{}
	natPort := &libvirtxml.NetworkForwardNATPort{
		Start: 1024,
		End: 65535,
	}
	forward.NAT = forwardNat
	forwardNat.Ports = append(forwardNat.Ports, *natPort)
	netcfg.Forward = forward

	// bridge
	bridge := &libvirtxml.NetworkBridge{}
	bridge.Name = "virbr11"
	bridge.STP = "on"
	bridge.Delay = "0"
	netcfg.Bridge = bridge

	// mac
	// ip/dhcp
	ip := &libvirtxml.NetworkIP{}
	ip.Address = "11.0.0.1"
	ip.Netmask = "255.255.255.0"
	dhcp := &libvirtxml.NetworkDHCP{}
	dhcpRange := &libvirtxml.NetworkDHCPRange{
		Start: "11.0.0.2",
		End: "11.0.0.254",
	}
	dhcp.Ranges = append(dhcp.Ranges, *dhcpRange)
	ip.DHCP = dhcp
	netcfg.IPs = append(netcfg.IPs, *ip)

	fmt.Println("netcfg:", netcfg)
	output, _:= xml.MarshalIndent(netcfg, "  ", "    ")
	//os.Stdout.Write(output)
	fmt.Println(string(output))
	res, err := conn.NetworkDefineXML(string(output))
	resName, err := res.GetName()
	fmt.Println("res: ", resName)
	res.SetAutostart(true)
	res.Create()
}

func GetXmlNetwork() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
	defer conn.Close()
	networks, err := conn.ListAllNetworks(0)
	for index, net := range networks {
		name, _ := net.GetName()
		fmt.Println(index, ": ", name, "------------")
		xmlstr, _ := net.GetXMLDesc(0)
		fmt.Println(index, ": ", xmlstr)
		netcfg := &libvirtxml.Network{}
		err = netcfg.Unmarshal(xmlstr)
		fmt.Println("domain", netcfg.Domain)
		fmt.Println("name", netcfg.Name)
		fmt.Println("forward", netcfg.Forward.Mode)
	}
}

