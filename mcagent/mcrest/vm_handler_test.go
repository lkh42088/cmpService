package mcrest

import (
	"cmpService/mcagent/config"
	"testing"
)

func TestGetMgoServer(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	GetMgoServer()
}