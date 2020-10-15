package ktrest

import (
	"cmpService/mcagent/config"
	"flag"
	"fmt"
	"testing"
)

func TestGetUserAuth(t *testing.T) {
	r := PostAuthTokens()
	GetStorageAccount(r)
}

func TestGetKtStorageTempUrl(t *testing.T) {
	GetStorageTempUrl()
}

func TestGetStorageContainer(t *testing.T) {
	_ = PostAuthTokens()
	err := GetStorageContainer("nubes-test")
	fmt.Println(err)
}

func TestPutStorageContainer(t *testing.T) {
	_ = PostAuthTokens()
	err := PutStorageContainer(GlobalToken, "nubes-test")
	fmt.Println(err)
}

func TestDeleteStorageContainer(t *testing.T) {
	_ = PostAuthTokens()
	err := DeleteStorageContainer( "nubes-test")
	fmt.Println(err)
}

func TestDivisionVmSnapshotFile(t *testing.T) {
	conf := flag.String("file", "/home/nubes/go/src/cmpService/mcagent/etc/mcagent.lkh.conf","Input configuration file")
	flag.Parse()
	config.ApplyGlobalConfig(*conf)
	DivisionVmSnapshotFile("windows10-100G-0.qcow2")
}

func TestPutStorageObject(t *testing.T) {
	conf := flag.String("file", "/home/nubes/go/src/cmpService/mcagent/etc/mcagent.lkh.conf","Input configuration file")
	flag.Parse()
	config.ApplyGlobalConfig(*conf)
	_ = PostAuthTokens()
	err := PutStorageObject("nubes-test", "a.txt")
	fmt.Println(err)
}

func TestPutDynamicLargeObjects(t *testing.T) {
	conf := flag.String("file", "/home/nubes/go/src/cmpService/mcagent/etc/mcagent.lkh.conf","Input configuration file")
	flag.Parse()
	config.ApplyGlobalConfig(*conf)
	_ = PostAuthTokens()
	//err := PutDynamicLargeObjects("nubes-test", "windows10-100G-0.qcow2",
	//	"vm_win10.z01")
	//err := PutDynamicLargeObjects("nubes-test", "windows10-100G-0.qcow2",
	//	"vm_win10.z02")
	err := PutDynamicLargeObjects("nubes-test", "windows10-100G-0.qcow2",
		"vm_win10.zip")
	fmt.Println(err)
}

func TestPutDLOManifest(t *testing.T) {
	conf := flag.String("file", "/home/nubes/go/src/cmpService/mcagent/etc/mcagent.lkh.conf","Input configuration file")
	flag.Parse()
	config.ApplyGlobalConfig(*conf)
	_ = PostAuthTokens()
	err := PutDLOManifest("nubes-test", "windows10-100G-0.qcow2")
	fmt.Println(err)
}

func TestGetStorageObject(t *testing.T) {
	_ = PostAuthTokens()
	//err := GetStorageObject("nubes-test", "windows10-100G-0.qcow2")
	err := GetStorageObject("nubes-test", "a.txt")
	fmt.Println(err)
}