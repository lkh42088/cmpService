package agent

import (
	"cmpService/mcagent/config"
	"testing"
)

func TestGetMgoImageByName(t *testing.T) {
	//dir := "/opt/images/"
	name := "/opt/images/windows10-100G.qcow2"
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	GetMgoImageByName(name)
	InitImages()
}
