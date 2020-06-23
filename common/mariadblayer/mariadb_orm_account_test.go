package mariadblayer

import (
	"cmpService/common/db"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateAccountTable(t *testing.T) {
	db := db.Connect(getTestJbhHomeConfig())
	if db == nil {
		return
	}
	defer db.Close()
	CreateAccountTable(db)
}

func TestDeleteAccountTable(t *testing.T) {
	db := db.Connect(getTestJbhHomeConfig())
	if db == nil {
		return
	}
	defer db.Close()
	DeleteAccountTable(db)
}

func TestGetFile(t *testing.T) {
	dirName, _ := os.Getwd()
	fmt.Printf("current directory: %s\n", dirName)
	getFileByPath("./mariadb_orm.go")
}

func getFileByPath(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to get file: ", err)
		return
	}
	fmt.Println("Success: size ", len(b))
}
