package convert

import (
	"cmpService/dbmigrator/config"
	"testing"
)

func TestJbhHomeCreateNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbhhome.conf"))
	CreateNewMariadbTable()
}

func TestJbhHomeConvertDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbhhome.conf"))
	RunConvertDb()
}

func TestJbhHomeDeleteDpb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbhhome.conf"))
	DeleteDeviceTb()
}

func TestJbhHomeClearNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.jbhhome.conf"))
	DropNewMariadbTable()
}
