package conf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"nubes/collector/lib"
	"os"
	"reflect"
)

type CollectorConfig map[string]string

type Config struct {
	Mongoip string		`json:"mongoip"`
	Mongodb string		`json:"mongodb"`
	Mongotable string	`json:"mongotable"`
	Influxip string		`json:"influxip"`
	Influxdb string		`json:"influxdb"`
	Svcmgrip string		`json:"svcmgrip"`
	Restip string		`json:"restip"`
	Restport string		`json:"restport"`
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

//need to change default conf
func GetDefaultConfig() Config {
	var c Config
	c.Mongoip =    "127.0.0.1"
	c.Mongodb =    "collector"
	c.Mongotable = "devices"
	c.Influxip =   "192.168.10.19"
	c.Influxdb =   "snmp_nodes"
	c.Svcmgrip =   "127.0.0.1"
	c.Restip =     "127.0.0.1"
	c.Restport =   "8884"
	return c
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
func UpdateConfig(key string, value string) error {
	var f *os.File
	var err error

	config := ReadConfig()
	//if collectorConfig == nil {
	//	collectorConfig = GetDefaultConfig()
	//}

	// No param
	if key == "" || value == "" {
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

	// Value Change
	if SetConfigByField(key, value, &config) < 0 {
		lib.LogWarn("Invalid key name.\n")
		return nil
	}

	// JSON transform
	b, err := json.Marshal(config)

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
func ReadConfig() (config Config) {
	// read file
	b, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		lib.LogWarn("Fail to Read rest conf file.(%s)\n", err)
		//return nil
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

func ProcessConfig(config string) {
	if IsFileExists(config) == false {
		ConfigPath = CreateDefaultConfig()
	} else {
		ConfigPath = config
	}
	fmt.Println(ConfigPath)
}

func SetConfigByField(key string, config string, c *Config) int {
	if key == "" {
		lib.LogWarn("Key or config string is empty.\n")
		return -1
	}
	// find json field and set value
	elements := reflect.ValueOf(c).Elem()
	target := elements.Type()
	// Find json field
	for i := 0; i < target.NumField(); i++ {
		tag :=  target.Field(i).Tag
		if tag.Get("json") == key {
			// Set Value
			elements.Field(i).SetString(config)
			return i
		}
	}
	return -1
}

