package convert

import (
	"cmpService/dbmigrator/config"
	"cmpService/dbmigrator/insert"
	"fmt"
	"os"
	"testing"
)

func getDbConfig(name string) string {
	dirName, _ := os.Getwd()
	configPath := fmt.Sprintf("%s/../etc/%s", dirName, name)
	fmt.Println(configPath)
	return configPath
}

func TestCreateNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.conf"))
	CreateNewMariadbTable()
}

func TestConvertDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.conf"))
	RunConvertDb()
}

func TestInsertItem(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.conf"))
	insert.InsertNewObject()
}

func TestDeleteDpb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.conf"))
	DeleteDeviceTb()
}

func TestClearNewMariadbDb(t *testing.T) {
	config.SetConfig(getDbConfig("dbmigrator.conf"))
	DropNewMariadbTable()
}
