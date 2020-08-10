package agent

import (
	"cmpService/mcagent/config"
	"fmt"
	"testing"
)

func TestGetMgoImageByName(t *testing.T) {
	//dir := "/opt/images/"
	name := "/opt/images/windows10-100G.qcow2"
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	GetMgoImageByName(name)
	InitImages()
}

func TestGetImages(t *testing.T) {
	config.ApplyGlobalConfig("../etc/mcagent.conf")
	list := GetImages()
	fmt.Println("list: ", list)
}