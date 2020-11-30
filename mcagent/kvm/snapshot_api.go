package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"fmt"
	"github.com/robfig/cron/v3"
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
	case 1:
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
		id, err = AddSnapshotCronWeekly(vm.Name)
	case 30:
		 /* Monthly */
		id, err = AddSnapshotCronMonthly(vm.Name)
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
	CronSch.DeleteSnapVm(config.VmName)

	var id cron.EntryID
	id = -1
	if config.Type == "true" {
		days, err := strconv.Atoi(config.Days)
		switch days {
		case 1:
			/* Daily */
			id, err = AddSnapshotCronDailyTime(config.VmName, config.Hours, config.Minutes)
			if err != nil {
				fmt.Println("AddCronSchFromVmSnapshot: error", err)
				return
			}
		case 7:
			/* Weekly */
			id, err = AddSnapshotCronWeekly(config.VmName)
		case 30:
			/* Monthly */
			id, err = AddSnapshotCronMonthly(config.VmName)
		default:
			/* etc */
			if config.Hours == "0" && config.Minutes == "0"  {
				fmt.Println("AddCronSchFromVmSnapshot: hours 0, minutes 0 --> skip")
				return
			} else if config.Hours == "0" {
				id, err = AddSnapshotCronByMin(config.VmName, config.Minutes)
			} else {
				id, err = AddSnapshotCronByHoursMin(config.VmName, config.Hours, config.Minutes)
			}
		}
	}

	if id != -1 {
		entry := SnapVm{
			VmName: config.VmName,
			CronId: id,
			Timer: fmt.Sprintf("%s days %s hours %s minutes",
				config.Days, config.Hours, config.Minutes),
		}
		CronSch.SnapVms = append(CronSch.SnapVms, entry)
	}
	for i, v := range CronSch.SnapVms {
		fmt.Printf("%d CronSch : %+v\n", i, v)
	}
}
