package kvm

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func AddBackupCronMonthly(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@monthly")
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run daily"))
		SafeBackup(vmName, GetTimeWord(), "Monthly, Backup by Cron")
	})
	return id, err
}

func AddBackupCronWeekly(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@weekly")
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		SafeBackup(vmName, GetTimeWord(), "Weekly, Backup by Cron")
	})
	return id, err
}

func AddBackupCronDaily(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@daily")
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run daily"))
		SafeBackup(vmName, GetTimeWord(), "Daily, Backup by Cron")
	})
	return id, err
}

func AddBackupCronDailyTime(vmName, hour, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("CRON_TZ=Asia/Seoul %s %s * * *", hour, min)
	fmt.Println("AddBackupCronDailyTime:", vmName, "-", cronTime)
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run at %s:%s Seoul time every day", hour, min))
		SafeBackup(vmName, GetTimeWord(), "Daily Time, Backup by Cron")
	})
	return id, err
}

func AddBackupCronPeriodically(vmName, hour, min string) (id cron.EntryID, err error) {
	cronTime := fmt.Sprintf("@every %sh%sm0s", hour, min)
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		SafeBackup(vmName, GetTimeWord(), "Periodically, Backup by Cron")
	})
	return id, err
}

func AddBackupCronByHoursMin(vmName, hour, min string) (cron.EntryID, error){
	cronTime := fmt.Sprintf("@every %sh%sm0s", hour, min)
	id, err := CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Printf("Backup: %s - %s\n", vmName, cronTime)
		SafeBackup(vmName, GetTimeWord(), cronTime)
	})
	return id, err
}

func AddBackupCronByMin(vmName, min string) (cron.EntryID, error){
	cronTime := fmt.Sprintf("@every 0h%sm0s", min)
	id, err := CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Printf("Backup: %s - %s\n", vmName, cronTime)
		SafeBackup(vmName, GetTimeWord(), cronTime)
	})
	return id, err
}
