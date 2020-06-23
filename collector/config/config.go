package config

import (
	"cmpService/common/config"
	"cmpService/common/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

type CollectorConfig struct {
	config.MongoDbConfig
	config.InfluxDbConfig
	SvcmgrIp       string `json:"svcmgr_ip"`
	RestServerIp   string `json:"rest_server_ip"`
	RestServerPort string `json:"rest_server_port"`
}

const collectorConfigName = "collector.conf"

var CollectorConfigPath string

//need to change default config
func GetDefaultConfig() *CollectorConfig {
	mongo := config.MongoDbConfig{
		"127.0.0.1",
		"collector",
		"devices",
	}
	influx := config.InfluxDbConfig{
		"192.168.10.19",
		"snmp_nodes",
		"cmpService",
		"nubes1510",
	}
	return &CollectorConfig{
		mongo,
		influx,
		"127.0.0.1",
		"127.0.0.1",
		"8884",
	}
}

func CreateDefaultConfig(cfgPath string) (config string) {

	file, err := os.OpenFile(cfgPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0777))
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

func UpdateConfig(path string, key string, value string) error {
	var f *os.File
	var err error

	conf := ReadConfig(path)

	// No param
	if key == "" || value == "" {
		lib.LogWarn("Not found change config param.\n")
		return nil
	}

	// file open
	if f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		os.FileMode(0644)); err != nil {
		lib.LogWarn("REST API Server can't create config file.\n")
		return err
	}
	defer f.Close()

	// Value Change
	if SetConfigByField(key, value, &conf) < 0 {
		lib.LogWarn("Invalid key name.\n")
		return nil
	}

	// JSON transform
	var b []byte
	if b, err = json.Marshal(conf); err != nil {
		lib.LogWarn("Failed Marshal!\n")
		return err
	}

	b, _ = lib.PrettyPrint(b)

	// write file
	_, err = f.WriteString(string(b))
	if err != nil {
		lib.LogWarn("Fail to write collector config.(%s)\n", err)
	}

	return nil
}

// Read config file and return config object
func ReadConfigByPath(path string) (config CollectorConfig) {
	// read file
	b, err := ioutil.ReadFile(path)
	if err != nil {
		lib.LogWarnln("Fail to Read rest config file:", err)
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

func ReadConfig(path string) (config CollectorConfig) {
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
		filepath := fmt.Sprintf("%s/%s", dirName, collectorConfigName)
		CollectorConfigPath = CreateDefaultConfig(filepath)
	} else {
		CollectorConfigPath = config
	}
	lib.LogWarnln("Set CollectorConfig:", CollectorConfigPath)
}

func SetConfigByField(key string, config string, c *CollectorConfig) int {
	if key == "" {
		lib.LogWarn("Key or config string is empty.\n")
		return -1
	}
	// find json field and set value
	elements := reflect.ValueOf(c).Elem()
	target := elements.Type()
	// Find json field
	for i := 0; i < target.NumField(); i++ {
		tag := target.Field(i).Tag
		if tag.Get("json") == key {
			// Set Value
			elements.Field(i).SetString(config)
			return i
		}
	}
	return -1
}

func SetConfigInfluxdb(cfg config.InfluxDbConfig) {
	isChanged := false
	conf := ReadConfig(CollectorConfigPath)
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
		ConfigureInfluxDB()
	}
}

func SetConfigMongodb(cfg config.MongoDbConfig) {
	isChanged := false
	conf := ReadConfig(CollectorConfigPath)
	if conf.MongoIp != cfg.MongoIp {
		conf.MongoIp = cfg.MongoIp
		isChanged = true
	}
	if conf.MongoDb != cfg.MongoDb {
		conf.MongoDb = cfg.MongoDb
		isChanged = true
	}
	if conf.MongoCollection != cfg.MongoCollection {
		conf.MongoCollection = cfg.MongoCollection
		isChanged = true
	}
	if isChanged {
		// Update Mongodb
		ConfigureMongoDB()
	}
}
