package mariadblayer

import (
	"fmt"
	db2 "cmpService/common/db"
	"cmpService/common/models"
	"testing"
)

func TestCodeAddEntry(t *testing.T) {
	config := getTestConfig()
	options := db2.GetDataSourceName(config)
	fmt.Println("config:", config)
	fmt.Println("options:", options)
	db, err := NewDBORM(config.DBDriver, options)
	if err != nil {
		return
	}
	code := models.Code{
		Type:    "type1",
		SubType: "subtype1",
		Name:    "codename1",
		Order:   1,
	}
	code, err = db.AddCode(code)
	fmt.Println("code: ", code, "err:", err)

	code = models.Code{
		Type:    "type2",
		SubType: "subtype2",
		Name:    "codename2",
		Order:   2,
	}
	code, err = db.AddCode(code)
	fmt.Println("code: ", code, "err:", err)
}

func TestCodeGetEntry(t *testing.T) {
	config := getTestConfig()
	options := db2.GetDataSourceName(config)
	fmt.Println("config:", config)
	fmt.Println("options:", options)
	db, err := NewDBORM(config.DBDriver, options)
	if err != nil {
		return
	}
	codes, err := db.GetAllCodes()
	fmt.Println("codes:", codes)
	fmt.Println("err:", err)
}

func TestCodeDeleteLastEntry(t *testing.T) {
	config := getTestConfig()
	options := db2.GetDataSourceName(config)
	fmt.Println("config:", config)
	fmt.Println("options:", options)
	db, err := NewDBORM(config.DBDriver, options)
	if err != nil {
		return
	}
	codes, err := db.GetAllCodes()
	fmt.Println("codes:", codes)
	fmt.Println("err:", err)

	var last models.Code
	for num, code := range codes {
		fmt.Println("num:", num)
		fmt.Println("code:", code)
		last = code
	}
	if last.CodeID != 0 {
		last, err = db.DeleteCode(last)
		fmt.Println("-->Delete last entry :", last)
		fmt.Println("-->err:", err)
	}
}

func TestCodeDeletesEntry(t *testing.T) {
	config := getTestConfig()
	options := db2.GetDataSourceName(config)
	fmt.Println("config:", config)
	fmt.Println("options:", options)
	db, err := NewDBORM(config.DBDriver, options)
	if err != nil {
		return
	}

	err = db.DeleteCodes()
	fmt.Println("err:", err)
}

func TestSubCodeAddEntry(t *testing.T) {
	config := getTestConfig()
	options := db2.GetDataSourceName(config)
	fmt.Println("config:", config)
	fmt.Println("options:", options)
	db, err := NewDBORM(config.DBDriver, options)
	if err != nil {
		return
	}
	code := models.SubCode{
		ID:     0,
		Code:   models.Code{},
		CodeID: 0,
		Name:   "",
		Order:  0,
	}
	code, err = db.AddSubCode(code)
	fmt.Println("code: ", code, "err:", err)
	code = models.SubCode{
		//Type:    "type2",
		//SubType: "subtype2",
		Name:    "codename2",
		Order:   2,
	}
	code, err = db.AddSubCode(code)
	fmt.Println("code: ", code, "err:", err)
}

func TestSubCodeGetEntry(t *testing.T) {
	config := getTestConfig()
	options := db2.GetDataSourceName(config)
	fmt.Println("config:", config)
	fmt.Println("options:", options)
	db, err := NewDBORM(config.DBDriver, options)
	if err != nil {
		return
	}
	codes, err := db.GetAllSubCodes()
	fmt.Println("codes:", codes)
	fmt.Println("err:", err)
}

func TestSubCodeDeleteLastEntry(t *testing.T) {
	config := getTestConfig()
	options := db2.GetDataSourceName(config)
	fmt.Println("config:", config)
	fmt.Println("options:", options)
	db, err := NewDBORM(config.DBDriver, options)
	if err != nil {
		return
	}
	codes, err := db.GetAllSubCodes()
	fmt.Println("codes:", codes)
	fmt.Println("err:", err)

	var last models.SubCode
	for num, code := range codes {
		fmt.Println("num:", num)
		fmt.Println("code:", code)
		last = code
	}

	if last.ID != 0 {
		last, err = db.DeleteSubCode(last)
		fmt.Println("-->Delete last entry :", last)
		fmt.Println("-->err:", err)
	}
}

func TestSubCodeDeletesEntry(t *testing.T) {
	config := getTestConfig()
	options := db2.GetDataSourceName(config)
	fmt.Println("config:", config)
	fmt.Println("options:", options)
	db, err := NewDBORM(config.DBDriver, options)
	if err != nil {
		return
	}

	err = db.DeleteSubCodes()
	fmt.Println("err:", err)
}
