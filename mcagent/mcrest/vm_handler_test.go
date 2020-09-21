package mcrest

import (
	"cmpService/common/utils"
	"cmpService/mcagent/config"
	"cmpService/mcagent/mcinflux"
	"testing"
)

func TestGetMgoServer(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	GetMcServer()
}

func TestGetVmInterfaceTrafficByMac(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	mcinflux.ConfigureInfluxDB()
	//GetVmInterfaceTrafficByMac("fe:54:00:d9:f7:6c")
}

func TestGetMyPublicIp(t *testing.T) {
	utils.GetMyPublicIp()
}