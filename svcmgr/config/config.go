package config

import (
	"encoding/json"
	"fmt"
	client "github.com/influxdata/influxdb1-client"
	"io/ioutil"
	"log"
	"nubes/common/config"
	"nubes/common/lib"
	"nubes/common/mariadblayer"
	"os"
)

type global_config struct {
	MariadbConfig config.DBConfig
	Mariadb       mariadblayer.DBORM

	InfluxdbConfig config.DBConfig
	InfluxdbBp     client.BatchPoints
	InfluxdbClient client.Client
	RestServer     string
}

var SvcmgrGlobalConfig = &global_config{}

type SvcmgrConfig struct {
	config.MariaDbConfig
	config.InfluxDbConfig
	RestServerIp string `json:"rest_server_ip"`
	RestServerPort string `json:"rest_server_port"`
}

const svcmgrConfigName = "svcmgr.conf"
var SvcmgrConfigPath string

func GetDefaultConfig() *SvcmgrConfig {
	maria := config.MariaDbConfig{
		"127.0.0.1",
		"nubes",
		"nubes",
		"nubes1510",
	}
	influx := config.InfluxDbConfig{
		"192.168.10.19",
		"snmp_nodes",
		"nubes",
		"nubes1510",
	}
	return &SvcmgrConfig{
		maria,
		influx,
		"0.0.0.0",
		"8081",
	}
}

func CreateDefaultConfig(cfgPath string) (config string) {

	file, err := os.OpenFile(cfgPath, os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	if err != nil {
		log.Fatal("Failed to create collector default config file", err)
	}
	defer file.Close()

	// Get default config
	defaultConfig := GetDefaultConfig()
	b, err := json.Marshal(defaultConfig)

	b, _ = lib.PrettyPrint(b)

	// write file
	_, err = file.WriteString(string(b))

	return cfgPath
}

func ReadConfig(path string) (config SvcmgrConfig) {
	// read file
	b, err := ioutil.ReadFile(path)
	if err != nil {
		lib.LogWarn("Fail to Read rest config file.(%s)\n", err)
		return config
	}
	// JSON transform
	err = json.Unmarshal(b, &config)
	fmt.Println(config)
	if err != nil {
		fmt.Println(err)
	}

	return config
}

func SetConfig(config string) {
	if config == "" || lib.IsFileExists(config) == false {
		// It create default config at current directory
		dirName, _ := os.Getwd()
		filepath := fmt.Sprintf("%s/%s", dirName, svcmgrConfigName)
		SvcmgrConfigPath = CreateDefaultConfig(filepath)
	} else {
		SvcmgrConfigPath = config
	}
	lib.LogWarnln("Set SvcmgrConfig:", SvcmgrConfigPath)
}

func SetConfigInfluxdb(cfg config.InfluxDbConfig) {
	isChanged := false
	conf := ReadConfig(SvcmgrConfigPath)
	if conf.InfluxPassword != cfg.InfluxPassword {
		conf.InfluxPassword = cfg.InfluxPassword
		isChanged = true
	}
	if conf.InfluxUser != cfg.InfluxUser {
		conf.InfluxUser = cfg.InfluxUser
		isChanged = true
	}
	if conf.InfluxDb != cfg.InfluxDb {
		conf.InfluxDb = cfg.InfluxDb
		isChanged = true
	}
	if conf.InfluxIp != cfg.InfluxIp {
		conf.InfluxIp = cfg.InfluxIp
		isChanged = true
	}
	if isChanged {
		// Update Mongodb
		//ConfigureInfluxDB()
	}
}

func SetConfigMariadb(cfg config.MariaDbConfig) {
	isChanged := false
	conf := ReadConfig(SvcmgrConfigPath)
	if conf.MariaIp != cfg.MariaIp {
		conf.MariaIp = cfg.MariaIp
		isChanged = true
	}
	if conf.MariaDb != cfg.MariaDb {
		conf.MariaDb = cfg.MariaDb
		isChanged = true
	}
	if conf.MariaUser != cfg.MariaUser {
		conf.MariaUser = cfg.MariaUser
		isChanged = true
	}
	if conf.MariaPassword != cfg.MariaPassword {
		conf.MariaPassword = cfg.MariaPassword
		isChanged = true
	}
	if isChanged {
		// Update Mariadb
		//ConfigureMariaDB()
	}
}

