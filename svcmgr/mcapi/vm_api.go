package mcapi

import (
	"bytes"
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendAddVm(vm mcmodel.McVm, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(vm)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlCreateVm)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendAddVm: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendAddVm: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendDeleteVm(vm mcmodel.McVm, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(vm)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlDeleteVm)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendDeleteVm: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendDeleteVm: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendGetVmById(vm mcmodel.McVm, server mcmodel.McServerDetail) bool {
	url := fmt.Sprintf("http://%s:8082%s%s/%d",server.IpAddr, lib.McUrlPrefix, lib.McUrlVm, vm.Idx)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("SendGetVmById: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendGetVmById: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendGetVmAll(server mcmodel.McServerDetail) bool {
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlVm)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("SendGetVmAll: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendGetVmAll: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}