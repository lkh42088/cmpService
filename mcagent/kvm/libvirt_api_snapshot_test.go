package kvm

import "testing"

var vmName = "vm01"

func TestCreateSnapshot(t *testing.T) {
	CreateSnapshot(vmName, "snap 02", "test snap 02")
}

func TestGetSnapshot(t *testing.T) {
	GetAllSnapshots(vmName)
}

func TestGetSnapshotListName(t *testing.T) {
	GetSnapshotsListName(vmName)
}

func TestDeleteAllSnapshot(t *testing.T) {
	DeleteAllSnapshot(vmName)
}
