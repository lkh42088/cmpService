package rest

import (
	"bufio"
	"bytes"
	"cmpService/common/lib"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/teamwork/reload"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
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
		UpdateSvcmgrConf(data)
	case lib.MC_AGENT, lib.WIN_AGENT:
		SendToMcAgent(data, lib.McUrlSystemModifyConf)
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
		SendToMcAgent(data, lib.McUrlAgentRestart)
	}

	c.JSON(http.StatusOK, "OK")
}

func UpdateSvcmgrConf(data lib.ConfVariable) {
	SetEnvValue(data.FieldName, data.Value)
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

func SetEnvValue(field string, newVal string) bool {
	// ORIGIN FILE OPEN
	current, _ := os.Getwd()
	origin_file := current + "/etc/svcmgr.conf"
	fd, err := os.Open(origin_file)
	if err != nil {
		fmt.Println("SetEnvValue: error 1", err)
		return false
	}
	defer fd.Close()

	// BACKUP FILE CREATE
	backup_file := origin_file + ".backup"
	backup_fd, err := os.Create(backup_file)
	if err != nil {
		fmt.Println("SetEnvValue: error 2", err)
		return false
	}
	defer backup_fd.Close()

	// UPDAATE CONF FILE
	w := bufio.NewWriter(backup_fd)
	if err != nil {
		fmt.Println("SetEnvValue: error 3", err)
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