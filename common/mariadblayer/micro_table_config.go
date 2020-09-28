package mariadblayer

import (
	"cmpService/common/mcmodel"
	"github.com/jinzhu/gorm"
)

// For mcagent
func CreateMcPcTable(db *gorm.DB) {
	if db.HasTable(&mcmodel.McVm{}) == false {
		db.AutoMigrate(&mcmodel.McVm{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McVm{})
	}

	if db.HasTable(&mcmodel.McServer{}) == false {
		db.AutoMigrate(&mcmodel.McServer{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McServer{})
	}

	if db.HasTable(&mcmodel.McVmSnapshot{}) == false {
		db.AutoMigrate(&mcmodel.McVmSnapshot{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McVmSnapshot{})
	}

	if db.HasTable(&mcmodel.McVmBackup{}) == false {
		db.AutoMigrate(&mcmodel.McVmBackup{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McVmSnapshot{})
	}
}

func DropMcPcTable(db *gorm.DB) {
	if db.HasTable(&mcmodel.McVm{}) {
		db.DropTable(&mcmodel.McVm{})
	}

	if db.HasTable(&mcmodel.McServer{}) {
		db.DropTable(&mcmodel.McServer{})
	}

	if db.HasTable(&mcmodel.McVmSnapshot{}) == false {
		db.AutoMigrate(&mcmodel.McVmSnapshot{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McVmSnapshot{})
	}

	if db.HasTable(&mcmodel.McVmBackup{}) == false {
		db.AutoMigrate(&mcmodel.McVmBackup{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McVmSnapshot{})
	}

}

// For svcmgr
func CreateMicroCloudTable(db *gorm.DB) {
	if db.HasTable(&mcmodel.McServer{}) == false {
		db.AutoMigrate(&mcmodel.McServer{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McServer{})
	}

	if db.HasTable(&mcmodel.McImages{}) == false {
		db.AutoMigrate(&mcmodel.McImages{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McImages{})
	}

	if db.HasTable(&mcmodel.McNetworks{}) == false {
		db.AutoMigrate(&mcmodel.McNetworks{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McNetworks{})
	}

	if db.HasTable(&mcmodel.McNetHost{}) == false {
		db.AutoMigrate(&mcmodel.McNetHost{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McNetHost{})
	}

	if db.HasTable(&mcmodel.McVm{}) == false {
		db.AutoMigrate(&mcmodel.McVm{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McVm{})
	}

	if db.HasTable(&mcmodel.McVmSnapshot{}) == false {
		db.AutoMigrate(&mcmodel.McVmSnapshot{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McVmSnapshot{})
	}

	if db.HasTable(&mcmodel.McVmBackup{}) == false {
		db.AutoMigrate(&mcmodel.McVmBackup{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.McVmSnapshot{})
	}

	// Baremetal system info
	if db.HasTable(&mcmodel.SysInfo{}) == false {
		db.AutoMigrate(&mcmodel.SysInfo{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&mcmodel.SysInfo{})
	}
}

func DropMicroCloudTable(db *gorm.DB) {
	if db.HasTable(&mcmodel.McServer{}) {
		db.DropTable(&mcmodel.McServer{})
	}

	if db.HasTable(&mcmodel.McImages{}) {
		db.DropTable(&mcmodel.McImages{})
	}

	if db.HasTable(&mcmodel.McNetworks{}) {
		db.DropTable(&mcmodel.McNetworks{})
	}

	if db.HasTable(&mcmodel.McNetHost{}) {
		db.DropTable(&mcmodel.McNetHost{})
	}

	if db.HasTable(&mcmodel.McVm{}) {
		db.DropTable(&mcmodel.McVm{})
	}

	if db.HasTable(&mcmodel.McVmSnapshot{}) {
		db.DropTable(&mcmodel.McVmSnapshot{})
	}

	if db.HasTable(&mcmodel.McVmBackup{}) {
		db.DropTable(&mcmodel.McVmBackup{})
	}

	if db.HasTable(&mcmodel.SysInfo{}) {
		db.DropTable(&mcmodel.SysInfo{})
	}
}
