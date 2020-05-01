package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nubes/common/models"
	"testing"
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

func TestRegisterUser(t *testing.T) {
	url := restServer + "/v1/register"
	user := models.User{
		ID:"andrew",
		Password: "andrew1510",
		Email: "andrew@nubes-bridge.com",
		Name:"anrew",
		Level: 1,
	}
	pbytes, _ := json.Marshal(user)
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

func TestLoginUser(t *testing.T) {
	url := restServer + "/v1/login"
	user := models.User{
		ID:"andrew",
		Password: "andrew1510",
	}
	pbytes, _ := json.Marshal(user)
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

