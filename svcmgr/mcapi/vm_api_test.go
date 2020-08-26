package mcapi

import (
	"cmpService/common/mcmodel"
	"testing"
)

func getData() (vm mcmodel.McVm, server mcmodel.McServerDetail){
	vm = mcmodel.McVm{
		Idx: 1,
		McServerIdx: 1,
		CompanyIdx: 1,
		Name: "win10-bhjung",
		Cpu: 4,
		Ram: 8192,
		Hdd: 100,
		OS: "win10",
		Image: "windows10-100G",
	}
	server = mcmodel.McServerDetail{
		McServer: mcmodel.McServer{
			Idx: 1,
			SerialNumber: "NB_SN_001",
			CompanyIdx: 191,
			Type: "Standard",
			IpAddr: "192.168.0.89",
		},
		CompanyName: "Nubes-Bridge",
	}
	return vm, server
}

func TestSendAddVm(t*testing.T) {
	vm, server := getData()
	SendAddVm(vm, server)
}

func TestSendDeleteVm(t*testing.T) {
	vm, server := getData()
	SendDeleteVm(vm, server)
}

func TestSendGetVmById(t*testing.T) {
	vm, server := getData()
	SendGetVmById(vm, server)
}

func TestSendGetServer(t*testing.T) {
	_, server := getData()
	SendGetMcServer(server)
}

func TestSendGetVmAll(t*testing.T) {
	_, server := getData()
	SendGetVmAll(server)
}