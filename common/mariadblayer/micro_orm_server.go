package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/models"
)

func (db *DBORM) GetMcServersPage(paging models.Pagination, cpName string) (
	servers mcmodel.McServerPage, err error) {
	var query string
	if cpName == "all" {
		query = ""
	} else {
		query = "c.cp_name = '" + cpName + "'"
	}
	err = db.Debug().
		Table("mc_server_tb").
		Select("mc_server_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_server_tb.mc_cp_idx").
		Order(servers.GetOrderBy(paging.OrderBy, paging.Order)).
		/*Limit(paging.RowsPerPage).*/
		Offset(paging.Offset).
		//Where("c.cp_name = ?", cpName).
		Where(query).
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

// Baremetal system info
func (db *DBORM) GetSystemInfoByMac(mac string) (info mcmodel.SysInfo, err error) {
	err = db.Where(mcmodel.SysInfo{IfMac: mac}).Find(&info).Error

	return info, err
}

func (db *DBORM) AddSystemInfo(obj mcmodel.SysInfo) (mcmodel.SysInfo, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) UpdateSystemInfo(obj mcmodel.SysInfo) (mcmodel.SysInfo, error) {
	return obj, db.Model(&obj).
		Update(map[string]interface{}{
			"hostname":         obj.Hostname,
			"os":               obj.OS,
			"uptime":           obj.Uptime,
			"boottime":         obj.BootTime,
			"cpu_core":         obj.CpuCore,
			"cpu_model":        obj.CpuModel,
			"platform":         obj.Platform,
			"platform_version": obj.PlatformVersion,
			"kernel_arch":      obj.KernelArch,
			"kernel_version":   obj.KernelVersion,
			"ip":               obj.IP,
			"if_name":          obj.IfName,
			"if_mac":           obj.IfMac,
			"mem_total":        obj.MemTotal,
			"disk_total":       obj.DiskTotal,
			"update_time":		obj.UpdateTime,
		}).Error
}
