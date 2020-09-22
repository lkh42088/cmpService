package config

import (
	"cmpService/common/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type McServerStatus struct {
	init         bool   `json:"-"`
	Enable       bool   `json:"enable"`
	SerialNumber string `json:"serialNumber"`
	CompanyName  string `json:"companyName"`
	CompanyIdx   int    `json:"companyIdx"`
}

var serverStatus McServerStatus

func GetServerStatus() *McServerStatus {
	cfg := GetMcGlobalConfig()

	if serverStatus.init {
		return &serverStatus
	}

	info, err := os.Stat(cfg.ServerStatusRepo)
	if os.IsNotExist(err) {
		fmt.Println("GetServerStatus: dose not exist config!")
		return &serverStatus
	}
	if info.IsDir() {
		fmt.Println("GetServerStatus: the config is directory!")
		return &serverStatus
	}
	b, err := ioutil.ReadFile(cfg.ServerStatusRepo)
	if err != nil {
		fmt.Println("GetServerStatus : err ", err)
		return &serverStatus
	}
	err = json.Unmarshal(b, &serverStatus)
	if err != nil {
		fmt.Println("GetServerStatus : err 2 ", err)
		return &serverStatus
	}
	serverStatus.init = true
	return &serverStatus
}

func GetSerialNumber() string {
	cfg := GetMcGlobalConfig()
	if cfg.SerialNumber == "" {
		serverStatus := GetServerStatus()
		SetSerialNumber2GlobalConfig(serverStatus.SerialNumber)
	}
	return GetMcGlobalConfig().SerialNumber
}

func WriteServerStatus(sn, cpName string, cpIdx int, isEnable bool) {
	cfg := GetMcGlobalConfig()
	serverStatus.CompanyName = cpName
	serverStatus.CompanyIdx = cpIdx
	serverStatus.SerialNumber = sn
	serverStatus.Enable = isEnable
	lib.WriteJsonFile(cfg.ServerStatusRepo, &serverStatus)
}

func DeleteServerStatus() {
	cfg := GetMcGlobalConfig()
	serverStatus.Enable = false
	lib.WriteJsonFile(cfg.ServerStatusRepo, &serverStatus)
}

