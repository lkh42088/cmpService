package mcrest

import (
	"bytes"
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/utils"
	"cmpService/mcagent/config"
	"cmpService/mcagent/mcinflux"
	config2 "cmpService/svcmgr/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetServer(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	GetMcServer()
}

func TestDeleteVm(t *testing.T) {
	url := fmt.Sprintf("http://%s:8082%s%s",
		ipaddr, lib.McUrlPrefix, lib.McUrlDeleteVm)
	msg := mcmodel.McVm{
		Name: "vm01",
	}
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendDeleteVm: error 1 ", err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendDeleteVm: error 2 ", err)
		return
	}
	fmt.Println("response: ", string(data))
}

func TestGetVmInterfaceTrafficByMac(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	mcinflux.ConfigureInfluxDB()
	//GetVmInterfaceTrafficByMac("fe:54:00:d9:f7:6c")
}

func TestGetMyPublicIp(t *testing.T) {
	utils.GetMyPublicIp()
}

func TestDeleteVmBackup(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.lkh.conf")
	cfg := config.GetMcGlobalConfig()
	db, _ := config2.SetMariaDB(cfg.MariaUser, cfg.MariaPassword, cfg.MariaDb,
		cfg.MariaIp, 3306)
	config.SetDbOrm(db)
	DeleteVmBackup("SN87-VM-01")
}
