package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/models"
)

// User, Companies, Auth
func (db *DBORM) GetCompaniesByName(name string) (companies []models.CompanyResponse, err error) {
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

func (db *DBORM) GetCompanyByName(name string) (company models.Company, err error) {
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

func (db *DBORM) AddCompany(company models.Company) (models.Company, error) {
	return company, db.Create(&company).Error
}

func (db *DBORM) DeleteAllCompany() error {
	return db.Delete(&models.Company{}).Error
}
