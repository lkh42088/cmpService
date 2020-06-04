package models

import "time"

type IpMgmt struct {
	Idx				uint		`gorm:"primary_key;column:ip_idx;unsigned;auto_increment;comment:'INDEX'" json:"idx"`
	DeviceCode		string		`gorm:"type:varchar(12);not null;column:device_code;comment:'장비 코드'" json:"deviceCode"`
	DevTag			string		`gorm:"type:varchar(50);not null;column:ip_dev_tag;comment:'DEVICE TAG'" json:"devTag"`
	SubnetMgmt		SubnetMgmt	`gorm:"foreignkey:SubIdx" json:"-"`
	SubIdx			uint		`gorm:"type:int(10);column:sub_idx;comment:'SUBNET INDEX'" json:"subIdx"`
	DevIp			string		`gorm:"type:varchar(15);not null;column:ip;comment:'IP'" json:"devIp"`
	NatIp			string		`gorm:"type:varchar(15);column:ip_nat;comment:'NAT IP'" json:"natIp"`
	DevDesc			string		`gorm:"type:varchar(30);column:ip_pdesc;comment:'DEVICE 상세'" json:"devDesc"`
	PortDesc		string		`gorm:"type:varchar(30);column:ip_port_desc;comment:'PORT 상세'" json:"portDesc"`
	DevMemo			string		`gorm:"type:text;column:ip_memo;comment:'IP MEMO'" json:"devMemo"`
	DevIpRegDate	time.Time	`gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:ip_reg_date;comment:'IP 등록일'" json:"devIpRegDate"`
	DevCommunity	string		`gorm:"type:varchar(50);not null;column:ip_comm;comment:'COMMUNITY'" json:"devCommunity"`
	DevUse			bool		`gorm:"type:tinyint(1);not null;default:1;column:ip_use_flag;comment:'사용 여부'" json:"devUse"`
}

func (IpMgmt) TableName() string {
	return "ip_tb"
}

type SubnetMgmt struct {
	Idx				uint		`gorm:"primary_key;unsigned;auto_increment;column:sub_idx;comment:'INDEX'" json:"idx"`
	DeviceCode		string		`gorm:"type:varchar(12);not null;column:device_code;comment:'장비 코드'" json:"deviceCode"`
	SubnetTag		string		`gorm:"type:varchar(255);not null;column:sub_tag;comment:'SUBNET TAG'" json:"subnetTag"`
	SubnetStart		string		`gorm:"type:varchar(15);not null;column:sub_ip_start;comment:'SUBNET START'" json:"subnetStart"`
	SubnetEnd		string		`gorm:"type:varchar(15);not null;column:sub_ip_end;comment:'SUBNET END'" json:"subnetEnd"`
	SubnetMask		string		`gorm:"type:varchar(15);not null;column:sub_mask;comment:'SUBNET MASK'" json:"subnetMask"`
	Gateway			string		`gorm:"type:varchar(15);not null;column:sub_gateway;comment:'게이트웨이'" json:"gateway"`
}

func (SubnetMgmt) TableName() string {
	return "subnet_tb"
}
