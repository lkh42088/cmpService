package convert

import "testing"

func TestCreateNewMariadbDb(t *testing.T) {
	CreateNewMariadbTable()
}

func TestConvertDb(t *testing.T) {
	RunConvertDb()
}

func TestDeleteDpb(t *testing.T) {
	DeleteDeviceTb()
}

func TestClearNewMariadbDb(t *testing.T) {
	DropNewMariadbTable()
}
