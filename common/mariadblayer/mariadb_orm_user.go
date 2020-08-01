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
		Select("user_tb.*, c.cp_name, a.auth_tag").
		Joins("INNER JOIN company_tb c ON c.cp_idx = user_tb.cp_idx").
		Joins("INNER JOIN auth_tb a ON a.auth_level = user_tb.user_auth_level").
		Where("user_id = ?", id).
		Find(&userDetail).Error
}

func (db *DBORM) GetUserById(id string) (user models.User, err error) {
	return user, db.Where("user_id = ?", id).Find(&user).Error
}

func (db *DBORM) GetUserDetailsByCpIdx(cpIdx int) (userDetails []models.UserDetail, err error) {
	return userDetails, db.
		Table("user_tb").
		//Select("user_tb.*, c.cp_name").
		//Joins("INNER JOIN company_tb c ON c.cp_idx = user_tb.cp_idx").
		Where("cp_idx = ?", cpIdx).
		Find(&userDetails).Error
}

func (db *DBORM) GetUserByEmail(email string) (user models.User, err error) {
	return user, db.Where("user_email = ?", email).Find(&user).Error
}

func (db *DBORM) AddUser(user models.User) (models.User, error) {
	return user, db.Create(&user).Error
}

func (db *DBORM) UpdateUserPassword(user models.User) (models.User, error) {
	return user, db.Model(&user).
		Updates(map[string]interface{}{
			"user_idx":      user.Idx,
			"user_password": user.Password,
		}).Error
}

func (db *DBORM) UpdateUser(user models.User) (models.User, error) {
	// exept: Avata
	return user, db.Model(&user).
		Updates(map[string]interface{}{
			"user_idx":                   user.Idx,
			"user_id":                    user.UserId,
			"user_password":              user.Password,
			"user_is_cp_account":         user.IsCompanyAccount,
			//"cp_idx":                     user.CompanyIdx,
			"user_auth_level":            user.AuthLevel,
			"user_tel":                   user.Tel,
			"user_hp":                    user.HP,
			"user_zip":                   user.Zipcode,
			"user_addr":                  user.Address,
			"user_addr_detail":           user.AddressDetail,
			"user_termination_date":      user.TermDate,
			"user_block_date":            user.BlockDate,
			"user_memo":                  user.Memo,
			"user_work_scope":            user.WorkScope,
			"user_department":            user.Department,
			"user_position":              user.Position,
			"user_email_auth_flag":       user.EmailAuth,
			"user_group_email_auth_flag": user.GroupEmailAuth,
			"user_register_date":         user.RegisterDate,
			"user_last_access_date":      user.LastAccessDate,
			"user_last_access_ip":        user.LastAccessIp,
		}).Error
}

func (db *DBORM) DeleteUser(user models.User) (models.User, error) {
	return user, db.Delete(&user).Error
}

func (db *DBORM) GetUsersPage(paging models.Pagination) (users models.UserPage, err error) {
	db.Model(&users.Users).
		Joins("INNER JOIN company_tb c ON c.cp_idx = user_tb.cp_idx").
		Count(&paging.TotalCount)
	err = db.
		Table("user_tb").
		Select("user_tb.*, c.cp_name, a.auth_tag").
		Joins("INNER JOIN company_tb c ON c.cp_idx = user_tb.cp_idx").
		Joins("INNER JOIN auth_tb a ON a.auth_level = user_tb.user_auth_level").
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

func (db *DBORM) GetUsersPageBySearch(paging models.Pagination, query string) (users models.UserPage, err error) {
	err = db.
		Debug().
		Table("user_tb").
		Select("user_tb.*, c.cp_name, a.auth_tag").
		Order(users.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Where(query).
		Joins("INNER JOIN company_tb c ON c.cp_idx = user_tb.cp_idx").
		Joins("INNER JOIN auth_tb a ON a.auth_level = user_tb.user_auth_level").
		Find(&users.Users).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

	db.Model(&users.Users).
		Joins("INNER JOIN company_tb c ON c.cp_idx = user_tb.cp_idx").
		Where(query).Count(&paging.TotalCount)
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

func (db *DBORM) GetAuth() (auths []models.Auth, err error) {
	return auths, db.Find(&auths).Error
}

func (db *DBORM) DeleteAllUserMember() error {
	return db.Delete(&models.User{}).Error
}

func (db *DBORM) DeleteAllAuth() error {
	return db.Delete(&models.Auth{}).Error
}
