package kvm

import "testing"

var Name = "vm01"

func TestBackupVmImage(t *testing.T) {
	BackupVmImage(Name)
}