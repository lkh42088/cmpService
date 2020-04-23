package mariadblayer

import (
	"github.com/jinzhu/gorm"
	"nubes/common/models"
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

	//Device
	if db.HasTable(&models.Device{}) == false {
		db.AutoMigrate(&models.Device{})
		db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	if db.HasTable(&models.DeviceComment{}) == false {
		db.AutoMigrate(&models.DeviceComment{})
		db.Model(&models.DeviceComment{}).AddForeignKey("dv_idx", "device_tb(dv_idx)", "RESTRICT", "RESTRICT")
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.DeviceComment{})
	}

}

func DropTable(db *gorm.DB) {
	if db.HasTable(&models.SubCode{}) {
		db.DropTable(&models.SubCode{})
	}
	if db.HasTable(&models.Code{}) {
		db.DropTable(&models.Code{})
	}

	// Device
	if db.HasTable(&models.DeviceComment{}) {
		db.DropTable(&models.DeviceComment{})
	}
	if db.HasTable(&models.Device{}) {
		db.DropTable(&models.Device{})
	}
}
