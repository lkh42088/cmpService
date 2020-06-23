package mysqllayer

import (
	"cmpService/dbmigrator/cbmodels"
	"github.com/jinzhu/gorm"
)

// 192.168.10.33
// cmpService/Nubes1510!
// database: db_db
func CreateCbTable(db *gorm.DB) {
	if db.HasTable(&cbmodels.Item{}) == false {
		db.AutoMigrate(&cbmodels.Item{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&cbmodels.Item{})
	}
	if db.HasTable(&cbmodels.SubItem{}) == false {
		db.AutoMigrate(&cbmodels.SubItem{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&cbmodels.SubItem{})
	}
	if db.HasTable(&cbmodels.CbDevice{}) == false {
		db.AutoMigrate(&cbmodels.CbDevice{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&cbmodels.CbDevice{})
	}
}

func DropCbTable(db *gorm.DB) {
	if db.HasTable(&cbmodels.Item{}) {
		db.DropTable(&cbmodels.Item{})
	}
	if db.HasTable(&cbmodels.SubItem{}) {
		db.DropTable(&cbmodels.SubItem{})
	}
	if db.HasTable(&cbmodels.CbDevice{}) {
		db.DropTable(&cbmodels.CbDevice{})
	}
}
