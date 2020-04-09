package mariadblayer

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"nubes/common/config"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, dataSource string) (*DBORM, error) {
	db, err := gorm.Open(dbname, dataSource)
	return &DBORM {
		DB: db,
	}, err
}

func Init(config *config.DBConfig) {
	arg := GetDataSourceName(config)
	fmt.Println(arg)
	db, err := sql.Open("mysql", arg)
	if err != nil {
		fmt.Println("ERROR...")
		fmt.Println(err)
		return
	}
	fmt.Println("Connect...")
	defer db.Close()

	var version string
	db.QueryRow("SELECT version()").Scan(&version)
	fmt.Println("Connected to:", version)
}

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
	options := GetDataSourceName(config)
	db, err := gorm.Open(config.DBDriver, options)
	if err != nil {
		fmt.Println("Connect: Error", err)
		fmt.Println("options:", options)
		return nil
	}
	return db
}

