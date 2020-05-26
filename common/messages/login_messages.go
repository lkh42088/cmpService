package messages

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

