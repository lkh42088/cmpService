package mariadblayer

import (
	"cmpService/common/models"
)

const defaultFieldName = "device_code"

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

func (db *DBORM) DeleteComments(code string) error {
	dc := models.DeviceComment{}
	where := GetWhereString(defaultFieldName)
	return db.Where(where, code).Delete(dc).Error
}


