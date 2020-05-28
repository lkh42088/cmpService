package utils

import (
	"cmpService/common/messages"
	"cmpService/common/models"
	"regexp"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

func ValidateUserbyMsg(user messages.UserRegisterMessage, err []string) []string {
	emailCheck := regexp.MustCompile(emailRegex).MatchString(user.Email)
	if emailCheck != true {
		err = append(err, "Invalid email")
	}
	if len(user.Password) < 4 {
		err = append(err, "Invalid password, Password should be more than 4 characters")
	}
	if len(user.Id) < 1 {
		err = append(err, "Invalid id, please enter a name")
	}
	if len(user.Name) < 1 {
		err = append(err, "Invalid name, please enter a name")
	}
	return err
}

func ValidateUser(user models.User, err []string) []string {
	emailCheck := regexp.MustCompile(emailRegex).MatchString(user.Email)
	if emailCheck != true {
		err = append(err, "Invalid email")
	}
	if len(user.Password) < 4 {
		err = append(err, "Invalid password, Password should be more than 4 characters")
	}
	if len(user.UserId) < 1 {
		err = append(err, "Invalid id, please enter a name")
	}
	if len(user.Name) < 1 {
		err = append(err, "Invalid name, please enter a name")
	}
	return err
}