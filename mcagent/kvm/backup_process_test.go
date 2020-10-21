package kvm

import (
	"fmt"
	"testing"
)

var Name = "vm1"

func TestBackupVmImage(t *testing.T) {
	output := BackupVmImage(Name)
	fmt.Println("Result: ", output)
}