package models

import (
	"cmpService/common/lib"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

const (
	UserOrmIdx              = "user_idx"
	UserOrmUserId           = "user_id"
	UserOrmPassword         = "user_password"
	UserOrmName             = "user_name"
	UserOrmIsCompanyAccount = "user_is_cp_account"
	UserOrmCompanyIdx       = "cp_idx"
	UserOrmEmail            = "user_email"
	UserOrmAuthLevel        = "user_auth_level"
	UserOrmTel              = "user_tel"
	UserOrmHP               = "user_hp"
	UserOrmZipcode          = "user_zip" // 5
	UserOrmAddress          = "user_addr"
	UserOrmAddressDetail    = "user_addr_detail"
	UserOrmTermDate         = "user_termination_date"
	UserOrmBlockDate        = "user_block_date"
	UserOrmMemo             = "user_memo" // 10
	UserOrmWorkScope        = "user_work_scope"
	UserOrmDepartment       = "user_department"
	UserOrmPosition         = "user_position"
	UserOrmEmailAuth        = "user_email_auth_flag"
	UserOrmGroupEmailAuth   = "user_group_email_auth_flag" // 15
	UserOrmRegisterDate     = "user_register_date"
	UserOrmLastAccessDate   = "user_last_access_date"
	UserOrmLastAccessIp     = "user_last_access_ip"
)

const (
	UserJsonIdx              = "idx"
	UserJsonUserId           = "userId"
	UserJsonPassword         = "password"
	UserJsonName             = "name"
	UserJsonIsCompanyAccount = "isCompanyAccount"
	UserJsonCompanyIdx       = "companyIdx" // 5
	UserJsonEmail            = "email"
	UserJsonAuthLevel        = "authLevel"
	UserJsonTel              = "tel"
	UserJsonHP               = "hp"
	UserJsonZipcode          = "zipcode" // 10
	UserJsonAddress          = "address"
	UserJsonAddressDetail    = "addressDetail"
	UserJsonTermDate         = "termDate"
	UserJsonBlockDate        = "blockDate"
	UserJsonMemo             = "memo" // 15
	UserJsonWorkScope        = "workScope"
	UserJsonDepartment       = "department"
	UserJsonPosition         = "position"
	UserJsonEmailAuth        = "emailAuth"
	UserJsonGroupEmailAuth   = "groupEmailAuth" //20
	UserJsonRegisterDate     = "registerDate"
	UserJsonLastAccessDate   = "lastAccessDate"
	UserJsonLastAccessIp     = "lastAccessIp"
)

var UserOrmMap = map[string]string{
	UserOrmIdx:              UserJsonIdx,
	UserOrmUserId:           UserJsonUserId,
	UserOrmPassword:         UserJsonPassword,
	UserOrmName:             UserJsonName,
	UserOrmIsCompanyAccount: UserJsonIsCompanyAccount,
	UserOrmCompanyIdx:       UserJsonCompanyIdx, /* 5*/
	UserOrmEmail:            UserJsonEmail,
	UserOrmAuthLevel:        UserJsonAuthLevel,
	UserOrmTel:              UserJsonTel,
	UserOrmHP:               UserJsonHP,
	UserOrmZipcode:          UserJsonZipcode, /* 10 */
	UserOrmAddress:          UserJsonAddress,
	UserOrmAddressDetail:    UserJsonAddressDetail,
	UserOrmTermDate:         UserJsonTermDate,
	UserOrmBlockDate:        UserJsonBlockDate,
	UserOrmMemo:             UserJsonMemo, /* 15 */
	UserOrmWorkScope:        UserJsonWorkScope,
	UserOrmDepartment:       UserJsonDepartment,
	UserOrmPosition:         UserJsonPosition,
	UserOrmEmailAuth:        UserJsonEmailAuth,
	UserOrmGroupEmailAuth:   UserJsonGroupEmailAuth, /* 20 */
	UserOrmRegisterDate:     UserJsonRegisterDate,
	UserOrmLastAccessDate:   UserJsonLastAccessDate,
	UserOrmLastAccessIp:     UserJsonLastAccessIp,
}

var UserJsonMap = map[string]string{
	UserJsonIdx:              UserOrmIdx,
	UserJsonUserId:           UserOrmUserId,
	UserJsonPassword:         UserOrmPassword,
	UserJsonName:             UserOrmName,
	UserJsonIsCompanyAccount: UserOrmIsCompanyAccount,
	UserJsonCompanyIdx:       UserOrmCompanyIdx, // 5
	UserJsonEmail:            UserOrmEmail,
	UserJsonAuthLevel:        UserOrmAuthLevel,
	UserJsonTel:              UserOrmTel,
	UserJsonHP:               UserOrmHP,
	UserJsonZipcode:          UserOrmZipcode, // 10
	UserJsonAddress:          UserOrmAddress,
	UserJsonAddressDetail:    UserOrmAddressDetail,
	UserJsonTermDate:         UserOrmTermDate,
	UserJsonBlockDate:        UserOrmBlockDate,
	UserJsonMemo:             UserOrmMemo, // 15
	UserJsonWorkScope:        UserOrmWorkScope,
	UserJsonDepartment:       UserOrmDepartment,
	UserJsonPosition:         UserOrmPosition,
	UserJsonEmailAuth:        UserOrmEmailAuth,
	UserJsonGroupEmailAuth:   UserOrmGroupEmailAuth, // 20
	UserJsonRegisterDate:     UserOrmRegisterDate,
	UserJsonLastAccessDate:   UserOrmLastAccessDate,
	UserJsonLastAccessIp:     UserOrmLastAccessIp,
}

type User struct {
	Idx              uint      `gorm:"primary_key;column:user_idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	UserId           string    `gorm:"unique;type:varchar(50);column:user_id;comment:'유저 ID'" json:"userId"`
	Password         string    `gorm:"type:varchar(255);not null;column:user_password;comment:'패스워드'" json:"password"`
	Name             string    `gorm:"type:varchar(255);not null;column:user_name;comment:'유저 이름'" json:"name"`
	IsCompanyAccount bool      `gorm:"type:tinyint(1);default:0;column:user_is_cp_account;comment:'회사 대표계정 여부'" json:"isCompanyAccount"`
	CompanyIdx       int       `gorm:"type:int(11);column:cp_idx;comment:'회사 고유값'" json:"companyIdx"`
	Email            string    `gorm:"type:varchar(255);column:user_email;comment:'이메일'" json:"email"`
	AuthLevel        int       `gorm:"type:int(11);default:10;column:user_auth_level;comment:'회원 권한 등급'" json:"authLevel"`
	Tel              string    `gorm:"type:varchar(15);column:user_tel;comment:'전화 번호'" json:"tel"`
	HP               string    `gorm:"type:varchar(15);column:user_hp;comment:'핸드폰 번호'" json:"hp"`
	Zipcode          string    `gorm:"type:varchar(15);column:user_zip;comment:'우편 번호'" json:"zipcode"`
	Address          string    `gorm:"type:varchar(255);column:user_addr;comment:'주소'" json:"address"`
	AddressDetail    string    `gorm:"type:varchar(255);column:user_addr_detail;comment:'상세 주소'" json:"addressDetail"`
	TermDate         time.Time `gorm:"type:datetime;column:user_termination_date;comment:'퇴직 일자'" json:"termDate"`
	BlockDate        time.Time `gorm:"type:datetime;column:user_block_date;comment:'접근 차단 일자'" json:"blockDate"`
	Memo             string    `gorm:"type:text;column:user_memo;comment:'메모'" json:"memo"`
	WorkScope        string    `gorm:"type:varchar(15);column:user_work_scope;comment:'업무 구분'" json:"workScope"`
	Department       string    `gorm:"type:varchar(15);column:user_department;comment:'부서'" json:"department"`
	Position         string    `gorm:"type:varchar(15);column:user_position;comment:'직급'" json:"position"`
	EmailAuth        bool      `gorm:"type:tinyint(1);default:0;column:user_email_auth_flag;comment:'개인 이메일 인증'" json:"emailAuth"`
	GroupEmailAuth   bool      `gorm:"type:tinyint(1);default:0;column:user_group_email_auth_flag;comment:'그룹 이메일 인증'" json:"groupEmailAuth"`
	Avata            []byte    `gorm:"type:blob;column:user_avata;comment:'아바타 데이터'" json:"avata"`
	RegisterDate     time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:user_register_date;comment:'등록일'" json:"registerDate"`
	LastAccessDate   time.Time `gorm:"type:datetime;column:user_last_access_date;comment:'최근 로그인 날짜'" json:"lastAccessDate"`
	LastAccessIp     string    `gorm:"type:varchar(15);column:user_last_access_ip;comment:'최근 접속 IP'" json:"lastAccessIp"`
}

type UserDetail struct {
	User
	CompanyName string `gorm:"type:varchar(50);column:cp_name" json:"cpName"`
}

func (User) TableName() string {
	return "user_tb"
}

type UserEmailAuth struct {
	UserEmailAuthID  uint   `gorm:"primary_key;column:uea_idx;not null"`
	User             User   `gorm:"foreignkey:Idx"`
	UserIdx          uint   `gorm:"column:user_idx"`
	UserId           string `gorm:"type:varchar(32);column:uea_user_id"`
	Email            string `gorm:"type:varchar(64);column:uea_email"`
	EmailAuthConfirm bool   `gorm:"column:uea_confirm" json:"-"`
	EmailAuthStore   string `gorm:"type:varchar(255);column:uea_store" json:"-"`
}

func (UserEmailAuth) TableName() string {
	return "user_email_auth_tb"
}

type LoginAuth struct {
	LoginAuthID      int    `gorm:"primary_key;column:la_idx;not null"`
	User             User   `gorm:"foreignkey:Idx"`
	UserIdx          uint   `gorm:"column:user_idx"`
	UserId           string `gorm:"type:varchar(32);column:la_user_id"`
	AuthUserId       string `gorm:"type:varchar(32);column:la_auth_user_id"`
	AuthEmail        string `gorm:"type:varchar(64);column:la_auth_email"`
	EmailAuthConfirm bool   `gorm:"column:la_email_auth_confirm" json:"-"`
}

func (LoginAuth) TableName() string {
	return "login_auth_tb"
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

type UserPage struct {
	Page  Pagination   `json:"page"`
	Users []UserDetail `json:"data"`
}

func (u UserPage) GetOrderBy(orderby, order string) string {
	val, exists := UserJsonMap[orderby]
	if !exists {
		val = "user_idx"
	}
	order = strings.ToLower(order)
	if !(order == "asc" || order == "desc") {
		order = "desc"
	}
	return val + " " + order
}

type UserEmailAuthMsg struct {
	Id     string `json:"id"`
	Email  string `json:"email"`
	Secret string `json:"secret"`
}
