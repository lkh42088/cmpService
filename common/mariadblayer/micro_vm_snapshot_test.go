package mariadblayer

import (
	"cmpService/common/config"
	"cmpService/common/db"
	"fmt"
	"testing"
)

func getSnapConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"Nubes1510!",
		"nubes",
		"192.168.0.40",
		3306,
	}
	return &config
}

func getSnapDb() (*DBORM, error) {
	config := getSnapConfig()
	options := db.GetDataSourceName(config)
	db, err := NewDBORM(config.DBDriver, options)
	return db, err
}

func TestGetSnapListAll(t *testing.T) {
	db, err := getSnapDb()
	if err != nil {
		fmt.Println(err)
		return
	}

	snapList, err := db.GetMcVmSnapshotByVmName("vm01")
	fmt.Println(len(snapList))
	for index, obj := range snapList {
		fmt.Printf("%2d: %4d, %s\n", index, obj.Idx, obj.Name)
	}

}