package svcmgrapi

import (
	"cmpService/common/mcmodel"
	"testing"
)

func TestSendUpdateVm2Svcmgr(t *testing.T) {
	vm := mcmodel.MgoVm{
		Idx:         1,
		McServerIdx: 1,
		CompanyIdx:  1,
		Name:        "vm1",
		Cpu:         4,
		Ram:         8192,
		Hdd:         100,
		OS:          "window10",
		Image:       "window10.qcow2",
	}
	SendUpdateVm2Svcmgr(vm, "192.168.0.72:8081")
}