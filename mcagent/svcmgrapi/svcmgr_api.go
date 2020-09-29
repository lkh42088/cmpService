package svcmgrapi

import (
	"bytes"
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendMcVmSnapshot2Svcmgr(obj mcmodel.McVmSnapshot, addr string) bool {
	pbytes, _ := json.Marshal(obj)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s%s", addr, lib.SvcmgrApiMicroMcAgentVmAddSnapshot)
	fmt.Println("Notify: ", url)
	obj.Dump()
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error: ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2: ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendUpdateServer2Svcmgr(obj mcmodel.McServerMsg, addr string) bool {
	pbytes, _ := json.Marshal(obj)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s%s", addr, lib.SvcmgrApiMicroServerResource)
	fmt.Println("Notify: ", url)
	obj.Dump()
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error: ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2: ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendUpdateVm2Svcmgr(vm mcmodel.McVm, addr string) bool {
	pbytes, _ := json.Marshal(vm)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s%s", addr, lib.SvcmgrApiMicroVmUpdateFromMc)
	fmt.Println("Notify: ", url)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error: ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2: ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

// Baremetal system info
func SendSysInfoToSvcmgr(info mcmodel.SysInfo, addr string) bool {
	pbytes, _ := json.Marshal(info)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s%s", addr, lib.SvcmgrApiMicroSystemInfo)
	//fmt.Println("Notify: ", url)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("post error: ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("response error: ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}
