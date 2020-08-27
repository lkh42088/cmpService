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

func (db *DBORM) GetMcServerBySerialNumber(sn string) (server mcmodel.McServerDetail, err error) {
	err = db.
		Table("mc_server_tb").
		Select("mc_server_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_server_tb.mc_cp_idx").
		Where(mcmodel.McServer{SerialNumber: sn}).
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

func (db *DBORM) UpdateMcServer(obj mcmodel.McServer) (mcmodel.McServer, error) {
	return obj, db.Model(&obj).
		Update(map[string]interface{}{
			"mc_status":         obj.Status,
			"mc_port":           obj.Port,
			"mc_mac":            obj.Mac,
			"mc_vm_count":       obj.VmCount,
			"mc_ip_addr":        obj.IpAddr,
			"mc_public_ip_addr": obj.PublicIpAddr,
		}).Error
}

func (db *DBORM) DeleteMcServer(obj mcmodel.McServer) (mcmodel.McServer, error) {
	return obj, db.Delete(&obj).Error
}

func (db *DBORM) GetMcVmsByServerIdx(serverIdx int) (obj []mcmodel.McVm, err error) {
	return obj, db.Table("mc_vm_tb").
		Where(mcmodel.McVm{McServerIdx: serverIdx}).
		Find(&obj).Error
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

func (db *DBORM) GetMcVmByIdx(idx uint) (vm mcmodel.McVm, err error) {
	return vm, db.Table("mc_vm_tb").
		Where(mcmodel.McVm{Idx: idx}).
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
	return obj, db.Model(&obj).
		Updates(map[string]interface{}{
			"vm_filename":       obj.Filename,
			"vm_full_path":      obj.FullPath,
			"vm_vmIndex":        obj.VmIndex,
			"vm_ip_addr":        obj.IpAddr,
			"vm_mac":            obj.Mac,
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

func (db *DBORM) AddMcImage(obj mcmodel.McImages) (mcmodel.McImages, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) GetMcImagesByServerIdx(serverIdx int) (obj []mcmodel.McImages, err error) {
	return obj, db.Table("mc_image_tb").
		Where(mcmodel.McImages{McServerIdx: serverIdx}).
		Find(&obj).Error
}

func (db *DBORM) DeleteMcImage(obj mcmodel.McImages) (mcmodel.McImages, error) {
	return obj, db.Delete(&obj).Error
}

func (db *DBORM) GetMcImagesPage(paging models.Pagination) (images mcmodel.McImagePage, err error) {
	err = db.
		Table("mc_image_tb").
		Select("mc_image_tb.*, c.cp_name, m.mc_serial_number").
		Joins("LEFT JOIN mc_server_tb m ON m.mc_idx = mc_image_tb.img_server_idx").
		Joins("LEFT JOIN company_tb c ON c.cp_idx = m.mc_cp_idx").
		//Order(images.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Find(&images.Images).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	paging.TotalCount = len(images.Images)
	images.Page = paging
	return images, err
}

func (db *DBORM) AddMcNetwork(obj mcmodel.McNetworks) (mcmodel.McNetworks, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) GetMcNetworksByServerIdx(serverIdx int) (obj []mcmodel.McNetworks, err error) {
	return obj, db.Table("mc_network_tb").
		Where(mcmodel.McNetworks{McServerIdx: serverIdx}).
		Find(&obj).Error
}

func (db *DBORM) DeleteMcNetwork(obj mcmodel.McNetworks) (mcmodel.McNetworks, error) {
	return obj, db.Delete(&obj).Error
}

func (db *DBORM) GetMcNetworksPage(paging models.Pagination) (networks mcmodel.McNetworkPage, err error) {
	err = db.
		Table("mc_network_tb").
		Select("mc_network_tb.*, c.cp_name, m.mc_serial_number").
		Joins("LEFT JOIN mc_server_tb m ON m.mc_idx = mc_network_tb.net_server_idx").
		Joins("LEFT JOIN company_tb c ON c.cp_idx = m.mc_cp_idx").
		//Order(networks.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Find(&networks.Networks).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	paging.TotalCount = len(networks.Networks)
	networks.Page = paging
	return networks, err
}
