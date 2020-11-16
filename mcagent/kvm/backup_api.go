package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"fmt"
	"github.com/robfig/cron/v3"
	"strconv"
)

/*********************************************************************************
 * Backup
 *********************************************************************************/
func AddCronSchForVmBackup(vm *mcmodel.McVm) {
	if vm.BackupType == false {
		return
	}

	if CronSch.LookupBackupVm(vm.Name) {
		fmt.Println("AddCronSchForVmBackup: Already have Backup config!")
		return
	}

	var id cron.EntryID
	var err error
	switch (vm.BackupDays) {
	case 7:
		/* Weekly */
	case 30:
		/* Monthly */
	default:
		/* etc */
		if vm.BackupHours == 0 && vm.BackupMinutes == 0  {
			fmt.Println("AddCronSchForVmBackup: hours 0, minutes 0 --> skip")
			return
		} else if vm.BackupHours == 0 {
			id, err = AddBackupCronByMin(vm.Name,
				strconv.Itoa(vm.BackupMinutes))
		} else {
			id, err = AddBackupCronByHoursMin(vm.Name,
				strconv.Itoa(vm.BackupHours),
				strconv.Itoa(vm.BackupMinutes))
		}
		if err != nil {
			fmt.Println("AddCronSchFromVmBackup error: ", err)
			return
		}
	}

	entry := BackupVm{
		VmName: vm.Name,
		CronId: id,
		Timer: fmt.Sprintf("%d days %d hours %d minutes",
			vm.BackupDays, vm.BackupHours, vm.BackupMinutes),
	}
	CronSch.BackupVms = append(CronSch.BackupVms, entry)
	fmt.Printf("\n## cron backup : %+v\n\n", CronSch.BackupVms)
}

func UpdateVmBackupByConfig(config *messages.BackupConfigMsg) {
	var configType string
	CronSch.DeleteBackupVm(config.VmName)
	configType = ""

	var id cron.EntryID
	id = -1
	var err error
	if config.Type == "true" {
		switch configType {
		case "designatedTime":
			id, err = AddBackupCronDailyTime(config.VmName, config.Hours, config.Minutes)
		case "periodically":
			if config.Days == "30" {
				// monthly
				id, err = AddBackupCronMonthly(config.VmName)
			} else if config.Days == "7" {
				// weekly
				id, err = AddBackupCronWeekly(config.VmName)
			} else if config.Days == "1" {
				// daily
				id, err = AddBackupCronDaily(config.VmName)
			} else {
				// hourly
				id, err = AddBackupCronPeriodically(config.VmName, config.Hours, "0")
			}
		default:
			if config.Hours == "0" && config.Minutes == "0"  {
				fmt.Println("AddCronSchForVmBackup: hours 0, minutes 0 --> skip")
				return
			} else if config.Hours == "0" {
				id, err = AddBackupCronByMin(config.VmName,	config.Minutes)
			} else {
				id, err = AddBackupCronByHoursMin(config.VmName, config.Hours, config.Minutes)
			}
			if err != nil {
				fmt.Println("AddCronSchFromVmBackup error: ", err)
				return
			}
		}
	}

	if err != nil {
		fmt.Println("UpdateVmSnapshotByConfig: error", err)
		return
	}
	if id != -1 {
		entry := BackupVm{
			VmName: config.VmName,
			CronId: id,
			Timer: fmt.Sprintf("%s days %s hours %s minutes",
				config.Days, config.Hours, config.Minutes),
		}
		CronSch.BackupVms = append(CronSch.BackupVms, entry)
	}
	for i, v := range CronSch.BackupVms {
		fmt.Printf("%d CronSch : %+v\n", i, v)
	}
}

