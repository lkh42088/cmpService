package models

import (
	"cmpService/common/lib"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
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

type UserMember struct {
	Idx				uint		`gorm:"primary_key;column:idx;auto_increment;comment:'INDEX'"`
	UserId 			string		`gorm:"primary_key;type:varchar(50);column:user_id;comment:'유저 ID'"`
	Password		string		`gorm:"type:varchar(255);column:user_password;comment:'패스워드'"`
	Name			string		`gorm:"type:varchar(255);column:user_name;comment:'유저 이름'"`
	Company			string		`gorm:"type:varchar(255);column:user_company;comment:'유저 회사명'"`
	Email			string		`gorm:"type:varchar(255);column:user_email;comment:'이메일'"`
	Homepage		string		`gorm:"type:varchar(255);column:user_homepage;comment:'홈페이지'"`
	AuthLevel		int			`gorm:"type:int(11);column:user_auth_level;comment:'회원 권한 등급'"`
	Tel				string		`gorm:"type:varchar(15);column:user_tel;comment:'전화 번호'"`
	HP				string		`gorm:"type:varchar(15);column:user_hp;comment:'핸드폰 번호'"`
	Zipcode			string		`gorm:"type:varchar(15);column:user_zip;comment:'우편 번호'"`
	Address			string		`gorm:"type:varchar(255);column:user_addr;comment:'주소'"`
	AddressDetail	string		`gorm:"type:varchar(255);column:user_addr_detail;comment:'상세 주소'"`
	IP				string		`gorm:"type:varchar(15);column:user_ip;comment:'IP'"`
	TermDate		time.Time	`gorm:"type:datetime;column:user_termination_date;comment:'퇴직 일자'"`
	BlockDate		time.Time	`gorm:"type:datetime;column:user_block_date;comment:'접근 차단 일자'"`
	Memo			string		`gorm:"type:varchar(255);column:user_memo;comment:'메모'"`
	AccumulateStats	bool		`gorm:"type:tinyint;column:user_accumulate_stats_flag;comment:'누적 통계 여부'"`
	WorkScope		string		`gorm:"type:varchar(15);column:user_work_scope;comment:'업무 구분'"`
	Department		string		`gorm:"type:varchar(15);column:user_department;comment:'부서'"`
	Position		string		`gorm:"type:varchar(15);column:user_position;comment:'직급'"`
	RegisterDate	time.Time	`gorm:"type:datetime;column:user_register_date;comment:'등록일'"`
	LastAccessDate	time.Time	`gorm:"type:datetime;column:user_last_access_date;comment:'최근 로그인 날짜'"`
	LastAccessIp	string		`gorm:"type:varchar(15);column:user_last_access_ip;comment:'최근 접속 IP'"`
}

func (UserMember) TableName() string {
	return "user_member_tb"
}

type Customer struct {
	Idx				uint		`gorm:"primary_key;column:idx;auto_increment;comment:'INDEX'"`
	UserId 			string		`gorm:"primary_key;type:varchar(50);column:user_id;comment:'회사 ID'"`
	Password		string		`gorm:"type:varchar(255);not null;column:user_password;comment:'패스워드'"`
	Company			string		`gorm:"type:varchar(255);not null;column:user_company;comment:'회사명'"`
	Email			string		`gorm:"type:varchar(255);column:user_email;comment:'이메일'"`
	Homepage		string		`gorm:"type:varchar(255);column:user_homepage;comment:'홈페이지'"`
	AuthLevel		int			`gorm:"type:int(11);column:user_auth_level;default:0;comment:'고객 권한 등급'"`
	Tel				string		`gorm:"type:varchar(15);column:user_tel;comment:'전화 번호'"`
	HP				string		`gorm:"type:varchar(15);column:user_hp;comment:'핸드폰 번호'"`
	Zipcode			string		`gorm:"type:varchar(15);not null;column:user_zip;comment:'우편 번호'"`
	Address			string		`gorm:"type:varchar(255);not null;column:user_addr;comment:'주소'"`
	AddressDetail	string		`gorm:"type:varchar(255);not null;column:user_addr_detail;comment:'상세 주소'"`
	IP				string		`gorm:"type:varchar(15);column:user_ip;comment:'IP'"`
	TermDate		time.Time	`gorm:"type:datetime;column:user_termination_date;comment:'해지 일자'"`
	BlockDate		time.Time	`gorm:"type:datetime;column:user_block_date;comment:'접근 차단 일자'"`
	Memo			string		`gorm:"type:varchar(255);column:user_memo;comment:'메모'"`
	AccumulateStats	bool		`gorm:"type:tinyint;column:user_accumulate_stats_flag;default:0;comment:'누적 통계 여부'"`
	RegisterDate	time.Time	`gorm:"type:datetime;column:user_register_date;default:CURRENT_TIMESTAMP;comment:'등록일'"`
	LastAccessDate	time.Time	`gorm:"type:datetime;column:user_last_access_date;comment:'최근 로그인 날짜'"`
	LastAccessIp	string		`gorm:"type:varchar(15);column:user_last_access_ip;comment:'최근 접속 IP'"`
}

func (Customer) TableName() string {
	return "customer_tb"
}

type Auth struct {
	Level 			int			`gorm:"primary_key;type:int(11);column:auth_level;comment:'권한 레벨'"`
	Tag 			string		`gorm:"type:varchar(255);column:auth_tag;comment:'권한 이름'"`
	Login			int			`gorm:"type:int(11);column:auth_login;default:0;comment:'로그인 권한'"`
	Device 			int			`gorm:"type:int(11);column:auth_device;default:0;comment:'장비 권한'"`
	Stat 			int			`gorm:"type:int(11);column:auth_stat;default:0;comment:'통계 권한'"`
	Monitor 		int			`gorm:"type:int(11);column:auth_monitor;default:0;comment:'모니터링 권한'"`
	Billing 		int			`gorm:"type:int(11);column:auth_billing;default:0;comment:'빌링 권한'"`
	Cloud 			int			`gorm:"type:int(11);column:auth_cloud;default:0;comment:'Cloud 권한'"`
	Board 			int			`gorm:"type:int(11);column:auth_board;default:0;comment:'게시판 권한'"`
	Manager			int			`gorm:"type:int(11);column:auth_manager;default:0;comment:'담당자 권한'"`
	WorkBoard		int			`gorm:"type:int(11);column:auth_work_board;default:0;comment:'작업 권한'"`
	TempStart		time.Time	`gorm:"type:datetime;column:auth_temp_start;comment:'임시 권한 시작일'"`
	TempEnd			time.Time	`gorm:"type:datetime;column:auth_temp_end;comment:'임시 권한 만료일'"`
}

func (Auth) TableName() string {
	return "auth_tb"
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

