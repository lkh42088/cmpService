package cron

import (
	"cmpService/common/messages"
	"cmpService/mcagent/kvm"
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func NewCronScheduler(interval int) *CronScheduler {
	return &CronScheduler{
		Interval: interval,
	}
}

func SetCronScheduler(c *CronScheduler) {
	CronSch = c
	CronSch.Cr = cron.New()
	CronSch.Cr.AddFunc("*/5 * * * *", func() { fmt.Println("CRON Scheduler- every second 5.")})
}

func (c *CronScheduler) Run() {
	//fmt.Printf("Entry %+v\n", c.Cr.Entries())
}

func AddSnapshotCronMonthly(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@monthly")
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run daily"))
		SafeSnapshot(vmName, GetTimeWord(), "Monthly, snapshot by Cron")
	})
	return id, err
}

func AddSnapshotCronWeekly(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@weekly")
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		SafeSnapshot(vmName, GetTimeWord(), "Weekly, snapshot by Cron")
	})
	return id, err
}

func AddSnapshotCronDaily(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@daily")
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run daily"))
		SafeSnapshot(vmName, GetTimeWord(), "Daily, snapshot by Cron")
	})
	return id, err
}

func AddSnapshotCronDailyTime(vmName, hour, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("CRON_TZ=Asia/Seoul %s %s * * *", hour, min)
	fmt.Println("AddSnapshotCronDailyTime:", vmName, "-", cronTime)
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run at %s:%s Seoul time every day", hour, min))
		SafeSnapshot(vmName, GetTimeWord(), "Daily Time, snapshot by Cron")
	})
	return id, err
}

func AddSnapshotCronPeriodically(vmName, hour, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@every %sh%sm0s", hour, min)
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		SafeSnapshot(vmName, GetTimeWord(), "Periodically, snapshot by Cron")
	})
	return id, err
}

func DeleteVmSnapshotByConfig(config *messages.SnapshotConfigMsg) {
	if CronSch.LookupSnapVm(config.VmName) == false {
		fmt.Println("DeleteVmSnapshotByConfig: dosen't have snapshot config!")
		return
	}

	CronSch.DeleteSnapVm(config.VmName)
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
	CronSch.Cr.AddFunc(cronTime, func() {
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
	return fmt.Sprintf("%d-%s-%d-%d-%d-%d",
		t.Year(),
		t.Month().String()[:3],
		t.Day(),
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
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Printf("Snapshot: %s - %s\n", vmName, cronTime)
		SafeSnapshot(vmName, GetTimeWord(), cronTime)
	})
	return id, err
}

func AddSnapshotCronByMin(vmName, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@every 0h%sm0s", min)
	id, err = CronSch.Cr.AddFunc(cronTime, func() {
		fmt.Printf("Snapshot: %s - %s\n", vmName, cronTime)
		SafeSnapshot(vmName, GetTimeWord(), cronTime)
		/*************************************
		 * Send snapshot entry to svcmgr
		 *************************************/
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
		kvm.LibvirtStartVm(msg.VmName)
	case "stop":
		kvm.LibvirtDestroyVm(msg.VmName)
	case "restart":
		kvm.LibvirtResumeVm(msg.VmName)
	case "suspend":
		kvm.LibvirtSuspendVm(msg.VmName)
	case "resume":
		kvm.LibvirtResumeVm(msg.VmName)
	}
}
