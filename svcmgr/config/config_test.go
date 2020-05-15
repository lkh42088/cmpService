package config

import (
	"fmt"
	"cmpService/common/config"
	"cmpService/common/lib"
	"os"
	"testing"
)

func getJungbhConfig() *SvcmgrConfig {
	maria := config.MariaDbConfig{
		"192.168.10.115",
		"cmpService",
		"cmpService",
		"nubes1510",
	}
	influx := config.InfluxDbConfig{
		"192.168.10.74",
		"snmp_nodes",
		"cmpService",
		"nubes1510",
	}
	return &SvcmgrConfig{
		maria,
		influx,
		"0.0.0.0",
		"8081",
	}
}

func TestWriteJungbhConfig(t *testing.T) {
	dirName, _ := os.Getwd()
	path := fmt.Sprintf("%s/../etc/%s", dirName, "svcmgr.jbh.conf")
	var cfg = getJungbhConfig()
	fmt.Println(cfg)
	config := lib.CreateConfig(path, cfg)
	fmt.Println(config)
}


func getJeeebConfig() *SvcmgrConfig {
	maria := config.MariaDbConfig{
		"192.168.227.129",
		"cmpService",
		"cmpService",
		"nubes1510!",
	}
	// influx 는 사용안함...
	influx := config.InfluxDbConfig{
		"192.168.10.74",
		"snmp_nodes",
		"cmpService",
		"nubes1510",
	}
	return &SvcmgrConfig{
		maria,
		influx,
		"0.0.0.0",
		"8081",
	}
}

func TestWriteJeeebConfig(t *testing.T) {
	dirName, _ := os.Getwd()
	path := fmt.Sprintf("%s/../etc/%s", dirName, "svcmgr.jeb.conf")
	var cfg = getJeeebConfig()
	fmt.Println(cfg)
	config := lib.CreateConfig(path, cfg)
	fmt.Println(config)
}

func TestWriteDefaultConfig(t *testing.T) {
	dirName, _ := os.Getwd()
	path := fmt.Sprintf("%s/../etc/%s", dirName, svcmgrConfigName)
	fmt.Println(path)
	var cfg = GetDefaultConfig()
	fmt.Println(cfg)
	config := lib.CreateConfig(path, cfg)
	fmt.Println(config)
}