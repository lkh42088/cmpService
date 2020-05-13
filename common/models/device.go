package models

import (
	"time"
)

type DeviceServer struct {
	Idx              uint      `gorm:"primary_key;column:idx;not null;unsigned;auto_increment"`
	OutFlag          bool      `gorm:"type:tinyint(1);column:out_flag;default:0"`
	Num              int       `gorm:"column:num"`
	CommentCnt       int       `gorm:"column:comment_cnt"`
	CommentLastDate  time.Time `gorm:"column:comment_last_date"`
	Option           string    `gorm:"type:varchar(50);column:option"`
	Hit              int       `gorm:"column:hit"`
	RegisterId       string    `gorm:"type:varchar(50);column:register_id"`
	Password         string    `gorm:"type:varchar(255);column:register_password"`
	RegisterName     string    `gorm:"type:varchar(50);column:register_name"`
	RegisterEmail    string    `gorm:"type:varchar(255);column:register_email"`
	RegisterDate     time.Time `gorm:"column:register_date;default:CURRENT_TIMESTAMP"`
	DeviceCode       string    `gorm:"type:varchar(255);column:device_code"`
	Model            int       `gorm:"column:model_cd"`
	Contents         string    `gorm:"type:text;column:contents"`
	Customer         int       `gorm:"column:customer_cd"`
	Manufacture      int       `gorm:"column:manufacture_cd"`
	DeviceType       int       `gorm:"column:device_type_cd"`
	WarehousingDate  string    `gorm:"type:varchar(10);column:warehousing_date"`
	RentDate         string    `gorm:"type:varchar(20);column:rent_date;default:'|'"`
	Ownership        string    `gorm:"type:varchar(10);column:ownership_cd;default:'|'"`
	OwnerCompany     string    `gorm:"type:varchar(255);column:owner_company"`
	HwSn             string    `gorm:"type:varchar(255);column:hw_sn"`
	IDC              int       `gorm:"column:idc_cd"`
	Rack             int       `gorm:"column:rack_cd"`
	Cost             string    `gorm:"type:varchar(255);column:cost"`
	Purpos           string    `gorm:"type:varchar(255);column:purpos"`
	Ip               string    `gorm:"type:varchar(255);column:ip;default:'|'"`
	Size             int       `gorm:"column:size_cd"`
	Spla             string    `gorm:"column:spla_cd;default:'|'"`
	Cpu              string    `gorm:"type:varchar(255);column:cpu"`
	Memory           string    `gorm:"type:varchar(255);column:memory"`
	Hdd              string    `gorm:"type:varchar(255);column:hdd"`
	MonitoringFlag   int       `gorm:"column:monitoring_flag"`
	MonitoringMethod int       `gorm:"column:monitoring_method"`
}

func (DeviceServer) TableName() string {
	return "device_server_tb"
}

type DeviceNetwork struct {
	Idx              uint      `gorm:"primary_key;column:idx;not null;unsigned;auto_increment"`
	OutFlag          bool      `gorm:"type:tinyint(1);column:out_flag;default:0"`
	Num              int       `gorm:"column:num"`
	CommentCnt       int       `gorm:"column:comment_cnt"`
	CommentLastDate  time.Time `gorm:"column:comment_last_date"`
	Option           string    `gorm:"type:varchar(50);column:option"`
	Hit              int       `gorm:"column:hit"`
	RegisterId       string    `gorm:"type:varchar(50);column:register_id"`
	Password         string    `gorm:"type:varchar(255);column:register_password"`
	RegisterName     string    `gorm:"type:varchar(50);column:register_name"`
	RegisterEmail    string    `gorm:"type:varchar(255);column:register_email"`
	RegisterDate     time.Time `gorm:"column:register_date;default:CURRENT_TIMESTAMP"`
	DeviceCode       string    `gorm:"type:varchar(255);column:device_code"`
	Model            int       `gorm:"column:model_cd"`
	Contents         string    `gorm:"type:text;column:contents"`
	Customer         int       `gorm:"column:customer_cd"`
	Manufacture      int       `gorm:"column:manufacture_cd"`
	DeviceType       int       `gorm:"column:device_type_cd"`
	WarehousingDate  string    `gorm:"type:varchar(10);column:warehousing_date"`
	RentDate         string    `gorm:"type:varchar(20);column:rent_date;default:'|'"`
	Ownership        string    `gorm:"type:varchar(10);column:ownership_cd;default:'|'"`
	OwnerCompany     string    `gorm:"type:varchar(255);column:owner_company"`
	HwSn             string    `gorm:"type:varchar(255);column:hw_sn"`
	IDC              int       `gorm:"column:idc_cd"`
	Rack             int       `gorm:"column:rack_cd"`
	Cost             string    `gorm:"type:varchar(255);column:cost"`
	Purpos           string    `gorm:"type:varchar(255);column:purpos"`
	Ip               string    `gorm:"type:varchar(255);column:ip;default:'|'"`
	Size             int       `gorm:"column:size_cd"`
	FirmwareVersion  string    `gorm:"type:varchar(50);column:firmware_version"`
	MonitoringFlag   int       `gorm:"column:monitoring_flag"`
	MonitoringMethod int       `gorm:"column:monitoring_method"`
}

func (DeviceNetwork) TableName() string {
	return "device_network_tb"
}

type DevicePart struct {
	Idx              uint      `gorm:"primary_key;column:idx;not null;unsigned;auto_increment"`
	OutFlag          bool      `gorm:"type:tinyint(1);column:out_flag;default:0"`
	Num              int       `gorm:"column:num"`
	CommentCnt       int       `gorm:"column:comment_cnt"`
	CommentLastDate  time.Time `gorm:"column:comment_last_date"`
	Option           string    `gorm:"type:varchar(50);column:option"`
	Hit              int       `gorm:"column:hit"`
	RegisterId       string    `gorm:"type:varchar(50);column:register_id"`
	Password         string    `gorm:"type:varchar(255);column:register_password"`
	RegisterName     string    `gorm:"type:varchar(50);column:register_name"`
	RegisterEmail    string    `gorm:"type:varchar(255);column:register_email"`
	RegisterDate     time.Time `gorm:"column:register_date; default:CURRENT_TIMESTAMP"`
	DeviceCode       string    `gorm:"type:varchar(255);column:device_code"`
	Model            int       `gorm:"column:model_cd"`
	Contents         string    `gorm:"type:text;column:contents"`
	Customer         int       `gorm:"column:customer_cd"`
	Manufacture      int       `gorm:"column:manufacture_cd"`
	DeviceType       int       `gorm:"column:device_type_cd"`
	WarehousingDate  string    `gorm:"type:varchar(10);column:warehousing_date"`
	RentDate         string    `gorm:"type:varchar(20);column:rent_date;default:'|'"`
	Ownership        string    `gorm:"type:varchar(10);column:ownership_cd;default:'|'"`
	OwnerCompany     string    `gorm:"type:varchar(255);column:owner_company"`
	HwSn             string    `gorm:"type:varchar(255);column:hw_sn"`
	IDC              int       `gorm:"column:idc_cd"`
	Rack             int       `gorm:"column:rack_cd"`
	Cost             string    `gorm:"type:varchar(255);column:cost"`
	Purpos           string    `gorm:"type:varchar(255);column:purpos"`
	Warranty         string    `gorm:"type:varchar(255);column:warranty"`
	MonitoringFlag   int       `gorm:"column:monitoring_flag"`
	MonitoringMethod int       `gorm:"column:monitoring_method"`
}

func (DevicePart) TableName() string {
	return "device_part_tb"
}

type DeviceComment struct {
	Idx          uint      `gorm:"primary_key;column:idx;not null;unsigned;auto_increment"`
	ParentTable  string    `gorm:"column:parent_table;not null"`
	ForeignIdx   int       `gorm:"column:fk_idx;not null"`
	Depth        int       `gorm:"column:depth"`
	Contents     string    `gorm:"column:contents"`
	RegisterId   string    `gorm:"type:varchar(50);column:register_id"`
	RegisterName string    `gorm:"type:varchar(50);column:register_name"`
	RegisterDate time.Time `gorm:"column:register_date;default:CURRENT_TIMESTAMP"`
}

func (DeviceComment) TableName() string {
	return "device_comment_tb"
}

type PageCreteria struct {
	Count		int
	TotalPage	int
	CurPage		int
	Size 		int
	OutFlag		string
	OrderKey	string
	Direction	int
	DeviceType	string
}

type DeviceServerPage struct {
	Page			PageCreteria
	Devices 		[]DeviceServer
}

type DeviceNetworkPage struct {
	Page			PageCreteria
	Devices 		[]DeviceNetwork
}

type DevicePartPage struct {
	Page			PageCreteria
	Devices 		[]DevicePart
}


