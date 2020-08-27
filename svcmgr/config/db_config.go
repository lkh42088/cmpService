package config

import (
	"cmpService/collector/influx"
	"cmpService/common/config"
	db2 "cmpService/common/db"
	"cmpService/common/lib"
	"cmpService/common/mariadblayer"
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
)

func SetMariaDB(user, passwd, dbname, ip string, port int) (db *mariadblayer.DBORM, err error) {
	dbconfig, err := config.NewDBConfig("mysql",
		user, passwd, dbname, ip, port)
	if err != nil {
		fmt.Println("[SetMariaDB] Error:", err)
		return
	}
	SvcmgrGlobalConfig.MariadbConfig = *dbconfig
	dataSource := db2.GetDataSourceName(dbconfig)
	db, err = mariadblayer.NewDBORM(dbconfig.DBDriver, dataSource)
	if err != nil {
		fmt.Println("[SetMariaDB] Error:", err)
		return
	}
	SvcmgrGlobalConfig.Mariadb = *db
	return db, err
}

func SetInfluxDB() {
	var influxConfig *influx.InfluxAccessor
	var dbConfig config.DBConfig
	config := ReadConfig(SvcmgrConfigPath)
	if config.InfluxIp == "" || config.InfluxDb == "" {
		lib.LogWarn("Failed InfluxDb configuration!\n")
		return
	}
	path := "http://" + config.InfluxIp + ":8086"
	influxConfig = NewInfluxCfg(
		path,
		config.InfluxUser,
		config.InfluxPassword,
		config.InfluxDb)

	dbConfig.DBDriver = ""
	dbConfig.Address = influxConfig.Url
	dbConfig.Password = influxConfig.Password
	dbConfig.DBName = influxConfig.Database
	dbConfig.Port = 8806

	SvcmgrGlobalConfig.InfluxdbConfig = dbConfig
	SvcmgrGlobalConfig.InfluxdbBp = influxConfig.Bp
	SvcmgrGlobalConfig.InfluxdbClient = influxConfig.Client
}

func NewInfluxCfg(url string, user string, passwd string, db string) *influx.InfluxAccessor {
	config := influx.InfluxAccessor{
		Url:      url,
		Username: user,
		Password: passwd,
		Database: db,
	}
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "ms",
		Database:  config.Database,
	})
	if err != nil {
		lib.LogWarn("InfluxDB NewInfluxCfg: Failed to get BatchPoints!!\n")
	} else {
		config.Bp = bp
	}

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.Url,
		Username: config.Username,
		Password: config.Password,
	})
	config.Client = c
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return &config
}