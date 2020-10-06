package mcapi

import (
	"bytes"
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendMcVmAction(msg messages.McVmActionMsg, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlApplyVmAction)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendMcVmAction: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendMcVmAction: error 2 ", err)
		return false
	}
	fmt.Println("SendMcVmAction: success - ", string(data))
	return true
}

func SendAddVmSnapshot(msg messages.SnapshotConfigMsg, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlAddVmSnapshot)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendAddVmSnapshot: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendAddVmSnapshot: error 2 ", err)
		return false
	}
	fmt.Println("SendAddVmSnapshot: success - ", string(data))
	return true
}

func SendDeleteVmSnapshot(msg messages.SnapshotConfigMsg, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlDeleteVmSnapshot)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendDeleteVmSnapshot: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendDeleteVmSnapshot: error 2 ", err)
		return false
	}
	fmt.Println("SendDeleteVmSnapshot: success - ", string(data))
	return true
}

func SendDeleteVmSnapshotList(msg messages.SnapshotEntryMsg, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlDeleteVmSnapshotList)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendDeleteVmSnapshotList: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendDeleteVmSnapshotList: error 2 ", err)
		return false
	}
	fmt.Println("SendDeleteVmSnapshotList : success - ", string(data))
	return true
}

func SendUpdateVmSnapshot(msg messages.SnapshotConfigMsg, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlUpdateVmSnapshot)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendUpdateVmSnapshot: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendUpdateVmSnapshot: error 2 ", err)
		return false
	}
	fmt.Println("SendUpdateVmSnapshot: success - ", string(data))
	return true
}

func SendUpdateVmStatus(msg messages.VmStatusActionMsg, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlUpdateVmSnapshot)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendUpdateVmStatus: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendUpdateVmStatus: error 2 ", err)
		return false
	}
	fmt.Println("SendUpdateVmStatus: success - ", string(data))
	return true
}

func SendRecoverySnapshot(msg messages.SnapshotEntry, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlRecoveryVmSnapshot)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendRecoverySnapshot: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendRecoverySnapshot: error 2 ", err)
		return false
	}
	fmt.Println("SendRecoverySnapshot: success - ", string(data))
	return true
}
