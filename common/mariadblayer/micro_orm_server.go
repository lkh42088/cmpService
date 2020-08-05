package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/models"
	"fmt"
)

func (db *DBORM) GetMcServersPage(paging models.Pagination) (servers mcmodel.McServerPage, err error) {
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

func (db *DBORM) GetMcServerByServerIdx(idx uint) (server mcmodel.McServerDetail, err error) {
	err = db.
		Table("mc_server_tb").
		Select("mc_server_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_server_tb.mc_cp_idx").
		Where(mcmodel.McServer{Idx: idx}).
		Find(&server).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

	return server, err
}

func (db *DBORM) GetMcServersByCpIdx(cpIdx int) (servers []mcmodel.McServerDetail, err error) {
	err = db.
		Table("mc_server_tb").
		Select("mc_server_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_server_tb.mc_cp_idx").
		Where(mcmodel.McServer{CompanyIdx: cpIdx}).
		Find(&servers).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

	return servers, err
}

func (db *DBORM) AddMcServer(obj mcmodel.McServer) (mcmodel.McServer, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) DeleteMcServer(obj mcmodel.McServer) (mcmodel.McServer, error) {
	return obj, db.Delete(&obj).Error
}

func (db *DBORM) GetMcVmsPage(paging models.Pagination) (vms mcmodel.McVmPage, err error) {
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

func (db *DBORM) updateVmCount(vm mcmodel.McVm, isAdd bool) {
	var server mcmodel.McServer
	err := db.Where(mcmodel.McServer{Idx: uint(vm.McServerIdx)}).
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

func (db *DBORM) AddMcVm(obj mcmodel.McVm) (vm mcmodel.McVm, err error) {
	err = db.Create(&obj).Error
	vm = obj
	if err == nil {
		db.updateVmCount(vm, true)
	}

	return vm, err
}

func (db *DBORM) DeleteMcVm(obj mcmodel.McVm) (vm mcmodel.McVm, err error) {
	err = db.Delete(&obj).Error
	vm = obj
	if err == nil {
		db.updateVmCount(vm, false)
	}

	return vm, err
}


