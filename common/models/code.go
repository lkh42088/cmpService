package models

type Code struct {
	CodeID  uint   `gorm:"primary_key;column:c_idx;not null;auto_increment"`
	Type    string `gorm:"type:varchar(50);column:c_type"`
	SubType string `gorm:"type:varchar(50);column:c_type_sub"`
	Name    string `gorm:"type:varchar(200);column:c_name"`
	Order   int    `gorm:"column:c_order"`
}

func (Code) TableName() string {
	return "code_tb"
}

type SubCode struct {
	ID     uint   `gorm:"primary_key;column:csub_idx;not null;auto_increment"`
	Code   Code   `gorm:"foreignkey:CodeID"`
	CodeID uint   `gorm:"column:c_idx"`
	Name   string `gorm:"type:varchar(200);column:csub_name"`
	Order  int    `gorm:"column:csub_order"`
}

func (SubCode) TableName() string {
	return "code_sub_tb"
}

