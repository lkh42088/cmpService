package rest

import (
	"bytes"
	config2 "cmpService/common/config"
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"cmpService/svcmgr/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func getServer() *mcmodel.McServer{
	return &mcmodel.McServer{
		SerialNumber: "NB_SN_001",
		CompanyIdx: 191,
		Type: "Standard",
		IpAddr: "192.168.0.73",
	}
}

func getVm() *mcmodel.McVm {
	return &mcmodel.McVm{
		McServerIdx: 1,
		CompanyIdx: 191,
		Name: "win10-bhjung",
		Cpu: 4,
		Ram: 8192,
		Hdd: 100,
		OS: "win10",
		Image: "windowns10-100G",
	}
}

var cfg = config2.MariaDbConfig{
	"192.168.0.40",
	"nubes",
	"nubes",
	"Nubes1510!",
}

var svcmgrAddr = "192.168.0.72"

func SendSvcAddServer(server mcmodel.McServer, addr string) bool {
	pbytes, _ := json.Marshal(server)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8081%s", addr, lib.SvcmgrApiMicroServerRegister)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error 1:", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2:", err)
		return false
	}
	fmt.Println("response:", string(data))
	return true
}

func TestAddMcServer(t*testing.T) {
	server := getServer()
	SendSvcAddServer(*server, svcmgrAddr)
}

func SendSvcAddVm(vm mcmodel.McVm, addr string) bool {
	pbytes, _ := json.Marshal(vm)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8081%s", addr, lib.SvcmgrApiMicroVmRegister)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error 1:", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2:", err)
		return false
	}
	fmt.Println("response:", string(data))
	return true
}

func SendSvcDeleteVm(vm mcmodel.McVm, addr string) bool {
	var msg messages.DeleteDataMessage
	msg.IdxList = append(msg.IdxList, int(vm.Idx))
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8081%s", addr, lib.SvcmgrApiMicroVmUnRegister)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error 1:", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2:", err)
		return false
	}
	fmt.Println("response:", string(data))
	return true
}

func TestAddMcVm(t*testing.T) {
	vm := getVm()
	SendSvcAddVm(*vm, svcmgrAddr)
}

func TestDeleteMcServer(t*testing.T) {
}


func TestDeleteMcVm(t*testing.T) {
	vm := getVm()
	db, err := config.SetMariaDB(
		cfg.MariaUser, cfg.MariaPassword, cfg.MariaDb,
		cfg.MariaIp, 3306)
	if err != nil {
		fmt.Println("db error: ", err)
		return
	}
	fmt.Printf("1. vm %v\n", vm)
	rvm, _ := db.GetMcVmByNameAndCpIdx(vm.Name, vm.CompanyIdx)
	fmt.Printf("2. vm %v\n", rvm)
	SendSvcDeleteVm(rvm, svcmgrAddr)
}


