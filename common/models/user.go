package models

import (
	"cmpService/common/lib"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Idx				uint		`gorm:"primary_key;column:user_idx;auto_increment;comment:'INDEX'" json:"type:uint"`
	UserId 			string		`gorm:"unique;type:varchar(50);column:user_id;comment:'유저 ID'"`
	Password		string		`gorm:"type:varchar(255);not null;column:user_password;comment:'패스워드'"`
	Name			string		`gorm:"type:varchar(255);not null;column:user_name;comment:'유저 이름'"`
	CompanyIdx		int			`gorm:"type:int(11);column:cp_idx;comment:'회사 고유값'" json:"type:int"`
	Email			string		`gorm:"type:varchar(255);column:user_email;comment:'이메일'"`
	AuthLevel		int			`gorm:"type:int(11);default:10;column:user_auth_level;comment:'회원 권한 등급'" json:"type:int"`
	Tel				string		`gorm:"type:varchar(15);column:user_tel;comment:'전화 번호'"`
	HP				string		`gorm:"type:varchar(15);column:user_hp;comment:'핸드폰 번호'"`
	Zipcode			string		`gorm:"type:varchar(15);column:user_zip;comment:'우편 번호'"`
	Address			string		`gorm:"type:varchar(255);column:user_addr;comment:'주소'"`
	AddressDetail	string		`gorm:"type:varchar(255);column:user_addr_detail;comment:'상세 주소'"`
	TermDate		time.Time	`gorm:"type:datetime;column:user_termination_date;comment:'퇴직 일자'" json:"type:datetime"`
	BlockDate		time.Time	`gorm:"type:datetime;column:user_block_date;comment:'접근 차단 일자'" json:"type:datetime"`
	Memo			string		`gorm:"type:text;column:user_memo;comment:'메모'"`
	WorkScope		string		`gorm:"type:varchar(15);column:user_work_scope;comment:'업무 구분'"`
	Department		string		`gorm:"type:varchar(15);column:user_department;comment:'부서'"`
	Position		string		`gorm:"type:varchar(15);column:user_position;comment:'직급'"`
	EmailAuth 		bool 		`gorm:"type:tinyint(1);default:0;column:user_email_auth_flag;comment:'개인 이메일 인증'" json:"type:bool"`
	GroupEmailAuth 	bool 		`gorm:"type:tinyint(1);default:0;column:user_group_email_auth_flag;comment:'그룹 이메일 인증'" json:"type:bool"`
	RegisterDate	time.Time	`gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:user_register_date;comment:'등록일'" json:"type:datetime"`
	LastAccessDate	time.Time	`gorm:"type:datetime;column:user_last_access_date;comment:'최근 로그인 날짜'" json:"type:datetime"`
	LastAccessIp	string		`gorm:"type:varchar(15);column:user_last_access_ip;comment:'최근 접속 IP'"`
}

func (User) TableName() string {
	return "user_tb"
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

func (UserEmailAuth) TableName() string {
	return "user_email_auth_tb"
}

type Company struct {
	Idx				uint		`gorm:"primary_key;column:cp_idx;auto_increment;comment:'INDEX'" json:"type:uint"`
	Name			string		`gorm:"type:varchar(255);not null;column:cp_name;comment:'회사명'"`
	Email			string		`gorm:"type:varchar(255);column:cp_email;comment:'이메일'"`
	Homepage		string		`gorm:"type:varchar(255);column:cp_homepage;comment:'홈페이지'"`
	Tel				string		`gorm:"type:varchar(15);column:cp_tel;comment:'전화 번호'"`
	HP				string		`gorm:"type:varchar(15);column:cp_hp;comment:'핸드폰 번호'"`
	Zipcode			string		`gorm:"type:varchar(15);not null;column:cp_zip;comment:'우편 번호'"`
	Address			string		`gorm:"type:varchar(255);not null;column:cp_addr;comment:'주소'"`
	AddressDetail	string		`gorm:"type:varchar(255);not null;column:cp_addr_detail;comment:'상세 주소'"`
	TermDate		time.Time	`gorm:"type:datetime;column:cp_termination_date;comment:'해지 일자'" json:"type:datetime"`
	IsCompany		bool 		`gorm:"type:tinyint(1);column:cp_is_company;default:1;comment:'회사 여부'" json:"type:bool"`
	Memo			string		`gorm:"type:text;column:cp_memo;comment:'메모'"`
}

func (Company) TableName() string {
	return "company_tb"
}

type Auth struct {
	Level 			int			`gorm:"primary_key;type:int(11);column:auth_level;comment:'권한 레벨'" json:"type:int"`
	Tag 			string		`gorm:"type:varchar(255);column:auth_tag;comment:'권한 이름'"`
	Login			int			`gorm:"type:int(11);column:auth_login;default:0;comment:'로그인 권한'" json:"type:int"`
	Device 			int			`gorm:"type:int(11);column:auth_device;default:0;comment:'장비 권한'" json:"type:int"`
	Stat 			int			`gorm:"type:int(11);column:auth_stat;default:0;comment:'통계 권한'" json:"type:int"`
	Monitor 		int			`gorm:"type:int(11);column:auth_monitor;default:0;comment:'모니터링 권한'" json:"type:int"`
	Billing 		int			`gorm:"type:int(11);column:auth_billing;default:0;comment:'빌링 권한'" json:"type:int"`
	Cloud 			int			`gorm:"type:int(11);column:auth_cloud;default:0;comment:'Cloud 권한'" json:"type:int"`
	Board 			int			`gorm:"type:int(11);column:auth_board;default:0;comment:'게시판 권한'" json:"type:int"`
	Manager			int			`gorm:"type:int(11);column:auth_manager;default:0;comment:'담당자 권한'" json:"type:int"`
	WorkBoard		int			`gorm:"type:int(11);column:auth_work_board;default:0;comment:'작업 권한'" json:"type:int"`
	TempStart		time.Time	`gorm:"type:datetime;column:auth_temp_start;comment:'임시 권한 시작일'" json:"type:datetime"`
	TempEnd			time.Time	`gorm:"type:datetime;column:auth_temp_end;comment:'임시 권한 만료일'" json:"type:datetime"`
}

func (Auth) TableName() string {
	return "auth_tb"
}

type CompanyResponse struct {
	Company
	UserId 			string		`gorm:"unique;type:varchar(50);column:user_id" json:"userId"`
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


