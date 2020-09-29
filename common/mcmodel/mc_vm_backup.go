package mcmodel

import (
	"cmpService/common/models"
	"encoding/json"
	"fmt"
	"strings"
)

var McVmSnapOrmMap = map[string]string{
	"snap_idx": "idx",
	"snap_server_idx": "serverIdx",
	"snap_cp_idx": "cpIdx",
	"snap_vm_name": "vmName",
	"snap_name": "name",
}

var McVmSnapJsonMap = map[string]string{
	"idx": "snap_idx",
	"serverIdx": "snap_server_idx",
	"cpIdx": "snap_cp_idx",
	"vmName": "snap_vm_name",
	"name": "snap_name",
}

type McVmSnapshot struct {
	Idx         uint   `gorm:"primary_key;column:snap_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	McServerIdx int    `gorm:"type:int(11);column:snap_server_idx;comment:'서버 고유값'" json:"serverIdx"`
	CompanyIdx  int    `gorm:"type:int(11);column:snap_cp_idx;comment:'회사 고유값'" json:"cpIdx"`
	VmName      string `gorm:"type:varchar(50);column:snap_vm_name;comment:'vm 이름'" json:"vmName"`
	Name        string `gorm:"type:varchar(50);column:snap_name;comment:'snapshot 이름'" json:"name"`
	Desc        string `gorm:"type:varchar(50);column:snap_desc;comment:'snapshot description'" json:"desc"`
	Year        int    `gorm:"type:int(11);column:snap_year;comment:'year'" json:"year"`
	Month       int    `gorm:"type:int(11);column:snap_month;comment:'month'" json:"month"`
	Day         int    `gorm:"type:int(11);column:snap_day;comment:'day'" json:"day"`
	Hour        int    `gorm:"type:int(11);column:snap_hour;comment:'hour'" json:"hour"`
	Minute      int    `gorm:"type:int(11);column:snap_minute;comment:'minute'" json:"minute"`
	Second      int    `gorm:"type:int(11);column:snap_second;comment:'second'" json:"second"`
	Current     bool   `gorm:"type:tinyint(1);column:snap_current;comment:'current vm'" json:"current"`
	ServerSn    string `gorm:"-" json:"serverSn"`
}

func (McVmSnapshot) TableName() string {
	return "mc_vm_snapshot_tb"
}

func (s *McVmSnapshot) Dump() {
	pretty, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("VM Snapshot : %s\n", s.VmName)
	fmt.Printf("%s\n", string(pretty))
}

func DumpMcVmSnapList(list []McVmSnapshot) {
	pretty, _ := json.MarshalIndent(list, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("VM Snapshot List: %d\n", len(list))
	fmt.Printf("%s\n", string(pretty))
}

type McVmSnapDetail struct {
	McVmSnapshot
	CompanyName  string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
	SerialNumber string `gorm:"type:varchar(50);column:mc_serial_number" json:"serialNumber"`
}

type McVmSnapPage struct {
	Page models.Pagination `json:"page"`
	Data []McVmSnapDetail  `json:"data"`
}

func (o McVmSnapPage) GetOrderBy(orderby, order string) string {
	val, exists := McVmSnapJsonMap[orderby]
	if !exists {
		val = "snap_idx"
	}
	order = strings.ToLower(order)
	if !(order == "asc" || order == "desc") {
		order = "desc"
	}
	return val + " " + order
}

type McVmBackup struct {
	Idx         uint   `gorm:"primary_key;column:snap_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	McServerIdx int    `gorm:"type:int(11);column:snap_server_idx;comment:'서버 고유값'" json:"serverIdx"`
	CompanyIdx  int    `gorm:"type:int(11);column:snap_cp_idx;comment:'회사 고유값'" json:"cpIdx"`
	VmName      string `gorm:"type:varchar(50);column:snap_vm_name;comment:'vm 이름'" json:"vmName"`
	Name        string `gorm:"type:varchar(50);column:snap_name;comment:'snapshot 이름'" json:"name"`
	Desc        string `gorm:"type:varchar(50);column:snap_desc;comment:'snapshot description'" json:"desc"`
	Month       int    `gorm:"type:int(11);column:snap_month;comment:'month'" json:"month"`
	Day         int    `gorm:"type:int(11);column:snap_day;comment:'day'" json:"day"`
	Hour        int    `gorm:"type:int(11);column:snap_hour;comment:'hour'" json:"hour"`
	Minute      int    `gorm:"type:int(11);column:snap_minute;comment:'minute'" json:"minute"`
}

func (McVmBackup) TableName() string {
	return "mc_vm_backup_tb"
}

type McVmBackupDetail struct {
	McVmBackup
	CompanyName  string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
	SerialNumber string `gorm:"type:varchar(50);column:mc_serial_number" json:"serialNumber"`
}

func DumpMcVmBackupList(list []McVmBackup) {
	pretty, _ := json.MarshalIndent(list, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("VM Backup List: %d\n", len(list))
	fmt.Printf("%s\n", string(pretty))
}
