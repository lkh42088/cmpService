package insert

import (
	"cmpService/common/db"
	"cmpService/common/mariadblayer"
	"cmpService/dbmigrator/config"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var companyIdx = 0

func InsertNewObject() {
	insertSubCodeItems()
	insertCompanies()
	insertUsers()
}

func insertCodeItem() {

}

func insertSubCodeItems() {
	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()

	var data = newSubCodeData
	for num, subCode := range data {
		fmt.Printf("insertSubCodeItem (%d)\n", num)
		newDb.AddSubCode(subCode)
	}
}


func insertCompanies() {
	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()

	var data = newCompanies
	for num, company := range data {
		fmt.Printf("insertCompanies (%d)\n", num)
		newCompany, _ := newDb.AddCompany(company)
		companyIdx = int(newCompany.Idx)
	}
}

func insertUsers() {
	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()

	var data = newUsers
	for num, user := range data {
		pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(pass)
		user.CompanyIdx = companyIdx
		fmt.Printf("insertUsers (%d)\n", num)
		newDb.AddUser(user)
	}
}


