package mysqllayer

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"nubes/common/config"
	config2 "nubes/dbmigrator/config"
	"testing"
)

func GetDataSourceName(config *config.DBConfig) string {
	options := fmt.Sprint("?charset=utf8mb4&parseTime=True&loc=Local")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		config.Username,
		config.Password,
		config.Address,
		config.Port,
		config.DBName,
		options)
}

func Connect(config *config.DBConfig) *gorm.DB {
	fmt.Println(config)
	options := GetDataSourceName(config)
	db, err := gorm.Open(config.DBDriver, options)
	if err != nil {
		fmt.Println("Connect: Error", err)
		fmt.Println("options:", options)
		return nil
	}
	return db
}

func TestCreateCbTable(t *testing.T) {
	config := config2.GetTestCbDatabaseConfig()
	fmt.Println("TEST: ", config)
	db := Connect(config)
	if db == nil {
		return
	}
	defer db.Close()
	CreateCbTable(db)
}
