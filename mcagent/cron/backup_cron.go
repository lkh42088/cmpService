package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func AddBackupCronByHoursMin(vmName, hour, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@every %sh%sm0s - Backup", hour, min)
	err = CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Printf("Backup: %s - %s\n", vmName, cronTime)
		SafeSnapshot(vmName, GetTimeWord(), cronTime)
		SafeBackup(vmName, GetTimeWord(), cronTime)
	})
	return id, err
}

func AddBackupCronByMin(vmName, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@every 0h%sm0s - Backup", min)
	err = CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Printf("Backup: %s - %s\n", vmName, cronTime)
		SafeSnapshot(vmName, GetTimeWord(), cronTime)
		SafeBackup(vmName, GetTimeWord(), cronTime)
	})
	return id, err
}
