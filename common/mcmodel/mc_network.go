package mcmodel

import (
	"cmpService/common/models"
	"encoding/json"
	"fmt"
	"strings"
)

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
	Uuid        string       `gorm:"-" json:"-"`
	Name        string       `gorm:"type:varchar(50);column:net_name;comment:'network 이름'" json:"name"`
	Bridge      string       `gorm:"type:varchar(50);column:net_bridge;comment:'bridge name'" json:"bridge"`
	Mode        string       `gorm:"type:varchar(50);column:net_mode;comment:'forward mode'" json:"mode"`
	Mac         string       `gorm:"type:varchar(50);column:net_mac;comment:'forward mode'" json:"mac"`
	DhcpStart   string       `gorm:"type:varchar(50);column:net_dhcp_start;comment:'dhcp start'" json:"dhcpStart"`
	DhcpEnd     string       `gorm:"type:varchar(50);column:net_dhcp_end;comment:'dhcp end'" json:"dhcpEnd"`
	Ip          string       `gorm:"type:varchar(50);column:net_ip;comment:'ip address'" json:"ip"`
	Netmask     string       `gorm:"type:varchar(50);column:net_netmask;comment:'netmask'" json:"netmask"`
	Prefix      uint         `gorm:"type:int(11);column:net_prefix;comment:'prefix'" json:"prefix"`
	Host        []McNetHost  `gorm:"-" json:"host"`
}

func (n *McNetworks) Dump() string {
	pretty, _ := json.MarshalIndent(n, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func (McNetworks) TableName() string {
	return "mc_network_tb"
}

type McNetworkDetail struct {
	McNetworks
	CompanyName  string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
	SerialNumber string `gorm:"unique;type:varchar(50);column:mc_serial_number;comment:'시리얼넘버'" json:"serialNumber"`
}

func (n *McNetworkDetail) Dump() string {
	pretty, _ := json.MarshalIndent(n, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

type McNetworkPage struct {
	Page     models.Pagination `json:"page"`
	Networks []McNetworkDetail `json:"data"`
}

func (n *McNetworkPage) Dump() string {
	pretty, _ := json.MarshalIndent(n, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
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
