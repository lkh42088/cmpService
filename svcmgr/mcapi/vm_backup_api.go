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

func SendAddVmBackup(msg messages.BackupConfigMsg, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlAddVmSnapshot)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendAddVmBackup: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendAddVmBackup: error 2 ", err)
		return false
	}
	fmt.Println("SendAddVmBackup: success - ", string(data))
	return true
}

func SendDeleteVmBackup(msg messages.BackupConfigMsg, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlDeleteVmSnapshot)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendDeleteVmBackup: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendDeleteVmBackup: error 2 ", err)
		return false
	}
	fmt.Println("SendDeleteVmBackup: success - ", string(data))
	return true
}

func SendDeleteVmBackupList(msg messages.BackupEntryMsg, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlDeleteVmSnapshotList)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendDeleteVmBackupList: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendDeleteVmBackupList: error 2 ", err)
		return false
	}
	fmt.Println("SendDeleteVmBackupList : success - ", string(data))
	return true
}

func SendUpdateVmBackup(msg messages.BackupConfigMsg, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlUpdateVmSnapshot)

	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendUpdateVmBackup: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendUpdateVmBackup: error 2 ", err)
		return false
	}
	fmt.Println("SendUpdateVmBackup: success - ", string(data))
	return true
}

func SendRecoveryBackup(msg messages.BackupEntry, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(msg)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlRecoveryVmSnapshot)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendRecoveryBackup: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendRecoveryBackup: error 2 ", err)
		return false
	}
	fmt.Println("SendRecoveryBackup: success - ", string(data))
	return true
}