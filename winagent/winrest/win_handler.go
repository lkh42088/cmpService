package winrest

import (
	"bufio"
	"cmpService/common/lib"
	"cmpService/winagent/common"
	"cmpService/winagent/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
	"net/http"
	"os"
	"reflect"
	"strings"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "Health chec ok!")
}

func ModifyConfVariable(c *gin.Context) {
	var data lib.ConfVariable
	c.ShouldBind(&data)

	if data.AgentType != lib.WIN_AGENT || data.FieldName == "" {
		c.JSON(http.StatusBadRequest, "Parameter isn't valid.\n")
		return
	}

	//if UpdateWinAgentConf(data.FieldName, data.Value) {
	if SetEnvValue(data.FieldName, data.Value) {
		if data.FieldName == "influxdb_ip" ||
			data.FieldName == "influxdb_port" {
			UpdateTelegrafConf(data.FieldName, data.Value)
			common.RestartTelegraf()
		}
	}

	c.JSON(http.StatusOK, "OK")
}

func RestartAgent(c *gin.Context) {
	fmt.Println("\n\nAgent Restart...\n\n")
	//reload.Exec()

	scMgr, err := mgr.Connect()
	defer scMgr.Disconnect()
	sc, err := scMgr.OpenService("CMPWindowService")
	defer sc.Close()

	_, err = sc.Control(svc.Shutdown)
	sc.Start()
	if err != nil {
		fmt.Println("[Service Restart Error] ", err)
	}

	c.JSON(http.StatusOK, "OK")
}

func UpdateWinAgentConf(field string, newVal string) bool {
	var f *os.File
	var err error

	// Get Conf
	conf := config.GetGlobalConfig()
	path := conf.WinAgentPath + "\\winagent.conf"

	// Change Value
	val := reflect.ValueOf(&conf).Elem()
	for i := 0; i < val.Type().NumField(); i++ {
		confField := val.Type().Field(i).Tag.Get("json")
		if strings.Contains(confField, field) {
			val.Field(i).SetString(newVal)
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

// Only can change to influxdb ip & port
func UpdateTelegrafConf(field string, value string) bool {
	influxdbTags := "[[outputs.influxdb]]"
	orgin_file := "c:\\Program files\\Telegraf\\telegraf.conf"
	fd, err := os.Open(orgin_file)
	if err != nil {
		fmt.Println("UpdateTelegrafConf: error", err)
		return false
	}
	defer fd.Close()

	backup_file := orgin_file +".cronsch"
	backup_fd, err := os.Create(backup_file)
	if err != nil {
		fmt.Println("UpdateTelegrafConf: error", err)
		return false
	}
	defer backup_fd.Close()

	w := bufio.NewWriter(backup_fd)
	if err != nil {
		fmt.Println("UpdateTelegrafConf: error", err)
		return false
	}

	isFind := false
	influxdbTagsArea := false
	findStr := "urls = [\"http"
	reader := bufio.NewReader(fd)
	for {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix  || err != nil {
			fmt.Println(isPrefix, "error", err)
			break
		}
		lineStr := string(line)
		if isFind == true {
			w.WriteString(lineStr+"\n")
		} else if strings.Contains(lineStr, influxdbTags) == true {
			w.WriteString(lineStr+"\n")
			influxdbTagsArea = true
		} else if isFind == false && influxdbTagsArea == true &&
			strings.Contains(lineStr, findStr) == true {
			if field == "influxdb_ip" {
				w.WriteString(fmt.Sprintf("  urls = [\"http://%s:%s\"]\n", value, config.GlobalConfig.InfluxDbPort))
			} else if field == "influxdb_port" {
				w.WriteString(fmt.Sprintf("  urls = [\"http://%s:%s\"]\n", config.GlobalConfig.InfluxDbIp, value))
			}
			w.WriteString("\n")
			isFind = true
		} else {
			w.WriteString(lineStr+"\n")
		}
	}
	w.Flush()
	backup_fd.Sync()

	err = common.CopyFile(backup_file, orgin_file)
	if err != nil {
		return false
	}

	return true
}

func SetEnvValue(field string, newVal string) bool {
	// ORIGIN FILE OPEN
	current := config.GetGlobalConfig()
	origin_file := current.WinAgentPath + "\\winagent.conf"
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

	// UPDAATE CONF FILE
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
	common.CopyFile(backup_file, origin_file)

	return true
}

