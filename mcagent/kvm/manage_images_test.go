package kvm

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

func TestList(t *testing.T) {
	var list []int
	list = append(list, 1)
	list = append(list, 2)
	fmt.Println(list)
	list = list[:0]
	fmt.Println(list)
}