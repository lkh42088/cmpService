package models

import "time"

type Auth struct {
	Level     int       `gorm:"primary_key;type:int(11);column:auth_level;comment:'권한 레벨'" json:"level"`
	Tag       string    `gorm:"type:varchar(255);column:auth_tag;comment:'권한 이름'" json:"tag"`
	Login     int       `gorm:"type:int(11);column:auth_login;default:0;comment:'로그인 권한'" json:"login"`
	Device    int       `gorm:"type:int(11);column:auth_device;default:0;comment:'장비 권한'" json:"device"`
	Platform  int       `gorm:"type:int(11);column:auth_platform;default:0;comment:'플랫폼 권한'" json:"platform"`
	Stat      int       `gorm:"type:int(11);column:auth_stat;default:0;comment:'통계 권한'" json:"stat"`
	Monitor   int       `gorm:"type:int(11);column:auth_monitor;default:0;comment:'모니터링 권한'" json:"monitor"`
	Billing   int       `gorm:"type:int(11);column:auth_billing;default:0;comment:'빌링 권한'" json:"billing"`
	Cloud     int       `gorm:"type:int(11);column:auth_cloud;default:0;comment:'Cloud 권한'" json:"cloud"`
	Board     int       `gorm:"type:int(11);column:auth_board;default:0;comment:'게시판 권한'" json:"board"`
	Manager   int       `gorm:"type:int(11);column:auth_manager;default:0;comment:'담당자 권한'" json:"manager"`
	WorkBoard int       `gorm:"type:int(11);column:auth_work_board;default:0;comment:'작업 권한'" json:"workBoard"`
	TempStart time.Time `gorm:"type:datetime;column:auth_temp_start;comment:'임시 권한 시작일'" json:"tempStart"`
	TempEnd   time.Time `gorm:"type:datetime;column:auth_temp_end;comment:'임시 권한 만료일'" json:"tempEnd"`
}

func (Auth) TableName() string {
	return "auth_tb"
}
