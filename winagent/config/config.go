package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type WinAgentConfig struct {
	WinAgentIp		   string			  `json:"winagent_ip"`
	WinAgentPort       string             `json:"winagent_port"`
	WinAgentPath	   string 			  `json:"winagent_path"`
	McAgentIp          string             `json:"mcagent_ip"`
	McAgentPort        string             `json:"mcagent_port"`
	SvcmgrIp           string             `json:"svcmgr_ip"`
	SvcmgrPort         string             `json:"svcmgr_port"`
	MonitoringInterval int                `json:"monitoring_interval"`
}

var GlobalConfig WinAgentConfig

func GetGlobalConfig() WinAgentConfig {
	return GlobalConfig
}

func ApplyGlobalConfig(file string) bool {
	fmt.Println("VM agent ApplyGlobalConfig: ", file)
	info, err := os.Stat(strings.TrimSpace(file))
	if os.IsNotExist(err) {
		fmt.Printf("ApplyGlobalConfig : doesn't exist config! (%s:%s)\n", err, file)
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

	err = json.Unmarshal(b, &GlobalConfig)
	if err != nil {
		fmt.Println("ApplyGlobalConfig : err 2 ", err)
		return false
	}

	return true
}

