package mysqllayer

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Contents Bridge Database ORM
type CBORM struct {
	*gorm.DB
}

func NewCBORM(dbname, dataSource string) (*CBORM, error) {
	db, err := gorm.Open(dbname, dataSource)
	return &CBORM{
		DB: db,
	}, err
}
