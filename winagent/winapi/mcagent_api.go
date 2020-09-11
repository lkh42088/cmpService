package winapi

import (
	"bytes"
	"cmpService/winagent/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendMsgToMcAgent(data interface{}, uri string) bool {
	conf := config.GetGlobalConfig()

	pbytes, _ := json.Marshal(data)
	buff := bytes.NewBuffer(pbytes)

	url := fmt.Sprintf("http://%s%s%s", conf.McAgentIp, ":" + conf.McAgentPort, uri)
	//fmt.Println("URL: ", url)

	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error: ", err)
		return false
	}
	defer response.Body.Close()

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2: ", err)
		return false
	}
	fmt.Println("response: ", string(res))
	return true
}
