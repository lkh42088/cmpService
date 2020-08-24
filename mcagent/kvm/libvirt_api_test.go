package kvm

import (
	"cmpService/mcagent/config"
	"testing"
)

func TestGetMcVirtInfo(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	GetMcVirtInfo()
}