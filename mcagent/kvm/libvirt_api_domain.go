package kvm

import (
	"fmt"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

/******************************************************************************
 * Domain
 ******************************************************************************/
func GetDomainListAll() (doms []libvirt.Domain, err error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
		return doms, err
	}
	doms, err = conn.ListAllDomains(0)
	if err != nil {
		fmt.Println("error2")
		return doms, err
	}
	//for index, dom := range doms {
	//	name, _ := dom.GetName()
	//	fmt.Println(index, ":", name)
	//	info, _ := dom.GetInfo()
	//	fmt.Println("info: ", info)
	//	addr, _ := dom.ListAllInterfaceAddresses(0)
	//	fmt.Println("addr: ", addr)
	//}
	return doms, err
}

func GetDomainByName(name string) (dom *libvirt.Domain, err error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
		return dom, err
	}
	dom, err = conn.LookupDomainByName(name)
	if err != nil {
		fmt.Println("error2")
		return dom, err
	}
	return dom, err
}

func GetXmlDomainByName(name string) *libvirtxml.Domain {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("error1")
		return nil
	}
	dom, err := conn.LookupDomainByName(name)
	if err != nil {
		fmt.Println("error2")
		return nil
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

