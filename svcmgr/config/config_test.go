package config

import (
	"fmt"
	"nubes/common/config"
	"nubes/common/lib"
	"os"
	"testing"
)

func getJungbhConfig() *SvcmgrConfig {
	maria := config.MariaDbConfig{
		"192.168.10.115",
		"nubes",
		"nubes",
		"nubes1510",
	}
	influx := config.InfluxDbConfig{
		"192.168.10.74",
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

func TestWriteJungbhConfig(t *testing.T) {
	dirName, _ := os.Getwd()
	path := fmt.Sprintf("%s/../etc/%s", dirName, "svcmgr.jbh.conf")
	var cfg = getJungbhConfig()
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