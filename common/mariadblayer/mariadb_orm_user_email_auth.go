package mariadblayer

import "cmpService/common/models"

func (db *DBORM) GetAllUserEmailAuth() (objs []models.UserEmailAuth, err error) {
	return objs, db.Find(&objs).Error
}

func (db *DBORM) GetUserEmailAuthByIdAndEmail(id, email string) (obj models.UserEmailAuth, err error) {
	return obj, db.Where(&models.UserEmailAuth{UserId: id, Email: email}).Find(&obj).Error
}

func (db *DBORM) GetUserEmailAuthByIdAndStore(id, store string) (obj models.UserEmailAuth, err error) {
	return obj, db.Where(&models.UserEmailAuth{UserId: id, EmailAuthStore: store}).Find(&obj).Error
}

func (db *DBORM) AddUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) DeleteUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error) {
	return obj, db.Delete(&obj).Error
}

func (db *DBORM) DeleteUserEmailAuthByUserId(id string) (obj []models.UserEmailAuth, err error) {
	return obj, db.Where(&models.UserEmailAuth{UserId: id}).Delete(&obj).Error
}

func (db *DBORM) UpdateUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error) {
	return obj, db.Model(&obj).UpdateColumns(&models.UserEmailAuth{EmailAuthConfirm: obj.EmailAuthConfirm,
		EmailAuthStore: obj.EmailAuthStore}).Error
}

//
func (db *DBORM) AddLoginAuth(obj models.LoginAuth) (models.LoginAuth, error) {
	return obj, db.Create(&obj).Error
}

func (db *DBORM) UpdateLoginAuth(obj models.LoginAuth) (models.LoginAuth, error) {
	return obj, db.Model(&obj).
		UpdateColumns(&models.LoginAuth{
			EmailAuthConfirm: obj.EmailAuthConfirm,
			EmailAuthStore: obj.EmailAuthStore,
		}).Error
}

func (db *DBORM) DeleteLoginAuthsByUserIdx(userIdx uint) (obj []models.LoginAuth, err error) {
	return obj, db.Where(&models.LoginAuth{UserIdx: userIdx}).Delete(&obj).Error
}

//
func (db *DBORM) GetLoginAuthByMySelfAuth(userId string) (obj models.LoginAuth, err error) {
	return obj, db.Where(&models.LoginAuth{UserId: userId, AuthUserId: userId}).Find(&obj).Error
}

func (db *DBORM) GetLoginAuthsByUserIdx(userIdx uint) (obj []models.LoginAuth, err error) {
	return obj, db.Where(&models.LoginAuth{UserIdx: userIdx}).Find(&obj).Error
}

func (db *DBORM) GetLoginAuthsByUserId(userId string) (obj []models.LoginAuth, err error) {
	return obj, db.Where(&models.LoginAuth{UserId: userId}).Find(&obj).Error
}

func (db *DBORM) GetLoginAuthByUserIdAndTargetId(userId, targetId string) (obj models.LoginAuth, err error) {
	return obj, db.Where(&models.LoginAuth{UserId: userId, AuthUserId: targetId}).Find(&obj).Error
}

func (db *DBORM) GetLoginAuthByUserIdAndTargetEmail(userId, targetEmail string) (obj models.LoginAuth, err error) {
	return obj, db.Where(&models.LoginAuth{UserId: userId, AuthEmail: targetEmail}).Find(&obj).Error
}

func (db *DBORM) GetLoginAuthsByAuthUserId(authUserId string) (obj []models.LoginAuth, err error) {
	return obj, db.Where(&models.LoginAuth{AuthUserId: authUserId}).Find(&obj).Error
}

func (db *DBORM) DeleteLoginAuth(obj models.LoginAuth) (models.LoginAuth, error) {
	return obj, db.Delete(&obj).Error
}

