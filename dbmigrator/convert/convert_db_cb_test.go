package convert

import (
	"cmpService/dbmigrator/config"
	"testing"
)

func TestCbCreateNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.cb.conf"))
	CreateNewMariadbTable()
}

func TestCbConvertDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.cb.conf"))
	RunConvertDb()
}

func TestCbDeleteDpb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.cb.conf"))
	DeleteDeviceTb()
}

func TestCbClearNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.cb.conf"))
	DropNewMariadbTable()
}
