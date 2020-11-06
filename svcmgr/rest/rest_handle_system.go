package rest

import (
	"bytes"
	"cmpService/common/lib"
	"cmpService/svcmgr/config"
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

	if data.AgentType == 0 || data.IpAddr == "" || data.FieldName == "" {
		c.JSON(http.StatusBadRequest, "Parameter isn't valid.\n")
		return
	}

	switch data.AgentType {
	case lib.SVCMGR_AGENT:
		UpdateSvcmgrConf(data.FieldName, data.Value)
	case lib.MC_AGENT, lib.WIN_AGENT:
		SendToMcAgent(data, lib.McUrlPrefix + lib.McUrlSystemModifyConf)
	}

	c.JSON(http.StatusOK, "OK")
}

func RestartAgent(c *gin.Context) {
	var data lib.ConfVariable
	c.ShouldBind(&data)

	if data.AgentType == 0 || data.IpAddr == "" {
		c.JSON(http.StatusBadRequest, "Parameter isn't valid.\n")
		return
	}

	switch data.AgentType {
	case lib.SVCMGR_AGENT:
		RestartSvcmgr()
	case lib.MC_AGENT, lib.WIN_AGENT:
		SendToMcAgent(data, lib.McUrlPrefix + lib.McUrlAgentRestart)
	}

	c.JSON(http.StatusOK, "OK")
}

func UpdateSvcmgrConf(field string, newVal string) bool {
	var f *os.File
	var err error

	// Get Conf
	conf := *config.SvcmgrConfigStore
	current, _ := os.Getwd()
	path := current + "/svcmgr.conf"

	// Change Value
	val := reflect.ValueOf(&conf).Elem()
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Field(i).Type().String() == "config.MariaDbConfig" {
			for j := 0; j < val.Field(i).Type().NumField(); j++ {
				confField := val.Field(i).Type().Field(j).Tag.Get("json")
				if strings.Contains(confField, field) {
					val.Field(i).Field(j).SetString(newVal)
				}
			}
		} else if val.Field(i).Type().String() == "config.InfluxDbConfig" {
			for j := 0; j < val.Field(i).Type().NumField(); j++ {
				confField := val.Field(i).Type().Field(j).Tag.Get("json")
				if strings.Contains(confField, field) {
					val.Field(i).Field(j).SetString(newVal)
				}
			}
		} else {
			confField := val.Type().Field(i).Tag.Get("json")
			if strings.Contains(confField, field) {
				val.Field(i).SetString(newVal)
			}
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

func SendToMcAgent(data lib.ConfVariable, uri string) bool {
	pbytes, _ := json.Marshal(data)
	buff := bytes.NewBuffer(pbytes)

	var serverIp string
	if data.AgentType == lib.WIN_AGENT {
		ipSet := strings.Split(data.IpAddr, "|")
		serverIp = ipSet[0]
	} else {
		serverIp = data.IpAddr
	}

	url := fmt.Sprintf("http://%s:8082%s", serverIp, uri)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendToMcAgent: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendToMcAgent: error 2 ", err)
		return false
	}
	fmt.Println("SendToMcAgent: success - ", string(resp))
	return true
}

func RestartSvcmgr() {
	fmt.Println("\n\nAgent Restart...\n\n")
	reload.Exec()
}
