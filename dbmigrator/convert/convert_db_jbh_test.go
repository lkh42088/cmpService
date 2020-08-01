package convert

import (
	"cmpService/common/db"
	"cmpService/common/mariadblayer"
	"cmpService/dbmigrator/config"
	"cmpService/dbmigrator/insert"
	"fmt"
	"testing"
)

func TestJbhCreateNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbh.conf"))
	CreateNewMariadbTable()
}

func TestJbhConvertDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbh.conf"))
	RunConvertDb()
}

func TestJbhInsertItem(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbh.conf"))
	insert.InsertNewObject()
}

func TestJbhDeleteDpb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbh.conf"))
	DeleteDeviceTb()
}

func TestJbhClearNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbh.conf"))
	DropNewMariadbTable()
}

//***************************************************************************
// Micro Cloud
//***************************************************************************
func TestMicroCloudCreateDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbh.conf"))

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

func TestMicroCloudDropDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbh.conf"))

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
