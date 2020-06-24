package mariadblayer

import (
	"cmpService/common/models"
)

const defaultFieldName = "device_code"
const idxFieldName = "comment_idx"
const contentsFieldName = "contents"

func (db *DBORM) GetAllComments() (comments []models.DeviceComment, err error) {
	return comments, db.Find(&comments).Error
}

func (db *DBORM) GetComments(code string) (comments []models.DeviceComment, err error) {
	return comments, db.Where(models.DeviceComment{DeviceCode: code}).Find(&comments).Error
}

func (db *DBORM) GetCommentByIdx(idx int) (comment models.DeviceComment, err error) {
	return comment, db.Where(models.DeviceComment{Idx: uint(idx)}).Find(&comment).Error
}

func (db *DBORM) UpdateComment(comment models.DeviceComment) error {
	return db.
		Model(&comment).
		Where(models.DeviceComment{Idx: comment.Idx}).
		Update(models.DeviceComment{Contents: comment.Contents}).Error
}

func (db *DBORM) AddComment(comment models.DeviceComment) error {
	return db.Create(&comment).Error
}

func (db *DBORM) DeleteAllComments() error {
	return db.Delete(&models.DeviceComment{}).Error
}

func (db *DBORM) DeleteComments(idx int) error {
	dc := models.DeviceComment{}
	return db.Where(models.DeviceComment{Idx: uint(idx)}).Delete(&dc).Error
}
