package agent

import (
	"flag"
	"testing"
)

func TestInsertMacInTelegrafConf(t *testing.T) {
	data := GetSysInfo()
	InsertMacInTelegrafConf(data.IfMac)
}

func TestRestartTelegraf(t *testing.T) {
	RestartTelegraf()
}

func TestRestServer(t *testing.T) {
	conf := flag.String("file", "c:\\Users\\user\\go\\src\\cmpService\\winagent\\winagent.conf",
		"Input configuration file")
	flag.Parse()
	Start(*conf)
}
