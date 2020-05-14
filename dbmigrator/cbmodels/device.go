package cbmodels

import "time"

type CbDeviceCommon struct {
	CbDeviceID int    `gorm:"primary_key;column:wr_id;not null"`
	WrNum      int    `gorm:"column:wr_num;not null;default:0"`
	WrReply    string `gorm:"type:varchar(10);column:wr_reply;not null;default:''"`
	WrParent   int    `gorm:"column:wr_parent;not null;default:0"`
	// type tinyint(4)
	WrIsComment    int    `gorm:"type:tinyint(4);column:wr_is_comment;not null;default:0"`
	WrComment      int    `gorm:"column:wr_comment;not null;default:0"`
	WrCommentReply string `gorm:"type:varchar(5);column:wr_comment_reply;not null;default:''"`
	CaName         string `gorm:"type:varchar(255);column:ca_name;not null;default:''"`
	// type set('html1','html2','secret','mail')
	WrOption  string `gorm:"type:set('html1','html2','secret','mail');column:wr_option;not null;default:''"`
	WrSubject string `gorm:"type:varchar(255);column:wr_subject;not null;default:''"`
	// type text
	WrContent string `gorm:"type:text;column:wr_content;not null"`
	// type text
	WrLink1 string `gorm:"type:text;column:wr_link1"`
	// type text
	WrLink2     string `gorm:"type:text;column:wr_link2"`
	WrLink1Hit  string `gorm:"column:wr_link1_hit;not null;default:0"`
	WrLink2Hit  int    `gorm:"column:wr_link2_hit;not null;default:0"`
	WrTrackback string `gorm:"type:varchar(255);column:wr_trackback;not null;default:''"`
	WrHit       int    `gorm:"column:wr_hit;not null;default:0"`
	WrGood      int    `gorm:"column:wr_good;not null;default:0"`
	WrNogood    int    `gorm:"column:wr_nogood;not null;default:0"`
	MbId        string `gorm:"type:varchar(255);column:mb_id;not null;default:''"`
	WrPassword  string `gorm:"type:varchar(255);column:wr_password;not null;default:''"`
	WrName      string `gorm:"type:varchar(255);column:wr_name;not null;default:''"`
	WrEmail     string `gorm:"type:varchar(255);column:wr_email;not null;default:''"`
	WrHomepage  string `gorm:"type:varchar(255);column:wr_homepage;not null;default:''"`
	// type datetime
	WrDatetime time.Time `gorm:"column:wr_datetime;not null"`
	WrLast     string    `gorm:"type:varchar(255);column:wr_last;not null;default:''"`
}

type CbDevice struct {
	CbDeviceCommon
	WrIp string `gorm:"type:varchar(255);column:wr_ip;not null;default:''"`
	Wr1  string `gorm:"type:varchar(255);column:wr_1;not null;default:''"`
	Wr2  string `gorm:"type:varchar(255);column:wr_2;not null;default:''"`
	Wr3  string `gorm:"type:varchar(255);column:wr_3;not null;default:''"`
	Wr4  string `gorm:"type:varchar(255);column:wr_4;not null;default:''"`
	Wr5  string `gorm:"type:varchar(255);column:wr_5;not null;default:''"`
	Wr6  string `gorm:"type:varchar(255);column:wr_6;not null;default:''"`
	Wr7  string `gorm:"type:varchar(255);column:wr_7;not null;default:''"`
	Wr8  string `gorm:"type:varchar(255);column:wr_8;not null;default:''"`
	Wr9  string `gorm:"type:varchar(255);column:wr_9;not null;default:''"`
	Wr10 string `gorm:"type:varchar(255);column:wr_10;not null;default:''"`
	Wr11 string `gorm:"type:varchar(255);column:wr_11;not null;default:''"`
	Wr12 string `gorm:"type:varchar(255);column:wr_12;not null;default:''"`
	Wr13 string `gorm:"type:varchar(255);column:wr_13;not null;default:''"`
	Wr14 string `gorm:"type:varchar(255);column:wr_14;not null;default:''"`
	Wr15 string `gorm:"type:varchar(255);column:wr_15;not null;default:''"`
}

func (CbDevice) TableName() string {
	return "g4_write_ndevice"
}

type ServerDevice struct {
	CbDevice
}

func (ServerDevice) TableName() string {
	return "g4_write_device"
}

type NetworkDevice struct {
	CbDevice
}

func (NetworkDevice) TableName() string {
	return "g4_write_ndevice"
}

type PartDevice struct {
	CbDevice
}

func (PartDevice) TableName() string {
	return "g4_write_pdevice"
}

