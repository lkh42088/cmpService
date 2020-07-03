package convert

import (
	"cmpService/dbmigrator/config"
	"testing"
)

func TestNubesCreateNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.nubes.conf"))
	CreateNewMariadbTable()
}

func TestNubesConvertDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.nubes.conf"))
	RunConvertDb()
}

func TestNubesDeleteDpb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.nubes.conf"))
	DeleteDeviceTb()
}

func TestNubesClearNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.nubes.conf"))
	DropNewMariadbTable()
}
