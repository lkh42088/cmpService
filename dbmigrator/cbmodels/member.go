package cbmodels

import (
	"time"
)

type CbMember struct {
	No				int			`gorm:"column:mb_no"`
	Id				string		`gorm:"column:mb_id"`
	Password		string		`gorm:"column:mb_password"`
	Name			string		`gorm:"column:mb_name"`
	Nick 			string		`gorm:"column:mb_nick"`
	NickDate		time.Time	`gorm:"column:mb_nick_date"`
	Email 			string		`gorm:"column:mb_email"`
	Homepage		string		`gorm:"column:mb_homepage"`
	PasswordQ		string		`gorm:"column:mb_password_q"`
	PasswordA		string		`gorm:"column:mb_password_a"`
	Level 			int			`gorm:"column:mb_level"`
	Jumin			string		`gorm:"column:mb_jumin"`
	Sex 			string		`gorm:"column:mb_sex"`
	Birth 			string		`gorm:"column:mb_birth"`
	Tel				string		`gorm:"column:mb_tel"`
	HP 				string		`gorm:"column:mb_hp"`
	ZIP1			string		`gorm:"column:mb_zip1"`
	ZIP2 			string		`gorm:"column:mb_zip2"`
	Addr1			string		`gorm:"column:mb_addr1"`
	Addr2			string		`gorm:"column:mb_addr2"`
	Signature 		string		`gorm:"column:mb_signature"`
	Recommend 		string		`gorm:"column:mb_recommend"`
	Point 			int			`gorm:"column:mb_point"`
	TodayLogin		time.Time	`gorm:"column:mb_today_login"`
	LoginIp			string		`gorm:"column:mb_login_ip"`
	Datetime 		time.Time	`gorm:"column:mb_datetime"`
	IP 				string		`gorm:"column:mb_ip"`
	LeaveDate		string		`gorm:"column:mb_leave_date"`
	InterceptDate	string		`gorm:"column:mb_intercept_date"`
	EmailCertify	time.Time	`gorm:"column:mb_email_certify"`
	Memo 			string		`gorm:"column:mb_memo"`
	LostCertify		string		`gorm:"column:mb_lost_certify"`
	Mailling		int			`gorm:"column:mb_mailling"`
	SMS 			int			`gorm:"column:mb_sms"`
	Open 			int			`gorm:"column:mb_open"`
	OpenDate 		time.Time	`gorm:"column:mb_open_date"`
	Profile 		string		`gorm:"column:mb_profile"`
	MemoCall 		string		`gorm:"column:mb_memo_call"`
	Mb1				string		`gorm:"column:mb_1"`
	Mb2				string		`gorm:"column:mb_2"`
	Mb3				string		`gorm:"column:mb_3"`
	Mb4				string		`gorm:"column:mb_4"`
	Mb5				string		`gorm:"column:mb_5"`
	Mb6				string		`gorm:"column:mb_6"`
	Mb7				string		`gorm:"column:mb_7"`
	Mb8				string		`gorm:"column:mb_8"`
	Mb9				string		`gorm:"column:mb_9"`
	Mb10			string		`gorm:"column:mb_10"`
}

func (CbMember) TableName() string {
	return "g4_member"
}