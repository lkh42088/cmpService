package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"cmpService/mcagent/config"
	"cmpService/mcagent/ktrest"
	"cmpService/mcagent/repo"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"github.com/robfig/cron/v3"
	"strconv"
	"sync"
	"time"
)

type CronBackup struct {
	Interval int
	Cr *cron.Cron
	Vms []BackupVm
}

type BackupVm struct {
	VmName string
	CronId cron.EntryID
	Timer string
}

var CronBack *CronBackup

//func (c *CronBackup) Start(parentwg *sync.WaitGroup) {
//	loop := 1
//	c.Cr.Start()
//	for {
//		c.Run()
//		time.Sleep(time.Duration(c.Interval * int(time.Second)))
//		loop += 1
//	}
//	parentwg.Done()
//}

//func (c *CronBackup) Run() {
//	//fmt.Printf("Entry %+v\n", c.Cr.Entries())
//}

var backupMutex sync.Mutex

func (c *CronBackup) LookupVm(vmName string) bool {
	for _, vm := range c.Vms {
		if vm.VmName == vmName {
			return true
		}
	}
	return false
}

func (c *CronBackup) DeleteVm(vmName string) bool {
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

func DeleteCronBackupEntry(id cron.EntryID) {
	CronBack.Cr.Remove(id)
}

func NewCronBackup(interval int) *CronBackup {
	return &CronBackup{
		Interval: interval,
	}
}

func ConfigCronBack() {
	c := NewCronBackup(5)
	SetCronBackup(c)
}

func SetCronBackup(c *CronBackup) {
	CronBack = c
	CronBack.Cr = cron.New()
	CronBack.Cr.AddFunc("*/5 * * * *", func() { fmt.Println("CRON BACKUP - every second 5.")})
}

func (c *CronBackup) Start(parentwg *sync.WaitGroup) {
	loop := 1
	c.Cr.Start()
	for {
		c.Run()
		time.Sleep(time.Duration(c.Interval * int(time.Second)))
		loop += 1
	}
	parentwg.Done()
}

func (c *CronBackup) Run() {
	//fmt.Printf("Entry %+v\n", c.Cr.Entries())
}

func AddBackupCronMonthly(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@monthly")
	id, err = CronBack.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run daily"))
		//SafeBackup(vmName, GetTimeWordForBackup(), "Monthly, backup by Cron")
	})
	return id, err
}

func AddBackupCronWeekly(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@weekly")
	id, err = CronBack.Cr.AddFunc(cronTime, func() {
		//SafeBackup(vmName, GetTimeWordForBackup(), "Weekly, backup by Cron")
	})
	return id, err
}

func AddBackupCronDaily(vmName string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@daily")
	id, err = CronBack.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run daily"))
		//SafeBackup(vmName, GetTimeWordForBackup(), "Daily, backup by Cron")
	})
	return id, err
}

func AddBackupCronDailyTime(vmName, hour, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("CRON_TZ=Asia/Seoul %s %s * * *", hour, min)
	fmt.Println("AddBackupCronDailyTime:", vmName, "-", cronTime)
	id, err = CronBack.Cr.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run at %s:%s Seoul time every day", hour, min))
		//SafeBackup(vmName, GetTimeWordForBackup(), "Daily Time, backup by Cron")
	})
	return id, err
}

func AddBackupCronPeriodically(vmName, hour, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@every %sh%sm0s", hour, min)
	id, err = CronBack.Cr.AddFunc(cronTime, func() {
		//SafeBackup(vmName, GetTimeWordForBackup(), "Periodically, backup by Cron")
	})
	return id, err
}

func AddBackupByMcVm(vm *mcmodel.McVm) {
	if vm.BackupType == ktrest.BackupDisable {
		return
	}

	if CronBack.LookupVm(vm.Name) {
		fmt.Println("AddBackupByMcVm: Already have backup config!")
		return
	}

	var id cron.EntryID
	var err error
	switch (vm.BackupDays) {
	case 2:
		/**************
		 * Daily
		 **************/
		id, err = AddBackupCronDailyTime(vm.Name,
			strconv.Itoa(vm.BackupHours),
			strconv.Itoa(vm.BackupMinutes))
		if err != nil {
			fmt.Println("AddBackupByMcVm: error", err)
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
		if vm.BackupHours == 0 && vm.BackupMinutes == 0  {
			fmt.Println("AddSnapshotByMcVm: hours 0, minutes 0 --> skip")
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
	CronBack.Vms = append(CronBack.Vms, entry)
}

func AddVmBackupByConfig(config *messages.BackupConfigMsg) {
	if CronBack.LookupVm(config.VmName) {
		fmt.Println("AddVmBackupByConifg: Already have backup config!")
		return
	}
	var id cron.EntryID
	var err error
	switch config.Type {
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
		fmt.Println("AddBackupConDailyTime: error", err)
		return
	}
	entry := BackupVm{
		VmName: config.VmName,
		CronId: id,
	}
	CronBack.Vms = append(CronBack.Vms, entry)
}

func UpdateVmBackupByConfig(config *messages.BackupConfigMsg) {
	if CronBack.LookupVm(config.VmName) == false {
		fmt.Println("UpdateVmBackupByConfig: dosn'thave backup config!")
		return
	}

	CronBack.DeleteVm(config.VmName)

	var id cron.EntryID
	var err error
	switch config.Type {
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
		fmt.Println("UpdateVmBackupByConfig: error", err)
		return
	}
	entry := BackupVm{
		VmName: config.VmName,
		CronId: id,
	}
	CronBack.Vms = append(CronBack.Vms, entry)
}

func DeleteVmBackupByConfig(config *messages.BackupConfigMsg) {
	if CronBack.LookupVm(config.VmName) == false {
		fmt.Println("DeleteVmBackupByConfig: dosen't have Backup config!")
		return
	}

	CronBack.DeleteVm(config.VmName)
}

func GetVmBackupAll() *[]messages.BackupEntry {
	var arr []messages.BackupEntry
	return &arr
}

func GetVmBackupByVmName(vmName string) *[]messages.BackupEntry {
	var arr []messages.BackupEntry
	return &arr
}

func AddBackupCronSecond(min, domName string) {
	var i int
	cronTime := fmt.Sprintf("@every 0h0m%ss", min)
	CronBack.Cr.AddFunc(cronTime, func() {
		//var backupName = fmt.Sprintf("%s-%d", domName, i)
		i += 1
		fmt.Println(fmt.Sprintf("every second - %s", min))
		//dom, err := GetDomainByName(domName)
		//if err != nil {
		//	fmt.Println("error get vm:", err)
		//	return
		//}
		//name, _ := dom.GetName()
		//fmt.Println("vm:", name)
		//SafeBackup(domName, GetTimeWordForBackup(), backupName)
	})
}

func GetTimeWordForBackup() string {
	t := time.Now()
	fmt.Println(t)
	return fmt.Sprintf("%d-%s-%d-%d-%d-%d",
		t.Year(),
		t.Month().String()[:3],
		t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func AddBackupCron(hour, min string) {
	c := cron.New()
	cronTime := fmt.Sprintf("CRON_TZ=Asia/Seoul %s %s * * *", hour, min)
	c.AddFunc(cronTime, func() {
		fmt.Println(fmt.Sprintf("run at %s:%s Seoul time every day", hour, min))
	})
}

func AddBackupCronByHoursMin(vmName, hour, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@every %sh%sm0s", hour, min)
	id, err = CronBack.Cr.AddFunc(cronTime, func() {
		fmt.Printf("Backup: %s - %s\n", vmName, cronTime)
		//SafeBackup(vmName, GetTimeWordForBackup(), cronTime)
		/*************************************
		 * Send Backup entry to svcmgr
		 *************************************/
	})
	return id, err
}

func AddBackupCronByMin(vmName, min string) (id cron.EntryID, err error){
	cronTime := fmt.Sprintf("@every 0h%sm0s", min)
	id, err = CronBack.Cr.AddFunc(cronTime, func() {
		fmt.Printf("Backup: %s - %s\n", vmName, cronTime)
		//SafeBackup(vmName, GetTimeWordForBackup(), cronTime)
		/*************************************
		 * Send Backup entry to svcmgr
		 *************************************/
	})
	return id, err
}

func AddBackupCronBySecond(num string) {
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

func RegisterRegularMsgForBackup() {
	cronTime := fmt.Sprintf("@every 0h0m%ss", "30")
	id, err := CronBack.Cr.AddFunc(cronTime, func() {
		var msg messages.ServerRegularMsg
		cfg := config.GetMcGlobalConfig()
		addr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)
		msg.SerialNumber = cfg.SerialNumber
		msg.Enable = repo.GlobalServerRepo.Enable
		msg.PrivateIp = cfg.ServerIp
		msg.PublicIp = cfg.ServerPublicIp
		msg.Port = cfg.McagentPort
		if repo.GlobalServerRepo.Enable {
			// Case 1: Send KeepAlive
			fmt.Printf("** RegisterRegularMsg(Cron ID:%d): Send keepalive msg to svcmgr\n", RegularSendToSvcmgrCronId)
			svcmgrapi.SendRegularMsg2Svcmgr(msg, addr, repo.GlobalServerRepo.Enable)

		} else {
			// Case 2: Send Registration
			fmt.Printf("** RegisterRegularMsg(Cron ID:%d): Send Registration msg to svcmgr\n", RegularSendToSvcmgrCronId)
			res := svcmgrapi.SendRegularMsg2Svcmgr(msg, addr, repo.GlobalServerRepo.Enable)
			if res == true {
				// Send ServerMsg
				svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)
				serverInfo := GetMcServerInfo()
				svcmgrapi.SendUpdateServer2Svcmgr(serverInfo, svcmgrRestAddr)
			}
		}
	})
	if err != nil {
		fmt.Println("RegisterRegularMsg: error ", err)
		return
	}
	RegularSendToSvcmgrCronId = id
	fmt.Println(">>> RegisterRegularMsg: Cron Id -", RegularSendToSvcmgrCronId)
}

