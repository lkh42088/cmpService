package mariadblayer

import "cmpService/common/models"

func (db *DBORM) GetAllUserEmailAuth() (objs []models.UserEmailAuth, err error) {
	return objs, db.Find(&objs).Error
}

func (db *DBORM) GetUserEmailAuthByIdAndEmail(id, email string) (obj models.UserEmailAuth, err error){
	return obj, db.Where(&models.UserEmailAuth{UserId:id, Email:email}).Find(&obj).Error
}

func (db *DBORM) GetUserEmailAuthByIdAndStore(id, store string) (obj models.UserEmailAuth, err error){
	return obj, db.Where(&models.UserEmailAuth{UserId:id, EmailAuthStore:store}).Find(&obj).Error
}

func (db *DBORM) AddUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) DeleteUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error) {
	return obj, db.Delete(&obj).Error
}

func (db *DBORM) DeleteUserEmailAuthByUserId(id string) (obj []models.UserEmailAuth, err error) {
	return obj, db.Where(&models.UserEmailAuth{UserId:id}).Delete(&obj).Error
}

func (db *DBORM) UpdateUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error) {
	return obj, db.Model(&obj).UpdateColumns(&models.UserEmailAuth{EmailAuthConfirm:obj.EmailAuthConfirm,
		EmailAuthStore:obj.EmailAuthStore}).Error
}
