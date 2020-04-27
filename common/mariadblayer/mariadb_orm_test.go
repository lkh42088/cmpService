package mariadblayer

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"nubes/common/config"
	"nubes/common/db"
	"testing"
)

func getTestDefaultConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"nubes1510!",
		"nubes",
		"127.0.0.1",
		3306,
	}

	return &config
}

func getTestJebConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"nubes1510!",
		"nubes",
		"192.168.227.129",
		3306,
	}

	return &config
}

func getTestJbhConfig() *config.DBConfig {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"nubes1510",
		"nubes",
		"192.168.10.115",
		3306,
	}

	return &config
}

func getTestConfig() *config.DBConfig {
	return getTestJbhConfig()
}

func Migration(conf config.DBConfig) {
	type Product2 struct {
		gorm.Model
		Code string
		Price uint
	}
	db, err := gorm.Open(conf.DBDriver, db.GetDataSourceName(&conf))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Product2{})

	db.Create(&Product2{Code: "L1212", Price: 1000})

	var product Product2
	db.First(&product, 1)
	db.First(&product, "code = ?", "L1212")

	db.Model(&product).Update("Price", 2000)

	db.Delete(&product)
}

func TestInit(t *testing.T) {
	config := getTestConfig()
	fmt.Println(config)
	db.Init(config)
}

func TestMigrtion(t *testing.T) {
	config := config.DBConfig{
		"mysql",
		"nubes",
		"nubes1510",
		"nubes",
		"192.168.122.127",
		3306,
	}
	fmt.Println(config)
	Migration(config)
}

