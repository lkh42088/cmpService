package models

import (
	"time"
)


/////
// DEVICE TABLE
/////
type DeviceCommon struct {
	Idx              uint      `gorm:"primary_key;column:device_idx;not null;unsigned;auto_increment;comment:'INDEX'"`
	OutFlag          bool      `gorm:"type:tinyint(1);column:out_flag;default:0;comment:'반출 여부'"`
	CommentCnt       int       `gorm:"type:int(11);column:comment_cnt;comment:'Comment 개수'"`
	CommentLastDate  time.Time `gorm:"type:datetime;column:comment_last_date;comment:'마지막 Comment 등록'"`
	RegisterId       string    `gorm:"type:varchar(50);column:register_id;comment:'등록자 ID'"`
	RegisterDate     time.Time `gorm:"type:datetime;column:register_date;default:CURRENT_TIMESTAMP;comment:'등록일'"`
	DeviceCode       string    `gorm:"unique;type:varchar(12);column:device_code;comment:'장비 코드'"`
	Model            int       `gorm:"type:int(11);column:model_cd;comment:'모델 코드'"`
	Contents         string    `gorm:"type:text;column:contents;comment:'내용'"`
	Customer         string    `gorm:"type:varchar(50);column:user_id;comment:'고객사명'"`
	Manufacture      int       `gorm:"type_int(11);column:manufacture_cd;comment:'제조사'"`
	DeviceType       int       `gorm:"type:int(11);column:device_type_cd;comment:'장비 구분'"`
	WarehousingDate  string    `gorm:"type:varchar(10);column:warehousing_date;comment:'입고일'"`
	RentDate         string    `gorm:"type:varchar(20);column:rent_date;default:'|';comment:'임대 기간'"`
	Ownership	     string    `gorm:"type:varchar(10);column:ownership_cd;comment:'소유권'"`
	OwnershipDiv     string    `gorm:"type:varchar(10);column:ownership_div_cd;comment:'소유 구분'"`
	OwnerCompany     string    `gorm:"type:varchar(255);column:owner_company;comment:'소유 회사'"`
	HwSn             string    `gorm:"type:varchar(255);column:hw_sn;comment:'하드웨어 S/N'"`
	IDC              int       `gorm:"type:int(11);column:idc_cd;comment:'IDC'"`
	Rack             int       `gorm:"type:int(11);column:rack_cd;comment:'Rack'"`
	Cost             string    `gorm:"type:varchar(255);column:cost;comment:'장비 원가'"`
	Purpos           string    `gorm:"type:varchar(255);column:purpos;comment:'장비 용도'"`
	MonitoringFlag   int       `gorm:"type:int(11);column:monitoring_flag;comment:'모니터링 여부'"`
	MonitoringMethod int       `gorm:"type:int(11);column:monitoring_method;comment:'모니터링 방식'"`
}

type DeviceServer struct {
	DeviceCommon
	Ip               string    `gorm:"type:varchar(255);column:ip;default:'|';comment:'IP'"`
	Size             int       `gorm:"column:size_cd;comment:'크기'"`
	Spla             string    `gorm:"column:spla_cd;default:'|';comment:'SPLA'"`
	Cpu              string    `gorm:"type:varchar(255);column:cpu;comment:'CPU'"`
	Memory           string    `gorm:"type:varchar(255);column:memory;comment:'MEMORY'"`
	Hdd              string    `gorm:"type:varchar(255);column:hdd;comment:'HDD'"`
	RackCode		 int	   `gorm:"type:int(11);column:rack_code_cd;comment:'Rack 코드'"`
	RackTag			 string	   `gorm:"type:varchar(255);column:rack_tag;comment:'Rack 태그'"`
	RackLoc			 int	   `gorm:"type:int(11);column:rack_loc;comment:'Rack 내 위치 번호'"`
}

func (DeviceServer) TableName() string {
	return "device_server_tb"
}

type DeviceNetwork struct {
	DeviceCommon
	Ip               string    `gorm:"type:varchar(255);column:ip;default:'|';comment:'IP'"`
	Size             int       `gorm:"column:size_cd;comment:'크기'"`
	FirmwareVersion  string    `gorm:"type:varchar(50);column:firmware_version;comment:'펌웨어 버전'"`
	RackCode		 int	   `gorm:"type:int(11);column:rack_code_cd;comment:'Rack 코드'"`
	RackTag			 string	   `gorm:"type:varchar(255);column:rack_tag;comment:'Rack 태그'"`
	RackLoc			 int	   `gorm:"type:int(11);column:rack_loc;comment:'Rack 내 위치 번호'"`
}

func (DeviceNetwork) TableName() string {
	return "device_network_tb"
}

type DevicePart struct {
	DeviceCommon
	Warranty         string    `gorm:"type:varchar(255);column:warranty;comment:'WARRANTY'"`
}

func (DevicePart) TableName() string {
	return "device_part_tb"
}

type DeviceCommonResponse struct {
	Idx              uint      `gorm:"primary_key;column:device_idx;not null;unsigned;auto_increment"`
	OutFlag          bool      `gorm:"type:tinyint(1);column:out_flag;default:0"`
	CommentCnt       int       `gorm:"type:int(11);column:comment_cnt;comment"`
	CommentLastDate  time.Time `gorm:"type:datetime;column:comment_last_date"`
	RegisterId       string    `gorm:"type:varchar(50);column:register_id"`
	RegisterDate     time.Time `gorm:"type:datetime;column:register_date;default:CURRENT_TIMESTAMP"`
	DeviceCode       string    `gorm:"unique;type:varchar(12);column:device_code"`
	Model            string    `gorm:"column:model_cd"`
	Contents         string    `gorm:"type:text;column:contents"`
	Customer         string    `gorm:"column:user_id"`
	Manufacture      string    `gorm:"column:manufacture_cd"`
	DeviceType       string    `gorm:"column:device_type_cd"`
	WarehousingDate  string    `gorm:"type:varchar(10);column:warehousing_date"`
	RentDate         string    `gorm:"type:varchar(20);column:rent_date;default:'|'"`
	Ownership	     string    `gorm:"type:varchar(10);column:ownership_cd"`
	OwnershipDiv     string    `gorm:"type:varchar(10);column:ownership_div_cd"`
	OwnerCompany     string    `gorm:"type:varchar(255);column:owner_company"`
	HwSn             string    `gorm:"type:varchar(255);column:hw_sn"`
	IDC              string    `gorm:"column:idc_cd"`
	Rack             string    `gorm:"column:rack_cd"`
	Cost             string    `gorm:"type:varchar(255);column:cost"`
	Purpos           string    `gorm:"type:varchar(255);column:purpos"`
	MonitoringFlag   int       `gorm:"type:tinyint;column:monitoring_flag"`
	MonitoringMethod int       `gorm:"type:int(11);column:monitoring_method"`
}

type DeviceServerResponse struct {
	DeviceCommonResponse
	Ip               string    `gorm:"type:varchar(255);column:ip;default:'|'"`
	Size             string    `gorm:"column:size_cd"`
	Spla             string    `gorm:"column:spla_cd;default:'|'"`
	Cpu              string    `gorm:"type:varchar(255);column:cpu"`
	Memory           string    `gorm:"type:varchar(255);column:memory"`
	Hdd              string    `gorm:"type:varchar(255);column:hdd"`
	RackCode		 int	   `gorm:"type:int(11);column:rack_code_cd"`
	RackTag			 string	   `gorm:"type:varchar(255);column:rack_tag"`
	RackLoc			 int	   `gorm:"type:int(11);column:rack_loc"`
}

type DeviceNetworkResponse struct {
	DeviceCommonResponse
	Ip               string    `gorm:"type:varchar(255);column:ip;default:'|'"`
	Size             string    `gorm:"column:size_cd"`
	FirmwareVersion  string    `gorm:"type:varchar(50);column:firmware_version"`
}

type DevicePartResponse struct {
	DeviceCommonResponse
	Warranty         string    `gorm:"type:varchar(255);column:warranty"`
}

/////
// COMMENT TABLE
/////
type DeviceComment struct {
	Idx          uint      `gorm:"primary_key;column:comment_idx;not null;unsigned;auto_increment;comment:'INDEX'"`
	DeviceCode   string    `gorm:"type:varchar(12);column:device_code;not null;comment:'장비 코드'"`
	Contents     string    `gorm:"type:text;column:comment_contents;comment:'내용'"`
	RegisterId   string    `gorm:"type:varchar(50);column:comment_register_id;comment:'작성자 ID'"`
	RegisterName string    `gorm:"type:varchar(50);column:comment_register_name;comment:'작성자 이름'"`
	RegisterDate time.Time `gorm:"column:comment_register_date;default:CURRENT_TIMESTAMP;comment:'작성일'"`
}

func (DeviceComment) TableName() string {
	return "device_comment_tb"
}


/////
// Management LOG TABLE
/////
type DeviceLog struct {
	Idx          uint      `gorm:"primary_key;column:log_idx;not null;unsigned;auto_increment;comment:'INDEX'"`
	DeviceCode   string    `gorm:"column:device_code;not nul;comment:'장비 코드'"`
	WorkCode     int	   `gorm:"column:log_work_code;not null;comment:'작업 코드'"`
	Field 		 string    `gorm:"column:log_field;comment:'변경 필드'"`
	OldStatus    string    `gorm:"column:log_old_status;comment:'이전 상태'"`
	NewStatus    string	   `gorm:"column:log_new_status;comment:'변경 상태'"`
	LogLevel	 int	   `gorm:"type:int(11);not null;column:log_level_cd;comment:'로그 레벨'"`
	RegisterDate time.Time `gorm:"column:log_register_date;default:CURRENT_TIMESTAMP;comment:'로그 발생일'"`
}

func (DeviceLog) TableName() string {
	return "device_mgmt_log_tb"
}

type PageCreteria struct {
	Count		int
	TotalPage	int
	CheckCnt	int
	Size 		int
	OutFlag		string
	OrderKey	string
	Direction	int
	DeviceType	string
}

type DeviceServerPage struct {
	Page			PageCreteria
	Devices 		[]DeviceServerResponse
}

type DeviceNetworkPage struct {
	Page			PageCreteria
	Devices 		[]DeviceNetworkResponse
}

type DevicePartPage struct {
	Page			PageCreteria
	Devices 		[]DevicePartResponse
}


