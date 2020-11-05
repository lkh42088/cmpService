package mcrest

import (
	"bytes"
	"cmpService/common/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var ipaddr = "192.168.254.18"

func TestGetResource(t *testing.T) {
	url := fmt.Sprintf("http://%s:8082%s%s",
		ipaddr, lib.McUrlPrefix, lib.McUrlResource)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("SendGetVmById: error 1 ", err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendGetVmById: error 2 ", err)
		return
	}
	//fmt.Println("response: ", string(data))
	var resource McResourceMsg
	json.Unmarshal(data, &resource)
	resource.Dump()
}

func TestDeleteResource(t *testing.T) {
	url := fmt.Sprintf("http://%s:8082%s%s",
		ipaddr, lib.McUrlPrefix, lib.McUrlResource)
	msg := ResourceMsg{
		Command: "clear",
	}
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendGetVmById: error 1 ", err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendGetVmById: error 2 ", err)
		return
	}
	fmt.Println("response: ", string(data))
	//var resource McResourceMsg
	//json.Unmarshal(data, &resource)
	//resource.Dump()
}