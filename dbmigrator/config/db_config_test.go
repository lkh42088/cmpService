package config

import (
	"cmpService/common/lib"
	"fmt"
	"os"
	"testing"
)

func getDbMigConfig() *MigratorConfig{
	return &MigratorConfig{
		NewDbIp:       "127.0.0.1",
		NewDbName:     "nubes",
		NewDbUser:     "nubes",
		NewDbPassword: "Nubes1510!",
		OldDbIp:       "127.0.0.1",
		OldDbName:     "cdn_db_2020",
		OldDbUser:     "nubes",
		OldDbPassword: "Nubes1510!",
	}
}

func TestWriteConfig(t *testing.T) {
	dirName, _ := os.Getwd()
	path := fmt.Sprintf("%s/../etc/%s", dirName, "dbmigrator.conf")
	var cfg = getDbMigConfig()
	fmt.Println(cfg)
	conf := lib.CreateConfig(path, cfg)
	fmt.Println(conf)
}

func getJbhDbMigConfig() *MigratorConfig{
	return &MigratorConfig{
		NewDbIp:       "192.168.32.130",
		NewDbName:     "nubes",
		NewDbUser:     "nubes",
		NewDbPassword: "Nubes1510!",
		OldDbIp:       "192.168.32.149",
		OldDbName:     "cdn_db_2020",
		OldDbUser:     "nubes",
		OldDbPassword: "Nubes1510!",
	}
}

func TestJbhWriteConfig(t *testing.T) {
	dirName, _ := os.Getwd()
	path := fmt.Sprintf("%s/../etc/%s", dirName, "dbmigrator.jbh.conf")
	var cfg = getJbhDbMigConfig()
	fmt.Println(cfg)
	conf := lib.CreateConfig(path, cfg)
	fmt.Println(conf)
}
