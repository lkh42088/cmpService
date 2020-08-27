package mariadblayer

import (
	"cmpService/common/models"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const typeField = "c_type"
const typeSubField = "c_type_sub"

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

func (db *DBORM) GetSubCodeList(cIdx []string) (subCodes []models.SubCodeResponse, err error) {
	return subCodes, db.
		//Debug().
		Table(CodeSubRawTable).
		Select(CodeAndSubCodeSelectQuery).
		Joins(CodeAndSubCodeJoinQuery).
		Where("c.c_idx IN (?)", cIdx).
		Find(&subCodes).Error
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

func (db *DBORM) GetCodeList(code string, subCode string) (codes []models.Code, err error) {
	return codes, db.
		Where(GetWhereString(typeSubField), subCode).
		Where(GetWhereString(typeField), code).
		Find(&codes).Error
}

func (db *DBORM) GetCodeByIdx(codeIdx string) (code models.Code, err error) {
	//reScan := db.Select("SELECT c_name FROM code_tb WHERE c_idx = ?", code).Scan(&cName)
	//err = db.Select("SELECT c_name FROM code_tb WHERE c_idx = ?", code).Error
	/*
		result := db.Table("code_tb").Select("c_name").Where("c_idx = ?", codeIdx)

		fmt.Println("cName : ", cName)
		fmt.Println("result : ", result)
		test := db.Raw("select * from code_tb").Scan(&cName)
		fmt.Println("test : ", test)*/

	/*return returnVal, db.Where("c_idx=?", code).Find(models.Code{}).Error*/
	return code, db.Where("c_idx = ?", codeIdx).Find(&code).Error
}

func (db *DBORM) GetSubCodeByIdx(codeIdx string) (subCode models.SubCode, err error) {
	return subCode, db.Where("csub_idx = ?", codeIdx).Find(&subCode).Error
}

/*
val, err := db.Where("c_idx=?", code).Find(models.Code{}).Error
err := db.Raw("SELECT c_name FROM code_tb WHERE c_idx = ?", code).Scan(&result)

*/
/*----------------------------------------------------------------------------------------------------------*/

func (db *DBORM) GetCodeTagList() (codes []models.Code, err error) {
	return codes, db.Select("DISTINCT c_type").Find(&codes).Error
}

func (db *DBORM) GetCodesMainByType(code string, subCode string) (codes []models.Code, err error) {
	return codes, db.Debug().
		Where(GetWhereString(typeSubField), subCode).
		Where(GetWhereString(typeField), code).
		Find(&codes).Error
}

func (db *DBORM) GetCodesSubByIdx(idx string) (codes []models.SubCode, err error) {
	return codes, db.Debug().Where("c_idx = ?", idx).Find(&codes).Error
}
