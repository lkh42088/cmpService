package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"io/ioutil"
	"net/http"
	"nubes/collector/device"
	"sync"
	"testing"
)

func TestRestRouter(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	RestRouter(&wg)
	wg.Wait()
}

func TestRestGet(t *testing.T) {
	req, err := http.NewRequest("GET",
		"http://localhost:8884" + apiPathPrefix + apiDevice, nil)
	if err != nil {
		fmt.Println("NewRequest err:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("response err:", err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("resp:", string(data))
}

func TestRestGet2(t *testing.T) {
	resp, err := http.Get("http://localhost:8884" + apiPathPrefix + apiDevice)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("resp:", string(data))
}

func TestRestPost(t *testing.T) {
	dev := device.Device{
		Id:            "",
		Ip:            "192.168.122.19",
		Port:          161,
		SnmpCommunity: "nubes",
	}
	pbytes, _ := json.Marshal(dev)
	buff := bytes.NewBuffer(pbytes)
	url := "http://localhost:8884" + apiPathPrefix + apiDevice
	resp, err := http.Post(url, "application/json", buff)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("resp:", string(data))
}

func TestRestDelete(t *testing.T) {
	req, err := http.NewRequest("DELETE",
		"http://localhost:8884" + apiPathPrefix + apiDevice + "/all", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("response err:", err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("resp:", string(data))
}

func TestId(t *testing.T) {
	objID := bson.NewObjectId()
	id := device.ID(fmt.Sprintf("%x",string(objID)))
	fmt.Printf("%s\n", id)
	fmt.Printf("%s\n", string(objID))
}
