package convert

import (
	"fmt"
	"nubes/common/db"
	"nubes/common/mariadblayer"
	"nubes/dbmigrator/config"
	"nubes/dbmigrator/mysqllayer"
)

func CreateNewMariadbTable() {
	// New Database: Mariadb
	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()
	mariadblayer.CreateTable(newDb.DB)
}

func DropNewMariadbTable() {
	// New Database: Mariadb
	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()
	mariadblayer.DropTable(newDb.DB)
}

func CreateTestCbMysqlTable() {
	// New Database: Mariadb
	newConfig := config.GetTestCbDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()
	mysqllayer.CreateCbTable(newDb.DB)
}

func DropTestCbMysqlTable() {
	// New Database: Mariadb
	newConfig := config.GetTestCbDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()
	mysqllayer.DropCbTable(newDb.DB)
}
