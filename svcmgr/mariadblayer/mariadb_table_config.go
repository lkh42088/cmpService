package mariadblayer

import (
	"github.com/jinzhu/gorm"
	"nubes/svcmgr/models"
)

func CreateTable(db *gorm.DB) {
	if db.HasTable(&models.Code{}) == false {
		db.AutoMigrate(&models.Code{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.SubCode{})
	}
	if db.HasTable(&models.SubCode{}) == false {
		db.AutoMigrate(&models.SubCode{})
		db.Model(&models.SubCode{}).AddForeignKey("c_idx",
			"code_tb(c_idx)", "RESTRICT", "RESTRICT") // or CASCADE
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.SubCode{})
	}
}

func DropTable(db *gorm.DB) {
	if db.HasTable(&models.SubCode{}) {
		db.DropTable(&models.SubCode{})
	}
	if db.HasTable(&models.Code{}) {
		db.DropTable(&models.Code{})
	}
}
