package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"fmt"
	"github.com/robfig/cron/v3"
	"strconv"
	"sync"
	"time"
)

type CronSnapshot struct {
	Interval int
	Cr *cron.Cron
	Vms []SnapVm
}

var CronSnap *CronSnapshot

type SnapVm struct {
	VmName string
	CronId cron.EntryID
	Timer string
}

var vmSliceMutex sync.Mutex

func (c *CronSnapshot) LookupVm(vmName string) bool {
	for _, vm := range c.Vms {
		if vm.VmName == vmName {
			return true
		}
	}
	return false
}

func (c *CronSnapshot) DeleteVm(vmName string) bool {
	findIt := -1
	for index, vm := range c.Vms {
		if vm.VmName == vmName {
			DeleteCronEntry(vm.CronId)
			findIt = index
		}
	}
	if findIt > 0 {
		vmSliceMutex.Lock()
		c.Vms = append(c.Vms[:findIt], c.Vms[findIt+1:]...)
		vmSliceMutex.Unlock()
		return true
	}
	return false
}

func DeleteCronEntry(id cron.EntryID) {
	CronSnap.Cr.Remove(id)
}

func NewCronSnapshot(interval int) *CronSnapshot {
	return &CronSnapshot{
		Interval: interval,
	}
}

func ConfigCron() {
	c := NewCronSnapshot(5)
	SetCronSnapshot(c)
}

func SetCronSnapshot(c *CronSnapshot) {
	CronSnap = c
	CronSnap.Cr = cron.New()
	CronSnap.Cr.AddFunc("*/5 * * * *", func() { fmt.Println("CRON SNAPSHOT - every second 5.")})
}

func (c *CronSnapshot) Start(parentwg *sync.WaitGroup) {
	loop := 1
	c.Cr.Start()
	for {
		c.Run()
		time.Sleep(time.Duration(c.Interval * int(time.Second)))
		loop += 1
	}
	parentwg.Done()
}

func (c *CronSnapshot) Run() {
	//fmt.Printf("Entry %+v\n", c.Cr.Entries())
}

func AddSnapshotCronMonthly(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@monthly")
	id, err = CronSnap.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run daily"))
		SafeSnapshot(vmName, GetTimeWord(), "Monthly, snapshot by Cron")
	})
	return id, err
}

func AddSnapshotCronWeekly(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@weekly")
	id, err = CronSnap.Cr.AddFunc(cronTime, func() {
		SafeSnapshot(vmName, GetTimeWord(), "Weekly, snapshot by Cron")
	})
	return id, err
}

func AddSnapshotCronDaily(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@daily")
	id, err = CronSnap.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run daily"))
		SafeSnapshot(vmName, GetTimeWord(), "Daily, snapshot by Cron")
	})
	return id, err
}

func AddSnapshotCronDailyTime(vmName, hour, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("CRON_TZ=Asia/Seoul %s %s * * *", hour, min)
	fmt.Println("AddSnapshotCronDailyTime:", vmName, "-", cronTime)
	id, err = CronSnap.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run at %s:%s Seoul time every day", hour, min))
		SafeSnapshot(vmName, GetTimeWord(), "Daily Time, snapshot by Cron")
	})
	return id, err
}

func AddSnapshotCronPeriodically(vmName, hour, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@every %sh%sm0s", hour, min)
	id, err = CronSnap.Cr.AddFunc(cronTime, func() {
		SafeSnapshot(vmName, GetTimeWord(), "Periodically, snapshot by Cron")
	})
	return id, err
}

func AddSnapshotByMcVm(vm *mcmodel.McVm) {
	if vm.SnapType == false {
		return
	}

	if CronSnap.LookupVm(vm.Name) {
		fmt.Println("AddSnapshotByMcVm: Already have snapshot config!")
		return
	}

	var id cron.EntryID
	var err error
	switch (vm.SnapDays) {
	case 2:
		/**************
		 * Daily
		 **************/
		id, err = AddSnapshotCronDailyTime(vm.Name,
			strconv.Itoa(vm.SnapHours),
			strconv.Itoa(vm.SnapMinutes))
		if err != nil {
			fmt.Println("AddSnapshotByMcVm: error", err)
			return
		}
	case 7:
		/**************
		 * Weekly
		 **************/
	case 30:
		/**************
		 * Monthly
		 **************/
	default:
		/**************
		 * etc
		 **************/
		if vm.SnapHours == 0 && vm.SnapMinutes == 0  {
			fmt.Println("AddSnapshotByMcVm: hours 0, minutes 0 --> skip")
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
		Timer: fmt.Sprintf("%s days %d hours %d minutes",
			vm.SnapDays, vm.SnapHours, vm.SnapMinutes),
	}
	CronSnap.Vms = append(CronSnap.Vms, entry)
}

func AddVmSnapshotByConfig(config *messages.SnapshotConfigMsg) {
	if CronSnap.LookupVm(config.VmName) {
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
	CronSnap.Vms = append(CronSnap.Vms, entry)
}

func UpdateVmSnapshotByConfig(config *messages.SnapshotConfigMsg) {
	if CronSnap.LookupVm(config.VmName) == false {
		fmt.Println("UpdateVmSnapshotByConfig: dosn'thave snapshot config!")
		return
	}

	CronSnap.DeleteVm(config.VmName)

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
	CronSnap.Vms = append(CronSnap.Vms, entry)
}

func DeleteVmSnapshotByConfig(config *messages.SnapshotConfigMsg) {
	if CronSnap.LookupVm(config.VmName) == false {
		fmt.Println("DeleteVmSnapshotByConfig: dosen't have snapshot config!")
		return
	}

	CronSnap.DeleteVm(config.VmName)
}

func GetVmSnapshotAll() *[]messages.SnapshotEntry {
	var arr []messages.SnapshotEntry
	return &arr
}

func GetVmSnapshotByVmName(vmName string) *[]messages.SnapshotEntry {
	var arr []messages.SnapshotEntry
	return &arr
}

func AddSnapshotCronSecond(min, domName string) {
	var i int
	cronTime := fmt.Sprintf("@every 0h0m%ss", min)
	CronSnap.Cr.AddFunc(cronTime, func() {
		var snapName = fmt.Sprintf("%s-%d", domName, i)
		i += 1
		fmt.Println(fmt.Sprintf("every second - %s", min))
		//dom, err := GetDomainByName(domName)
		//if err != nil {
		//	fmt.Println("error get vm:", err)
		//	return
		//}
		//name, _ := dom.GetName()
		//fmt.Println("vm:", name)
		SafeSnapshot(domName, GetTimeWord(), snapName)
	})
}

func GetTimeWord() string {
	t := time.Now()
	fmt.Println(t)
	return fmt.Sprintf("%d%s%d-%d-%d-%d",
		t.Year(), t.Month().String()[:3], t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func AddSnapshotCron(hour, min string) {
	c := cron.New()
	cronTime := fmt.Sprintf("CRON_TZ=Asia/Seoul %s %s * * *", hour, min)
	c.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run at %s:%s Seoul time every day", hour, min))
	})
}

func AddSnapshotCronByHoursMin(vmName, hour, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@every %sh%sm0s", hour, min)
	id, err = CronSnap.Cr.AddFunc(cronTime, func() {
		fmt.Printf("Snapshot: %s - %s\n", vmName, cronTime)
		SafeSnapshot(vmName, GetTimeWord(), cronTime)
	})
	return id, err
}

func AddSnapshotCronByMin(vmName, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@every 0h%sm0s", min)
	id, err = CronSnap.Cr.AddFunc(cronTime, func() {
		fmt.Printf("Snapshot: %s - %s\n", vmName, cronTime)
		SafeSnapshot(vmName, GetTimeWord(), cronTime)
	})
	return id, err
}

func AddSnapshotCronBySecond(num string) {
	c := cron.New()
	cronTime := fmt.Sprintf("@every 0h0m%ss", num)
	c.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("every second - %s", num))
	})
	fmt.Println("start cron")
	fmt.Printf("Entry %+v\n", c.Entries())
	c.Start()
	time.Sleep(1 * time.Minute)
	fmt.Println("Stop cron")
	c.Stop()
	time.Sleep(1 * time.Minute)
}

func UpdateVmStatus(msg *messages.VmStatusActionMsg) {
	switch(msg.Status) {
	case "start":
		LibvirtStartVm(msg.VmName)
	case "stop":
		LibvirtDestroyVm(msg.VmName)
	case "restart":
		LibvirtResumeVm(msg.VmName)
	case "suspend":
		LibvirtSuspendVm(msg.VmName)
	case "resume":
		LibvirtResumeVm(msg.VmName)
	}
}