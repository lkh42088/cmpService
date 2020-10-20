package convert

import (
	"cmpService/common/db"
	"cmpService/common/mariadblayer"
	"cmpService/dbmigrator/config"
	"cmpService/dbmigrator/insert"
	"fmt"
	"testing"
)

func TestLkhCreateNewMariaDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.lkh.conf"))
	CreateNewMariadbTable()
}

func TestLkhConvertDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.lkh.conf"))
	RunConvertDb()
}

func TestLkhInsertItem(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.lkh.conf"))
	insert.InsertNewObject()
}

func TestLkhDeleteDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.lkh.conf"))
	DeleteDeviceTb()
}

func TestLkhClearNewMariaDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.lkh.conf"))
	DropNewMariadbTable()
}

//***************************************************************************
// Micro Cloud
//***************************************************************************
func TestLkhMicroCloudCreateDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.lkh.conf"))

	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()
	mariadblayer.CreateMicroCloudTable(newDb.DB)
}

func TestLkhMicroCloudDropDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.lkh.conf"))

	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()
	mariadblayer.DropMicroCloudTable(newDb.DB)
}

//***************************************************************************
// Micro Cloud PC
//***************************************************************************
func TestLkhMcPcCreateDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.mc.lkh.conf"))

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

func TestLkhMcPcDropDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.mc.lkh.conf"))

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
