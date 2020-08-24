package kvm

import "testing"

func TestCreateSnapshot(t *testing.T) {
	CreateSnapshot("win10-01", "test snap")
}

func TestGetSnapshot(t *testing.T) {
	GetAllSnapshots("win10-01")
}
