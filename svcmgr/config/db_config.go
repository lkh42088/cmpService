package config

import (
	"cmpService/common/config"
	db2 "cmpService/common/db"
	"cmpService/common/mariadblayer"
	"fmt"
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

