package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/models"
)

func (db *DBORM) AddMcVmSnapshot(obj mcmodel.McVmSnapshot) (mcmodel.McVmSnapshot, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) DeleteMcVmSnapshot(obj mcmodel.McVmSnapshot) (mcmodel.McVmSnapshot, error) {
	return obj, db.Delete(&obj).Error
}

func (db *DBORM) GetMcVmSnapshotByMcServerIdx(idx int) (obj []mcmodel.McVmSnapshot, err error) {
	return obj, db.Table("mc_vm_snapshot_tb").
		Where(mcmodel.McVmSnapshot{McServerIdx: idx}).Find(&obj).Error
}

func (db *DBORM) GetMcVmSnapshotByCpIdx(idx int) (obj []mcmodel.McVmSnapshot, err error) {
	return obj, db.Table("mc_vm_snapshot_tb").
		Where(mcmodel.McVmSnapshot{CompanyIdx: idx}).Find(&obj).Error
}

func (db *DBORM) UpdateMcVmSnapshotCurrent(obj mcmodel.McVmSnapshot) (mcmodel.McVmSnapshot, error) {
	return obj, db.Debug().
		Table("mc_vm_snapshot_tb").
		Where(mcmodel.McVmSnapshot{Idx: obj.Idx}).
		Updates(map[string]interface{}{
			"snap_current": obj.Current,
	}).Error
}

func (db *DBORM) GetMcVmSnapshotCurrentByVmName(vmName string) (obj []mcmodel.McVmSnapshot, err error) {
	return obj, db.Table("mc_vm_snapshot_tb").
		Where(mcmodel.McVmSnapshot{VmName: vmName, Current: true}).Find(&obj).Error
}

func (db *DBORM) GetMcVmSnapshotByName(name string) (obj mcmodel.McVmSnapshot, err error) {
	return obj, db.Table("mc_vm_snapshot_tb").
		Where(mcmodel.McVmSnapshot{Name: name}).Find(&obj).Error
}

func (db *DBORM) GetMcVmSnapshotByIdx(idx uint) (obj mcmodel.McVmSnapshot, err error) {
	return obj, db.Table("mc_vm_snapshot_tb").
		Where(mcmodel.McVmSnapshot{Idx: idx}).Find(&obj).Error
}

func (db *DBORM) GetMcVmSnapshotPage(paging models.Pagination, cpName string) (obj mcmodel.McVmSnapPage, err error) {
	var query string
	if cpName == "all" {
		query = ""
	} else {
		query = "c.cp_name = '" + cpName + "'"
	}
	err = db.Table("mc_vm_snapshot_tb").
		Select("mc_vm_snapshot_tb.*, c.cp_name, m.mc_serial_number").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_snapshot_tb.snap_cp_idx").
		Joins("INNER JOIN mc_server_tb m ON m.mc_idx = mc_vm_snapshot_tb.snap_server_idx").
		Order(obj.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Where(query).
		Find(&obj.Data).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	db.Table("mc_vm_snapshot_tb").
		Select("mc_vm_snapshot_tb.*, c.cp_name, m.mc_serial_number").
		Joins("INNER JOIN company_tb c ON c.cp_idx = mc_vm_snapshot_tb.snap_cp_idx").
		Where(query).
		Count(&paging.TotalCount)
	obj.Page = paging

	return obj, err
}