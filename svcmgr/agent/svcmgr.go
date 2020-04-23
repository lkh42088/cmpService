package main

import (
	"fmt"
	"nubes/common/config"
	db2 "nubes/common/db"
	"nubes/common/mariadblayer"
	config2 "nubes/svcmgr/config"
	"nubes/svcmgr/rest"
)

func main () {
	db, err := SetDatabase()
	if err != nil {
		fmt.Println("Main: ERROR - ", err)
		return
	}
	SetRestServer(db)
/*	http.HandleFunc("/", serveStatic)
	http.ListenAndServe(Port, nil)*/
}

func SetDatabase() (db *mariadblayer.DBORM, err error) {
	dbconfig, err := config.NewDBConfig("mysql", "nubes", "nubes1510!",
		"nubes","192.168.227.129", 3306)
	/*dbconfig, err := config.NewDBConfig("mysql", "nubes", "nubes1510",
		"nubes","192.168.122.127", 3306)*/
	if err != nil {
		fmt.Println("[SetDatabase] Error:", err)
		return
	}
	config2.SvcmgrConfig.MariadbConfig = *dbconfig
	dataSource := db2.GetDataSourceName(dbconfig)
	db, err = mariadblayer.NewDBORM(dbconfig.DBDriver, dataSource)
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

/*func serveStatic(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("src/main/template/public/index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}*/