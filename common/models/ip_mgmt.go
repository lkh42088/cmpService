package models

import "time"

type IpMgmt struct {
	Idx				uint		`gorm:"primary_key;column:ip_idx;unsigned;auto_increment;comment:'INDEX'"`
	DeviceCode		string		`gorm:"type:varchar(12);not null;column:device_code;comment:'장비 코드'"`
	DevTag			string		`gorm:"type:varchar(50);not null;column:ip_dev_tag;comment:'DEVICE TAG'"`
	NetworkId		int			`gorm:"type:int(11);not null;column:net_id_cd;comment:'네트워크 코드'"`
	DevIp			string		`gorm:"type:varchar(15);not null;column:ip;comment:'IP'"`
	NatIp			string		`gorm:"type:varchar(15);column:ip_nat;comment:'NAT IP'"`
	DevDesc			string		`gorm:"type:varchar(30);column:ip_pdesc;comment:'DEVICE 상세'"`
	PortDesc		string		`gorm:"type:varchar(30);column:ip_port_desc;comment:'PORT 상세'"`
	DevMemo			string		`gorm:"type:text;column:ip_memo;comment:'IP MEMO'"`
	DevIpRegDate	time.Time	`gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:ip_reg_date;comment:'IP 등록일'"`
	DevCommunity	string		`gorm:"type:varchar(50);not null;column:ip_comm;comment:'COMMUNITY'"`
	DevUse			bool		`gorm:"type:tinyint(1);not null;default:1;column:ip_use_flag;comment:'사용 여부'"`
}

func (IpMgmt) TableName() string {
	return "ip_tb"
}

type SubnetMgmt struct {
	Idx				uint		`gorm:"primary_key;unsigned;auto_increment;column:sub_idx;comment:'INDEX'"`
	DeviceCode		string		`gorm:"type:varchar(12);not null;column:device_code;comment:'장비 코드'"`
	NetworkId		int			`gorm:"type:int(11);not null;column:net_id_cd;comment:'네트워크 코드'"`
	SubnetTag		string		`gorm:"type:varchar(255);not null;column:sub_tag;comment:'SUBNET TAG'"`
	SubnetStart		string		`gorm:"type:varchar(15);not null;column:sub_ip_start;comment:'SUBNET START'"`
	SubnetEnd		string		`gorm:"type:varchar(15);not null;column:sub_ip_end;comment:'SUBNET END'"`
	SubnetMask		string		`gorm:"type:varchar(15);not null;column:sub_mask;comment:'SUBNET MASK'"`
	Gateway			string		`gorm:"type:varchar(15);not null;column:sub_gateway;comment:'게이트웨이'"`
}

func (SubnetMgmt) TableName() string {
	return "subnet_tb"
}
