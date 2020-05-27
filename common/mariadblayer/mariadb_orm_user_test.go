package mariadblayer

import (
	"cmpService/common/db"
	"cmpService/common/models"
	"fmt"
	"testing"
)

func getJungbhDb() (*DBORM, error){
	config := getTestJbhConfig()
	options := db.GetDataSourceName(config)
	db, err := NewDBORM(config.DBDriver, options)
	return db, err
}

func TestAddUser(t *testing.T) {
	db, err := getJungbhDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	user := models.User {
		//ID:"jungbh",
		//Password: "nubes1510",
		//Email: "jungbh@cmpService-bridge.com",
		//Name: "jungbh",
		//Level: 10,
	}
	user, err = db.AddUser(user)
	fmt.Println(user)
}

func TestGetUser(t *testing.T) {
	db, err := getJungbhDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err := db.GetUserById("jungbh")
	fmt.Println(user)
}
