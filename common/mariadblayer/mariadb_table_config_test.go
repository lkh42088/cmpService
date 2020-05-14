package mariadblayer

import (
	"nubes/common/db"
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
