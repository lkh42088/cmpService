package convert

import (
	"cmpService/dbmigrator/config"
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

func TestLkhDeleteDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.lkh.conf"))
	DeleteDeviceTb()
}

func TestLkhClearNewMariaDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.lkh.conf"))
	DropNewMariadbTable()
}
