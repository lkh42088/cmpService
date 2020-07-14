package convert

import (
	"cmpService/dbmigrator/config"
	"cmpService/dbmigrator/insert"
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
