package mariadblayer

import "cmpService/common/mcmodel"

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
