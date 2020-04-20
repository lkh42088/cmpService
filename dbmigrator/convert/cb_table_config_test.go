package convert

import "testing"

func TestCreateCbMysqlDb(t *testing.T) {
	CreateTestCbMysqlTable()
}

func TestClearCbMysqlDb(t *testing.T) {
	DropTestCbMysqlTable()
}

func TestReCreateCbMysqlDb(t *testing.T) {
	DropTestCbMysqlTable()
	CreateTestCbMysqlTable()
}
