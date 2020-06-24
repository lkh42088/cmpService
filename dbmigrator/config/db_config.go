package config

import (
	"cmpService/common/config"
	"cmpService/common/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type global_config struct {
	NewDbConfig *config.DBConfig
	OldDbConfig *config.DBConfig
}

type MigratorConfig struct {
	NewDbIp       string `json:"new_db_ip"`
	NewDbName     string `json:"new_db_name"`
	NewDbUser     string `json:"new_db_user"`
	NewDbPassword string `json:"new_db_password"`
	OldDbIp       string `json:"old_db_ip"`
	OldDbName     string `json:"old_db_name"`
	OldDbUser     string `json:"old_db_user"`
	OldDbPassword string `json:"old_db_password"`
}

var DbMigratorGlobalConfig = &global_config{}

func SetOldDb(dbConfig *config.DBConfig) {
	DbMigratorGlobalConfig.OldDbConfig = dbConfig
}

func SetNewDb(dbConfig *config.DBConfig) {
	DbMigratorGlobalConfig.NewDbConfig = dbConfig
}

// Test Database of Contents Bridge
func GetTestCbDatabaseConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"root",
		"Nubes1510!",
		"test_cdn_db",
		"192.168.121.154",
		3306,
	}
	return &config
}

// Mariadb of Customizing Contents Bridge
func GetNewDatabaseConfig() *config.DBConfig {
	return DbMigratorGlobalConfig.NewDbConfig
}

// Mysql database of Contents Bridge
func GetOldDatabaseConfig() *config.DBConfig {
	return DbMigratorGlobalConfig.OldDbConfig
}

func SetConfig(configPath string) {
	if configPath == "" || lib.IsFileExists(configPath) == false {
		dirName, _ := os.Getwd()
		configPath = fmt.Sprintf("%s/etc/%s", dirName, "dbmigrator.conf")
		if configPath == "" || lib.IsFileExists(configPath) == false {
			log.Fatal("%% You MUST input config file path!")
			return
		}
	}

	cfg := ReadDbMigratorConfig(configPath)
	newDb := &config.DBConfig{
		DBDriver: "mysql",
		Username: cfg.NewDbUser,
		DBName:   cfg.NewDbName,
		Password: cfg.NewDbPassword,
		Address:  cfg.NewDbIp,
		Port:     3306,
	}
	oldDb := &config.DBConfig{
		DBDriver: "mysql",
		Username: cfg.OldDbUser,
		Password: cfg.OldDbPassword,
		DBName:   cfg.OldDbName,
		Address:  cfg.OldDbIp,
		Port:     3306,
	}

	SetNewDb(newDb)
	SetOldDb(oldDb)
}

func ReadDbMigratorConfig(path string) (config MigratorConfig) {
	// read file
	b, err := ioutil.ReadFile(path)
	if err != nil {
		lib.LogWarn("Fail to Read rest config file.(%s)\n", err)
		return config
	}
	// JSON transform
	err = json.Unmarshal(b, &config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}
