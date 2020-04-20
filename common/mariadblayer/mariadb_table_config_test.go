package mariadblayer

import (
	db2 "nubes/common/db"
	"testing"
)

func TestCreateTable(t *testing.T) {
	db := db2.Connect(getTestConfig())
	if db == nil {
		return
	}
	defer db.Close()
	CreateTable(db)
}

func TestDropTable(t *testing.T) {
	db := db2.Connect(getTestConfig())
	if db == nil {
		return
	}
	defer db.Close()
	DropTable(db)
}
