package rest

import (
	"bytes"
	"cmpService/common/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

var restServer = "http://localhost:8081"

func TestRestAddDeviceMonitoring(t *testing.T) {
	url := restServer + "/v1/devices/monitoring"

	msg := DeviceMonitoringRequest{
		"id-0001",
		"agent",
	}
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)

	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("response:", string(data))
}

func TestRestAddCode(t *testing.T) {
	url := restServer + "/v1/code"

	code := models.Code{
		CodeID:  0,
		Type:    "type1",
		SubType: "subtype1",
		Name:    "name1",
		Order:   1,
	}
	pbytes, _ := json.Marshal(code)
	buff := bytes.NewBuffer(pbytes)

	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error 1: ", err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2: ", err)
		return
	}
	fmt.Println("response:", string(data))
}

func TestRestGetCode(t *testing.T) {
	url := restServer + "/v1/codes"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("response:", string(data))
}

func TestRestDeleteCode(t *testing.T) {
	url := restServer + "/v1/codes"
	resquest, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	client := &http.Client{}
	response, err := client.Do(resquest)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("response", string(data))
}


func TestAddDevice(t *testing.T) {
	commentlast, _ := time.Parse(time.RFC3339, "2019-12-22T10:28:44+09:00")
	data := models.DeviceServer{
		DeviceCommon: models.DeviceCommon{
			OutFlag:false,
			CommentCnt:0,
			CommentLastDate:commentlast,
			RegisterId:"hjt0601",
			DeviceCode:"CBS09999",
			Model:83,
			Contents:"\u003cbr /\u003e",
			Customer:"NB",
			Manufacture:56,
			DeviceType:7,
			WarehousingDate:"0",
			RentDate:"|",
			Ownership:"2|3",
			OwnerCompany:"hddigital",
			HwSn:"MY11353392",
			IDC:16,
			Rack:163,
			Cost:"",
			Purpos:"",
			MonitoringFlag:0,
			MonitoringMethod:0},
		Ip:"220.90.201.198|",
		Size:19,
		Spla:"|",
		Cpu:"",
		Memory:"",
		Hdd:"",
		RackTag:  "",
		RackLoc:  0,
	}

	url := "http://0.0.0.0:8081/v1/device/create/server"
	pbytes, _ := json.Marshal(data)
	buff := bytes.NewBuffer(pbytes)

	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error 1: ", err)
		return
	}
	defer response.Body.Close()
	d, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2: ", err)
		return
	}
	fmt.Println("response:", string(d))
}
