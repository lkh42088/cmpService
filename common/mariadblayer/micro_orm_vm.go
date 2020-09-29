package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/models"
	"fmt"
)

func (db *DBORM) GetMcVmsByServerIdx(serverIdx int) (obj []mcmodel.McVm, err error) {
	return obj, db.Table("mc_vm_tb").
		Where(mcmodel.McVm{McServerIdx: serverIdx}).
		Find(&obj).Error
}

func (db *DBORM) GetMcVmByMac(mac string) (obj mcmodel.McVm, err error) {
	return obj, db.
		Table("mc_vm_tb").
		Where(mcmodel.McVm{Mac: mac}).
		Find(&obj).Error
}

func (db *DBORM) GetMcVmsPage(paging models.Pagination, cpName string) (vms mcmodel.McVmPage, err error) {
	var query string
	if cpName == "all" {
		query = ""
	} else {
		query = "c.cp_name = '" + cpName + "'"
	}
	err = db.Debug().
		Table("mc_vm_tb").
		Select("mc_vm_tb.*, c.cp_name, m.mc_serial_number").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_tb.vm_cp_idx").
		Joins("INNER JOIN mc_server_tb m ON m.mc_idx = mc_vm_tb.vm_server_idx").
		Order(vms.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Where(query).
		Find(&vms.Vms).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	db.Table("mc_vm_tb").
		Select("mc_vm_tb.*, c.cp_name, m.mc_serial_number").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_tb.vm_cp_idx").
		Joins("INNER JOIN mc_server_tb m ON m.mc_idx = mc_vm_tb.vm_server_idx").
		Offset(paging.Offset).
		Where(query).Count(&paging.TotalCount)
	vms.Page = paging

	return vms, err
}


func (db *DBORM) GetMcVmsCount(cpName string) (total int, operate int, vm int, err error) {
	var query string
	if cpName == "all" {
		query = ""
	} else {
		query = "c.cp_name = '" + cpName + "'"
	}
	err = db.Debug().
		Table("mc_vm_tb").
		Select("count(distinct(vm_mac))").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_tb.vm_cp_idx").
		Joins("INNER JOIN mc_server_tb m ON m.mc_idx = mc_vm_tb.vm_server_idx").
		Where(query).
		Count(&vm).Error
		//Find(&vms.Vms).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

	return total, operate, vm, err
}
/*
func (db *DBORM) GetMcWinGraphs(mac string) (obj mcmodel.McWinVmGraph, err error) {
	fmt.Println("McWinVmGraph ~ DB 조회 시작!! : ", mac)

	return obj, err
}
*/
func (db *DBORM) GetMcVmByIdx(idx uint) (vm mcmodel.McVm, err error) {
	return vm, db.Table("mc_vm_tb").
		Where(mcmodel.McVm{Idx: idx}).
		Find(&vm).Error
}

func (db *DBORM) GetAllMcVm() (vms []mcmodel.McVm, err error) {
	return vms, db.Table("mc_vm_tb").
		Find(&vms).Error
}

func (db *DBORM) GetMcVmByName(name string) (vm mcmodel.McVm, err error) {
	return vm, db.Table("mc_vm_tb").
		Where(mcmodel.McVm{Name: name}).
		Find(&vm).Error
}

func (db *DBORM) GetMcVmByNameAndCpIdx(name string, cpidx int) (vm mcmodel.McVm, err error) {
	return vm, db.Table("mc_vm_tb").
		Where(mcmodel.McVm{Name: name, CompanyIdx: cpidx}).
		Find(&vm).Error
}

func (db *DBORM) UpdateVmCount(vm mcmodel.McVm, isAdd bool) {
	var server mcmodel.McServer
	err := db.Where(mcmodel.McServer{Idx: uint(vm.McServerIdx)}).
		Find(&server).Error
	fmt.Printf("UpdateVmCount: vm %v \n", vm)
	fmt.Printf("UpdateVmCount: server %v \n", server)

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
		Where(mcmodel.McServer{Idx: server.Idx}).
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
		db.UpdateVmCount(vm, true)
	}

	return vm, err
}

func (db *DBORM) UpdateMcVm(obj mcmodel.McVm) (mcmodel.McVm, error) {
	return obj, db.
		Model(&obj).
		Where(mcmodel.McVm{Idx: obj.Idx}).
		Updates(map[string]interface{}{
			"vm_filename":       obj.Filename,
			"vm_full_path":      obj.FullPath,
			"vm_vmIndex":        obj.VmIndex,
			"vm_ip_addr":        obj.IpAddr,
			"vm_mac":            obj.Mac,
			"vm_vnc_port":       obj.VncPort,
			"vm_current_status": obj.CurrentStatus,
			"vm_remote_addr":    obj.RemoteAddr,
		}).Error
}

func (db *DBORM) DeleteMcVm(obj mcmodel.McVm) (vm mcmodel.McVm, err error) {
	err = db.Delete(&obj).Error
	vm = obj
	if err == nil {
		db.UpdateVmCount(vm, false)
	}

	return vm, err
}

func (db *DBORM) GetVmTotalCount() (total int, operate int, err error) {
	// VM total count
	err = db.
		Table("mc_vm_tb").
		Select("count(distinct(vm_mac))").
		Count(&total).Error
	// Operating VM count
	err = db.
		Table("mc_vm_tb").
		Select("count(distinct(vm_mac))").
		Where(mcmodel.McVm{CurrentStatus: "running"}).
		Count(&operate).Error
	return total, operate, err
}
