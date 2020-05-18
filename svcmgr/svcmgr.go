package main

import (
	"flag"
	"fmt"
	"cmpService/common/config"
	db2 "cmpService/common/db"
	"cmpService/common/mariadblayer"
	config2 "cmpService/svcmgr/config"
	"cmpService/svcmgr/rest"
)

func main () {
	configFile := flag.String("file", "/etc/svcmgr/svcmgr.conf",
		"Input configuration file")
	flag.Parse()
	config2.SetConfig(*configFile)
	db, err := SetMariaDB()
	if err != nil {
		fmt.Println("Main: ERROR - ", err)
		return
	}
	SetRestServer(db)
}

func SetMariaDB() (db *mariadblayer.DBORM, err error) {
	cfg := config2.ReadConfig(config2.SvcmgrConfigPath)
	dbconfig, err := config.NewDBConfig("mysql",
		cfg.MariaUser, cfg.MariaPassword, cfg.MariaDb,
		cfg.MariaIp, 3306)
	if err != nil {
		fmt.Println("[SetMariaDB] Error:", err)
		return
	}
	config2.SvcmgrGlobalConfig.MariadbConfig = *dbconfig
	dataSource := db2.GetDataSourceName(dbconfig)
	db, err = mariadblayer.NewDBORM(dbconfig.DBDriver, dataSource)
	if err != nil {
		fmt.Println("[SetMariaDB] Error:", err)
		return
	}
	config2.SvcmgrGlobalConfig.Mariadb = *db
	return db, err
}

func SetRestServer(db *mariadblayer.DBORM) {
	cfg := config2.ReadConfig(config2.SvcmgrConfigPath)
	restServer := fmt.Sprintf("%s:%s", cfg.RestServerIp, cfg.RestServerPort)
	config2.SvcmgrGlobalConfig.RestServer = restServer
	rest.RunAPI(restServer, db)
}
