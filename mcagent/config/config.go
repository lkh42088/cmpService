package config

import (
	"bufio"
	"cmpService/common/config"
	"cmpService/common/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const MAX_VM_COUNT = 10

type McAgentConfig struct {
	config.MongoDbConfig
	config.InfluxDbConfig
	McagentIp          string             `json:"mcagent_ip"`
	McagentPort        string             `json:"mcagent_port"`
	SvcmgrIp           string             `json:"svcmgr_ip"`
	SvcmgrPort         string             `json:"svcmgr_port"`
	VmImageDir         string             `json:"vm_image_dir"`
	VmInstanceDir      string             `json:"vm_instance_dir"`
	ServerPort         string             `json:"server_port"`
	ServerMac          string             `json:"server_mac"`
	ServerIp           string             `json:"server_ip"`
	ServerPublicIp     string             `json:"server_public_ip"`
	ServerStatusRepo   string             `json:"server_status_repo"`
	MonitoringInterval int                `json:"monitoring_interval"`
	DnatBasePortNum    int                `json:"dnat_base_port_num"`
	SerialNumber       string             `json:"-"`
	VmNumber           [MAX_VM_COUNT]uint `json:"-"`
}

var globalConfig McAgentConfig

func GetGlobalConfig() McAgentConfig {
	return globalConfig
}

func SetSerialNumber2GlobalConfig(sn string) {
	globalConfig.SerialNumber = sn
	if sn != "" {
		fmt.Println("config Telegraf...")
		SetTelegraf(sn)
		RestartTelegraf()
	}
}

func SetGlobalConfigByVmNumber(index, value uint) {
	globalConfig.VmNumber[index] = value
}

func ApplyGlobalConfig(file string) bool {
	fmt.Println("ApplyGlobalConfig: ", file)
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		fmt.Println("ApplyGlobalConfig : dose not exist config!")
		return false
	}
	if info.IsDir() {
		fmt.Println("ApplyGlobalConfig : the config is directory!")
		return false
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("ApplyGlobalConfig : err ", err)
		return false
	}

	err = json.Unmarshal(b, &globalConfig)
	if err != nil {
		fmt.Println("ApplyGlobalConfig : err 2 ", err)
		return false
	}

	// Default Number
	if globalConfig.DnatBasePortNum == 0 {
		globalConfig.DnatBasePortNum = 17000
	}

	globalConfig.ServerPublicIp = utils.GetMyPublicIp()
	return true
}

func SetTelegraf(sn string) bool {
	orgin_file := "/etc/telegraf/telegraf.conf"
	fd, err := os.Open(orgin_file)
	if err != nil {
		fmt.Println("SetTelegraf: error", err)
		return false
	}
	defer fd.Close()

	backup_file := orgin_file +".backup"
	backup_fd, err := os.Create(backup_file)
	if err != nil {
		fmt.Println("SetTelegraf: error", err)
		return false
	}
	defer backup_fd.Close()

	w := bufio.NewWriter(backup_fd)
	if err != nil {
		fmt.Println("SetTelegraf: error", err)
		return false
	}

	isFind := false
	global_tags_area := false
	findStr := "serial_number"
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
		} else if strings.Contains(lineStr, findStr) == true {
			w.WriteString(fmt.Sprintf("  %s = \"%s\"\n", findStr, sn))
			isFind = true
		} else if strings.Contains(lineStr, "[global_tags]") == true {
			w.WriteString(lineStr+"\n")
			global_tags_area = true
		} else if isFind == false &&
			global_tags_area == true &&
			len(strings.Trim(lineStr, " ")) == 0 {
			w.WriteString(fmt.Sprintf("  %s = \"%s\"\n", findStr, sn))
			w.WriteString("\n")
			isFind = true
		} else {
			w.WriteString(lineStr+"\n")
		}
	}
	w.Flush()
	backup_fd.Sync()

	args := []string{
		backup_file,
		orgin_file,
	}

	binary := "mv"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output:", string(output))

	return true
}

func RestartTelegraf () {
	args := []string{
		"restart",
		"telegraf",
	}

	binary := "systemctl"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output:", string(output))
}