package mcrest

import (
	"bufio"
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
	"os/exec"
	"reflect"
	"strconv"
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
	return SetEnvValue(field, newVal)
}

func UpdateMcAgentConf2(field string, newVal string) bool {
	// 구조체를 이용하여 conf 파일 적용하는 방식은 구조체에 따라 변동이 많으므로
	// 기존 conf 파일을 읽어 update 하는 방식으로 변경
	var f *os.File
	var err error

	// Get Conf
	conf := config.GetMcGlobalConfig()
	current, _ := os.Getwd()
	path := current + "/mcagent.conf"
	fmt.Println(path)

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
		} else if val.Field(i).Type().String() == "config.MongoDbConfig" {
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
				fmt.Println(val.Field(i).Type().String())
				if val.Field(i).Type().String() == "int" {
					v, _ := strconv.Atoi(newVal)
					val.Field(i).SetInt(int64(v))
				} else {
					val.Field(i).SetString(newVal)
				}
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
	fmt.Println("WIN URL: ", url)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendToWinAgent send error :", err)
		return false
	}
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendToWinAgent response error : ", err)
		return false
	}
	fmt.Println("SendToWinAgent: success - ", string(resp), response.StatusCode)
	return true
}

func RestartMcServer() {
	fmt.Println("\n\nAgent Restart...\n\n")
	reload.Exec()
}

func SetEnvValue(field string, newVal string) bool {
	// ORIGIN FILE OPEN
	current, _ := os.Getwd()
	origin_file := current + "/mcagent.conf"
	fd, err := os.Open(origin_file)
	if err != nil {
		fmt.Println("SetEnvValue file open error :", err)
		return false
	}
	defer fd.Close()

	// BACKUP FILE CREATE
	backup_file := origin_file + ".backup"
	backup_fd, err := os.Create(backup_file)
	if err != nil {
		fmt.Println("SetEnvValue file create error :", err)
		return false
	}
	defer backup_fd.Close()

	// UPDATE CONF FILE
	w := bufio.NewWriter(backup_fd)
	if err != nil {
		fmt.Println("SetEnvValue Writer error :", err)
		return false
	}

	isFind := false
	findStr := field
	reader := bufio.NewReader(fd)
	for {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix || err != nil {
			fmt.Println(isPrefix, "error", err)
			break
		}
		lineStr := string(line)
		if newVal != "" {
			if isFind == true {
				w.WriteString(lineStr + "\n")
			} else if strings.Contains(lineStr, findStr) == true {
				w.WriteString(fmt.Sprintf("  \"%s\": \"%s\",\n", findStr, newVal))
				isFind = true
			} else {
				w.WriteString(lineStr + "\n")
			}
		}
	}
	w.Flush()
	backup_fd.Sync()

	// FILE CHANGE
	args := []string{
		backup_file,
		origin_file,
	}

	binary := "mv"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("SetEnvValue output:", string(output))

	return true
}