package cron

import (
	"cmpService/common/mcmodel"
	"fmt"
	"github.com/robfig/cron"
	"strconv"
)

/*********************************************************************************
 * Backup
 *********************************************************************************/
func AddBackupByMcVm(vm *mcmodel.McVm) {
	if vm.SnapType == false {
		return
	}

	if CronSch.LookupBackupVm(vm.Name) {
		fmt.Println("AddBackupByMcVm: Already have Backup config!")
		return
	}

	var id cron.EntryID
	//var err error
	switch (vm.BackupDays) {
	case 7:
		/* Weekly */
	case 30:
		/* Monthly */
	default:
		/* etc */
		if vm.SnapHours == 0 && vm.SnapMinutes == 0  {
			fmt.Println("AddBackupByMcVm: hours 0, minutes 0 --> skip")
			return
		} else if vm.SnapHours == 0 {
			AddBackupCronByMin(vm.Name,
				strconv.Itoa(vm.SnapMinutes))
		} else {
			AddBackupCronByHoursMin(vm.Name,
				strconv.Itoa(vm.SnapHours),
				strconv.Itoa(vm.SnapMinutes))
		}
	}

	entry := BackupVm{
		VmName: vm.Name,
		CronId: id,
		Timer: fmt.Sprintf("%d days %d hours %d minutes",
			vm.BackupDays, vm.BackupHours, vm.BackupMinutes),
	}
	CronSch.BackupVms = append(CronSch.BackupVms, entry)
}

