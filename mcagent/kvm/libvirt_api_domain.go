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
	//conn, err := libvirt.NewConnect("qemu:///system")
	conn, err := GetQemuConnect()
	if err != nil {
		fmt.Println("GetDomainListAll error1", err)
		return doms, err
	}
	doms, err = conn.ListAllDomains(0)
	if err != nil {
		fmt.Println("GetDomainListAll error2", err)
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
	conn, err := GetQemuConnect()
	if err != nil {
		fmt.Println("GetDomainByName error1")
		return dom, err
	}
	dom, err = conn.LookupDomainByName(name)
	if err != nil {
		fmt.Println("GetDomainByName error2")
		return dom, err
	}
	return dom, err
}

func GetXmlDomainByName(name string) *libvirtxml.Domain {
	conn, err := GetQemuConnect()
	if err != nil {
		fmt.Println("GetXmlDomainByName error1")
		return nil
	}

	dom, err := conn.LookupDomainByName(name)
	if err != nil {
		fmt.Println("GetXmlDomainByName error2")
		return nil
	}

	xmldoc, err :=dom.GetXMLDesc(0)
	if err != nil {
		fmt.Println("GetXmlDomainByName error3")
		return nil
	}

	fmt.Println("xml:", xmldoc)
	domcfg := &libvirtxml.Domain{}
	err = domcfg.Unmarshal(xmldoc)
	if err != nil {
		fmt.Println("GetXmlDomainByName error4")
		return nil
	}
	return domcfg
}

func LibvirtShutdownVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("ShutdownVm error:", err)
		return
	}
	err = dom.Shutdown()
	if err != nil {
		fmt.Println("ShutdownVm error:", err)
		return
	}
}

func LibvirtSuspendVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("LibvirtSuspendVm error:", err)
		return
	}
	err = dom.Suspend()
	if err != nil {
		fmt.Println("LibvirtSuspendVm error:", err)
		return
	}
}

func LibvirtResumeVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("LibvirtResumeVm error:", err)
		return
	}
	err = dom.Resume()
	if err != nil {
		fmt.Println("LibvirtResumeVm error:", err)
		return
	}
}

func GetLibvirtVmState(name string) string {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("GetLibvirtVmState error:", err)
		return "unknown"
	}
	status, _, err := dom.GetState()
	if err != nil {
		fmt.Println("GetLibvirtVmState error:", err)
		return "unknown"
	}
	fmt.Println(name, ConvertVmStatus(status))
	return ConvertVmStatus(status)
}

func LibvirtDestroyVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("LibvirtDestroyVm error:", err)
		return
	}
	err = dom.Destroy()
	if err != nil {
		fmt.Println("LibvirtDestroyVm error:", err)
		return
	}
}

func LibvirtUndefineVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("KvmUndefine error:", err)
		return
	}
	err = dom.Undefine()
	if err != nil {
		fmt.Println("KvmUndefine error:", err)
		return
	}
}

func LibvirtStartVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("KvmStart error:", err)
		return
	}
	err = dom.Create()
	if err != nil {
		fmt.Println("KvmStart error:", err)
		return
	}
}

func LibvirtRebootVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("LibvirtRebootVm error:", err)
		return
	}
	err = dom.Reboot(0)
	if err != nil {
		fmt.Println("LibvirtRebootVm error:", err)
		return
	}
}

func LibvirtResetVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("LibvirtResetVm error:", err)
		return
	}
	err = dom.Reset(0)
	if err != nil {
		fmt.Println("LibvirtResetVm error:", err)
		return
	}
}
