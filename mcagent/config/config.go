package config

import (
	"cmpService/common/config"
	"cmpService/common/lib"
	"encoding/json"
	"io/ioutil"
	"os"
)

type McAgentConfig struct {
	config.MongoDbConfig
	McagentIp   string `json:"mcagent_ip"`
	McagentPort string `json:"mcagent_port"`
	SvcmgrIp    string `json:"svcmgr_ip"`
	SvcmgrPort  string `json:"svcmgr_port"`
}

var globalConfig McAgentConfig

func GetGlobalConfig () McAgentConfig {
	return globalConfig
}

func ApplyGlobalConfig(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	if info.IsDir() {
		return false
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		lib.LogWarnln(err)
		return false
	}
	err = json.Unmarshal(b, &globalConfig)
	if err != nil {
		lib.LogWarnln(err)
		return false
	}
	return true
}