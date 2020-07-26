package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/models"
)

// User, Companies, Auth
func (db *DBORM) GetCompaniesByCpName(name string) (companies []models.CompanyResponse, err error) {
	name = "%" + name + "%"
	return companies, db.
		Table(CompanyRawTable).
		Where("cp_name like ?", name).
		Find(&companies).Error
}

func (db *DBORM) GetCompaniesByLikeCpName(name string) (companies []models.CompanyResponse, err error) {
	name = "%" + name + "%"
	return companies, db.
		Table(CompanyRawTable).
		Where("cp_name like ?", name).
		Find(&companies).Error
}

func (db *DBORM) GetCompaniesWithUserByLikeCpName(name string) (companies []models.CompanyResponse, err error) {
	name = "%" + name + "%"
	return companies, db.
		Table(CompanyRawTable).
		Select(CompanyAndUserIdSelectQuery).
		Where("cp_name like ?", name).
		Joins(CompanyAndUserJoinQuery).
		Find(&companies).Error
}

func (db *DBORM) GetCompanies() (companies []models.CompanyResponse, err error) {
	return companies, db.
		Table(CompanyRawTable).
		Find(&companies).Error
}

func (db *DBORM) GetCompanyByCpName(name string) (company models.Company, err error) {
	return company, db.Where("cp_name = ?", name).Find(&company).Error
}

func (db *DBORM) GetCompaniesPage(paging models.Pagination) (companies models.CompanyPage, err error) {
	db.Model(&companies.Companies).Count(&paging.TotalCount)
	err = db.
		Order(companies.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Find(&companies.Companies).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	companies.Page = paging
	return companies, err
}

func (db *DBORM) GetCompaniesPageBySearch(paging models.Pagination, query string) (companies models.CompanyPage, err error) {
	err = db.
		Order(companies.GetOrderBy(paging.OrderBy, paging.Order)).
		Limit(paging.RowsPerPage).
		Offset(paging.Offset).
		Where(query).
		Find(&companies.Companies).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	paging.TotalCount = len(companies.Companies)
	companies.Page = paging
	return companies, err
}

func (db *DBORM) AddCompany(company models.Company) (models.Company, error) {
	return company, db.Create(&company).Error
}

func (db *DBORM) UpdateCompany(obj models.Company) (models.Company, error) {
	return obj, db.Model(&obj).
		Updates(map[string]interface{}{
			"cp_idx":obj.Idx,
			"cp_name":obj.Name,
			"cp_email":obj.Email,
			"cp_homepage":obj.Homepage,
			"cp_tel":obj.Tel,
			"cp_hp":obj.HP,
			"cp_zip":obj.Zipcode,
			"cp_addr":obj.Address,
			"cp_addr_detail":obj.AddressDetail,
			"cp_termination_date":obj.TermDate,
			"cp_is_company":obj.IsCompany,
			"cp_memo":obj.Memo,
		}).Error
}

func (db *DBORM) DeleteCompany(company models.Company) (models.Company, error) {
	return company, db.Delete(&company).Error
}

func (db *DBORM) DeleteAllCompany() error {
	return db.Delete(&models.Company{}).Error
}
