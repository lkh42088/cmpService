package models

import "time"

type Device struct {
	Idx              uint      `gorm:"primary_key;column:dv_idx;not null"`
	MenuType         string    `gorm:"type:varchar(50);column:dv_type"`
	OutFlag          bool      `gorm:"type:tinyint(1);column:dv_out_flag;default:0"`
	Num              int       `gorm:"column:dv_num"`
	CommentCnt       int       `gorm:"column:dv_comment_cnt"`
	CommentLastDate  time.Time `gorm:"column:dv_comment_last_date"`
	Option           string    `gorm:"type:varchar(50);column:dv_option"`
	Hit              int       `gorm:"column:dv_hit"`
	RegisterId       string    `gorm:"type:varchar(50);column:register_id"`
	Password         string    `gorm:"type:varchar(255);column:dv_register_password"`
	RegisterName     string    `gorm:"type:varchar(50);column:dv_register_name"`
	RegisterEmail    string    `gorm:"type:varchar(255);column:dv_register_email"`
	RegisterDate     time.Time `gorm:"column:register_date"`
	DeviceCode       string    `gorm:"type:varchar(255);column:dv_device_code"`
	Model            int       `gorm:"column:dv_model_cd"`
	Contents         string    `gorm:"type:text;column:dv_contents"`
	Customer         int       `gorm:"column:dv_customer_cd"`
	Manufacture      int       `gorm:"column:dv_manufacture_cd"`
	DeviceType       int       `gorm:"column:dv_device_type_cd"`
	WarehousingDate  string    `gorm:"type:varchar(10);column:dv_warehousing_date"`
	RentDate         string    `gorm:"type:varchar(20);column:dv_rent_date;default:'|'"`
	Ownership        string    `gorm:"type:varchar(10);column:dv_ownership_cd;default:'|'"`
	OwnerCompany     string    `gorm:"type:varchar(255);column:dv_owner_company"`
	HwSn             string    `gorm:"type:varchar(255);column:dv_hw_sn"`
	IDC              int       `gorm:"column:dv_idc_cd"`
	Rack             int       `gorm:"column:dv_rack_cd"`
	Cost             string    `gorm:"type:varchar(255);column:dv_cost"`
	Purpos           string    `gorm:"type:varchar(255);column:dv_purpos"`
	Ip               string    `gorm:"type:varchar(255);column:dv_ip;default:'|'"`
	Size             int       `gorm:"column:dv_size_cd"`
	Spla             int       `gorm:"column:dv_spla_cd"`
	Cpu              string    `gorm:"type:varchar(255);column:dv_cpu"`
	Memory           string    `gorm:"type:varchar(255);column:dv_memory"`
	Hdd              string    `gorm:"type:varchar(255);column:dv_hdd"`
	FirmwareVersion  string    `gorm:"type:varchar(50);column:dv_firmware_version"`
	Warranty         string    `gorm:"type:varchar(255);column:dv_warranty"`
	MonitoringFlag   int       `gorm:"column:dv_mornitoring_flag"`
	MonitoringMethod int       `gorm:"column:dv_mornitoring_method"`
}

func (Device) TableName() string {
	return "device_tb"
}

type DeviceComment struct {
	Idx          uint      `gorm:"primary_key;column:dvc_idx;not null"`
	ForeignIdx   Device    `gorm:"foreignkey:Idx;not null"`
	Depth        int       `gorm:"column:dvc_depth"`
	Contents     string    `gorm:"column:dvc_contents"`
	RegisterId   string    `gorm:"type:varchar(50);column:register_id"`
	RegisterName string    `gorm:"type:varchar(50);column:dvc_register_name"`
	RegisterDate time.Time `gorm:"column:register_date"`
}

func (DeviceComment) TableName() string {
	return "device_comment_tb"
}
