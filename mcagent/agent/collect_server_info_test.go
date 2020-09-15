package agent

import (
	"testing"
)

func TestGetSysInfo(t *testing.T) {
	//GetSysInfoAll()
	GetSysInfo()
}

func TestFindRouteInterface(t *testing.T) {
	FindRouteInterface("0.0.0.0")
}

func TestSendSysInfo(t *testing.T) {
	SendSysInfo()
}