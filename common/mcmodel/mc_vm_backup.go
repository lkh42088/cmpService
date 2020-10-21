package mcmodel

import (
	"cmpService/common/models"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var McVmSnapOrmMap = map[string]string{
	"snap_idx":        "idx",
	"snap_server_idx": "serverIdx",
	"snap_cp_idx":     "cpIdx",
	"snap_vm_name":    "vmName",
	"snap_name":       "name",
	"snap_desc":       "desc",
	"snap_year":       "year",
	"snap_month":      "month",
	"snap_day":        "day",
	"snap_hour":       "hour",
	"snap_minute":     "minute",
	"snap_second":     "second",
	"snap_current":    "current",
}

var McVmSnapJsonMap = map[string]string{
	"idx":       "snap_idx",
	"serverIdx": "snap_server_idx",
	"cpIdx":     "snap_cp_idx",
	"vmName":    "snap_vm_name",
	"name":      "snap_name",
	"desc":      "snap_desc",
	"year":      "snap_year",
	"month":     "snap_month",
	"day":       "snap_day",
	"hour":      "snap_hour",
	"minute":    "snap_minute",
	"second":    "snap_second",
	"current":   "snap_current",
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
	Command     string `gorm:"-" json:"command"`
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

/** BACKUP */
var McVmBackupOrmMap = map[string]string{
	"backup_idx":                  "idx",
	"backup_cp_idx":               "cpIdx",
	"backup_server_idx":           "serverIdx",
	"backup_server_serial_number": "serverSn",
	"backup_kt_auth_url":          "authUrl",
	"backup_nas_name":             "nasBackupName",
	"backup_kt_container_name":    "containerName",
	"backup_container_date":       "containerDate",
	"backup_name":                 "name",
	"backup_register_date":        "registerDate",
	"backup_size":                 "fileSize",
	"backup_vm_name":              "vmName",
	"backup_desc":                 "desc",
	"backup_month":                "month",
	"backup_day":                  "day",
	"backup_hour":                 "hour",
	"backup_minute":               "minute",
}

var McVmBackupJsonMap = map[string]string{
	"idx":           "backup_idx",
	"cpIdx":         "backup_cp_idx",
	"serverIdx":     "backup_server_idx",
	"serverSn":      "backup_server_serial_number",
	"authUrl":       "backup_kt_auth_url",
	"nasBackupName": "backup_nas_name",
	"containerName": "backup_kt_container_name",
	"containerDate": "backup_container_date",
	"name":          "backup_name",
	"registerDate":  "backup_register_date",
	"fileSize":      "backup_size",
	"vmName":        "backup_vm_name",
	"desc":          "backup_desc",
	"month":         "backup_month",
	"day":           "backup_day",
	"hour":          "backup_hour",
	"minute":        "backup_minute",
}

type McVmBackup struct {
	Idx             uint      `gorm:"primary_key;column:backup_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	CompanyIdx      int       `gorm:"type:int(11);not null;column:backup_cp_idx;comment:'회사 고유값'" json:"cpIdx"`
	McServerIdx     int       `gorm:"type:int(11);column:backup_server_idx;comment:'서버 고유값'" json:"serverIdx"`
	McServerSn      string    `gorm:"type:int(11);column:backup_server_serial_number;comment:'서버 시리얼 넘버'" json:"serverSn"`
	KtAuthUrl       string    `gorm:"type:int(11);column:backup_kt_auth_url;comment:'KT 사용자 인증 URL'" json:"authUrl"`
	NasBackupName   string    `gorm:"type:int(11);column:backup_nas_name;comment:'NAS 백업 파일 이름'" json:"nasBackupName"`
	KtContainerName string    `gorm:"type:int(11);column:backup_kt_container_name;comment:'컨테이너 이름'" json:"containerName"`
	KtContainerDate time.Time `gorm:"type:datetime;column:backup_container_date;comment:'컨테이너 생성일'" json:"containerDate"`
	Name            string    `gorm:"type:varchar(50);column:backup_name;comment:'백업 파일 이름'" json:"filename"`
	LastBackupDate  time.Time `gorm:"type:datetime;column:backup_register_date;comment:'최종 백업 날짜'" json:"registerDate"`
	BackupSize      int       `gorm:"type:int(11);column:backup_size;comment:'백업 이미지 크기'" json:"fileSize"`
	VmName          string    `gorm:"type:varchar(50);column:backup_vm_name;comment:'백업 VM 이름'" json:"vmName"`
	Desc            string    `gorm:"type:varchar(255);column:backup_desc;comment:'백업 상세'" json:"desc"`
	Year            int       `gorm:"type:int(11);column:snap_year;comment:'year'" json:"year"`
	Month           int       `gorm:"type:int(11);column:backup_month;comment:'month'" json:"month"`
	Day             int       `gorm:"type:int(11);column:backup_day;comment:'day'" json:"day"`
	Hour            int       `gorm:"type:int(11);column:backup_hour;comment:'hour'" json:"hour"`
	Minute          int       `gorm:"type:int(11);column:backup_minute;comment:'minute'" json:"minute"`
	Second          int       `gorm:"type:int(11);column:snap_second;comment:'second'" json:"second"`
	ServerSn        string    `gorm:"-" json:"serverSn"`
	Command         string    `gorm:"-" json:"command"`
}

func (McVmBackup) TableName() string {
	return "mc_vm_backup_tb"
}

type McVmBackupDetail struct {
	McVmBackup
	CompanyName  string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
	SerialNumber string `gorm:"type:varchar(50);column:mc_serial_number" json:"serialNumber"`
}

type McBackupPage struct {
	Page    models.Pagination  `json:"page"`
	Backups []McVmBackupDetail `json:"data"`
}

func (s *McVmBackup) Dump() {
	pretty, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("VM Backup: %s\n", s.VmName)
	fmt.Printf("%s\n", string(pretty))
}

func DumpMcVmBackupList(list []McVmBackup) {
	pretty, _ := json.MarshalIndent(list, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("VM Backup List: %d\n", len(list))
	fmt.Printf("%s\n", string(pretty))
}

func (m McBackupPage) GetOrderBy(orderby, order string) string {
	val, exists := McVmBackupJsonMap[orderby]
	if !exists {
		val = "mc_idx"
	}
	order = strings.ToLower(order)
	if !(order == "asc" || order == "desc") {
		order = "desc"
	}
	return val + " " + order
}
