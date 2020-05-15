package config

import (
	"fmt"
	"cmpService/common/config"
	"cmpService/common/lib"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	path := "/home/andrew/projects/go/src/cmpService/collector/etc/collector.jbh.config"
	config := ReadConfigByPath(path)
	fmt.Println(config)
}

//need to change default config
func getJungbhConfig() *CollectorConfig {
	mongo := config.MongoDbConfig{
		"192.168.10.115",
		"collector",
		"devices",
	}
	influx := config.InfluxDbConfig{
		"192.168.10.74",
		"snmp_nodes",
		"cmpService",
		"nubes1510",
	}
	return &CollectorConfig{
		mongo,
		influx,
		"127.0.0.1",
		"0.0.0.0",
		"8884",
	}
}

func TestWriteJungbhConfig(t *testing.T) {
	dirName, _ := os.Getwd()
	path := fmt.Sprintf("%s/../etc/%s", dirName, "collector.jbh.conf")
	var cfg = getJungbhConfig()
	fmt.Println(cfg)
	config := lib.CreateConfig(path, cfg)
	fmt.Println(config)
}

func TestWriteDefaultConfig(t *testing.T) {
	dirName, _ := os.Getwd()
	path := fmt.Sprintf("%s/../etc/%s", dirName, collectorConfigName)
	fmt.Println(path)
	var cfg = GetDefaultConfig()
	fmt.Println(cfg)
	config := lib.CreateConfig(path, cfg)
	fmt.Println(config)
}

func TestUpdateConfig(t *testing.T) {
	dirName, _ := os.Getwd()
	path := fmt.Sprintf("%s/../etc/%s", dirName, collectorConfigName)
	UpdateConfig(path, "mongo_ip", "127.0.0.2")
}