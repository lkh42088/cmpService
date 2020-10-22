package mariadblayer

import (
	"cmpService/common/config"
	"cmpService/common/db"
	"cmpService/common/models"
	"fmt"
	"testing"
)

func getTestLkhConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"Nubes1510!",
		"nubes",
		"192.168.0.44",
		3306,
	}
	return &config
}

func getTestLkhDb() (*DBORM, error) {
	config := getTestLkhConfig()
	options := db.GetDataSourceName(config)
	db, err := NewDBORM(config.DBDriver, options)
	return db, err
}

func TestDBORM_GetDevicesServerWithJoin(t *testing.T) {
	cri := models.PageCreteria{
		Count:      0,
		TotalPage:  0,
		CheckCnt:   1,
		Size:       1000,
		OutFlag:    "0",
		OrderKey:   "",
		Direction:  0,
		DeviceType: "",
	}
	db, err := getTestLkhDb()
	if err != nil {
		return
	}
	data, _ := db.GetDevicesServerWithJoin(cri)
	//data, _ := db.GetDevicesNetworkWithJoin(cri)
	//data, _ := db.GetDevicesPartWithJoin(cri)
	for i, v := range data.Devices {
		fmt.Println(i, "    ", v)
	}
}

