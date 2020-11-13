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
	if CronSch.LookupBackupVm(config.VmName) == false {
		fmt.Println("UpdateVmBakcupByConfig: dosn't have backup config!")
		return
	}

	CronSch.DeleteBackupVm(config.VmName)

	var id cron.EntryID
	var err error
	switch (config.Type) {
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
	}

	if err != nil {
		fmt.Println("UpdateVmSnapshotByConfig: error", err)
		return
	}
	entry := BackupVm{
		VmName: config.VmName,
		CronId: id,
	}
	CronSch.BackupVms = append(CronSch.BackupVms, entry)
}

