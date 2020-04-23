package conf

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"nubes/collector/lib"
	"os"
)

type CollectorConfig map[string]string

type Config struct {
	Mongoip string
	Mongodb string
	Mongotable string
	Influxip string
	Influxdb string
	Svcmgrip string
	Restip string
	Restport string
}

const defaultConfigName = "collector.conf"
const Mongoip = "mongoip"
const Mongodb = "mongodb"
const Mongotable = "mongotable"
const Influxip = "influxip"
const Influxdb = "influxdb"
const Svcmgrip = "svcmgrip"
const Restip = "restip"
const Restport = "restport"

var ConfigPath string

// need to change default conf
func GetDefaultConfig() CollectorConfig {
	return CollectorConfig{
		Mongoip:    "127.0.0.1",
		Mongodb:    "collector",
		Mongotable: "devices",
		Influxip:   "192.168.10.19",
		Influxdb:   "snmp_nodes",
		Svcmgrip:   "127.0.0.1",
		Restip:     "127.0.0.1",
		Restport:   "8884",
	}
}

func ConvertCollectorConfig(c Config) CollectorConfig {
	return CollectorConfig{
		Mongoip: c.Mongoip,
		Mongodb: c.Mongodb,
		Mongotable: c.Mongotable,
		Influxip: c.Influxip,
		Influxdb: c.Influxdb,
		Svcmgrip: c.Svcmgrip,
		Restip: c.Restip,
		Restport: c.Restport,
	}
}

func CreateDefaultConfig() (config string) {

	// It create default conf at current directory
	dirName, _ := os.Getwd()
	filepath := fmt.Sprintf("%s/%s", dirName, defaultConfigName)

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	if err != nil {
		log.Fatal("Failed to create collector default conf file (%s)\n", err)
	}
	defer file.Close()

	// Get default conf
	defaultConfig := GetDefaultConfig()
	b, err := json.Marshal(defaultConfig)

	b, _ = PrettyPrint(b)

	// write file
	_, err = file.WriteString(string(b))

	return filepath
}

// not exist param string : default conf
// exist param string : change conf
func UpdateConfig(key string, config string) error {
	var f *os.File
	var err error

	collectorConfig := ReadConfig()
	if collectorConfig == nil {
		collectorConfig = GetDefaultConfig()
	}

	// No param
	if key == "" || config == "" {
		lib.LogWarn("Not found change conf param.\n")
		return nil
	}

	// file open
	if f, err = os.OpenFile(ConfigPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		os.FileMode(0644)); err != nil {
		lib.LogWarn("REST API Server can't create conf file.\n")
		return err
	}
	defer f.Close()

	// map value change
	if _, ok := collectorConfig[key]; ok {
		collectorConfig[key] = config
		fmt.Println(collectorConfig)
	} else {
		return errors.New("Invalid key name.\n")
	}

	// JSON transform
	b, err := json.Marshal(collectorConfig)

	b, _ = PrettyPrint(b)

	// write file
	_, err = f.WriteString(string(b))
	if err != nil {
		lib.LogWarn("Fail to write collector conf.(%s)\n", err)
	}

	return nil
}

func PrettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Read conf file and return conf object
func ReadConfig() (config CollectorConfig) {
	// read file
	b, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		lib.LogWarn("Fail to Read rest conf file.(%s)\n", err)
		return nil
	}
	// JSON transform
	//err = json.Unmarshal(b, &config)
	var cfg Config
	err = json.Unmarshal(b, &cfg)
	fmt.Println(cfg)
	if err != nil {
		fmt.Println(err)
	}
	config = ConvertCollectorConfig(cfg)
	return config
}

func ProcessConfig(config string) {
	if IsFileExists(config) == false {
		ConfigPath = CreateDefaultConfig()
	} else {
		ConfigPath = config
	}
	fmt.Println(ConfigPath)
}

