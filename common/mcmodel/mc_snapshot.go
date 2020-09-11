package mcmodel

import "cmpService/common/models"

type McVmSnapshot struct {
	Idx    uint   `gorm:"primary_key;column:vs_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	VmIdx  uint   `gorm:"column:vs_vm_idx;not null;auto_increment;comment:'INDEX'" json:"vmIdx"`
	Name   string `gorm:"type:varchar(50);column:vs_name;comment:'vm snapshot 이름'" json:"snapName"`
}

func (McVmSnapshot) TableName() string {
	return "mc_vm_snapshot_tb"
}

type McVmSnapDetail struct {
	McVm
	CompanyName  string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
	SerialNumber string `gorm:"type:varchar(50);column:mc_serial_number" json:"serialNumber"`
}

type McVmSnapPage struct {
	Page models.Pagination `json:"page"`
	Vms  []McVmSnapDetail  `json:"data"`
}
