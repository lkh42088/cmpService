package mariadblayer

import (
	"nubes/common/models"
)

func (db *DBORM) GetAllCodes() (codes []models.Code, err error) {
	return codes, db.Find(&codes).Error
}

func (db *DBORM) AddCode(code models.Code) (models.Code, error) {
	return code, db.Create(&code).Error
}

func (db *DBORM) DeleteCode(code models.Code) (models.Code, error) {
	return code, db.Delete(&code).Error
}

func (db *DBORM) DeleteCodes() error {
	return db.Delete(&models.Code{}).Error
}

func (db *DBORM) GetAllSubCodes() (subCodes []models.SubCode, err error) {
	return subCodes, db.Find(&subCodes).Error
}

func (db *DBORM) AddSubCode(subCode models.SubCode) (models.SubCode, error) {
	return subCode, db.Create(&subCode).Error
}

func (db *DBORM) DeleteSubCode(subCode models.SubCode) (models.SubCode, error) {
	return subCode, db.Delete(&subCode).Error
}

func (db *DBORM) DeleteSubCodes() error {
	return db.Delete(&models.SubCode{}).Error
}

