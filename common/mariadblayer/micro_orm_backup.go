package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/models"
)

func (db *DBORM) GetMcBackupPage(paging models.Pagination, cpName string) (
	backups mcmodel.McBackupPage, err error) {
	var query string
	if cpName == "all" {
		query = ""
	} else {
		query = "c.cp_name = '" + cpName + "'"
	}
	err = db.Debug().
		Table("mc_vm_backup_tb").
		Select("mc_vm_backup_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_backup_tb.backup_cp_idx").
		Order(backups.GetOrderBy(paging.OrderBy, paging.Order)).
		Offset(paging.Offset).
		Where(query).
		Find(&backups.Backups).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	paging.TotalCount = len(backups.Backups)
	backups.Page = paging

	return backups, err
}

func (db *DBORM) GetMcBackupByServerIdx(idx int) (backup mcmodel.McVmBackupDetail, err error) {
	err = db.
		Table("mc_vm_backup_tb").
		Select("mc_vm_backup_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_backup_tb.backup_cp_idx").
		Where(mcmodel.McVmBackup{McServerIdx: idx}).
		Find(&backup).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

	return backup, err
}

func (db *DBORM) GetMcBackupsByCpIdx(cpIdx int) (backups []mcmodel.McVmBackupDetail, err error) {
	err = db.
		Table("mc_vm_backup_tb").
		Select("mc_vm_backup_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_backup_tb.backup_cp_idx").
		Where(mcmodel.McVmBackup{CompanyIdx: cpIdx}).
		Find(&backups).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

	return backups, err
}

func (db *DBORM) GetMcVmBackupByVmName(vmName string) (obj mcmodel.McVmBackup, err error) {
	return obj, db.Table("mc_vm_backup_tb").
		Where(mcmodel.McVmBackup{VmName: vmName}).Find(&obj).Error
}

func (db *DBORM) GetMcVmBackupByName(name string) (obj mcmodel.McVmBackup, err error) {
	return obj, db.Table("mc_vm_backup_tb").
		Where(mcmodel.McVmBackup{Name: name}).Find(&obj).Error
}

func (db *DBORM) GetMcVmBackupByIdx(idx uint) (obj mcmodel.McVmBackup, err error) {
	return obj, db.Table("mc_vm_backup_tb").
		Where(mcmodel.McVmBackup{Idx: idx}).Find(&obj).Error
}

func (db *DBORM) GetMcVmBackup() (backups []mcmodel.McVmBackup, err error) {
	return backups, db.Table("mc_vm_backup_tb").
		Select("mc_vm_backup_tb.*").
		Find(&backups).Error
}

func (db *DBORM) AddMcVmBackup(obj mcmodel.McVmBackup) (mcmodel.McVmBackup, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) UpdateMcBackup(obj mcmodel.McVmBackup) (mcmodel.McVmBackup, error) {
	return obj, db.
		Model(&obj).
		Where(mcmodel.McVmBackup{Idx: obj.Idx}).
		Update(map[string]interface{}{
			"backup_cp_idx":               obj.CompanyIdx,
			"backup_server_idx":           obj.McServerIdx,
			"backup_server_serial_number": obj.McServerSn,
			"backup_kt_auth_url":          obj.KtAuthUrl,
			"backup_nas_name":             obj.NasBackupName,
			"backup_kt_container_name":    obj.KtContainerName,
			"backup_container_date":       obj.KtContainerDate,
			"backup_name":                 obj.Name,
			"backup_register_date":        obj.LastBackupDate,
			"backup_size":                 obj.BackupSize,
			"backup_vm_name":              obj.VmName,
			"backup_desc":                 obj.Desc,
			"backup_month":                obj.Month,
			"backup_day":                  obj.Day,
			"backup_hour":                 obj.Hour,
			"backup_minute":               obj.Minute,
		}).Error
}

func (db *DBORM) UpdateKtAuthUrl(ip string, authUrl string) (mcmodel.McServer, error) {
	server := mcmodel.McServer{}
	return server, db.
		Model(&server).
		Where(mcmodel.McServer{IpAddr: ip}).
		Update(mcmodel.McServer{UcloudAuthUrl: authUrl}).
		Error
}

func (db *DBORM) DeleteMcVmBackup(obj mcmodel.McVmBackup) (mcmodel.McVmBackup, error) {
	return obj, db.Delete(&obj).Error
}
