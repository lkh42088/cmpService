package ktrest

import (
	"bytes"
	"cmpService/common/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// For Backup
func SendUpdateAuthUrl2Svcmgr(obj KtAuthUrl, addr string) bool {
	pbytes, _ := json.Marshal(obj)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s%s", addr, lib.SvcmgrApiMicroKtAuthUrl)
	fmt.Println("Notify: ", url)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error: ", err)
		return false
	} else {
		defer response.Body.Close()
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2: ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}
