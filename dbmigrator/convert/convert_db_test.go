package convert

import "testing"

func TestCreateNewMariadbDb(t *testing.T) {
	CreateNewMariadbTable()
}

func TestConvertDb(t *testing.T) {
	RunConvertDb()
}

func TestClearNewMariadbDb(t *testing.T) {
	DropNewMariadbTable()
}
