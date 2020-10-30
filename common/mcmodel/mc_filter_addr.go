package mcmodel

import (
	"cmpService/common/models"
	"encoding/json"
	"fmt"
	"strings"
)

var McFilterRuleOrmMap = map[string]string{
	"filter_idx":              "idx",
	"filter_mc_serial_number": "serialNumber",
	"filter_cp_idx":           "cpIdx",
	"filter_ip_addr":          "ipAddr",
}

var McFilterRuleJsonMap = map[string]string{
	"idx":          "filter_idx",
	"serialNumber": "filter_mc_serial_number",
	"cpIdx":        "filter_cp_idx",
	"ipAddr":       "filter_ip_addr",
}

type McFilterRule struct {
	Idx          uint   `gorm:"primary_key;column:filter_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	SerialNumber string `gorm:"unique;type:varchar(50);column:filter_mc_serial_number;comment:'시리얼넘버'" json:"serialNumber"`
	CompanyIdx   int    `gorm:"type:int(11);column:filter_cp_idx;comment:'회사 고유값'" json:"cpIdx"`
	IpAddr       string `gorm:"type:varchar(50);column:filter_ip_addr;comment:'IP Address'" json:"ipAddr"`
}

func (m *McFilterRule) Dump() string {
	pretty, _ := json.MarshalIndent(m, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func (McFilterRule) TableName() string {
	return "mc_filter_rule_tb"
}

type McFilterRuleDetail struct {
	McFilterRule
	CompanyName string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
}

func (m *McFilterRuleDetail) Dump() string {
	pretty, _ := json.MarshalIndent(m, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

type McFilterRulePage struct {
	Page    models.Pagination    `json:"page"`
	FilterRules []McFilterRuleDetail `json:"data"`
}

func (m *McFilterRulePage) Dump() string {
	pretty, _ := json.MarshalIndent(m, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func (m McFilterRulePage) GetOrderBy(orderby, order string) string {
	val, exists := McFilterRuleJsonMap[orderby]
	if !exists {
		val = "filter_idx"
	}
	order = strings.ToLower(order)
	if !(order == "asc" || order == "desc") {
		order = "desc"
	}
	return val + " " + order
}
