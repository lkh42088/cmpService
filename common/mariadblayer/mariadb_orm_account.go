package mariadblayer

import (
	. "cmpService/common/models"
	"github.com/jinzhu/gorm"
)

func (db *DBORM) GetAccount() (accounts []Account, err error) {
	return accounts, db.Find(&accounts).Error
}

func (db *DBORM) AddAccount(account Account) (Account, error) {
	return account, db.Create(&account).Error
}

func (db *DBORM) DeleteAccount(account Account) (Account, error) {
	return account, db.Delete(&account).Error
}

func CreateAccountTable(db *gorm.DB) {
	if db.HasTable(&Account{}) == false {
		db.AutoMigrate(&Account{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Account{})
	}
}

func DeleteAccountTable(db *gorm.DB) {
	if db.HasTable(&Account{}) {
		db.DropTable(&Account{})
	}
}
