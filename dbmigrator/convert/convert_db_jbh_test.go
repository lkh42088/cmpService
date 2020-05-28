package convert

import (
	"cmpService/dbmigrator/config"
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

func TestJbhDeleteDpb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbh.conf"))
	DeleteDeviceTb()
}

func TestJbhClearNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbh.conf"))
	DropNewMariadbTable()
}
