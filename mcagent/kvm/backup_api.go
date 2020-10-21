package kvm

import (
	"cmpService/common/mcmodel"
	"fmt"
	"github.com/robfig/cron"
	"strconv"
)

/*********************************************************************************
 * Backup
 *********************************************************************************/
func AddCronSchForVmBackup(vm *mcmodel.McVm) {
	if vm.BackupType == 0 {
		return
	}

	if CronSch.LookupBackupVm(vm.Name) {
		fmt.Println("AddCronSchForVmBackup: Already have Backup config!")
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
		if vm.BackupHours == 0 && vm.BackupMinutes == 0  {
			fmt.Println("AddCronSchForVmBackup: hours 0, minutes 0 --> skip")
			return
		} else if vm.BackupHours == 0 {
			AddBackupCronByMin(vm.Name,
				strconv.Itoa(vm.BackupMinutes))
		} else {
			AddBackupCronByHoursMin(vm.Name,
				strconv.Itoa(vm.BackupHours),
				strconv.Itoa(vm.BackupMinutes))
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

