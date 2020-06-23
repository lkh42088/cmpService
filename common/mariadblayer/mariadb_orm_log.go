package mariadblayer

import (
	"cmpService/common/models"
)

func (db *DBORM) GetAllLogs() (logs []models.DeviceLog, err error) {
	return logs, db.Find(&logs).Error
}

func (db *DBORM) GetLogs(code string) (logs []models.DeviceLog, err error) {
	where := GetWhereString(defaultFieldName)
	return logs, db.Where(where, code).Find(&logs).Error
}

func (db *DBORM) GetLogByIdx(idx int) (log models.DeviceLog, err error) {
	where := GetWhereString(idxFieldName)
	return log, db.Where(where, idx).Find(&log).Error
}

func (db *DBORM) UpdateLog(field string, change string, log models.DeviceLog) error {
	where := GetWhereString(idxFieldName)
	return db.Model(&log).Where(where, log.Idx).Update(field, change).Error
}

func (db *DBORM) AddLog(log models.DeviceLog) error {
	return db.Create(&log).Error
}

func (db *DBORM) DeleteAllLogs() error {
	return db.Delete(&models.DeviceLog{}).Error
}

func (db *DBORM) DeleteLog(idx int) error {
	dl := models.DeviceLog{}
	where := GetWhereString(idxFieldName)
	return db.Where(where, idx).Delete(&dl).Error
}
