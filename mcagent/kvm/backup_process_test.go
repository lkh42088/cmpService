package kvm

import (
	"fmt"
	"testing"
	"time"
)

var Name = "SN87-VM-01"

func TestBackupVmImage(t *testing.T) {
	output, size := BackupVmImage(Name)
	fmt.Println("Result: ", output, size)
}

func TestSafeBackup(t *testing.T) {
	SafeBackup(Name, time.Now().String(), time.Now().String())
}