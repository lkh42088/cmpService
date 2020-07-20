package models

import (
	"encoding/json"
	"fmt"
	"time"
)

/////
// DEVICE TABLE
/////
type DeviceCommon struct {
	Idx              uint      `gorm:"primary_key;column:device_idx;not null;unsigned;auto_increment;comment:'INDEX'" json:"idx,omitempty"`
	OutFlag          bool      `gorm:"type:tinyint(1);column:out_flag;default:0;comment:'반출 여부'" json:"outFlag"`
	CommentCnt       int       `gorm:"type:int(11);column:comment_cnt;comment:'Comment 개수'" json:"commentCnt"`
	CommentLastDate  time.Time `gorm:"type:datetime;column:comment_last_date;comment:'마지막 Comment 등록'" json:"-"`
	RegisterId       string    `gorm:"type:varchar(50);column:register_id;comment:'등록자 ID'" json:"registerId,omitempty"`
	RegisterDate     time.Time `gorm:"type:datetime;column:register_date;default:CURRENT_TIMESTAMP;comment:'등록일'" json:"-"`
	DeviceCode       string    `gorm:"unique;type:varchar(12);column:device_code;comment:'장비 코드'" json:"deviceCode"`
	Model            int       `gorm:"type:int(11);column:model_cd;comment:'모델 코드'" json:"model"`
	Contents         string    `gorm:"type:text;column:contents;comment:'내용'" json:"contents,omitempty"`
	Customer         string    `gorm:"type:varchar(50);column:user_id;comment:'고객사명'" json:"customer,omitempty"`
	Manufacture      int       `gorm:"type_int(11);column:manufacture_cd;comment:'제조사'" json:"manufacture"`
	DeviceType       int       `gorm:"type:int(11);column:device_type_cd;comment:'장비 구분'" json:"deviceType"`
	WarehousingDate  string    `gorm:"type:varchar(10);column:warehousing_date;comment:'입고일'" json:"warehousingDate,omitempty"`
	RentDate         string    `gorm:"type:varchar(20);column:rent_date;default:'|';comment:'임대 기간'" json:"rentDate,omitempty"`
	Ownership        string    `gorm:"type:varchar(10);column:ownership_cd;comment:'소유권'" json:"ownership,omitempty"`
	OwnershipDiv     string    `gorm:"type:varchar(10);column:ownership_div_cd;comment:'소유 구분'" json:"ownershipDiv,omitempty"`
	OwnerCompany     string    `gorm:"type:varchar(255);column:owner_company;comment:'소유 회사'" json:"ownerCompany,omitempty"`
	HwSn             string    `gorm:"type:varchar(255);column:hw_sn;comment:'하드웨어 S/N'" json:"hwSn,omitempty"`
	IDC              int       `gorm:"type:int(11);column:idc_cd;comment:'IDC'" json:"idc"`
	Rack             int       `gorm:"type:int(11);column:rack_cd;comment:'Rack'" json:"rack"`
	Cost             string    `gorm:"type:varchar(255);column:cost;comment:'장비 원가'" json:"cost"`
	Purpose          string    `gorm:"type:varchar(255);column:purpose;comment:'장비 용도'" json:"purpose,omitempty"`
	MonitoringFlag   bool      `gorm:"type:tinyint(1);column:monitoring_flag;comment:'모니터링 여부'" json:"monitoringFlag"`
	MonitoringMethod int       `gorm:"type:int(11);column:monitoring_method;comment:'모니터링 방식'" json:"monitoringMethod"`
}

type DeviceServer struct {
	DeviceCommon
	Ip      string `gorm:"type:varchar(255);column:ip;default:'|';comment:'IP'" json:"ip,omitempty"`
	Size    int    `gorm:"column:size_cd;comment:'크기'" json:"size"`
	Spla    string `gorm:"column:spla_cd;default:'|';comment:'SPLA'" json:"spla,omitempty"`
	Cpu     string `gorm:"type:varchar(255);column:cpu;comment:'CPU'" json:"cpu,omitempty"`
	Memory  string `gorm:"type:varchar(255);column:memory;comment:'MEMORY'" json:"memory,omitempty"`
	Hdd     string `gorm:"type:varchar(255);column:hdd;comment:'HDD'" json:"hdd,omitempty"`
	RackTag string `gorm:"type:varchar(255);column:rack_tag;comment:'Rack 태그'" json:"rackTag,omitempty"`
	RackLoc int    `gorm:"type:int(11);column:rack_loc;comment:'Rack 내 위치 번호'" json:"rackLoc"`
}

func (DeviceServer) TableName() string {
	return "device_server_tb"
}

type DeviceNetwork struct {
	DeviceCommon
	Ip              string `gorm:"type:varchar(255);column:ip;default:'|';comment:'IP'" json:"ip,omitempty"`
	Size            int    `gorm:"column:size_cd;comment:'크기'" json:"size"`
	FirmwareVersion string `gorm:"type:varchar(50);column:firmware_version;comment:'펌웨어 버전'" json:"firmwareVersion,omitempty"`
	RackTag         string `gorm:"type:varchar(255);column:rack_tag;comment:'Rack 태그'" json:"rackTag,omitempty"`
	RackLoc         int    `gorm:"type:int(11);column:rack_loc;comment:'Rack 내 위치 번호'" json:"rackLoc"`
}

func (DeviceNetwork) TableName() string {
	return "device_network_tb"
}

type DevicePart struct {
	DeviceCommon
	Warranty string `gorm:"type:varchar(255);column:warranty;comment:'WARRANTY'" json:"warranty,omitempty"`
	RackCode int    `gorm:"type:int(11);column:rack_code_cd;comment:'Rack 사이즈 코드'" json:"rackCode"`
}

func (DevicePart) TableName() string {
	return "device_part_tb"
}

type DeviceCommonResponse struct {
	Idx              uint      `gorm:"primary_key;column:device_idx;not null;unsigned;auto_increment" json:"idx"`
	OutFlag          bool      `gorm:"type:tinyint(1);column:out_flag;default:0" json:"outFlag"`
	CommentCnt       int       `gorm:"type:int(11);column:comment_cnt;comment" json:"commentCnt"`
	CommentLastDate  time.Time `gorm:"type:datetime;column:comment_last_date" json:"commentLastDate"`
	RegisterId       string    `gorm:"type:varchar(50);column:register_id" json:"registerId"`
	RegisterDate     time.Time `gorm:"type:datetime;column:register_date;default:CURRENT_TIMESTAMP" json:"registerDate"`
	DeviceCode       string    `gorm:"unique;type:varchar(12);column:device_code" json:"deviceCode"`
	Model            string    `gorm:"column:model_cd" json:"model"`
	Contents         string    `gorm:"type:text;column:contents" json:"contents"`
	Customer         string    `gorm:"column:user_id" json:"customer"`
	Manufacture      string    `gorm:"column:manufacture_cd" json:"manufacture"`
	DeviceType       string    `gorm:"column:device_type_cd" json:"deviceType"`
	WarehousingDate  string    `gorm:"type:varchar(10);column:warehousing_date" json:"warehousingDate"`
	RentDate         string    `gorm:"type:varchar(20);column:rent_date;default:'|'" json:"rentDate"`
	Ownership        string    `gorm:"type:varchar(10);column:ownership_cd" json:"ownership"`
	OwnershipDiv     string    `gorm:"type:varchar(10);column:ownership_div_cd" json:"ownershipDiv"`
	OwnerCompany     string    `gorm:"type:varchar(255);column:owner_company" json:"ownerCompany"`
	HwSn             string    `gorm:"type:varchar(255);column:hw_sn" json:"hwSn"`
	IDC              string    `gorm:"column:idc_cd" json:"idc"`
	Rack             string    `gorm:"column:rack_cd" json:"rack"`
	Cost             string    `gorm:"type:varchar(255);column:cost" json:"cost"`
	Purpose          string    `gorm:"type:varchar(255);column:purpose" json:"purpose"`
	MonitoringFlag   bool      `gorm:"type:tinyint(1);column:monitoring_flag" json:"monitoringFlag"`
	MonitoringMethod int       `gorm:"type:int(11);column:monitoring_method" json:"monitoringMethod"`
}

type DeviceServerResponse struct {
	DeviceCommonResponse
	Ip           string `gorm:"type:varchar(255);column:ip;default:'|'" json:"ip"`
	Size         string `gorm:"column:size_cd" json:"size"`
	Spla         string `gorm:"column:spla_cd;default:'|'" json:"spla"`
	Cpu          string `gorm:"type:varchar(255);column:cpu" json:"cpu"`
	Memory       string `gorm:"type:varchar(255);column:memory" json:"memory"`
	Hdd          string `gorm:"type:varchar(255);column:hdd" json:"hdd"`
	RackTag      string `gorm:"type:varchar(255);column:rack_tag" json:"rackTag"`
	RackLoc      int    `gorm:"type:int(11);column:rack_loc" json:"rackLoc"`
	Customer     string `gorm:"column:company_name" json:"customerName"`
	OwnerCompany string `gorm:"type:varchar(255);column:owner_company_name" json:"ownerCompanyName"`
}

type DeviceNetworkResponse struct {
	DeviceCommonResponse
	Ip              string `gorm:"type:varchar(255);column:ip;default:'|'" json:"ip"`
	Size            string `gorm:"column:size_cd" json:"size"`
	FirmwareVersion string `gorm:"type:varchar(50);column:firmware_version" json:"firmwareVersion"`
	RackTag         string `gorm:"type:varchar(255);column:rack_tag" json:"rackTag"`
	RackLoc         int    `gorm:"type:int(11);column:rack_loc" json:"rackLoc"`
	Customer        string `gorm:"column:company_name" json:"customerName"`
	OwnerCompany    string `gorm:"type:varchar(255);column:owner_company_name" json:"ownerCompanyName"`
}

type DevicePartResponse struct {
	DeviceCommonResponse
	Warranty     string `gorm:"type:varchar(255);column:warranty" json:"warranty"`
	RackCode     int    `gorm:"type:int(11);column:rack_code_cd" json:"rackCode"`
	Customer     string `gorm:"column:company_name" json:"customerName"`
	OwnerCompany string `gorm:"type:varchar(255);column:owner_company_name" json:"ownerCompanyName"`
}

/////
// COMMENT TABLE
/////
type DeviceComment struct {
	Idx          uint      `gorm:"primary_key;column:comment_idx;not null;unsigned;auto_increment;comment:'INDEX'" json:"idx"`
	DeviceCode   string    `gorm:"type:varchar(12);column:device_code;not null;comment:'장비 코드'" json:"deviceCode"`
	Contents     string    `gorm:"type:text;column:comment_contents;comment:'내용'" json:"contents"`
	RegisterId   string    `gorm:"type:varchar(50);column:comment_register_id;comment:'작성자 ID'" json:"registerId"`
	RegisterName string    `gorm:"type:varchar(50);column:comment_register_name;comment:'작성자 이름'" json:"registerName"`
	RegisterDate time.Time `gorm:"column:comment_register_date;default:CURRENT_TIMESTAMP;comment:'작성일'" json:"registerDate"`
}

func (DeviceComment) TableName() string {
	return "device_comment_tb"
}

/////
// Management LOG TABLE
/////
type DeviceLog struct {
	Idx          uint      `gorm:"primary_key;column:log_idx;not null;unsigned;auto_increment;comment:'INDEX'" json:"idx"`
	DeviceCode   string    `gorm:"type:varchar(12);column:device_code;not nul;comment:'장비 코드'" json:"deviceCode"`
	WorkCode     int       `gorm:"type:int(11);column:log_work_cd;not null;comment:'작업 코드'" json:"workCode"`
	Field        string    `gorm:"type:varchar(50);column:log_field;comment:'변경 필드'" json:"field"`
	OldStatus    string    `gorm:"type:varchar(255);column:log_old_status;comment:'이전 상태'" json:"oldStatus"`
	NewStatus    string    `gorm:"type:varchar(255);column:log_new_status;comment:'변경 상태'" json:"newStatus"`
	LogLevel     int       `gorm:"type:int(11);not null;column:log_level_cd;comment:'로그 레벨'" json:"logLevel"`
	RegisterId   string    `gorm:"type:varchar(50);column:log_register_id;comment:'로그 등록자 ID'" json:"registerId"`
	RegisterName string    `gorm:"type:varchar(50);column:log_register_name;comment:'로그 등록자 이름'" json:"registerName"`
	RegisterDate time.Time `gorm:"column:log_register_date;default:CURRENT_TIMESTAMP;comment:'로그 발생일'" json:"registerDate"`
}

func (DeviceLog) TableName() string {
	return "device_mgmt_log_tb"
}

type PageCreteria struct {
	Count          int    `json:"count"`     // 전체 row 개수 in DB
	TotalPage      int    `json:"totalPage"` // 전체 페이지
	CheckCnt       int    `json:"checkCnt"`  // Current row counter (offset)
	Size           int    `json:"size"`      // limit
	OutFlag        string `json:"outFlag"`   // 0: 반입, 1: 반출
	OrderKey       string `json:"orderKey"`  // order field
	Direction      int    `json:"direction"` // order : asc, desc
	DeviceType     string `json:"deviceType"`
	Row            int    `json:"row"`
	Page           int    `json:"page"`
	OffsetPage     int    `json:"offsetPage"`
	RentPeriodFlag string `json:"rentPeriod"`
}

func (p *PageCreteria) String() {
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("%s\n", data)
}

type DeviceServerPage struct {
	Page    PageCreteria
	Devices []DeviceServerResponse
}

type DeviceNetworkPage struct {
	Page    PageCreteria
	Devices []DeviceNetworkResponse
}

type DevicePartPage struct {
	Page    PageCreteria
	Devices []DevicePartResponse
}
