package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"nubes/collector/lib"
	"os"
	"os/user"
)

type RestConfig map[string]string

// config file path
const collecrot_filepath = "/nubes/collector/etc"
const collector_filename = "collector.conf"
const config_path = collecrot_filepath + "/" + collector_filename

// need to change default config
func NewRestConfig() RestConfig {
	return RestConfig{
		"mongoip"	:	"127.0.0.1",
		"mongodb"	:	"collector",
		"mongotable":	"devices",
		"influxip" 	:	"192.168.10.19",
		"influxdb" 	:	"snmp_nodes",
		"svcmgrip"	:	"127.0.0.1",
		"restip"	:	"127.0.0.1",
		"restport"	:	"8884",
	}
}

// Get '/HOME/USER' directory path
func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		lib.LogWarn("Fail to get user home directory path.\n")
		return ""
	}
	return usr.HomeDir
}

// Need to check file.Close()
func CreateConf() (*os.File, error) {
	var file *os.File
	filepath := GetHomeDir() + config_path

	// file exist check
	if _, err := os.Stat(filepath); err != nil {
		// file directory check
		if os.MkdirAll(GetHomeDir()+collecrot_filepath,
			os.FileMode(0777)) != nil {
			lib.LogWarn(
				"Fail to collector config directory create(%s)\n", err)
			return nil, nil
		}
		// file create
		file, err = os.OpenFile(filepath,
			os.O_CREATE|os.O_RDWR, os.FileMode(0777))
		if err != nil {
			lib.LogWarn(
				"Fail to collector config file create(%s)\n", err)
			return nil, nil
		}

		// Init config write
		restconfig := NewRestConfig()
		b, err := json.Marshal(restconfig)
		// write file
		_, err = file.WriteString(string(b))
		defer file.Close()

		return file, err
	}
	return file, nil
}

// Read config file and return config object
func ReadConf() (r RestConfig) {
	filepath := GetHomeDir() + config_path
	// read file
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		lib.LogWarn("Fail to Read rest config file.(%s)\n", err)
		return nil
	}
	// JSON transform
	err = json.Unmarshal(b, &r)
	if err != nil {
		fmt.Println(err)
	}
	// if content is empty, create new object
	//if restconfig == nil {
	//	restconfig = NewRestConfig()
	//}
	return r
}

// not exist param string : default config
// exist param string : change config
func WriteConf(key string, config string) error {
	var f *os.File
	var err error
	filepath := GetHomeDir() + config_path

	restconfig := ReadConf()
	if restconfig == nil {
		restconfig = NewRestConfig()
	}
	// No param
	if key == "" || config == "" {
		lib.LogWarn("Not found change config param.\n")
	}
	// file open
	if f, err = os.OpenFile(filepath,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		os.FileMode(0777)); err != nil {
		lib.LogWarn("REST API Server can't create config file.\n")
		return err
	}
	// map value change
	if _, ok := restconfig[key]; ok {
		restconfig[key] = config
		fmt.Println(restconfig)
	} else {
		return errors.New("Invalid key name.\n")
	}

	// JSON transform
	b, err := json.Marshal(restconfig)
	// write file
	_, err = f.WriteString(string(b))
	defer f.Close()
	if err != nil {
		lib.LogWarn("Fail to write collector config.(%s)\n", err)
	}
	return nil
}

func FindConfig() {
	filepath := GetHomeDir() + config_path
	// file exist check
	if _, err := os.Stat(filepath); err != nil {
		// if not file, create file
		if _, err := CreateConf(); err != nil {
			lib.LogWarn("REST Server config file is nothing.(%s)\n", err)
			return
		}
		lib.LogWarn("REST API Server config file is OK!\n")
	}
	return
}
