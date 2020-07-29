package models

import "strings"

const (
	McServerOrmIdx          = "mc_idx"
	McServerOrmSerialNumber = "mc_serial_number"
	mcServerOrmCompanyIdx   = "mc_cp_idx"
	mcServerOrmType         = "mc_type"
	mcServerOrmStatus       = "mc_status"
	mcServerOrmVmCount      = "mc_vm_count"
	mcServerOrmIpAddr       = "mc_ip_addr"
)

const (
	McServerJsonIdx          = "idx"
	McServerJsonSerialNumber = "serialNumber"
	mcServerJsonCompanyIdx   = "cpIdx"
	mcServerJsonType         = "type"
	mcServerJsonStatus       = "status"
	mcServerJsonVmCount      = "vmCount"
	mcServerJsonIpAddr       = "ipAddr"
)

var McServerOrmMap = map[string]string{
	McServerOrmIdx:          McServerJsonIdx,
	McServerOrmSerialNumber: McServerJsonSerialNumber,
	mcServerOrmCompanyIdx:   mcServerJsonCompanyIdx,
	mcServerOrmType:         mcServerJsonType,
	mcServerOrmStatus:       mcServerJsonStatus,
	mcServerOrmVmCount:      mcServerJsonVmCount,
	mcServerOrmIpAddr:       mcServerJsonIpAddr,
}

var McServerJsonMap = map[string]string{
	McServerJsonIdx:          McServerOrmIdx,
	McServerJsonSerialNumber: McServerOrmSerialNumber,
	mcServerJsonCompanyIdx:   mcServerOrmCompanyIdx,
	mcServerJsonType:         mcServerOrmType,
	mcServerJsonStatus:       mcServerOrmStatus,
	mcServerJsonVmCount:      mcServerOrmVmCount,
	mcServerJsonIpAddr:       mcServerOrmIpAddr,
}

var McVmOrmMap = map[string]string{
	"vm_idx":        "idx",
	"vm_server_idx": "serverIdx",
	"vm_cp_idx":     "cpIdx",
	"vm_name":       "name",
	"vm_cpu":        "cpu",
	"vm_ram":        "ram",
	"vm_hdd":        "hdd",
	"vm_os":         "os",
	"vm_network":    "network",
	"vm_ip_addr":    "ipAddr",
}

var McVmJsonMap = map[string]string{
	"idx":       "vm_idx",
	"serverIdx": "vm_server_idx",
	"cpIdx":     "vm_cp_idx",
	"name":      "vm_name",
	"cpu":       "vm_cpu",
	"ram":       "vm_ram",
	"hdd":       "vm_hdd",
	"os":        "vm_os",
	"network":   "vm_network",
	"ipAddr":    "vm_ip_addr",
}

type McServer struct {
	Idx          uint   `gorm:"primary_key;column:mc_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	SerialNumber string `gorm:"unique;type:varchar(50);column:mc_serial_number;comment:'시리얼넘버'" json:"serialNumber"`
	CompanyIdx   int    `gorm:"type:int(11);column:mc_cp_idx;comment:'회사 고유값'" json:"cpIdx"`
	Type         string `gorm:"type:varchar(50);column:mc_type;comment:'서버 타입'" json:"type"`
	Status       int    `gorm:"type:int(11);column:mc_status;comment:'서버 상태'" json:"status"`
	VmCount      int    `gorm:"type:int(11);column:mc_vm_count;comment:'vm 개수'" json:"vmCount"`
	IpAddr       string `gorm:"type:varchar(50);column:mc_ip_addr;comment:'IP Address'" json:"ipAddr"`
}

type McVm struct {
	Idx         uint   `gorm:"primary_key;column:vm_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	McServerIdx int    `gorm:"type:int(11);column:vm_server_idx;comment:'서버 고유값'" json:"serverIdx"`
	CompanyIdx  int    `gorm:"type:int(11);column:vm_cp_idx;comment:'회사 고유값'" json:"cpIdx"`
	Name        string `gorm:"type:varchar(50);column:vm_name;comment:'vm 이름'" json:"name"`
	Cpu         int    `gorm:"type:int(11);column:vm_cpu;comment:'vm cpu'" json:"cpu"`
	Ram         int    `gorm:"type:int(11);column:vm_ram;comment:'vm ram'" json:"ram"`
	Hdd         int    `gorm:"type:int(11);column:vm_hdd;comment:'vm hdd'" json:"hdd"`
	OS          string `gorm:"type:varchar(50);column:vm_os;comment:'vm os'" json:"os"`
	Network     string `gorm:"type:varchar(50);column:vm_network;comment:'vm network'" json:"network"`
	IpAddr      string `gorm:"type:varchar(50);column:vm_ip_addr;comment:'vm ip address'" json:"ipAddr"`
}

func (McServer) TableName() string {
	return "mc_server_tb"
}

func (McVm) TableName() string {
	return "mc_vm_tb"
}

type McServerDetail struct {
	McServer
	CompanyName string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
}

type McServerPage struct {
	Page    Pagination       `json:"page"`
	Servers []McServerDetail `json:"data"`
}

type McVmDetail struct {
	McVm
	CompanyName  string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
	SerialNumber string `gorm:"type:varchar(50);column:mc_serial_number" json:"serialNumber"`
}

type McVmPage struct {
	Page Pagination   `json:"page"`
	Vms  []McVmDetail `json:"data"`
}

func (m McServerPage) GetOrderBy(orderby, order string) string {
	val, exists := McServerJsonMap[orderby]
	if !exists {
		val = "mc_idx"
	}
	order = strings.ToLower(order)
	if !(order == "asc" || order == "desc") {
		order = "desc"
	}
	return val + " " + order
}

func (m McVmPage) GetOrderBy(orderby, order string) string {
	val, exists := McVmJsonMap[orderby]
	if !exists {
		val = "vm_idx"
	}
	order = strings.ToLower(order)
	if !(order == "asc" || order == "desc") {
		order = "desc"
	}
	return val + " " + order
}
