package messages

import (
	"cmpService/common/models"
	"encoding/json"
	"fmt"
)

type UserLoginMessage struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Result   int    `json:"result"`
	Comment  string `json:"comment"`
}

type DeleteDataMessage struct {
	IdxList []int `json:"idx"`
	AvataList []string `json:"avata"`
}

type UserRegisterMessage struct {
	CpName             string              `json:"cpName"`
	CpIdx              int                 `json:"cpIdx"`
	IsCompanyAccount   bool                `json:"isCompanyAccount"`
	Id                 string              `json:"id"`
	Password           string              `json:"password"`
	Email              string              `json:"email"`
	Name               string              `json:"name"`
	Tel                string              `json:"tel"`
	HP                 string              `json:"hp"`
	AuthLevel          int                 `json:"authLevel"`
	ZipCode            string              `json:"zipCode"`
	Address            string              `json:"address"`
	AddressDetail      string              `json:"addressDetail"`
	EmailAuthFlag      bool                `json:"emailAuthFlag"`
	EmailAuthGroupFlag bool                `json:"emailAuthGroupFlag"`
	EmailAuthGroupList []models.UserDetail `json:"emailAuthGroupList"`
	Memo               string              `json:"memo"`
	Avata              string              `json:"avata"`
}

func (u UserRegisterMessage) String() {
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("Msg: %s\n", data)
	for i, entry := range u.EmailAuthGroupList {
		data, _ := json.Marshal(entry)
		fmt.Printf("%d entry %s\n", i, data)
	}
}

type UserRegisterMessageBackup struct {
	CpName             string           `json:"cpName"`
	CpIdx              int              `json:"cpIdx"`
	Id                 string           `json:"id"`
	Password           string           `json:"password"`
	Email              string           `json:"email"`
	Name               string           `json:"name"`
	EmailAuthFlag      bool             `json:"emailAuthFlag"`
	EmailAuthGroupFlag bool             `json:"emailAuthGroupFlag"`
	EmailAuthGroupList []EmailAuthEntry `json:"emailAuthGroupList"`
}

type EmailAuthEntry struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type UserInfo struct {
	Id                 string `json:"id"`
	Password           string `json:"password"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	Level              int    `json:"level"`
	CpName             string `json:"cpName"`
	EmailAuthFlag      bool   `json:"emailAuthFlag"`
	EmailAuthGroupFlag bool   `json:"emailAuthGroupFlag"`
	AuthEmail          string `json:"authEmail"`
	Avata              string              `json:"avata"`
}

func GetUserEmailAuth(id, email string) (emailAuth models.UserEmailAuth) {
	emailAuth.UserId = id
	emailAuth.Email = email
	emailAuth.EmailAuthConfirm = false
	emailAuth.EmailAuthStore = ""
	return emailAuth
}

func (msg *UserRegisterMessage) Convert() (user models.User, emailAuthList []models.UserEmailAuth) {
	// user
	user.UserId = msg.Id
	user.Password = msg.Password
	user.Name = msg.Name
	user.Email = msg.Email
	user.HP = msg.HP
	user.AuthLevel = msg.AuthLevel
	user.Zipcode = msg.ZipCode
	user.Address = msg.Address
	user.AddressDetail = msg.AddressDetail
	user.GroupEmailAuth = msg.EmailAuthGroupFlag
	user.EmailAuth = msg.EmailAuthFlag

	// email auth
	if user.GroupEmailAuth {
		for _, email := range msg.EmailAuthGroupList {
			emailAuth := GetUserEmailAuth(msg.Id, email.Email)
			emailAuthList = append(emailAuthList, emailAuth)
		}
	} else if user.EmailAuth {
		emailAuth := GetUserEmailAuth(msg.Id, msg.Email)
		emailAuthList = append(emailAuthList, emailAuth)
	}
	return user, emailAuthList
}

func (msg *UserRegisterMessage) Translate() (user models.User, emailAuthList []models.LoginAuth) {
	// user
	user.Name = msg.Name
	user.Email = msg.Email
	user.UserId = msg.Id
	user.Password = msg.Password
	user.GroupEmailAuth = msg.EmailAuthGroupFlag
	user.EmailAuth = msg.EmailAuthFlag
	user.IsCompanyAccount = msg.IsCompanyAccount
	user.HP = msg.HP
	user.Tel = msg.Tel
	user.AuthLevel = msg.AuthLevel
	user.Zipcode = msg.ZipCode
	user.Address = msg.Address
	user.AddressDetail = msg.AddressDetail
	user.GroupEmailAuth = msg.EmailAuthGroupFlag
	user.EmailAuth = msg.EmailAuthFlag
	user.Memo = msg.Memo
	user.Avata = msg.Avata

	// email auth
	if user.GroupEmailAuth {
		for _, entry := range msg.EmailAuthGroupList {
			var loginAuth models.LoginAuth
			loginAuth.UserId = user.UserId
			loginAuth.AuthUserId = entry.UserId
			loginAuth.AuthEmail = entry.Email
			emailAuthList = append(emailAuthList, loginAuth)
		}
	} else if user.EmailAuth {
		var loginAuth models.LoginAuth
		loginAuth.UserId = user.UserId
		loginAuth.AuthUserId = user.UserId
		loginAuth.AuthEmail = user.Email
		emailAuthList = append(emailAuthList, loginAuth)
	}
	return user, emailAuthList
}
