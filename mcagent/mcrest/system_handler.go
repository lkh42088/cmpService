package mcrest

import (
	"bytes"
	"cmpService/common/lib"
	"cmpService/mcagent/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/teamwork/reload"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

func ModifyConfVariable(c *gin.Context) {
	var data lib.ConfVariable
	c.ShouldBind(&data)

	if data.IpAddr == "" || data.FieldName == "" {
		c.JSON(http.StatusBadRequest, "Parameter isn't valid.\n")
		return
	}

	switch data.AgentType {
	case lib.MC_AGENT:
		UpdateMcAgentConf(data.FieldName, data.Value)
	case lib.WIN_AGENT:
		SendToWinAgent(data, lib.McUrlWinSystemModifyConf)
	}

	c.JSON(http.StatusOK, "OK")
}

func RestartMcAgent(c *gin.Context) {
	var data lib.ConfVariable
	c.ShouldBind(&data)

	if data.IpAddr == "" {
		c.JSON(http.StatusBadRequest, "Parameter isn't valid.\n")
		return
	}

	switch data.AgentType {
	case lib.MC_AGENT:
		RestartMcServer()
	case lib.WIN_AGENT:
		SendToWinAgent(data, lib.McUrlWinAgentRestart)
	}

	c.JSON(http.StatusOK, "OK")
}

func UpdateMcAgentConf(field string, newVal string) bool {
	var f *os.File
	var err error

	// Get Conf
	conf := config.GetMcGlobalConfig()
	current, _ := os.Getwd()
	path := current + "/mcagent.conf"

	// Change Value
	val := reflect.ValueOf(conf)
	for i := 0; i < val.Type().NumField(); i++ {
		confField := val.Type().Field(i).Tag.Get("json")
		if confField == field {
			reflect.ValueOf(conf).Elem().Field(i).SetString(newVal)
		}
	}

	// File Open
	if f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		os.FileMode(0644)); err != nil {
		lib.LogWarn("REST API Server can't create config file.\n")
		return false
	}
	defer f.Close()

	// JSON transform
	var b []byte
	if b, err = json.Marshal(conf); err != nil {
		lib.LogWarn("Failed Marshal!\n")
		return false
	}

	b, _ = lib.PrettyPrint(b)

	// write file
	_, err = f.WriteString(string(b))
	if err != nil {
		lib.LogWarn("Fail to write collector config.(%s)\n", err)
	}

	return true
}

func SendToWinAgent(data lib.ConfVariable, uri string) bool {
	pbytes, _ := json.Marshal(data)
	buff := bytes.NewBuffer(pbytes)

	var serverIp string
	if data.AgentType == lib.WIN_AGENT {
		ipSet := strings.Split(data.IpAddr, "|")
		serverIp = ipSet[1]
	} else {
		return false
	}

	url := fmt.Sprintf("http://%s:8083%s", serverIp, uri)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendToWinAgent: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendToWinAgent: error 2 ", err)
		return false
	}
	fmt.Println("SendToWinAgent: success - ", string(resp))
	return true
}

func RestartMcServer() {
	fmt.Println("\n\nAgent Restart...\n\n")
	reload.Exec()
}
