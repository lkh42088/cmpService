package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"fmt"
	"github.com/robfig/cron"
	"strconv"
)

/*********************************************************************************
 * Snapshot
 *********************************************************************************/
/* When to start mcagent daemon, add snapshot cronsch config */
func AddCronSchFromVmSnapshot(vm *mcmodel.McVm) {
	if vm.SnapType == false {
		return
	}

	if CronSch.LookupSnapVm(vm.Name) {
		fmt.Println("AddCronSchFromVmSnapshot: Already have snapshot config!")
		return
	}

	var id cron.EntryID
	var err error
	switch (vm.SnapDays) {
	case 2:
		 /* Daily */
		id, err = AddSnapshotCronDailyTime(vm.Name,
			strconv.Itoa(vm.SnapHours),
			strconv.Itoa(vm.SnapMinutes))
		if err != nil {
			fmt.Println("AddCronSchFromVmSnapshot: error", err)
			return
		}
	case 7:
		 /* Weekly */
	case 30:
		 /* Monthly */
	default:
		 /* etc */
		if vm.SnapHours == 0 && vm.SnapMinutes == 0  {
			fmt.Println("AddCronSchFromVmSnapshot: hours 0, minutes 0 --> skip")
			return
		} else if vm.SnapHours == 0 {
			AddSnapshotCronByMin(vm.Name,
				strconv.Itoa(vm.SnapMinutes))
		} else {
			AddSnapshotCronByHoursMin(vm.Name,
				strconv.Itoa(vm.SnapHours),
				strconv.Itoa(vm.SnapMinutes))
		}
	}

	entry := SnapVm{
		VmName: vm.Name,
		CronId: id,
		Timer: fmt.Sprintf("%d days %d hours %d minutes",
			vm.SnapDays, vm.SnapHours, vm.SnapMinutes),
	}
	CronSch.SnapVms = append(CronSch.SnapVms, entry)
}

/* By Rest Apis */
func AddVmSnapshotByConfig(config *messages.SnapshotConfigMsg) {
	if CronSch.LookupSnapVm(config.VmName) {
		fmt.Println("AddVmSnapshotByConifg: Already have snapshot config!")
		return
	}
	var id cron.EntryID
	var err error
	switch (config.Type) {
	case "designatedTime":
		id, err = AddSnapshotCronDailyTime(config.VmName, config.Hours, config.Minutes)
	case "periodically":
		if config.Days == "30" {
			// monthly
			id, err = AddSnapshotCronMonthly(config.VmName)
		} else if config.Days == "7" {
			// weekly
			id, err = AddSnapshotCronWeekly(config.VmName)
		} else if config.Days == "1" {
			// daily
			id, err = AddSnapshotCronDaily(config.VmName)
		} else {
			// hourly
			id, err = AddSnapshotCronPeriodically(config.VmName, config.Hours, "0")
		}
	default:
	}
	if err != nil {
		fmt.Println("AddSnapshotConDailyTime: error", err)
		return
	}
	entry := SnapVm{
		VmName: config.VmName,
		CronId: id,
	}
	CronSch.SnapVms = append(CronSch.SnapVms, entry)
}

func UpdateVmSnapshotByConfig(config *messages.SnapshotConfigMsg) {
	if CronSch.LookupSnapVm(config.VmName) == false {
		fmt.Println("UpdateVmSnapshotByConfig: dosn'thave snapshot config!")
		return
	}

	CronSch.DeleteSnapVm(config.VmName)

	var id cron.EntryID
	var err error
	switch (config.Type) {
	case "designatedTime":
		id, err = AddSnapshotCronDailyTime(config.VmName, config.Hours, config.Minutes)
	case "periodically":
		if config.Days == "30" {
			// monthly
			id, err = AddSnapshotCronMonthly(config.VmName)
		} else if config.Days == "7" {
			// weekly
			id, err = AddSnapshotCronWeekly(config.VmName)
		} else if config.Days == "1" {
			// daily
			id, err = AddSnapshotCronDaily(config.VmName)
		} else {
			// hourly
			id, err = AddSnapshotCronPeriodically(config.VmName, config.Hours, "0")
		}
	default:
	}

	if err != nil {
		fmt.Println("UpdateVmSnapshotByConfig: error", err)
		return
	}
	entry := SnapVm{
		VmName: config.VmName,
		CronId: id,
	}
	CronSch.SnapVms = append(CronSch.SnapVms, entry)
}
