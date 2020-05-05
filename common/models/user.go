package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"nubes/common/lib"
)

type User struct {
	UserID int `gorm:"primary_key;column:idx" json:"-"`
	ID string `gorm:"type:varchar(32);column:id" json:"id"`
	Password string `gorm:"type:varchar(256);column:password" json:"password"`
	Email string `gorm:"type:varchar(64);column:email" json:"email"`
	Name string `gorm:"type:varchar(20);column:name" json:"name"`
	Level int `gorm:"column:level" json:"level"`
}

func (User) TableName() string {
	return "user_tb"
}

func (u User) String() string {
	return fmt.Sprintf("userid %d, id %s, password %s, email %s, level %d",
		u.UserID, u.ID, u.Password, u.Email, u.Level)
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
