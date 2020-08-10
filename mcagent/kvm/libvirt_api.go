package kvm

import (
	"cmpService/common/mcmodel"
	"fmt"
	"github.com/go-xmlfmt/xmlfmt"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func GetXmlDomain(name string) *libvirtxml.Domain{
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
	dom, err := conn.LookupDomainByName("win10-bhjung")
	if err != nil {
		fmt.Println("error2")
	}
	xmldoc, err :=dom.GetXMLDesc(0)
	if err != nil {
		fmt.Println("error3")
	}
	fmt.Println("xml:", xmldoc)

	domcfg := &libvirtxml.Domain{}
	err = domcfg.Unmarshal(xmldoc)
	if err != nil {
		fmt.Println("error4")
	}
	return domcfg
}

func GetDomain() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
	doms, err := conn.ListAllDomains(0)
	if err != nil {
		fmt.Println("error2")
	}
	for index, dom := range doms {
		name, _ := dom.GetName()
		fmt.Println(index, ":", name)
		info, _ := dom.GetInfo()
		fmt.Println("info: ", info)
		addr, _ := dom.ListAllInterfaceAddresses(0)
		fmt.Println("addr: ", addr)
	}
}

func GetNetworksFromXml() (list []mcmodel.MgoNetwork, err error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
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

func GetXmlNetwork() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
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

func GetAllNetwork() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
	}
	networks, err := conn.ListAllNetworks(0)
	for index, net := range networks {
		name, _ := net.GetName()
		bridge, _ := net.GetBridgeName()
		fmt.Println(index, ": ", name, bridge)
	}
}

func DumpXml(output string) {
	x := xmlfmt.FormatXML(output, "\t", "  ")
	print(x)
}