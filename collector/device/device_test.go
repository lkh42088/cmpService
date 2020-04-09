package device

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestDevice(t *testing.T) {
	d := Device{
		Ip:            "192.168.1.10",
		Port:          500,
		SnmpCommunity: "public",
	}
	b, err :=json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

func TestDevice2(t *testing.T) {
	b := []byte(`{"ip":"192.168.1.1","port":501,"community":"private"}`)
	a := Device{}
	err := json.Unmarshal(b, &a)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(a.Ip)
	fmt.Println(a.Port)
	fmt.Println(a.SnmpCommunity)
	fmt.Println(a)
}

func TestDevice3(t *testing.T) {
	device := Device{
		Ip:            "192.168.1.1",
		Port:          30,
		SnmpCommunity: "public",
	}
	pbyte, _ := json.Marshal(device)
	buff := bytes.NewBuffer(pbyte)
	data := url.Values{}
	data.Set("device", buff.String())
	req, err := http.PostForm( "http://127.0.0.1:8884/api/v1/device", data)
	if err != nil {
		fmt.Println("err is ", err)
	}
	f, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Println(string(f))
}

func TestSNMPDevice1(t *testing.T) {
	device := Device{
		Ip:            "127.0.0.1",
		Port:          161,
		SnmpCommunity: "nubes",
	}
	pbyte, _ := json.Marshal(device)
	buff := bytes.NewBuffer(pbyte)
	req, err := http.Post( "http://127.0.0.1:8884/api/v1/device/json", "application/json", buff)
	if err != nil {
		fmt.Println("err is ", err)
	}
	f, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Println(string(f))
}

func TestSNMPDevice2(t *testing.T) {
	device := Device{
		Ip:            "192.168.122.15",
		Port:          161,
		SnmpCommunity: "nubes",
	}
	pbyte, _ := json.Marshal(device)
	buff := bytes.NewBuffer(pbyte)
	req, err := http.Post( "http://127.0.0.1:8884/api/v1/device/json", "application/json", buff)
	if err != nil {
		fmt.Println("err is ", err)
	}
	f, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Println(string(f))
}

type ResponseError struct {
	Err error
}

type Response struct {
	ID     ID            `json:"id,omitempty"`
	Device Device        `json:"device"`
	Error  ResponseError `json:"error"`
}

func TestDevice5(t *testing.T) {
	req, err := http.Get( "http://127.0.0.1:8884/api/v1/device/1")
	if err != nil {
		fmt.Println("err is ", err)
	}
	body, _ :=ioutil.ReadAll(req.Body)
	req.Body.Close()
	res := Response{}
	json.Unmarshal(body, &res)
	//fmt.Println(string(f))
	fmt.Println(res)
}


