package kvm

import (
	config2 "cmpService/mcagent/config"
	"cmpService/mcagent/ktrest"
	"cmpService/svcmgr/config"
	"fmt"
	"testing"
	"time"
)

var Name = "SN87-VM-01"
//var Name = "ebjee"

func GetConfig() {
	config2.ApplyGlobalConfig("/home/nubes/go/src/cmpService/mcagent/etc/mcagent.lkh.conf")
	cfg := config2.GetMcGlobalConfig()
	db, _ := config.SetMariaDB(cfg.MariaUser, cfg.MariaPassword, cfg.MariaDb,
		cfg.MariaIp, 3306)
	config2.SetDbOrm(db)
	_ = ktrest.PostAuthTokens()
	ktrest.ConfigurationForKtContainer()

}

func TestBackupVmImage(t *testing.T) {
	output, size := BackupVmImage(Name)
	fmt.Println("Result: ", output, size)
}

func TestSafeBackup(t *testing.T) {
	GetConfig()
	SafeBackup(Name, GetTimeWord(), time.Now().String())
}

func TestMcVmBackup(t *testing.T) {
	GetConfig()
	McVmBackup(Name, "SN87-VM-01-cronsch.qcow2.decrease")
}

