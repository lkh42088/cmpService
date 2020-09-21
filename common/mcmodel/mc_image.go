package mcmodel

import (
	"cmpService/common/models"
	"encoding/json"
	"fmt"
	"strings"
)

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
	Variant     string `gorm:"type:varchar(50);column:img_variant;comment:'이미지 타입'" json:"variant"` // os: win10
	Name        string `gorm:"type:varchar(50);column:img_name;comment:'이미지 이름'" json:"name"`       // image : windows10-250G
	Hdd         int    `gorm:"type:int(11);column:img_hdd;comment:'image size'" json:"hdd"`
	Desc        string `gorm:"type:varchar(50);column:img_desc;comment:'이미지 desc'" json:"desc"`
	FullName    string `gorm:"type:varchar(50);column:img_full_name;comment:'이미지 fullname'" json:"fullName"`
}

func (n *McImages) Dump() string {
	pretty, _ := json.MarshalIndent(n, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func (McImages) TableName() string {
	return "mc_image_tb"
}

type McImageDetail struct {
	McImages
	CompanyName  string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
	SerialNumber string `gorm:"unique;type:varchar(50);column:mc_serial_number;comment:'시리얼넘버'" json:"serialNumber"`
}

func (n *McImageDetail) Dump() string {
	pretty, _ := json.MarshalIndent(n, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

type McImagePage struct {
	Page   models.Pagination `json:"page"`
	Images []McImageDetail   `json:"data"`
}

func (n *McImagePage) Dump() string {
	pretty, _ := json.MarshalIndent(n, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
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

