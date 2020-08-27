package config

import (
	"cmpService/common/config"
	"cmpService/common/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
