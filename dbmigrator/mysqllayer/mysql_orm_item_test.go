package mysqllayer

import (
	"fmt"
	"cmpService/common/db"
	"testing"
)

func TestItemGetEntry (t *testing.T) {
	config := getMysqlConfig()
	options := db.GetDataSourceName(config)
	db, err := NewCBORM(config.DBDriver, options)
	fmt.Println(options)
	if err != nil {
		fmt.Println(err)
		return
	}
	items, err := db.GetAllItems()
	if err != nil {
		fmt.Println(err)
	}
	for num, item := range items {
		fmt.Println(num, ":", item)
		//code := convert.GetCodeByItem(item)
		//fmt.Println("-->", code)
	}
	//fmt.Println(items)
}

func TestDeviceGetEntry (t *testing.T) {
	config := getMysqlConfig()
	options := db.GetDataSourceName(config)
	db, err := NewCBORM(config.DBDriver, options)
	fmt.Println(options)
	if err != nil {
		fmt.Println(err)
		return
	}
	devices, err := db.GetAllDevices()
	if err != nil {
		fmt.Println(err)
	}
	for num, device := range devices {
		fmt.Println(num, ":", device)
	}
}

