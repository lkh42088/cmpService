package ktrest

import (
	"bufio"
	"cmpService/common/download"
	"cmpService/mcagent/config"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/**
 * KT Storage File Upload/Download
 */
// File Division & Zip
func DivisionVmBackupFile(fileName string) (files []string, err error) {
	// Get file path
	dstPath := config.GetMcGlobalConfig().VmBackupDir
	lastPath := dstPath + "/" + fileName
	fmt.Println("PATH: ", lastPath)

	// file check
	if _, err := os.Stat(lastPath); os.IsNotExist(err) {
		fmt.Println(err)
		return files, err
	}

	// system call
	args := []string{
		"-s",
		"500m",
		"-o",
		//dstPath + "/" + fileName + ".zip",
		//lastPath,
		lastPath + ".zip",
		lastPath,
	}
	binary := "zip"
	cmd := exec.Command(binary, args...)
	if _, err := os.Stat(lastPath+".zip"); os.IsNotExist(err) {
		_, err = cmd.Output()
	}
	if err != nil {
		return files, err
	}

	allFiles, err := ioutil.ReadDir(config.GetMcGlobalConfig().VmBackupDir)
	for _, file := range allFiles {
		if strings.Contains(file.Name(), fileName + ".z") {
			files = append(files, file.Name())
		}
	}
	fmt.Println("## FILE LIST : ", files)
	if len(files) == 0 {
		return files, errors.New("Division is failed.")
	}

	return files, nil
}

// UnZip File
func UnZipVmBackupFile(fileName string, dstName string) error {
	// Get file path
	//conf := config.GetMcGlobalConfig()
	//path := conf.VmInstanceDir
	lastPath := "/home/nubes/go/src/cmpService/mcagent/ktrest/" + fileName
	fmt.Println("PATH: ", lastPath)

	// file check
	if _, err := os.Stat(lastPath); os.IsNotExist(err) {
		fmt.Println(err)
		return err
	}

	// system call
	args := []string{
		fileName,
		"-d",
		dstName,
	}
	binary := "unzip"
	cmd := exec.Command(binary, args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("output:", string(output))
	return nil
}

// Put Storage Object
func PutStorageObject(container string, fileName string) error {
	// Get file path
	conf := config.GetMcGlobalConfig()
	path := conf.VmInstanceDir
	lastPath := path + "/" + fileName
	fmt.Println("PATH: ", lastPath)

	// Get file
	fileInfo, err := os.Stat(lastPath)
	if err != nil {
		return fmt.Errorf("Error: Not find this file.\n")
	}
	file, _ := os.Open(lastPath)
	data := bufio.NewReader(file)

	// Request URL
	baseUrl := GlobalAccountUrl + "/" + container + "/" + fileName
	req, _ := http.NewRequest("PUT", baseUrl, data)
	// Request HEADER
	req.Header.Add("X-Auth-Token", GlobalToken)
	req.Header.Add("Content-Type", "test/plain; charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(int(fileInfo.Size())))
	fmt.Println("URL: ", req)

	//Send API Query
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer resp.Body.Close()
	}

	//Parsing data
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Error: %s\n", resp.Status)
	}

	return fmt.Errorf("Success\n")
}

// Upload Backup File (DLO)
func PutDynamicLargeObjects(container string, originFileName string, fileName string) error {
	// Get file path
	conf := config.GetMcGlobalConfig()
	path := conf.VmBackupDir
	lastPath := path + "/" + fileName
	fmt.Println("PATH: ", lastPath)

	// Get file
	fileInfo, err := os.Stat(lastPath)
	if err != nil {
		return fmt.Errorf("Error: Not find this file.\n")
	}
	file, _ := os.Open(lastPath)
	data := bufio.NewReader(file)

	// Request URL
	baseUrl := GlobalAccountUrl + "/" + container + "/" + originFileName + "/" + fileName
	req, _ := http.NewRequest("PUT", baseUrl, data)
	// Request HEADER
	req.Header.Add("X-Auth-Token", GlobalToken)
	//req.Header.Add("Content-Type", ContentTypeJson)
	req.Header.Add("Content-Type", ContentTypeBinary)
	req.Header.Add("Range", Range4096)
	req.Header.Add("Content-Length", strconv.Itoa(int(fileInfo.Size())))
	fmt.Println("URL: ", req)

	//Send API Query
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer resp.Body.Close()
	}

	//Parsing data
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Error: %s\n", resp.Status)
	}

	return nil
}

// Put Dynamic Large Object Manifest File
func PutDLOManifest(container string, originFileName string) error {
	// Request URL : https://ssproxy2.ucloudbiz.olleh.com/v1/AUTH_fa632a4a0d04488c93b7184be92df4c8/SN87_87/vm-01-cronsch.qcow2.decrease
	baseUrl := GlobalAccountUrl + "/" + container + "/" + originFileName
	req, _ := http.NewRequest("PUT", baseUrl, nil)
	// Request HEADER
	req.Header.Add("X-Auth-Token", GlobalToken)
	req.Header.Add("X-Object-Manifest", container + "/" + originFileName + "/")
	//req.Header.Add("Content-Type", ContentTypeJson)
	req.Header.Add("Content-Type", ContentTypeBinary)
	req.Header.Add("Content-Length", "0")
	fmt.Println("URL: ", req)

	//Send API Query
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return err
	} else {
		defer resp.Body.Close()
	}

	//Parsing data
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Error: %s\n", resp.Status)
	}

	return nil
}

// Get Storage Object (File download)
func GetStorageObject(container string, fileName string, ch chan int) error {
	conf := config.GetMcGlobalConfig()
	path := conf.VmInstanceDir
	lastPath := path + "/" + fileName
	fmt.Println("PATH: ", lastPath)

	// Request URL
	baseUrl := GlobalAccountUrl + "/" + container + "/" + fileName
	req, _ := http.NewRequest("GET", baseUrl, nil)
	// Request HEADER
	req.Header.Add("X-Auth-Token", GlobalToken)
	req.Header.Add("Content-Type", ContentTypeBinary)
	//req.Header.Add("Range", RANGE_4096)
	fmt.Println("URL: ", req)

	ch <- 1 // file transfer start

	//Send API Query
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer resp.Body.Close()
	}
	ch <- 5	// receive complete

	//Parsing data
	if resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Error: %s\n", resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error: %s\n", err)
	}
	//err = ioutil.WriteFile(lastPath, data, BACKUP_FILE_PERMISSION)
	err = ioutil.WriteFile(path + "/" + "test_win10.qcow2", data, BackupFilePermission)	// test code
	if err != nil {
		return fmt.Errorf("Error: %s\n", err)
	}

	ch <- 10	// success

	return fmt.Errorf("Success\n")
}

// Get Storage Object (File download)
func GetStorageObjectByDLO(container string, fileName string, ch chan int) error {
	conf := config.GetMcGlobalConfig()
	path := conf.VmInstanceDir
	lastPath := path + "/" + fileName
	fmt.Println("PATH: ", lastPath)

	// Request URL
	baseUrl := GlobalAccountUrl + "/" + container + "/" + fileName
	req, _ := http.NewRequest("GET", baseUrl, nil)
	// Request HEADER
	req.Header.Add("X-Auth-Token", GlobalToken)
	req.Header.Add("Content-Type", ContentTypeBinary)
	req.Header.Add("Range", Range4096)
	fmt.Println("URL: ", req)

	ch <- 1 // file transfer start
	download.DownloadLargeObjectForKtStorage(baseUrl, GlobalToken)
	ch <- 5	// receive complete

	return fmt.Errorf("Success\n")
}

func PrintDownloading(ch chan int) {
	v := <- ch
	switch v {
	case 1:
		fmt.Println("FILE TRANSFER START.")
	case 5:
		fmt.Println("FILE RECEIVE COMPLETE.")
	case 10:
		fmt.Println("FILE TRANSFER SUCCESS.")
	}
}

// DELETE Storage Object (File delete)
func DeleteStorageObject(container string, fileName string) error {
	// Request URL
	baseUrl := GlobalAccountUrl + "/" + container + "/" + fileName
	req, _ := http.NewRequest("DELETE", baseUrl, nil)
	// Request HEADER
	req.Header.Add("X-Auth-Token", GlobalToken)
	req.Header.Add("Content-Type", ContentTypeJson)
	//fmt.Println("URL: ", req)

	//Send API Query
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer resp.Body.Close()
	}

	//Parsing data
	if resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Error: %s\n", resp.Status)
	}

	return fmt.Errorf("Success\n")
}
