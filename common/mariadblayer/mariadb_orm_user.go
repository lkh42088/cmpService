package mariadblayer

import (
	"cmpService/common/models"
)

func (db *DBORM) GetAllUsers() (users []models.User, err error) {
	return users, db.Find(&users).Error
}

func (db *DBORM) GetUserById(id string) (user models.User, err error) {
	return user, db.Where("id = ?", id).Find(&user).Error
}

func (db *DBORM) GetUserByEmail(email string) (user models.User, err error) {
	return user, db.Where("email = ?", email).Find(&user).Error
}

func (db *DBORM) AddUser(user models.User) (models.User, error) {
	return user, db.Create(&user).Error
}

func (db *DBORM) DeleteUser(user models.User) (models.User, error) {
	return user, db.Delete(&user).Error
}

// User, Company, Auth
func (db *DBORM) GetCompaniesByName(name string) (companies []models.Company, err error) {
	name = "%" + name + "%"
	return companies, db.Debug().Where("cp_name like ?", name).Find(&companies).Error
}

func (db *DBORM) AddUserMember(user models.User) error {
	return db.Create(&user).Error
}

func (db *DBORM) AddCompany(company models.Company) (models.Company, error) {
	return company, db.Create(&company).Error
}

func (db *DBORM) AddAuth(auth models.Auth) error {
	return db.Create(&auth).Error
}

func (db *DBORM) DeleteAllUserMember() error {
	return db.Delete(&models.User{}).Error
}

func (db *DBORM) DeleteAllCompany() error {
	return db.Delete(&models.Company{}).Error
}

func (db *DBORM) DeleteAllAuth() error {
	return db.Delete(&models.Auth{}).Error
}


