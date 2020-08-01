package insert

import (
	"cmpService/common/models"
	"time"
)

var newSubCodeData = []models.SubCode{
	// SubCode Table : Rack 없음 subcode 추가
	{ID: 0, Code: models.Code{}, CodeID: 15, Name: "랙 없음", Order: 128},
	{0, models.Code{}, 16, "랙 없음", 189},
	{0, models.Code{}, 17, "랙 없음", 2},
	{0, models.Code{}, 18, "랙 없음", 2},
	{0, models.Code{}, 102, "랙 없음", 4},
	{0, models.Code{}, 438, "랙 없음", 3},
	{0, models.Code{}, 536, "랙 없음", 1},
	{0, models.Code{}, 544, "랙 없음", 1},
	{0, models.Code{}, 552, "랙 없음", 20},
	{0, models.Code{}, 574, "랙 없음", 2},
	{0, models.Code{}, 575, "랙 없음", 2},
	{0, models.Code{}, 613, "랙 없음", 2},
}

var newUsers = []models.User{
	// User Table : Nubes-Bridge Base User 추가
	{
		Idx:              0,
		UserId:           "bhjung",
		Password:         "nubes1510!#",
		Name:             "정병화",
		IsCompanyAccount: false,
		CompanyIdx:       0,
		Email:            "bhjung@nubes-bridge.com",
		AuthLevel:        5,
		Tel:              "02-1111-3333",
		HP:               "010-1111-3333",
		Zipcode:          "",
		Address:          "",
		AddressDetail:    "",
		TermDate:         time.Time{},
		BlockDate:        time.Time{},
		Memo:             "",
		WorkScope:        "",
		Department:       "",
		Position:         "",
		EmailAuth:        false,
		GroupEmailAuth:   false,
		Avata:            nil,
		RegisterDate:     time.Time{},
		LastAccessDate:   time.Time{},
		LastAccessIp:     "",
	},
	{
		Idx:              0,
		UserId:           "khlee",
		Password:         "nubes1510!#",
		Name:             "이경훈",
		IsCompanyAccount: false,
		CompanyIdx:       0,
		Email:            "khlee@nubes-bridge.com",
		AuthLevel:        10,
		Tel:              "02-1111-3333",
		HP:               "010-1111-3333",
		Zipcode:          "",
		Address:          "",
		AddressDetail:    "",
		TermDate:         time.Time{},
		BlockDate:        time.Time{},
		Memo:             "",
		WorkScope:        "",
		Department:       "",
		Position:         "",
		EmailAuth:        false,
		GroupEmailAuth:   false,
		Avata:            nil,
		RegisterDate:     time.Time{},
		LastAccessDate:   time.Time{},
		LastAccessIp:     "",
	},
	{
		Idx:              0,
		UserId:           "ebjee",
		Password:         "nubes1510!#",
		Name:             "지은빈",
		IsCompanyAccount: false,
		CompanyIdx:       0,
		Email:            "ebjee@nubes-bridge.com",
		AuthLevel:        10,
		Tel:              "02-1111-3333",
		HP:               "010-1111-3333",
		Zipcode:          "",
		Address:          "",
		AddressDetail:    "",
		TermDate:         time.Time{},
		BlockDate:        time.Time{},
		Memo:             "",
		WorkScope:        "",
		Department:       "",
		Position:         "",
		EmailAuth:        false,
		GroupEmailAuth:   false,
		Avata:            nil,
		RegisterDate:     time.Time{},
		LastAccessDate:   time.Time{},
		LastAccessIp:     "",
	},
	{
		Idx:              0,
		UserId:           "james",
		Password:         "nubes1510!#",
		Name:             "안종석",
		IsCompanyAccount: false,
		CompanyIdx:       0,
		Email:            "james@nubes-bridge.com",
		AuthLevel:        5,
		Tel:              "02-1111-3333",
		HP:               "010-1111-3333",
		Zipcode:          "",
		Address:          "",
		AddressDetail:    "",
		TermDate:         time.Time{},
		BlockDate:        time.Time{},
		Memo:             "",
		WorkScope:        "",
		Department:       "",
		Position:         "",
		EmailAuth:        false,
		GroupEmailAuth:   false,
		Avata:            nil,
		RegisterDate:     time.Time{},
		LastAccessDate:   time.Time{},
		LastAccessIp:     "",
	},
	// 콘텐츠 브릿지 기본 User 추가
	{
		Idx:              0,
		UserId:           "cbadmin",
		Password:         "qlenfrl!#24!#",
		Name:             "관리자",
		IsCompanyAccount: false,
		CompanyIdx:       6,
		Email:            "idc@conbridge.co.kr",
		AuthLevel:        5,
		Tel:              "02-562-0694",
		HP:               "02-562-0694",
		Zipcode:          "",
		Address:          "",
		AddressDetail:    "",
		TermDate:         time.Time{},
		BlockDate:        time.Time{},
		Memo:             "",
		WorkScope:        "",
		Department:       "",
		Position:         "",
		EmailAuth:        false,
		GroupEmailAuth:   false,
		Avata:            nil,
		RegisterDate:     time.Time{},
		LastAccessDate:   time.Time{},
		LastAccessIp:     "",
	},
}

var newCompanies = []models.Company{
	// Company Table : Nubes-Bridge 추가
	{
		Idx:           0,
		Name:          "Nubes-Bridge",
		Email:         "nubes@nubes-bridge.com",
		Homepage:      "nubes-bridge.com",
		Tel:           "02-111-2222",
		HP:            "",
		Zipcode:       "",
		Address:       "",
		AddressDetail: "",
		TermDate:      time.Time{},
		IsCompany:     false,
		UserId:        "nubes",
		Memo:          "",
	},
}

var newSubnets = []models.SubnetMgmt{
	{
		SubnetTag: "회사 내부망1", SubnetStart: "10.1.1.1", SubnetEnd: "10.1.1.128", SubnetMask: "255.255.255.128", Gateway: "10.1.1.1",
	}, {
		SubnetTag: "회사 내부망2", SubnetStart: "20.1.1.1", SubnetEnd: "20.1.1.128", SubnetMask: "255.255.255.128", Gateway: "10.1.1.1",
	}, {
		SubnetTag: "회사 내부망3", SubnetStart: "30.1.1.1", SubnetEnd: "30.1.1.128", SubnetMask: "255.255.255.128", Gateway: "10.1.1.1",
	}, {
		SubnetTag: "회사 내부망4", SubnetStart: "40.1.1.1", SubnetEnd: "40.1.1.128", SubnetMask: "255.255.255.128", Gateway: "10.1.1.1",
	}, {
		SubnetTag: "회사 내부망5", SubnetStart: "50.1.1.1", SubnetEnd: "50.1.1.128", SubnetMask: "255.255.255.128", Gateway: "10.1.1.1",
	}, {
		SubnetTag: "회사 내부망6", SubnetStart: "60.1.1.1", SubnetEnd: "60.1.1.128", SubnetMask: "255.255.255.128", Gateway: "10.1.1.1",
	}, {
		SubnetTag: "회사 내부망7", SubnetStart: "70.1.1.1", SubnetEnd: "70.1.1.128", SubnetMask: "255.255.255.128", Gateway: "10.1.1.1",
	}, {
		SubnetTag: "회사 내부망8", SubnetStart: "80.1.1.1", SubnetEnd: "80.1.1.128", SubnetMask: "255.255.255.128", Gateway: "10.1.1.1",
	}, {
		SubnetTag: "개발팀", SubnetStart: "100.1.1.1", SubnetEnd: "100.1.1.255", SubnetMask: "255.255.255.0", Gateway: "100.1.1.1",
	}, {
		SubnetTag: "기획팀", SubnetStart: "101.1.1.1", SubnetEnd: "101.1.1.255", SubnetMask: "255.255.255.0", Gateway: "101.1.1.1",
	}, {
		SubnetTag: "기술팀", SubnetStart: "102.1.1.1", SubnetEnd: "102.1.1.255", SubnetMask: "255.255.255.0", Gateway: "102.1.1.1",
	}, {
		SubnetTag: "영업팀", SubnetStart: "103.1.1.1", SubnetEnd: "103.1.1.255", SubnetMask: "255.255.255.0", Gateway: "103.1.1.1",
	},
}

var newAuth = []models.Auth{
	{
		Level: 1, Tag: "최고 관리자", Login: 2, Device: 2, Platform: 2, Stat: 2, Monitor: 2, Billing: 2,
		Cloud: 2, Board: 2, Manager: 2, WorkBoard: 2, TempStart: time.Time{}, TempEnd: time.Time{},
	},
	{
		Level: 2, Tag: "CB 관리자", Login: 2, Device: 2, Platform: 0, Stat: 2, Monitor: 2, Billing: 2,
		Cloud: 2, Board: 2, Manager: 2, WorkBoard: 2, TempStart: time.Time{}, TempEnd: time.Time{},
	},
	{
		Level: 3, Tag: "", Login: 0, Device: 0, Platform: 0, Stat: 0, Monitor: 0, Billing: 0,
		Cloud: 0, Board: 0, Manager: 0, WorkBoard: 0, TempStart: time.Time{}, TempEnd: time.Time{},
	},
	{
		Level: 4, Tag: "CB 작업자", Login: 2, Device: 2, Platform: 0, Stat: 0, Monitor: 0, Billing: 0,
		Cloud: 0, Board: 2, Manager: 2, WorkBoard: 2, TempStart: time.Time{}, TempEnd: time.Time{},
	},
	{
		Level: 5, Tag: "고객사 관리자", Login: 2, Device: 2, Platform: 0, Stat: 1, Monitor: 1, Billing: 1,
		Cloud: 2, Board: 2, Manager: 2, WorkBoard: 2, TempStart: time.Time{}, TempEnd: time.Time{},
	},
	{
		Level: 6, Tag: "고객사 사용자", Login: 2, Device: 2, Platform: 0, Stat: 1, Monitor: 1, Billing: 0,
		Cloud: 0, Board: 2, Manager: 0, WorkBoard: 2, TempStart: time.Time{}, TempEnd: time.Time{},
	},
	{
		Level: 7, Tag: "", Login: 0, Device: 0, Platform: 0, Stat: 0, Monitor: 0, Billing: 0,
		Cloud: 0, Board: 0, Manager: 0, WorkBoard: 0, TempStart: time.Time{}, TempEnd: time.Time{},
	},
	{
		Level: 8, Tag: "", Login: 0, Device: 0, Platform: 0, Stat: 0, Monitor: 0, Billing: 0,
		Cloud: 0, Board: 0, Manager: 0, WorkBoard: 0, TempStart: time.Time{}, TempEnd: time.Time{},
	},
	{
		Level: 9, Tag: "임시 사용자", Login: 2, Device: 1, Platform: 0, Stat: 1, Monitor: 0, Billing: 0,
		Cloud: 0, Board: 1, Manager: 0, WorkBoard: 0, TempStart: time.Time{}, TempEnd: time.Time{},
	},
	{
		Level: 10, Tag: "해지 고객", Login: 0, Device: 0, Platform: 0, Stat: 0, Monitor: 0, Billing: 0,
		Cloud: 0, Board: 0, Manager: 0, WorkBoard: 0, TempStart: time.Time{}, TempEnd: time.Time{},
	},
}