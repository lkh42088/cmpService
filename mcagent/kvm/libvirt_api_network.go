package kvm

import (
	"cmpService/common/mcmodel"
	"encoding/xml"
	"fmt"
	"github.com/go-xmlfmt/xmlfmt"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"strconv"
	"strings"
)

/******************************************************************************
 * Network
 ******************************************************************************/
func GetAllNetwork() (networks []libvirt.Network, err error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("GetAllNetwork: error", err)
		return networks, err
	}
	defer conn.Close()
	networks, err = conn.ListAllNetworks(0)
	//for index, net := range networks {
	//	name, _ := net.GetName()
	//	bridge, _ := net.GetBridgeName()
	//	fmt.Println(index, ": ", name, bridge)
	//}
	return networks, err
}

func GetNetworkByName(name string) (*libvirt.Network, error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
	defer conn.Close()
	return conn.LookupNetworkByName("net11")
}

func GetXmlNetworkByName() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
	defer conn.Close()
	net, err := conn.LookupNetworkByName("net11")
	name, _ := net.GetName()
	fmt.Println(name, "------------")
	xmlstr, _ := net.GetXMLDesc(0)
	fmt.Println(xmlstr)
	netcfg := &libvirtxml.Network{}
	err = netcfg.Unmarshal(xmlstr)
	fmt.Println("domain", netcfg.Domain)
	fmt.Println("name", netcfg.Name)
	fmt.Println("forward", netcfg.Forward.Mode)
	res := net.Undefine()
	res = net.Destroy()
	fmt.Println(res)
}

func GetMgoNetworksFromXmlNetwork() (list []mcmodel.MgoNetwork, err error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
	defer conn.Close()
	networks, err := conn.ListAllNetworks(0)
	for index, net := range networks {
		var entry mcmodel.MgoNetwork
		name, _ := net.GetName()
		fmt.Println(index, ": ", name, "------------")
		xmlstr, _ := net.GetXMLDesc(0)
		//fmt.Println(index, ": ", xmlstr)
		netcfg := &libvirtxml.Network{}
		err = netcfg.Unmarshal(xmlstr)
		fmt.Println("domain", netcfg.Domain)
		fmt.Println("name", netcfg.Name)
		fmt.Println("forward", netcfg.Forward.Mode)
		entry.Name = netcfg.Name
		entry.Mode = netcfg.Forward.Mode
		entry.Uuid = netcfg.UUID
		entry.Bridge = netcfg.Bridge.Name
		for index, Ip := range netcfg.IPs {
			if index == 0 {
				entry.Ip = Ip.Address
				entry.Netmask = Ip.Netmask
				entry.Prefix = Ip.Prefix
				break
			}
		}
		list = append(list, entry)
	}
	return list, err
}

func MakeXmlNetwork(name, bridgeName, ipAddr, netmask string) *libvirtxml.Network {
	netcfg := &libvirtxml.Network{}
	// name
	netcfg.Name = name

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
	bridge.Name = bridgeName
	bridge.STP = "on"
	bridge.Delay = "0"
	netcfg.Bridge = bridge

	// mac
	// ip/dhcp
	ip := &libvirtxml.NetworkIP{}
	ip.Address = ipAddr
	ip.Netmask = netmask
	addrArray := strings.Split(ipAddr, ".")
	lastNum, _ := strconv.Atoi(addrArray[3])
	dhcp := &libvirtxml.NetworkDHCP{}
	dhcpRange := &libvirtxml.NetworkDHCPRange{
		Start: fmt.Sprintf("%s.%s.%s.%d",
			addrArray[0],
			addrArray[1],
			addrArray[2],
			lastNum + 1),
		End: fmt.Sprintf("%s.%s.%s.254",
			addrArray[0],
			addrArray[1],
			addrArray[2]),
	}
	fmt.Println(dhcpRange)
	dhcp.Ranges = append(dhcp.Ranges, *dhcpRange)
	ip.DHCP = dhcp
	netcfg.IPs = append(netcfg.IPs, *ip)
	return netcfg
}

func CreateNetworkByMgoNetwork(net mcmodel.MgoNetwork) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
	defer conn.Close()
	netcfg := MakeXmlNetwork(net.Name, net.Bridge, net.Ip, net.Netmask)
	output, _:= xml.MarshalIndent(netcfg, "  ", "    ")
	fmt.Println(string(output))
	res, err := conn.NetworkDefineXML(string(output))
	resName, err := res.GetName()
	fmt.Println("res: ", resName)
	res.SetAutostart(true)
	res.Create()
}

func DeleteNetwork(name string) {
	net, err := GetNetworkByName(name)
	if err != nil {
		return
	}
	net.Undefine()
	net.Destroy()
}

func DumpXml(output string) {
	x := xmlfmt.FormatXML(output, "\t", "  ")
	print(x)
}

