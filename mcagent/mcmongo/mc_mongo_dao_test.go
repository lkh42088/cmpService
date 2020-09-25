package mcmongo

import (
	//"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"fmt"
	"os"
	"testing"
)

func configure() {
	path, _ := os.Getwd()
	fmt.Println("current path: ", path);
	file := "../etc/mcagent.conf"
	config.ApplyGlobalConfig(file)
	Configure()
}

//func TestMcMongoAddVm(t *testing.T) {
//	configure()
//	vm := mcmodel.MgoVm{
//		Idx:         1,
//		McServerIdx: 1,
//		CompanyIdx:  1,
//		Name:        "vm1",
//		Cpu:         4,
//		Ram:         8192,
//		Hdd:         100,
//		OS:          "window10",
//		Image:       "window10.qcow2",
//	}
//	fmt.Println("vm:", vm)
//	McMongo.AddVm(&vm)
//}
//
//func TestMcMongoUpdateVm(t *testing.T) {
//	configure()
//	vm := mcmodel.MgoVm{
//		Idx:         1,
//		McServerIdx: 1,
//		CompanyIdx:  1,
//		Name:        "vm1",
//		Cpu:         4,
//		Ram:         8192,
//		Hdd:         100,
//		OS:          "window10",
//		Image:       "window10.qcow2",
//		CurrentStatus: "running",
//	}
//	fmt.Println("vm:", vm)
//	McMongo.AddVm(&vm)
//}
//
//func TestMcMongoGetVmById(t *testing.T) {
//	configure()
//	vm, _ := McMongo.GetVmById(1)
//	fmt.Println("vm:", vm)
//}
//
//func TestMcMongoGetVmAll(t *testing.T) {
//	configure()
//	vms, err := McMongo.GetVmAll()
//	if err != nil {
//		fmt.Println("err:", err)
//	}
//	fmt.Println("vms:", vms)
//}

func TestMcMongoRemoveVm(t *testing.T) {
	configure()
	McMongo.DeleteVm(2)
}
