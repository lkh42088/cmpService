package mcmodel

import (
	"cmpService/common/models"
	"encoding/json"
	"fmt"
	"strings"
)

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

type McServer struct {
	Idx          uint   `gorm:"primary_key;column:mc_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	SerialNumber string `gorm:"unique;type:varchar(50);column:mc_serial_number;comment:'시리얼넘버'" json:"serialNumber"`
	CompanyIdx   int    `gorm:"type:int(11);column:mc_cp_idx;comment:'회사 고유값'" json:"cpIdx"`
	Type         string `gorm:"type:varchar(50);column:mc_type;comment:'서버 타입'" json:"type"`
	Status       int    `gorm:"type:int(11);column:mc_status;comment:'서버 상태'" json:"status"`
	VmCount      int    `gorm:"type:int(11);column:mc_vm_count;comment:'vm 개수'" json:"vmCount"`
	Port         string `gorm:"type:varchar(50);column:mc_port;comment:'Port'" json:"port"`
	Mac          string `gorm:"type:varchar(50);column:mc_mac;comment:'Mac Address'" json:"mac"`
	IpAddr       string `gorm:"type:varchar(50);column:mc_ip_addr;comment:'IP Address'" json:"ip"`
	PublicIpAddr string `gorm:"type:varchar(50);column:mc_public_ip_addr;comment:'Public IP Address'" json:"publicIp"`
}

func (McServer) TableName() string {
	return "mc_server_tb"
}

type McServerDetail struct {
	McServer
	CompanyName string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
}

type McServerPage struct {
	Page    models.Pagination `json:"page"`
	Servers []McServerDetail  `json:"data"`
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

var McNetworkJsonMap = map[string]string{
	"idx":       "net_idx",
	"serverIdx": "net_server_idx",
	"name":      "net_name",
	"bridge":    "net_bridge",
	"mode":      "net_mode",
	"ip":        "net_ip",
	"netmask":   "net_netmask",
	"prefix":    "net_prefix",
}

type McNetHost struct {
	Idx          uint   `gorm:"primary_key;column:nh_idx" json:"idx"`
	McNetworkIdx int    `gorm:"type:int(11);column:nh_net_idx" json:"networkIdx"`
	Mac          string `gorm:"type:varchar(50);column:nh_mac" json:"mac"`
	Ip           string `gorm:"type:varchar(50);column:nh_ip" json:"ip"`
	Hostname     string `gorm:"type:varchar(50);column:nh_hostname" json:"hostname"`
}

func (McNetHost) TableName() string {
	return "mc_net_host"
}

type McNetworks struct {
	Idx         uint         `gorm:"primary_key;column:net_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	McServerIdx int          `gorm:"type:int(11);column:net_server_idx;comment:'서버 고유값'" json:"serverIdx"`
	Name        string       `gorm:"type:varchar(50);column:net_name;comment:'network 이름'" json:"name"`
	Bridge      string       `gorm:"type:varchar(50);column:net_bridge;comment:'bridge name'" json:"bridge"`
	Mode        string       `gorm:"type:varchar(50);column:net_mode;comment:'forward mode'" json:"mode"`
	Ip          string       `gorm:"type:varchar(50);column:net_ip;comment:'ip address'" json:"ip"`
	Netmask     string       `gorm:"type:varchar(50);column:net_netmask;comment:'netmask'" json:"netmask"`
	Prefix      uint         `gorm:"type:int(11);column:net_prefix;comment:'prefix'" json:"prefix"`
	Host        *[]McNetHost `gorm:"-" json:"host"`
}

func (McNetworks) TableName() string {
	return "mc_network_tb"
}

type McNetworkDetail struct {
	McNetworks
	CompanyName  string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
	SerialNumber string `gorm:"unique;type:varchar(50);column:mc_serial_number;comment:'시리얼넘버'" json:"serialNumber"`
}

type McNetworkPage struct {
	Page     models.Pagination `json:"page"`
	Networks []McNetworkDetail `json:"data"`
}

func (m McNetworkPage) GetOrderBy(orderby, order string) string {
	val, exists := McNetworkJsonMap[orderby]
	if !exists {
		val = "net_idx"
	}
	order = strings.ToLower(order)
	if !(order == "asc" || order == "desc") {
		order = "desc"
	}
	return val + " " + order
}

var McImageJsonMap = map[string]string{
	"idx":       "img_idx",
	"serverIdx": "img_server_idx",
	"variant":   "img_variant",
	"name":      "img_name",
	"hdd":       "img_hdd",
}

type McImages struct {
	Idx         uint   `gorm:"primary_key;column:img_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	McServerIdx int    `gorm:"type:int(11);column:img_server_idx;comment:'서버 고유값'" json:"serverIdx"`
	Variant     string `gorm:"type:varchar(50);column:img_variant;comment:'이미지 타입'" json:"variant"`
	Name        string `gorm:"type:varchar(50);column:img_name;comment:'이미지 이름'" json:"name"`
	Hdd         int    `gorm:"type:int(11);column:img_hdd;comment:'image size'" json:"hdd"`
}

func (McImages) TableName() string {
	return "mc_image_tb"
}

type McImageDetail struct {
	McImages
	CompanyName  string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
	SerialNumber string `gorm:"unique;type:varchar(50);column:mc_serial_number;comment:'시리얼넘버'" json:"serialNumber"`
}

type McImagePage struct {
	Page   models.Pagination `json:"page"`
	Images []McImageDetail   `json:"data"`
}

func (m McImagePage) GetOrderBy(orderby, order string) string {
	val, exists := McImageJsonMap[orderby]
	if !exists {
		val = "img_idx"
	}
	order = strings.ToLower(order)
	if !(order == "asc" || order == "desc") {
		order = "desc"
	}
	return val + " " + order
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
	"idx":           "vm_idx",
	"serverIdx":     "vm_server_idx",
	"cpIdx":         "vm_cp_idx",
	"name":          "vm_name",
	"cpu":           "vm_cpu",
	"ram":           "vm_ram",
	"hdd":           "vm_hdd",
	"os":            "vm_os",
	"image":         "vm_image",
	"network":       "vm_network",
	"ipAddr":        "vm_ip_addr",
	"mac":           "vm_mac",
	"currentStatus": "vm_current_status",
	"remoteAddr":    "vm_remote_addr",
}

type McVm struct {
	Idx           uint   `gorm:"primary_key;column:vm_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	McServerIdx   int    `gorm:"type:int(11);column:vm_server_idx;comment:'서버 고유값'" json:"serverIdx"`
	CompanyIdx    int    `gorm:"type:int(11);column:vm_cp_idx;comment:'회사 고유값'" json:"cpIdx"`
	Name          string `gorm:"type:varchar(50);column:vm_name;comment:'vm 이름'" json:"name"`
	Cpu           int    `gorm:"type:int(11);column:vm_cpu;comment:'vm cpu'" json:"cpu"`
	Ram           int    `gorm:"type:int(11);column:vm_ram;comment:'vm ram'" json:"ram"`
	Hdd           int    `gorm:"type:int(11);column:vm_hdd;comment:'vm hdd'" json:"hdd"`
	Desc          string `gorm:"type:varchar(100);column:vm_desc;comment:'vm description'" json:"desc"`
	OS            string `gorm:"type:varchar(50);column:vm_os;comment:'vm os'" json:"os"`
	Image         string `gorm:"type:varchar(50);column:vm_image;comment:'vm image'" json:"image"`
	Filename      string `gorm:"type:varchar(50);column:vm_filename;comment:'vm image'" json:"filename"`
	VmIndex       int    `gorm:"type:int(11);column:vm_vmIndex;comment:'vm index'" json:"vmIndex"`
	FullPath      string `gorm:"type:varchar(50);column:vm_full_path;comment:'file full path'" json:"fullPath"`
	Network       string `gorm:"type:varchar(50);column:vm_network;comment:'vm network'" json:"network"`
	IpAddr        string `gorm:"type:varchar(50);column:vm_ip_addr;comment:'vm ip address'" json:"ipAddr"`
	RemoteAddr    string `gorm:"type:varchar(50);column:vm_remote_addr;comment:'Remote Address for RDP'" json:"remoteAddr"`
	Mac           string `gorm:"type:varchar(50);column:vm_mac;comment:'vm mac address'" json:"mac"`
	ConfigStatus  string `gorm:"type:varchar(50);column:vm_config_status;comment:'vm config status'" json:"configStatus"`
	CurrentStatus string `gorm:"type:varchar(50);column:vm_current_status;comment:'vm current status'" json:"currentStatus"`
}

func (McVm) TableName() string {
	return "mc_vm_tb"
}

type McVmDetail struct {
	McVm
	CompanyName  string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
	SerialNumber string `gorm:"type:varchar(50);column:mc_serial_number" json:"serialNumber"`
}

type McVmPage struct {
	Page models.Pagination `json:"page"`
	Vms  []McVmDetail      `json:"data"`
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

type McServerMsg struct {
	SerialNumber string `json:"serialNumber"`
	Port         string `json:"port"`
	Mac          string `json:"mac"`
	IpAddr       string `json:"ip"`
	PublicIpAddr string `json:"publicIp"`
	Vms          *[]McVm
	Networks     *[]McNetworks
	Images       *[]McImages
}

func (s *McServerMsg) Dump() string {
	pretty, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}
