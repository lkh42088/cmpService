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

func SendAddVm(vm mcmodel.McVm) bool {
	pbytes, _ := json.Marshal(vm)
	buff := bytes.NewBuffer(pbytes)
	response, err := http.Post(lib.McUrlCreateVm, "application/json", buff)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendDeleteVm(vm mcmodel.McVm) bool {
	pbytes, _ := json.Marshal(vm)
	buff := bytes.NewBuffer(pbytes)
	response, err := http.Post(lib.McUrlDeleteVm, "application/json", buff)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

