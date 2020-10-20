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

func (db *DBORM) GetMcServer() (servers []mcmodel.McServer, err error) {
	return servers, db.Table("mc_server_tb").
		Select("mc_server_tb.*").
		Find(&servers).Error
}

func (db *DBORM) GetMcServerByIp(ip string) (server mcmodel.McServer, err error) {
	return server, db.Table("mc_server_tb").
		Select("mc_server_tb.*").
		Where(mcmodel.McServer{IpAddr: ip}).
		Find(&server).Error
}

func (db *DBORM) AddMcServer(obj mcmodel.McServer) (mcmodel.McServer, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) UpdateMcServerAll(obj mcmodel.McServer) (mcmodel.McServer, error) {
	// Update Serial Number !!!
	return obj, db.
		Model(&obj).
		Where(mcmodel.McServer{Idx: obj.Idx}).
		Update(map[string]interface{}{
			"mc_cp_idx":          obj.CompanyIdx,
			"mc_serial_number":   obj.SerialNumber,
			"mc_type":            obj.Type,
			"mc_status":          obj.Status,
			"mc_enable":          obj.Enable,
			"mc_port":            obj.Port,
			"mc_mac":             obj.Mac,
			"mc_vm_count":        obj.VmCount,
			"mc_ip_addr":         obj.IpAddr,
			"mc_public_ip_addr":  obj.PublicIpAddr,
			"mc_l4_port":         obj.L4Port,
			"mc_register_type":   obj.RegisterType,
			"mc_domain_prefix":   obj.DomainPrefix,
			"mc_domain_id":       obj.DomainId,
			"mc_domain_password": obj.DomainPassword,
			"mc_kt_access_key":   obj.UcloudAccessKey,
			"mc_kt_secret_key":   obj.UcloudSecretKey,
			"mc_kt_project_id":   obj.UcloudProjectId,
			"mc_kt_domain_id":    obj.UcloudDomainId,
			"mc_kt_auth_url":	 obj.UcloudAuthUrl,
			"mc_nas_url":         obj.NasUrl,
			"mc_nas_id":          obj.NasId,
			"mc_nas_password":    obj.NasPassword,
		}).Error
}

func (db *DBORM) UpdateMcServer(obj mcmodel.McServer) (mcmodel.McServer, error) {
	return obj, db.
		Model(&obj).
		Where(mcmodel.McServer{Idx: obj.Idx}).
		Update(map[string]interface{}{
			"mc_status":          obj.Status,
			"mc_enable":          obj.Enable,
			"mc_port":            obj.Port,
			"mc_mac":             obj.Mac,
			"mc_vm_count":        obj.VmCount,
			"mc_ip_addr":         obj.IpAddr,
			"mc_public_ip_addr":  obj.PublicIpAddr,
			"mc_l4_port":         obj.L4Port,
			"mc_register_type":   obj.RegisterType,
			"mc_domain_prefix":   obj.DomainPrefix,
			"mc_domain_id":       obj.DomainId,
			"mc_domain_password": obj.DomainPassword,
			"mc_kt_access_key":   obj.UcloudAccessKey,
			"mc_kt_secret_key":   obj.UcloudSecretKey,
			"mc_kt_project_id":   obj.UcloudProjectId,
			"mc_kt_domain_id":    obj.UcloudDomainId,
			"mc_kt_auth_url":	 obj.UcloudAuthUrl,
			"mc_nas_url":         obj.NasUrl,
			"mc_nas_id":          obj.NasId,
			"mc_nas_password":    obj.NasPassword,
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
	return obj, db.
		Model(&obj).
		Where(mcmodel.SysInfo{IfMac: obj.IfMac}).
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
			"update_time":      obj.UpdateTime,
		}).Error
}

func (db *DBORM) GetServerTotalCount() (total int, operate int, err error) {
	// Server total count
	err = db.
		Table("mc_server_tb").
		Select("count(distinct(mc_mac))").
		Count(&total).Error
	// Operating server count
	err = db.
		Table("mc_server_tb").
		Select("count(distinct(mc_mac))").
		Where(mcmodel.McServer{Status: 1}).
		Count(&operate).Error
	return total, operate, err
}
