package main

import (
	"fmt"
	"nubes/common/config"
	config2 "nubes/svcmgr/config"
	"nubes/svcmgr/mariadblayer"
	"nubes/svcmgr/rest"
)

func main () {
	db, err := SetDatabase()
	if err != nil {
		fmt.Println("Main: ERROR - ", err)
		return
	}
	SetRestServer(db)
}

func SetDatabase() (db *mariadblayer.DBORM, err error) {
	dbconfig, err := config.NewDBConfig("mysql", "nubes", "nubes1510",
		"nubes","192.168.122.127", 3306)
	if err != nil {
		fmt.Println("[SetDatabase] Error:", err)
		return
	}
	config2.SvcmgrConfig.MariadbConfig = *dbconfig
	dataSource := mariadblayer.GetDataSourceName(dbconfig)
	db, err = mariadblayer.NewORM(dbconfig.DBDriver, dataSource)
	if err != nil {
		fmt.Println("[SetDatabase] Error:", err)
		return
	}
	config2.SvcmgrConfig.Mariadb = *db
	return db, err
}

func SetRestServer(db *mariadblayer.DBORM) {
	restServer := "0.0.0.0:8081"
	config2.SvcmgrConfig.RestServer = restServer
	rest.RunAPI(restServer, db)
}
