package mcmodel

import (
	"cmpService/common/models"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	McServerOrmIdx             = "mc_idx"
	McServerOrmSerialNumber    = "mc_serial_number"
	mcServerOrmCompanyIdx      = "mc_cp_idx"
	mcServerOrmType            = "mc_type"
	mcServerOrmStatus          = "mc_status"
	mcServerOrmVmCount         = "mc_vm_count"
	mcServerOrmIpAddr          = "mc_ip_addr"
	McServerOrmUcloudAccessKey = "mc_kt_access_key"
	McServerOrmUcloudSecretKey = "mc_kt_secret_key"
	McServerOrmUcloudProjectId = "mc_kt_project_id"
	McServerOrmUcloudDomainId  = "mc_kt_domain_id"
	McServerOrmUcloudAuthUrl   = "mc_kt_auth_url"
	McServerOrmNasUrl          = "mc_nas_url"
	McServerOrmNasId           = "mc_nas_id"
	McServerOrmNasPassword     = "mc_nas_password"
)

const (
	McServerJsonIdx             = "idx"
	McServerJsonSerialNumber    = "serialNumber"
	mcServerJsonCompanyIdx      = "cpIdx"
	mcServerJsonType            = "type"
	mcServerJsonStatus          = "status"
	mcServerJsonVmCount         = "vmCount"
	mcServerJsonIpAddr          = "ipAddr"
	McServerJsonUcloudAccessKey = "accessKey"
	McServerJsonUcloudSecretKey = "secretKey"
	McServerJsonUcloudProjectId = "projectId"
	McServerJsonUcloudDomainId  = "ktDomainId"
	McServerJsonUcloudAuthUrl   = "authUrl"
	McServerJsonNasUrl          = "nasUrl"
	McServerJsonNasId           = "nasId"
	McServerJsonNasPassword     = "nasPassword"
)

var McServerOrmMap = map[string]string{
	McServerOrmIdx:             McServerJsonIdx,
	McServerOrmSerialNumber:    McServerJsonSerialNumber,
	mcServerOrmCompanyIdx:      mcServerJsonCompanyIdx,
	mcServerOrmType:            mcServerJsonType,
	mcServerOrmStatus:          mcServerJsonStatus,
	mcServerOrmVmCount:         mcServerJsonVmCount,
	mcServerOrmIpAddr:          mcServerJsonIpAddr,
	McServerOrmUcloudAccessKey: McServerJsonUcloudAccessKey,
	McServerOrmUcloudSecretKey: McServerJsonUcloudSecretKey,
	McServerOrmUcloudProjectId: McServerJsonUcloudProjectId,
	McServerOrmUcloudDomainId:  McServerJsonUcloudDomainId,
	McServerOrmUcloudAuthUrl:   McServerJsonUcloudAuthUrl,
	McServerOrmNasUrl:          McServerJsonNasUrl,
	McServerOrmNasId:           McServerJsonNasId,
	McServerOrmNasPassword:     McServerJsonNasPassword,
}

var McServerJsonMap = map[string]string{
	McServerJsonIdx:             McServerOrmIdx,
	McServerJsonSerialNumber:    McServerOrmSerialNumber,
	mcServerJsonCompanyIdx:      mcServerOrmCompanyIdx,
	mcServerJsonType:            mcServerOrmType,
	mcServerJsonStatus:          mcServerOrmStatus,
	mcServerJsonVmCount:         mcServerOrmVmCount,
	mcServerJsonIpAddr:          mcServerOrmIpAddr,
	McServerJsonUcloudAccessKey: McServerOrmUcloudAccessKey,
	McServerJsonUcloudSecretKey: McServerOrmUcloudSecretKey,
	McServerJsonUcloudProjectId: McServerOrmUcloudProjectId,
	McServerJsonUcloudDomainId:  McServerOrmUcloudDomainId,
	McServerJsonUcloudAuthUrl:   McServerOrmUcloudAuthUrl,
	McServerJsonNasUrl:          McServerOrmNasUrl,
	McServerJsonNasId:           McServerOrmNasId,
	McServerJsonNasPassword:     McServerOrmNasPassword,
}

type McServer struct {
	Idx             uint   `gorm:"primary_key;column:mc_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	SerialNumber    string `gorm:"unique;type:varchar(50);column:mc_serial_number;comment:'시리얼넘버'" json:"serialNumber"`
	CompanyIdx      int    `gorm:"type:int(11);column:mc_cp_idx;comment:'회사 고유값'" json:"cpIdx"`
	Type            string `gorm:"type:varchar(50);column:mc_type;comment:'서버 타입'" json:"type"`
	Enable          bool   `gorm:"type:tinyint(1);column:mc_enable;comment:'registration status'" json:"enable"`
	Status          int    `gorm:"type:int(11);column:mc_status;comment:'서버 상태'" json:"status"`
	VmCount         int    `gorm:"type:int(11);column:mc_vm_count;comment:'vm 개수'" json:"vmCount"`
	Port            string `gorm:"type:varchar(50);column:mc_port;comment:'Port'" json:"port"`
	Mac             string `gorm:"type:varchar(50);column:mc_mac;comment:'Mac Address'" json:"mac"`
	IpAddr          string `gorm:"type:varchar(50);column:mc_ip_addr;comment:'IP Address'" json:"ipAddr"`
	PublicIpAddr    string `gorm:"type:varchar(50);column:mc_public_ip_addr;comment:'Public IP Address'" json:"publicIp"`
	L4Port          string `gorm:"type:varchar(50);column:mc_l4_port;comment:'L4 Port'" json:"l4Port"`
	RegisterType    int    `gorm:"type:int(11);column:mc_register_type;comment:'Register type'" json:"registerType"`
	DomainPrefix    string `gorm:"type:varchar(50);column:mc_domain_prefix;comment:'Domain prefix'" json:"domainPrefix"`
	DomainId        string `gorm:"type:varchar(50);column:mc_domain_id;comment:'Domain id'" json:"domainId"`
	DomainPassword  string `gorm:"type:varchar(50);column:mc_domain_password;comment:'Domain password'" json:"domainPassword"`
	UcloudAccessKey string `gorm:"type:varchar(50);column:mc_kt_access_key;comment:'KT Ucloud Access Key'" json:"accessKey"`
	UcloudSecretKey string `gorm:"type:varchar(50);column:mc_kt_secret_key;comment:'KT Ucloud Secret Key'" json:"secretKey"`
	UcloudProjectId string `gorm:"type:varchar(50);column:mc_kt_project_id;comment:'KT Ucloud Project ID'" json:"projectId"`
	UcloudDomainId  string `gorm:"type:varchar(50);column:mc_kt_domain_id;comment:'KT Ucloud Domain ID'" json:"ktDomainId"`
	UcloudAuthUrl   string `gorm:"type:varchar(255);column:mc_kt_auth_url;comment:'KT 사용자 인증 URL'" json:"authUrl"`
	NasUrl          string `gorm:"type:varchar(50);column:mc_nas_url;comment:'NAS URL'" json:"nasUrl"`
	NasId           string `gorm:"type:varchar(50);column:mc_nas_id;comment:'NAS ID'" json:"nasId"`
	NasPassword     string `gorm:"type:varchar(50);column:mc_nas_password;comment:'NAS PASSWORD'" json:"nasPassword"`
}

func (s *McServer) Dump() string {
	pretty, _ := json.MarshalIndent(s, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func (McServer) TableName() string {
	return "mc_server_tb"
}

type McServerDetail struct {
	McServer
	CompanyName string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
}

func (s *McServerDetail) Dump() string {
	pretty, _ := json.MarshalIndent(s, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

type McServerPage struct {
	Page    models.Pagination `json:"page"`
	Servers []McServerDetail  `json:"data"`
}

func (s *McServerPage) Dump() string {
	pretty, _ := json.MarshalIndent(s, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
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

type McServerMsg struct {
	SerialNumber    string        `json:"serialNumber"`
	Port            string        `json:"port"`
	Mac             string        `json:"mac"`
	Ip              string        `json:"ip"`
	L4Port          string        `json:"l4Port"`
	PublicIp        string        `json:"publicIp"`
	UcloudAccessKey string        `json:"accessKey"`
	UcloudSecretKey string        `json:"secretKey"`
	UcloudProjectId string        `json:"projectId"`
	UcloudDomainId  string        `json:"ktDomainId"`
	UcloudAuthUrl   string        `json:"authUrl"`
	NasUrl          string        `json:"nasUrl"`
	NasId           string        `json:"nasId"`
	NasPassword     string        `json:"nasPassword"`
	Vms             *[]McVm       `json:"vms"`
	Networks        *[]McNetworks `json:"networks"`
	Images          *[]McImages   `json:"images"`
}

func (s *McServerMsg) Dump() string {
	pretty, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func (n *McServerMsg) DumpSummary() {
	vmCount := 0
	netCount := 0
	imgCount := 0
	if n.Vms != nil {
		vmCount = len(*n.Vms)
	}
	if n.Networks != nil {
		netCount = len(*n.Networks)
	}
	if n.Images != nil {
		imgCount = len(*n.Images)
	}
	if n.SerialNumber == "" {
		fmt.Printf("SN(none), ")
	} else {
		fmt.Printf("SN(%s), ", n.SerialNumber)
	}
	fmt.Printf("server(%s/%s, %s/%s), vm(%d), network(%d), image(%d)\n",
		n.Ip,
		n.PublicIp,
		n.Port,
		n.Mac,
		vmCount, netCount, imgCount)
}

type McWinVmGraph struct {
	Cpu     WinCpuStat
	Mem     WinMemStat
	Disk    WinDiskStat
	Traffic []WinVmIfStat
}

func DumpVmList(list []McVm) {
	pretty, _ := json.MarshalIndent(list, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("VM List: %d\n", len(list))
	fmt.Printf("%s\n", string(pretty))
}

func DumpNetworkList(list []McNetworks) {
	pretty, _ := json.MarshalIndent(list, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("Network List: %d\n", len(list))
	fmt.Printf("%s\n", string(pretty))
}

func DumpImageList(list []McImages) {
	pretty, _ := json.MarshalIndent(list, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("image List: %d\n", len(list))
	fmt.Printf("%s\n", string(pretty))
}
