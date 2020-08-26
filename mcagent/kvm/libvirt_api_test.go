package kvm

import (
	"cmpService/mcagent/config"
	"testing"
)

func TestGetMcVirtInfo(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	GetMcVirtInfoDebug()
}

func TestDumpMcVirtInfo(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	DumpMcVirtInfo()
}

func TestGetMcServerInfo(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	server := GetMcServerInfo()
	server.Dump()
}