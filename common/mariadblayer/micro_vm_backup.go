package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/models"
)

func (db *DBORM) GetMcVmBackupPageParam(paging models.Pagination, cpIdx string, serverIdx string,
	name string) (obj mcmodel.McBackupPage,
	err error) {
	var query string

	query = "backup_cp_idx = '" + cpIdx + "' and backup_vm_name = '" + name + "' and backup_server_idx" +
		" = '" + serverIdx + "'"

	err = db.Table("mc_vm_backup_tb").Debug().
		Select("mc_vm_backup_tb.*, c.cp_name, m.mc_serial_number").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_backup_tb.backup_cp_idx").
		Joins("INNER JOIN mc_server_tb m ON m.mc_idx = mc_vm_backup_tb.backup_server_idx").
		Order(obj.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Where(query).
		Find(&obj.Backups).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	db.Table("mc_vm_backup_tb").Debug().
		Select("mc_vm_backup_tb.*, c.cp_name, m.mc_serial_number").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_backup_tb.backup_cp_idx").
		Where(query).
		Count(&paging.TotalCount)
	obj.Page = paging

	return obj, err
}
