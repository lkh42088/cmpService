package mariadblayer

import (
	"cmpService/common/models"
)

const defaultFieldName = "device_code"
const idxFieldName = "idx"

func (db *DBORM) GetAllComments() (comments []models.DeviceComment, err error) {
	return comments, db.Find(&comments).Error
}

func (db *DBORM) GetComments(code string) (comments []models.DeviceComment, err error) {
	where := GetWhereString(defaultFieldName)
	return comments, db.Where(where, code).Find(&comments).Error
}

func (db *DBORM) AddComment(comments models.DeviceComment) error {
	return db.Create(comments).Error
}

func (db *DBORM) DeleteAllComments() error {
	return db.Delete(&models.DeviceComment{}).Error
}

func (db *DBORM) DeleteComments(idx int) error {
	dc := models.DeviceComment{}
	where := GetWhereString(idxFieldName)
	return db.Where(where, idx).Delete(&dc).Error
}


