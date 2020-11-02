package convert

import (
	"cmpService/common/db"
	"cmpService/common/mariadblayer"
	"cmpService/dbmigrator/config"
	"cmpService/dbmigrator/insert"
	"fmt"
	"testing"
)

func TestJbhDevCreateNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbhdev.conf"))
	CreateNewMariadbTable()
}

func TestJbhDevConvertDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbhdev.conf"))
	RunConvertDb()
}

func TestJbhDevInsertItem(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbhdev.conf"))
	insert.InsertNewObject()
}

func TestJbhDevDeleteDpb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbhdev.conf"))
	DeleteDeviceTb()
}

func TestJbhDevClearNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbhdev.conf"))
	DropNewMariadbTable()
}

//***************************************************************************
// Micro Cloud
//***************************************************************************
func TestJbhDevMicroCloudCreateDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbhdev.conf"))

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

func TestJbhDevMicroCloudDropDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbhdev.conf"))

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
