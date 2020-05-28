package messages

import "cmpService/common/models"

type UserLoginMessage struct {
	Id string `json:"id"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type UserRegisterMessage struct {
	Id                 string   `json:"id"`
	Password           string   `json:"password"`
	Email              string   `json:"email"`
	Name               string   `json:"name"`
	EmailAuthFlag      bool     `json:"email_auth_flag"`
	EmailAuthGroupFlag bool     `json:"email_auth_group_flag"`
	EmailAuthGroupList []string `json:"email_auth_group_list"`
}

type UserInfo struct {
	Id string `json:"id"`
	Password string `json:"password"`
	Email string `json:"email"`
	Name string `json:"name"`
	EmailAuthFlag bool `json:"email_auth_flag"`
	EmailAuthGroupFlag bool `json:"email_auth_group_flag"`
}

func GetUserEmailAuth(id, email string) (emailAuth models.UserEmailAuth) {
	emailAuth.UserId = id
	emailAuth.Email = email
	emailAuth.EmailAuthConfirm = false
	emailAuth.EmailAuthStore = ""
	return emailAuth
}

func (msg *UserRegisterMessage)Convert() (user models.User, emailAuthList []models.UserEmailAuth) {
	// user
	user.Name = msg.Name
	user.Email = msg.Email
	user.UserId = msg.Id
	user.Password = msg.Password
	user.GroupEmailAuth = msg.EmailAuthGroupFlag
	user.EmailAuth = msg.EmailAuthFlag

	// email auth
	if user.GroupEmailAuth {
		for _, email := range msg.EmailAuthGroupList {
			emailAuth := GetUserEmailAuth(msg.Id, email)
			emailAuthList = append(emailAuthList, emailAuth)
		}
	} else if user.EmailAuth {
		emailAuth := GetUserEmailAuth(msg.Id, msg.Email)
		emailAuthList = append(emailAuthList, emailAuth)
	}
	return user, emailAuthList
}

