package agent

import (
	"cmpService/winagent/common"
	"flag"
	"testing"
)

func TestInsertMacInTelegrafConf(t *testing.T) {
	data := common.GetSysInfo()
	common.InsertMacInTelegrafConf(data.IfMac)
}

func TestRestartTelegraf(t *testing.T) {
	common.RestartTelegraf()
}

func TestRestServer(t *testing.T) {
	conf := flag.String("file", "c:\\Users\\user\\go\\src\\cmpService\\winagent\\winagent.conf",
		"Input configuration file")
	flag.Parse()
	Start(*conf)
}
