package mariadblayer

import "cmpService/common/models"

func (db *DBORM) GetAllUserEmailAuth() (objs []models.UserEmailAuth, err error) {
	return objs, db.Find(&objs).Error
}

func (db *DBORM) GetUserEmailAuthByUniqId(uniqId string) (obj models.UserEmailAuth, err error){
	return obj, db.Where("uniqId = ?", uniqId).Find(&obj).Error
}

func (db *DBORM) AddUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) DeleteUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error) {
	return obj, db.Delete(&obj).Error
}

func (db *DBORM) UpdateUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error) {
	return obj, db.Update(&obj).Error
}