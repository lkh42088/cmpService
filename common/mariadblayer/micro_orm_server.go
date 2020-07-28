package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/models"
)

func (db *DBORM) GetMcServersPage(paging models.Pagination) (servers models.McServerPage, err error) {
	err = db.
		Table("mc_server_tb").
		Select("mc_server_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_server_tb.mc_cp_idx").
		Order(servers.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Find(&servers.Servers).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	paging.TotalCount = len(servers.Servers)
	servers.Page = paging
	return servers, err
}

func (db *DBORM) AddMcServer(obj models.McServer) (models.McServer, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) DeleteMcServer(obj models.McServer) (models.McServer, error) {
	return obj, db.Delete(&obj).Error
}

func (db *DBORM) GetMcVmsPage(paging models.Pagination) (vms models.McVmPage, err error) {
	err = db.
		Table("mc_vm_tb").
		Select("mc_vm_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_tb.mc_cp_idx").
		Order(vms.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Find(&vms.Vms).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	paging.TotalCount = len(vms.Vms)
	vms.Page = paging
	return vms, err
}

func (db *DBORM) AddMcVm(obj models.McVm) (models.McVm, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) DeleteMcVm(obj models.McVm) (models.McVm, error) {
	return obj, db.Delete(&obj).Error
}


