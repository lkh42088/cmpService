package mariadblayer

import "nubes/svcmgr/models"

type DBLayer interface {
	// Code
	GetAllCodes() ([]models.Code, error)
	AddCode(code models.Code) (models.Code, error)
	DeleteCode(code models.Code) (models.Code, error)
	DeleteCodes() error
	// SubCode
	GetAllSubCodes() ([]models.SubCode, error)
	AddSubCode(subCode models.SubCode) (models.SubCode, error)
	DeleteSubCode(subCode models.SubCode) (models.SubCode, error)
	DeleteSubCodes() error
}

