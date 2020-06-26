package models

import (
	"strings"
	"time"
)

const (
	CompanyOrmIdx           = "cp_idx"
	CompanyOrmName          = "cp_name"
	CompanyOrmEmail         = "cp_email"
	CompanyOrmHomepage      = "cp_homepage"
	CompanyOrmTel           = "cp_tel"
	CompanyOrmHP            = "cp_hp"
	CompanyOrmZipcode       = "cp_zip"
	CompanyOrmAddress       = "cp_addr"
	CompanyOrmAddressDetail = "cp_addr_detail"
	CompanyOrmTermDate      = "cp_termination_date"
	CompanyOrmIsCompany     = "cp_is_company"
	CompanyOrmMemo          = "cp_memo"
)

const (
	CompanyJsonIdx           = "idx"
	CompanyJsonName          = "name"
	CompanyJsonEmail         = "email"
	CompanyJsonHomepage      = "homepage"
	CompanyJsonTel           = "tel"
	CompanyJsonHP            = "hp"
	CompanyJsonZipcode       = "zipcode"
	CompanyJsonAddress       = "address"
	CompanyJsonAddressDetail = "addressDetail"
	CompanyJsonTermDate      = "termDate"
	CompanyJsonIsCompany     = "isCompany"
	CompanyJsonMemo          = "memo"
)

var CompanyOrmMap = map[string]string{
	CompanyOrmIdx:           CompanyJsonIdx,
	CompanyOrmName:          CompanyJsonName,
	CompanyOrmEmail:         CompanyJsonEmail,
	CompanyOrmHomepage:      CompanyJsonHomepage,
	CompanyOrmTel:           CompanyJsonTel,
	CompanyOrmHP:            CompanyJsonHP,
	CompanyOrmZipcode:       CompanyJsonZipcode,
	CompanyOrmAddress:       CompanyJsonAddress,
	CompanyOrmAddressDetail: CompanyJsonAddressDetail,
	CompanyOrmTermDate:      CompanyJsonTermDate,
	CompanyOrmIsCompany:     CompanyJsonIsCompany,
	CompanyOrmMemo:          CompanyJsonMemo,
}

var CompanyJsonMap = map[string]string{
	CompanyJsonIdx:           CompanyOrmIdx,
	CompanyJsonName:          CompanyOrmName,
	CompanyJsonEmail:         CompanyOrmEmail,
	CompanyJsonHomepage:      CompanyOrmHomepage,
	CompanyJsonTel:           CompanyOrmTel,
	CompanyJsonHP:            CompanyOrmHP,
	CompanyJsonZipcode:       CompanyOrmZipcode,
	CompanyJsonAddress:       CompanyOrmAddress,
	CompanyJsonAddressDetail: CompanyOrmAddressDetail,
	CompanyJsonTermDate:      CompanyOrmTermDate,
	CompanyJsonIsCompany:     CompanyOrmIsCompany,
	CompanyJsonMemo:          CompanyOrmMemo,
}

type Company struct {
	Idx           uint      `gorm:"primary_key;column:cp_idx;auto_increment;comment:'INDEX'" json:"idx"`
	Name          string    `gorm:"type:varchar(255);not null;column:cp_name;comment:'회사명'" json:"name"`
	Email         string    `gorm:"type:varchar(255);column:cp_email;comment:'이메일'" json:"email"`
	Homepage      string    `gorm:"type:varchar(255);column:cp_homepage;comment:'홈페이지'" json:"homepage"`
	Tel           string    `gorm:"type:varchar(15);column:cp_tel;comment:'전화 번호'" json:"tel"`
	HP            string    `gorm:"type:varchar(15);column:cp_hp;comment:'핸드폰 번호'" json:"hp"`
	Zipcode       string    `gorm:"type:varchar(15);not null;column:cp_zip;comment:'우편 번호'" json:"zipcode"`
	Address       string    `gorm:"type:varchar(255);not null;column:cp_addr;comment:'주소'" json:"address"`
	AddressDetail string    `gorm:"type:varchar(255);not null;column:cp_addr_detail;comment:'상세 주소'" json:"addressDetail"`
	TermDate      time.Time `gorm:"type:datetime;column:cp_termination_date;comment:'해지 일자'" json:"termDate"`
	IsCompany     bool      `gorm:"type:tinyint(1);column:cp_is_company;default:1;comment:'회사 여부'" json:"isCompany"`
	Memo          string    `gorm:"type:text;column:cp_memo;comment:'메모'" json:"memo"`
}

type CompanyDetail struct {
	Company
	UserId string `gorm:"type:varchar(16); column:user_id" json:"userId"`
	UserPassword string `gorm:"type:varchar(16); column:user_password" json:"userPassword"`
}

func (Company) TableName() string {
	return "company_tb"
}

type CompanyResponse struct {
	Company
	UserId string `gorm:"unique;type:varchar(50);column:user_id" json:"userId"`
}

type CompanyPage struct {
	Page      Pagination `json:"page"`
	Companies []Company  `json:"data"`
}

func (c CompanyPage) GetOrderBy(orderby, order string) string {
	val, exists := CompanyJsonMap[orderby]
	if !exists {
		val = "cp_name"
	}
	order = strings.ToLower(order)
	if !(order == "asc" || order == "desc") {
		order = "asc"
	}
	return val + " " + order
}
