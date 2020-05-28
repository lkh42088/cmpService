package mariadblayer

import (
	"cmpService/common/models"
	"github.com/jinzhu/gorm"
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
	// -> collectdevice log
	if db.HasTable(&models.DeviceLog{}) == false {
		db.AutoMigrate(&models.DeviceLog{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.DeviceLog{})
	}

	// User
	if db.HasTable(&models.User{}) == false {
		db.AutoMigrate(&models.User{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
	}
	if db.HasTable(&models.UserEmailAuth{}) == false {
		db.AutoMigrate(&models.UserEmailAuth{})
		db.Model(&models.UserEmailAuth{}).AddForeignKey("user_idx",
			"user_tb(user_idx)", "RESTRICT", "RESTRICT") // or CASCADE
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.UserEmailAuth{})
	}

	// User(temp), Customer, Auth
	//if db.HasTable(&models.User{}) == false {
	//	db.AutoMigrate(&models.User{})
	//	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
	//}
	if db.HasTable(&models.Company{}) == false {
		db.AutoMigrate(&models.Company{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Company{})
	}
	if db.HasTable(&models.Auth{}) == false {
		db.AutoMigrate(&models.Auth{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Auth{})
	}

	// IP, Subnet
	if db.HasTable(&models.IpMgmt{}) == false {
		db.AutoMigrate(&models.IpMgmt{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.IpMgmt{})
	}
	if db.HasTable(&models.SubnetMgmt{}) == false {
		db.AutoMigrate(&models.SubnetMgmt{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.SubnetMgmt{})
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

	// -> collectdevice log
	if db.HasTable(&models.DeviceLog{}) {
		db.DropTable(&models.DeviceLog{})
	}

	// User
	if db.HasTable(&models.User{}) {
		db.DropTable(&models.User{})
	}
	if db.HasTable(&models.UserEmailAuth{}) {
		db.DropTable(&models.UserEmailAuth{})
	}

	// User, Customer, Auth
	if db.HasTable(&models.User{}) {
		db.DropTable(&models.User{})
	}
	if db.HasTable(&models.Company{}) {
		db.DropTable(&models.Company{})
	}
	if db.HasTable(&models.Auth{}) {
		db.DropTable(&models.Auth{})
	}

	// IP, Subnet
	if db.HasTable(&models.IpMgmt{}) {
		db.DropTable(&models.IpMgmt{})
	}
	if db.HasTable(&models.SubnetMgmt{}) {
		db.DropTable(&models.SubnetMgmt{})
	}

}
