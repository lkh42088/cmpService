package mariadblayer

import (
	"cmpService/common/mcmodel"
	"github.com/jinzhu/gorm"
)

func CreateMicroCloudTable(db *gorm.DB) {
	if db.HasTable(&mcmodel.McServer{}) == false {
		db.AutoMigrate(&mcmodel.McServer{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McServer{})
	}

	if db.HasTable(&mcmodel.McVm{}) == false {
		db.AutoMigrate(&mcmodel.McVm{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McVm{})
	}
}

func DropMicroCloudTable(db *gorm.DB) {
	if db.HasTable(&mcmodel.McServer{}) {
		db.DropTable(&mcmodel.McServer{})
	}

	if db.HasTable(&mcmodel.McVm{}) {
		db.DropTable(&mcmodel.McVm{})
	}
}
