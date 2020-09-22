package mcmodel

import (
	"cmpService/common/models"
	"encoding/json"
	"fmt"
	"strings"
)

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
	Idx            uint   `gorm:"primary_key;column:vm_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	McServerIdx    int    `gorm:"type:int(11);column:vm_server_idx;comment:'서버 고유값'" json:"serverIdx"`
	CompanyIdx     int    `gorm:"type:int(11);column:vm_cp_idx;comment:'회사 고유값'" json:"cpIdx"`
	Name           string `gorm:"type:varchar(50);column:vm_name;comment:'vm 이름'" json:"name"`
	Cpu            int    `gorm:"type:int(11);column:vm_cpu;comment:'vm cpu'" json:"cpu"`
	Ram            int    `gorm:"type:int(11);column:vm_ram;comment:'vm ram'" json:"ram"`
	Hdd            int    `gorm:"type:int(11);column:vm_hdd;comment:'vm hdd'" json:"hdd"`
	Desc           string `gorm:"type:varchar(100);column:vm_desc;comment:'vm description'" json:"desc"`
	OS             string `gorm:"type:varchar(50);column:vm_os;comment:'vm os'" json:"os"`
	Image          string `gorm:"type:varchar(50);column:vm_image;comment:'vm image'" json:"image"`
	Filename       string `gorm:"type:varchar(50);column:vm_filename;comment:'vm image'" json:"filename"`
	VmIndex        int    `gorm:"type:int(11);column:vm_vmIndex;comment:'vm index'" json:"vmIndex"`
	FullPath       string `gorm:"type:varchar(50);column:vm_full_path;comment:'file full path'" json:"fullPath"`
	Network        string `gorm:"type:varchar(50);column:vm_network;comment:'vm network'" json:"network"`
	IpAddr         string `gorm:"type:varchar(50);column:vm_ip_addr;comment:'vm ip address'" json:"ipAddr"`
	RemoteAddr     string `gorm:"type:varchar(50);column:vm_remote_addr;comment:'Remote Address for RDP'" json:"remoteAddr"`
	VncPort        string `gorm:"type:varchar(50);column:vm_vnc_port;comment:'vm vnc port'" json:"vncPort"`
	Mac            string `gorm:"type:varchar(50);column:vm_mac;comment:'vm mac address'" json:"mac"`
	ConfigStatus   string `gorm:"type:varchar(50);column:vm_config_status;comment:'vm config status'" json:"configStatus"`
	CurrentStatus  string `gorm:"type:varchar(50);column:vm_current_status;comment:'vm current status'" json:"currentStatus"`
	SnapType       bool   `gorm:"type:tinyint(1);column:vm_snap_type" json:"snapType"`
	SnapDays       int    `gorm:"type:int(11);column:vm_snap_days" json:"snapDays"`
	SnapHours      int    `gorm:"type:int(11);column:vm_snap_hours" json:"snapHours"`
	SnapMinutes    int    `gorm:"type:int(11);column:vm_snap_minutes" json:"snapMinutes"`
	IsCreated      bool   `gorm:"-" json:"isCreated"`
	IsProcess      bool   `gorm:"-" json:"isProcess"`
	IsChangeIpAddr bool   `gorm:"-" json:"-"`
}

func (v *McVm) Dump() string {
	pretty, _ := json.MarshalIndent(v, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func DumpMcVmList(list []McVm) {
	pretty, _ := json.MarshalIndent(list, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("VM List: %d\n", len(list))
	fmt.Printf("%s\n", string(pretty))
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

func (v McVm) LookupList(list *[]McVm) (bool, int, *McVm){
	for index, vm := range *list {
		if vm.Name == v.Name {
			return true, index, &vm
		}
	}
	return false, 0, nil
}