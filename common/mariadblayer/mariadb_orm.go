package mariadblayer

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewDBORM(dbname, dataSource string) (*DBORM, error) {
	db, err := gorm.Open(dbname, dataSource)
	return &DBORM{
		DB: db,
	}, err
}

