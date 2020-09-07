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
		fmt.Println("error1")
		return doms, err
	}
	defer conn.Close()
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
	conn, err := GetQemuConnect()
	if err != nil {
		fmt.Println("error1")
		return dom, err
	}
	defer conn.Close()
	dom, err = conn.LookupDomainByName(name)
	if err != nil {
		fmt.Println("error2")
		return dom, err
	}
	return dom, err
}

func GetXmlDomainByName(name string) *libvirtxml.Domain {
	conn, err := GetQemuConnect()
	if err != nil {
		fmt.Println("error1")
		return nil
	}
	defer conn.Close()
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

func KvmShutdownVm(name string) {
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

func KvmSuspendVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("KvmSuspendVm error:", err)
		return
	}
	err = dom.Suspend()
	if err != nil {
		fmt.Println("KvmSuspendVm error:", err)
		return
	}
}

func KvmResumeVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("KvmResumeVm error:", err)
		return
	}
	err = dom.Resume()
	if err != nil {
		fmt.Println("KvmResumeVm error:", err)
		return
	}
}

func GetKvmVmState(name string) string {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("GetKvmVmState error:", err)
		return "unknown"
	}
	status, _, err := dom.GetState()
	if err != nil {
		fmt.Println("GetKvmVmState error:", err)
		return "unknown"
	}
	fmt.Println(name, ConvertVmStatus(status))
	return ConvertVmStatus(status)
}

func KvmDestroyVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("KvmDestroyVm error:", err)
		return
	}
	err = dom.Destroy()
	if err != nil {
		fmt.Println("KvmDestroyVm error:", err)
		return
	}
}

func KvmUndefineVm(name string) {
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

func KvmStartVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("KvmUndefine error:", err)
		return
	}
	err = dom.Create()
	if err != nil {
		fmt.Println("KvmUndefine error:", err)
		return
	}
}

func KvmRebootVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("KvmUndefine error:", err)
		return
	}
	err = dom.Reboot(0)
	if err != nil {
		fmt.Println("KvmUndefine error:", err)
		return
	}
}

func KvmResetVm(name string) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("KvmUndefine error:", err)
		return
	}
	err = dom.Reset(0)
	if err != nil {
		fmt.Println("KvmUndefine error:", err)
		return
	}
}