package models

type Code struct {
	CodeID  uint   `gorm:"primary_key;column:c_idx;not null;auto_increment" json:"codeId"`
	Type    string `gorm:"type:varchar(50);column:c_type" json:"type"`
	SubType string `gorm:"type:varchar(50);column:c_type_sub" json:"subType"`
	Name    string `gorm:"type:varchar(200);column:c_name" json:"name"`
	Order   int    `gorm:"column:c_order" json:"order"`
}

func (Code) TableName() string {
	return "code_tb"
}

type SubCode struct {
	ID     uint   `gorm:"primary_key;column:csub_idx;not null;auto_increment" json:"id"`
	Code   Code   `gorm:"foreignkey:CodeID" json:"code"`
	CodeID uint   `gorm:"column:c_idx" json:"codeId"`
	Name   string `gorm:"type:varchar(200);column:csub_name" json:"name"`
	Order  int    `gorm:"column:csub_order" json:"order"`
}

func (SubCode) TableName() string {
	return "code_sub_tb"
}

type SubCodeResponse struct {
	ID    uint   `gorm:"primary_key;column:csub_idx;not null;auto_increment" json:"id"`
	Name  string `gorm:"type:varchar(200);column:csub_name" json:"name"`
	Order int    `gorm:"column:csub_order" json:"order"`
	Code  `json:"code"`
}
