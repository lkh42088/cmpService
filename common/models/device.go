package models

import "time"

type DeviceServer struct {
	Idx              uint      `gorm:"primary_key;column:dvs_idx;not null"`
	OutFlag          bool      `gorm:"type:tinyint(1);column:dvs_out_flag;default:0"`
	Num              int       `gorm:"column:dvs_num"`
	CommentCnt       int       `gorm:"column:dvs_comment_cnt"`
	CommentLastDate  time.Time `gorm:"column:dvs_comment_last_date"`
	Option           string    `gorm:"type:varchar(50);column:dvs_option"`
	Hit              int       `gorm:"column:dvs_hit"`
	RegisterId       string    `gorm:"type:varchar(50);column:register_id"`
	Password         string    `gorm:"type:varchar(255);column:register_password"`
	RegisterName     string    `gorm:"type:varchar(50);column:register_name"`
	RegisterEmail    string    `gorm:"type:varchar(255);column:register_email"`
	RegisterDate     time.Time `gorm:"column:register_date"`
	DeviceCode       string    `gorm:"type:varchar(255);column:dvs_device_code"`
	Model            int       `gorm:"column:dvs_model_cd"`
	Contents         string    `gorm:"type:text;column:dvs_contents"`
	Customer         int       `gorm:"column:dvs_customer_cd"`
	Manufacture      int       `gorm:"column:dvs_manufacture_cd"`
	DeviceType       int       `gorm:"column:dvs_device_type_cd"`
	WarehousingDate  string    `gorm:"type:varchar(10);column:dvs_warehousing_date"`
	RentDate         string    `gorm:"type:varchar(20);column:dvs_rent_date;default:'|'"`
	Ownership        string    `gorm:"type:varchar(10);column:dvs_ownership_cd;default:'|'"`
	OwnerCompany     string    `gorm:"type:varchar(255);column:dvs_owner_company"`
	HwSn             string    `gorm:"type:varchar(255);column:dvs_hw_sn"`
	IDC              int       `gorm:"column:dvs_idc_cd"`
	Rack             int       `gorm:"column:dvs_rack_cd"`
	Cost             string    `gorm:"type:varchar(255);column:dvs_cost"`
	Purpos           string    `gorm:"type:varchar(255);column:dvs_purpos"`
	Ip               string    `gorm:"type:varchar(255);column:dvs_ip;default:'|'"`
	Size             int       `gorm:"column:dvs_size_cd"`
	Spla             int       `gorm:"column:dvs_spla_cd"`
	Cpu              string    `gorm:"type:varchar(255);column:dvs_cpu"`
	Memory           string    `gorm:"type:varchar(255);column:dvs_memory"`
	Hdd              string    `gorm:"type:varchar(255);column:dvs_hdd"`
	MonitoringFlag   int       `gorm:"column:dvs_mornitoring_flag"`
	MonitoringMethod int       `gorm:"column:dvs_mornitoring_method"`
}

func (DeviceServer) TableName() string {
	return "device_server_tb"
}

type DeviceNetwork struct {
	Idx              uint      `gorm:"primary_key;column:dvn_idx;not null"`
	OutFlag          bool      `gorm:"type:tinyint(1);column:dvn_out_flag;default:0"`
	Num              int       `gorm:"column:dvn_num"`
	CommentCnt       int       `gorm:"column:dvn_comment_cnt"`
	CommentLastDate  time.Time `gorm:"column:dvn_comment_last_date"`
	Option           string    `gorm:"type:varchar(50);column:dvn_option"`
	Hit              int       `gorm:"column:dvn_hit"`
	RegisterId       string    `gorm:"type:varchar(50);column:register_id"`
	Password         string    `gorm:"type:varchar(255);column:register_password"`
	RegisterName     string    `gorm:"type:varchar(50);column:register_name"`
	RegisterEmail    string    `gorm:"type:varchar(255);column:register_email"`
	RegisterDate     time.Time `gorm:"column:register_date"`
	DeviceCode       string    `gorm:"type:varchar(255);column:dvn_device_code"`
	Model            int       `gorm:"column:dvn_model_cd"`
	Contents         string    `gorm:"type:text;column:dvn_contents"`
	Customer         int       `gorm:"column:dvn_customer_cd"`
	Manufacture      int       `gorm:"column:dvn_manufacture_cd"`
	DeviceType       int       `gorm:"column:dvn_device_type_cd"`
	WarehousingDate  string    `gorm:"type:varchar(10);column:dvn_warehousing_date"`
	RentDate         string    `gorm:"type:varchar(20);column:dvn_rent_date;default:'|'"`
	Ownership        string    `gorm:"type:varchar(10);column:dvn_ownership_cd;default:'|'"`
	OwnerCompany     string    `gorm:"type:varchar(255);column:dvn_owner_company"`
	HwSn             string    `gorm:"type:varchar(255);column:dvn_hw_sn"`
	IDC              int       `gorm:"column:dvn_idc_cd"`
	Rack             int       `gorm:"column:dvn_rack_cd"`
	Cost             string    `gorm:"type:varchar(255);column:dvn_cost"`
	Purpos           string    `gorm:"type:varchar(255);column:dvn_purpos"`
	Ip               string    `gorm:"type:varchar(255);column:dvn_ip;default:'|'"`
	Size             int       `gorm:"column:dvn_size_cd"`
	FirmwareVersion  string    `gorm:"type:varchar(50);column:dvn_firmware_version"`
	Warranty         string    `gorm:"type:varchar(255);column:dvn_warranty"`
	MonitoringFlag   int       `gorm:"column:dvn_mornitoring_flag"`
	MonitoringMethod int       `gorm:"column:dvn_mornitoring_method"`
}

func (DeviceNetwork) TableName() string {
	return "device_network_tb"
}

type DevicePart struct {
	Idx              uint      `gorm:"primary_key;column:dvp_idx;not null"`
	OutFlag          bool      `gorm:"type:tinyint(1);column:dvp_out_flag;default:0"`
	Num              int       `gorm:"column:dvp_num"`
	CommentCnt       int       `gorm:"column:dvp_comment_cnt"`
	CommentLastDate  time.Time `gorm:"column:dvp_comment_last_date"`
	Option           string    `gorm:"type:varchar(50);column:dvp_option"`
	Hit              int       `gorm:"column:dvp_hit"`
	RegisterId       string    `gorm:"type:varchar(50);column:register_id"`
	Password         string    `gorm:"type:varchar(255);column:register_password"`
	RegisterName     string    `gorm:"type:varchar(50);column:register_name"`
	RegisterEmail    string    `gorm:"type:varchar(255);column:register_email"`
	RegisterDate     time.Time `gorm:"column:register_date"`
	DeviceCode       string    `gorm:"type:varchar(255);column:dvp_device_code"`
	Model            int       `gorm:"column:dvp_model_cd"`
	Contents         string    `gorm:"type:text;column:dvp_contents"`
	Customer         int       `gorm:"column:dvp_customer_cd"`
	Manufacture      int       `gorm:"column:dvp_manufacture_cd"`
	DeviceType       int       `gorm:"column:dvp_device_type_cd"`
	WarehousingDate  string    `gorm:"type:varchar(10);column:dvp_warehousing_date"`
	RentDate         string    `gorm:"type:varchar(20);column:dvp_rent_date;default:'|'"`
	Ownership        string    `gorm:"type:varchar(10);column:dvp_ownership_cd;default:'|'"`
	OwnerCompany     string    `gorm:"type:varchar(255);column:dvp_owner_company"`
	HwSn             string    `gorm:"type:varchar(255);column:dvp_hw_sn"`
	IDC              int       `gorm:"column:dvp_idc_cd"`
	Rack             int       `gorm:"column:dvp_rack_cd"`
	Cost             string    `gorm:"type:varchar(255);column:dvp_cost"`
	Purpos           string    `gorm:"type:varchar(255);column:dvp_purpos"`
	Warranty         string    `gorm:"type:varchar(255);column:dvp_warranty"`
	MonitoringFlag   int       `gorm:"column:dvp_mornitoring_flag"`
	MonitoringMethod int       `gorm:"column:dvp_mornitoring_method"`
}

func (DevicePart) TableName() string {
	return "device_part_tb"
}

type DeviceComment struct {
	Idx          uint      `gorm:"primary_key;column:dvc_idx;not null"`
	ParentTable  string    `gorm:"column:dvc_parent_table;not null"`
	ForeignIdx   int       `gorm:"column:fk_idx;not null"`
	Depth        int       `gorm:"column:dvc_depth"`
	Contents     string    `gorm:"column:dvc_contents"`
	RegisterId   string    `gorm:"type:varchar(50);column:register_id"`
	RegisterName string    `gorm:"type:varchar(50);column:dvc_register_name"`
	RegisterDate time.Time `gorm:"column:register_date"`
}

func (DeviceComment) TableName() string {
	return "device_comment_tb"
}
