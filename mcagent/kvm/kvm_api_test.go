package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"testing"
)
//
//func dumpVm() {
//	conn, err := libvirt.NewConnect("qemu:///system")
//	if err != nil {
//		fmt.Println("error 1", err)
//	}
//	defer conn.Close()
//
//	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
//	if err != nil {
//		fmt.Println("error 2", err)
//	}
//
//	fmt.Printf("%d running domains:\n", len(doms))
//	for _, dom := range doms {
//		name, err := dom.GetName()
//		if err == nil {
//			fmt.Printf("  %s\n", name)
//		}
//		dom.Free()
//	}
//}
//
//func TestDumpVm(t*testing.T){
//	dumpVm()
//}

func getData() (vm mcmodel.MgoVm, server mcmodel.McServerDetail){
	vm = mcmodel.MgoVm{
		Idx:         1,
		McServerIdx: 1,
		CompanyIdx:  1,
		Name:        "win10-bhjung",
		Cpu:         4,
		Ram:         8192,
		Hdd:         100,
		OS:          "win10",
		Image:       "windows10-100G",
	}
	server = mcmodel.McServerDetail{
		McServer: mcmodel.McServer{
			Idx: 1,
			IpAddr: "192.168.0.73",
		},
	}
	return vm, server
}

func TestCreateVmInstance(t *testing.T) {
	vm, _ := getData()

	fmt.Println("start...")
	CreateVmInstance(vm)
	fmt.Println("finished...")
}

func TestStartVm(t *testing.T) {
	vm, _ := getData()
	StartVm(vm)
}

func TestGetIpAddressOfVm(t *testing.T) {
	vm, _ := getData()
	GetIpAddressOfVm(vm)
}

func TestShutdownVm(t *testing.T) {
	vm, _ := getData()
	ShutdownVm(vm)
}

func TestUndefineVm(t *testing.T) {
	vm, _ := getData()
	UndefineVm(vm)
}

func TestStatusVm(t *testing.T) {
	vm, _ := getData()
	StatusVm(vm)
}

func TestDeleteVm(t *testing.T) {
	vm, _ := getData()
	ShutdownVm(vm)
	UndefineVm(vm)
}

func TestCopyVmInstance(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	vm, _ := getData()
	CopyVmInstance(&vm)
}

func TestGetVmFromLibvirt(t *testing.T) {
	GetVmFromLibvirt()
}

func TestLs(t *testing.T) {
	binary, _:= exec.LookPath("ls")
	args := []string{"-a", "-l", "-h"}
	fmt.Println("args:", args)
	env := os.Environ()
	syscall.Exec(binary, args, env)
}

