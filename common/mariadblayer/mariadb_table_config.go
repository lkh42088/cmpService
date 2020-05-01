package mariadblayer

import (
	"github.com/jinzhu/gorm"
	"nubes/common/models"
)

func CreateTable(db *gorm.DB) {
	if db.HasTable(&models.Code{}) == false {
		db.AutoMigrate(&models.Code{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Code{})
	}
	if db.HasTable(&models.SubCode{}) == false {
		db.AutoMigrate(&models.SubCode{})
		db.Model(&models.SubCode{}).AddForeignKey("c_idx",
			"code_tb(c_idx)", "RESTRICT", "RESTRICT") // or CASCADE
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.SubCode{})
	}

	// Device
	// -> collectdevice server
	if db.HasTable(&models.DeviceServer{}) == false {
		db.AutoMigrate(&models.DeviceServer{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.DeviceServer{})
	}
	// -> collectdevice network
	if db.HasTable(&models.DeviceNetwork{}) == false {
		db.AutoMigrate(&models.DeviceNetwork{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.DeviceNetwork{})
	}
	// -> collectdevice part
	if db.HasTable(&models.DevicePart{}) == false {
		db.AutoMigrate(&models.DevicePart{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.DevicePart{})
	}
	// -> collectdevice comment
	if db.HasTable(&models.DeviceComment{}) == false {
		db.AutoMigrate(&models.DeviceComment{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.DeviceComment{})
	}

	// Login
	if db.HasTable(&models.User{}) == false {
		db.AutoMigrate(&models.User{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
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
	// -> collectdevice server
	if db.HasTable(&models.DeviceServer{}) {
		db.DropTable(&models.DeviceServer{})
	}
	// -> collectdevice network
	if db.HasTable(&models.DeviceNetwork{}) {
		db.DropTable(&models.DeviceNetwork{})
	}
	// -> collectdevice part
	if db.HasTable(&models.DevicePart{}) {
		db.DropTable(&models.DevicePart{})
	}
	// -> collectdevice comment
	if db.HasTable(&models.DeviceComment{}) {
		db.DropTable(&models.DeviceComment{})
	}

	// Login
	if db.HasTable(&models.User{}) {
		db.DropTable(&models.User{})
	}
}
