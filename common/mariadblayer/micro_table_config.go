package mariadblayer

import (
	"cmpService/common/models"
	"github.com/jinzhu/gorm"
)

func CreateMicroCloudTable(db *gorm.DB) {
	if db.HasTable(&models.McServer{}) == false {
		db.AutoMigrate(&models.McServer{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.McServer{})
	}

	if db.HasTable(&models.McVm{}) == false {
		db.AutoMigrate(&models.McVm{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.McVm{})
	}
}

func DropMicroCloudTable(db *gorm.DB) {
	if db.HasTable(&models.McServer{}) {
		db.DropTable(&models.McServer{})
	}

	if db.HasTable(&models.McVm{}) {
		db.DropTable(&models.McVm{})
	}
}
