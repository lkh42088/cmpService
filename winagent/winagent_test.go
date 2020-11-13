package main

import (
	"cmpService/winagent/common"
	"testing"
)

func TestCMPWindowService_Execute(t *testing.T) {
	common.InsertMacInTelegrafConf("11:11:11:11:11:22")
}
