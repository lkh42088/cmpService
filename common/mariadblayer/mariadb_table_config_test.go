package mariadblayer

import (
	"cmpService/common/db"
	"testing"
)

func TestCreateTable(t *testing.T) {
	db := db.Connect(getTestConfig())
	if db == nil {
		return
	}
	defer db.Close()
	CreateTable(db)
}

func TestDropTable(t *testing.T) {
	db := db.Connect(getTestConfig())
	if db == nil {
		return
	}
	defer db.Close()
	DropTable(db)
}

// Jungbh CB
func TestCreateJbhCBTable(t *testing.T) {
	db := db.Connect(getTestJbhCBConfig())
	if db == nil {
		return
	}
	defer db.Close()
	CreateTable(db)
}

func TestDropJbhCBTable(t *testing.T) {
	db := db.Connect(getTestJbhCBConfig())
	if db == nil {
		return
	}
	defer db.Close()
	DropTable(db)
}
