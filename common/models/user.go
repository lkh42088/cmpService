package models

import (
	"cmpService/common/lib"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID int `gorm:"primary_key;column:idx" json:"-"`
	ID string `gorm:"type:varchar(32);column:id" json:"username"`
	Password string `gorm:"type:varchar(255);column:password" json:"password"`
	Email string `gorm:"type:varchar(64);column:email" json:"email"`
	Name string `gorm:"type:varchar(20);column:name" json:"name"`
	Level int `gorm:"column:level" json:"level"`
	HaveEmailAuth bool `gorm:"column:have_email_auth" json:"haveEmailAuth"`
	HaveGroupEmailAuth bool `gorm:"column:have_group_email_auth" json:"haveGroupEmailAuth"`
}

type UserEmailAuth struct {
	UserEmailAuthID int `gorm:"primary_key;column:idx"`
	// Unique Id : UserId + Email
	// e.g) UniqueId = adminhonggildong@conbridge.com
	//      UserID = admin
	//      Email = honggildong@conbridge.com
	UniqId string `gorm:"type:varchar(128);column:uniqid"`
	UserId string `gorm:"type:varchar(32);column:userid"`
	Email string `gorm:"type:varchar(64);column:email"`
	EmailAuthConfirm bool `gorm:"column:email_auth_confirm" json:"-"`
	// EmailAuthStore : token + secret key
	//   - token : when to login, generate JWT token
	//   - secrete key : when to check email authentication, generate UUID as secret key
	EmailAuthStore string `gorm:"type:varchar(255);column:email_auth_store" json:"-"`
}

type UserEmailAuthMsg struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Secret string `json:"secret"`
}

func (User) TableName() string {
	return "user_tb"
}

func (UserEmailAuth) TableName() string {
	return "user_email_auth_tb"
}

func (u User) String() string {
	return fmt.Sprintf("userid %d, id %s, password %s, email %s, name %s, level %d",
		u.UserID, u.ID, u.Password, u.Email, u.Name, u.Level)
}

func HashPassword(user *User) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		lib.LogWarnln(err)
		return
	}
	user.Password = string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

