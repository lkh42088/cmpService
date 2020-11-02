package convert

import (
	"cmpService/common/db"
	"cmpService/common/mariadblayer"
	"cmpService/dbmigrator/config"
	"fmt"
	"testing"
)

//***************************************************************************
// Micro Cloud PC
//***************************************************************************
func TestJbhDevMcPcCreateDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.mc.jbhdev.conf"))

	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()
	mariadblayer.CreateMcPcTable(newDb.DB)
}

func TestJbhDevMcPcDropDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.mc.jbhdev.conf"))

	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()
	mariadblayer.DropMcPcTable(newDb.DB)
}

