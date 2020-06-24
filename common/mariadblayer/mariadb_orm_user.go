package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/models"
)

func (db *DBORM) GetAllUsers() (users []models.User, err error) {
	return users, db.Find(&users).Error
}

func (db *DBORM) GetUserDetailById(id string) (userDetail models.UserDetail, err error) {
	return userDetail, db.
		Table("user_tb").
		Select("user_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = user_tb.cp_idx").
		Where("user_id = ?", id).
		Find(&userDetail).Error
}

func (db *DBORM) GetUserById(id string) (user models.User, err error) {
	return user, db.Where("user_id = ?", id).Find(&user).Error
}

func (db *DBORM) GetUserByEmail(email string) (user models.User, err error) {
	return user, db.Where("user_email = ?", email).Find(&user).Error
}

func (db *DBORM) AddUser(user models.User) (models.User, error) {
	return user, db.Create(&user).Error
}

func (db *DBORM) DeleteUser(user models.User) (models.User, error) {
	return user, db.Delete(&user).Error
}

func (db *DBORM) GetUsersPage(paging models.Pagination) (users models.UserPage, err error) {
	db.Model(&users.Users).Count(&paging.TotalCount)
	err = db.
		Table("user_tb").
		Select("user_tb.*, c.cp_name").
		Joins("INNER JOIN company_tb c ON c.cp_idx = user_tb.cp_idx").
		Order(users.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Find(&users.Users).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	users.Page = paging
	return users, err
}

func (db *DBORM) GetUserByUserId(userId string) (user models.User, err error) {
	return user, db.Where(models.User{UserId: userId}).Find(&user).Error
}

func (db *DBORM) AddUserMember(user models.User) error {
	return db.Create(&user).Error
}

func (db *DBORM) AddAuth(auth models.Auth) error {
	return db.Create(&auth).Error
}

func (db *DBORM) DeleteAllUserMember() error {
	return db.Delete(&models.User{}).Error
}

func (db *DBORM) DeleteAllAuth() error {
	return db.Delete(&models.Auth{}).Error
}


