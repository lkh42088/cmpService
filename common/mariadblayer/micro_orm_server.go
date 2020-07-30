package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"fmt"
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

func (db *DBORM) GetMcServersByCpIdx(cpIdx int) (servers []models.McServerDetail, err error) {
	err = db.
		Table("mc_server_tb").
		Select("mc_server_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_server_tb.mc_cp_idx").
		Where(models.McServer{CompanyIdx: cpIdx}).
		Find(&servers).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

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
		Select("mc_vm_tb.*, c.cp_name, m.mc_serial_number").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_tb.vm_cp_idx").
		Joins("INNER JOIN mc_server_tb m ON m.mc_idx = mc_vm_tb.vm_server_idx").
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

func (db *DBORM) updateVmCount(vm models.McVm, isAdd bool) {
	var server models.McServer
	err := db.Where(models.McServer{Idx: uint(vm.McServerIdx)}).
		Find(&server).Error
	fmt.Printf("updateVmCount: vm %v \n", vm)
	fmt.Printf("updateVmCount: server %v \n", server)

	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
		return
	}

	if isAdd {
		server.VmCount += 1
	} else if server.VmCount > 0 {
		server.VmCount -= 1
	}

	err = db.Model(&server).
		Updates(map[string]interface{}{
			"mc_vm_count": server.VmCount,
		}).Error
	fmt.Printf("update: server %v\n", server)
	fmt.Printf("update: err %v\n", err)
}

func (db *DBORM) AddMcVm(obj models.McVm) (vm models.McVm, err error) {
	err = db.Create(&obj).Error
	vm = obj
	if err == nil {
		db.updateVmCount(vm, true)
	}

	return vm, err
}

func (db *DBORM) DeleteMcVm(obj models.McVm) (vm models.McVm, err error) {
	err = db.Delete(&obj).Error
	vm = obj
	if err == nil {
		db.updateVmCount(vm, false)
	}

	return vm, err
}


