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

	// Device
	// -> device server
	if db.HasTable(&models.DeviceServer{}) == false {
		db.AutoMigrate(&models.DeviceServer{})
		db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	// -> device network
	if db.HasTable(&models.DeviceNetwork{}) == false {
		db.AutoMigrate(&models.DeviceNetwork{})
		db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	// -> device part
	if db.HasTable(&models.DevicePart{}) == false {
		db.AutoMigrate(&models.DevicePart{})
		db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	// -> device comment
	if db.HasTable(&models.DeviceComment{}) == false {
		db.AutoMigrate(&models.DeviceComment{})
		db.Set("gorm:table_options", "ENGINE=InnoDB")
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
	// -> device server
	if db.HasTable(&models.DeviceServer{}) {
		db.DropTable(&models.DeviceServer{})
	}
	// -> device network
	if db.HasTable(&models.DeviceNetwork{}) {
		db.DropTable(&models.DeviceNetwork{})
	}
	// -> device part
	if db.HasTable(&models.DevicePart{}) {
		db.DropTable(&models.DevicePart{})
	}
	// -> device comment
	if db.HasTable(&models.DeviceComment{}) {
		db.DropTable(&models.DeviceComment{})
	}
}
