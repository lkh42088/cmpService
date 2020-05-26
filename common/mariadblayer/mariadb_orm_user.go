package mariadblayer

import "cmpService/common/models"

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

// UserMember, Customer, Auth
func (db *DBORM) AddUserMember(user models.UserMember) error {
	return db.Create(&user).Error
}

func (db *DBORM) AddCustomer(customer models.Customer) error {
	return db.Create(&customer).Error
}

func (db *DBORM) AddAuth(auth models.Auth) error {
	return db.Create(&auth).Error
}

func (db *DBORM) DeleteAllUserMember() error {
	return db.Delete(&models.UserMember{}).Error
}

func (db *DBORM) DeleteAllCustomer() error {
	return db.Delete(&models.Customer{}).Error
}

func (db *DBORM) DeleteAllAuth() error {
	return db.Delete(&models.Auth{}).Error
}


