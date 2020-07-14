package insert

import (
	"cmpService/common/db"
	"cmpService/common/mariadblayer"
	"cmpService/dbmigrator/config"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func InsertNewObject() {
	//insertSubCodeItems()
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
		fmt.Printf("insertUsers (%d)\n", num)
		newDb.AddUser(user)
	}
}


