package convert

import (
	"cmpService/dbmigrator/config"
	"cmpService/dbmigrator/insert"
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
