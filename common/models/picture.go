package models

type Account struct {
	Idx uint `gorm:"primary_key;column:account_idx;not null;auto_increment;comment:'INDEX'"`
	Name string `gorm:"type:varchar(16);column:name;comment:'user name'"`
	Picture []byte `gorm:"column:picture;comment:'user picture'"`
}

func (Account) TableName() string {
	return "account_tb"
}

