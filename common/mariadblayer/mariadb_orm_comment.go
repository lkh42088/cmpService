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
	where := GetWhereString(defaultFieldName)
	return comments, db.
		//Debug().
		Select(CommentSelectQuery).
		Joins(CompanyAndCommentLeftJoinQuery).
		Where(where, code).
		Find(&comments).Error
}

func (db *DBORM) GetCommentByIdx(idx int) (comment models.DeviceComment, err error) {
	where := GetWhereString(idxFieldName)
	return comment, db.Where(where, idx).Find(&comment).Error
}

func (db *DBORM) UpdateComment(comment models.DeviceComment) error {
	where := GetWhereString(idxFieldName)
	return db.Model(&comment).Where(where, comment.Idx).Update(contentsFieldName, comment.Contents).Error
}

func (db *DBORM) AddComment(comment models.DeviceComment) error {
	return db.Create(&comment).Error
}

func (db *DBORM) DeleteAllComments() error {
	return db.Delete(&models.DeviceComment{}).Error
}

func (db *DBORM) DeleteComments(idx int) error {
	dc := models.DeviceComment{}
	where := GetWhereString(idxFieldName)
	return db.Where(where, idx).Delete(&dc).Error
}


