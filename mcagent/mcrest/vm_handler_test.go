package mcrest

import (
	"cmpService/mcagent/config"
	"cmpService/mcagent/mcinflux"
	"testing"
)

func TestGetMgoServer(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	GetMgoServer()
}

func TestGetVmInterfaceTrafficByMac(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	mcinflux.ConfigureInfluxDB()
	//GetVmInterfaceTrafficByMac("fe:54:00:d9:f7:6c")
}