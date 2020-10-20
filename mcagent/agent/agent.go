package agent

import (
	"cmpService/common/mcmodel"
	"cmpService/common/utils"
	config2 "cmpService/mcagent/config"
	"cmpService/mcagent/ddns"
	"cmpService/mcagent/ktrest"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/mcinflux"
	"cmpService/mcagent/mcrest"
	"cmpService/mcagent/repo"
	"cmpService/svcmgr/config"
	"fmt"
	"os"
	"sync"
)

func Start (config string) {
	var wg sync.WaitGroup

	wg.Add(5)

	if !config2.ApplyGlobalConfig(config) {
		return
	}

	if ! configure() {
		fmt.Println("Fatal: Failed configuration!")
		return
	}

	SetSysInfo()

	// Start Cron
	if kvm.CronSnap != nil {
		go kvm.CronSnap.Start(&wg)
	} else {
		wg.Done()
	}

	// Rest Api Server
	go mcrest.Start(&wg)

	// MonitorRoutine VMs
	if MonitorR != nil {
		go MonitorR.StartByVirsh(&wg)
	} else {
		wg.Done()
	}

	if kvm.CreateVmFsm != nil {
		go kvm.CreateVmFsm.Start(&wg)
	} else {
		wg.Done()
	}

	if kvm.LibvirtR != nil {
		go kvm.LibvirtR.Start(&wg)
	} else {
		wg.Done()
	}

	//if kvm.LibvirtS != nil {
	//	go kvm.LibvirtS.Start(&wg)
	//} else {
	//	wg.Done()
	//}

	/*********************************
	 * BareMetal system info
	 *********************************/
	SendSysInfo()

	/*********************************
	 * Apply Cron for snapshot
	 *********************************/
	ApplyCronForSnapshot()

	/*********************************
	 * cron for Register, health check
	 *********************************/
	kvm.RegisterRegularMsg()
	kvm.RegisterRegularMsgForBackup()

	/****************************************
	 * Check kt account & nas info for backup
	 ****************************************/
	CheckBackup()

	wg.Wait()
}

func processSerialNumberByConfigFile(sn string) {
	// 1. Get From DB
	server := repo.GetMcServer()
	if server != nil {
		config2.SetSerialNumber2GlobalConfig(server.SerialNumber)
		return
	}

	var newServer mcmodel.McServerDetail
	newServer.SerialNumber = sn
	mcrest.AddMcServer(newServer, false)
}

func processSerialNumber() bool {
	// 1. Get From DB
	server := repo.GetMcServer()
	if server != nil {
		config2.SetSerialNumber2GlobalConfig(server.SerialNumber)
		return true
	}

	// 2. Get From ETC file
	serverStatus := config2.GetServerStatus()
	if serverStatus.SerialNumber != "" {
		config2.SetSerialNumber2GlobalConfig(serverStatus.SerialNumber)
		// Write DB
		var msg mcmodel.McServerDetail
		msg.SerialNumber = serverStatus.SerialNumber
		msg.CompanyIdx = serverStatus.CompanyIdx
		msg.CompanyName = serverStatus.CompanyName
		repo.AddServer2Repo(&msg)
		return true
	}

	fmt.Println("processSerialNumber: it dose not Serial Number!")
	return false
}

func configure() bool {

	cfg := config2.GetMcGlobalConfig()
	/********************************
	 * Configure Mariadb
	 ********************************/
	db, err := config.SetMariaDB(cfg.MariaUser, cfg.MariaPassword, cfg.MariaDb,
		cfg.MariaIp, 3306)
	if err != nil {
		fmt.Println("Main: ERROR - ", err)
		return false
	}
	config2.SetDbOrm(db)

	/********************************
	 * Serial Number
	 ********************************/
	if cfg.SerialNumber == "" {
		processSerialNumber()
	} else {
		processSerialNumberByConfigFile(cfg.SerialNumber)
	}

	/********************************
	 * Apply DDNS
	 ********************************/
	mcServer := repo.GetMcServer()
	fmt.Println("configure: mcServer ")
	mcServer.Dump()
	if mcServer != nil && mcServer.Enable {
		ddns.ApplyDdns(mcServer.McServer)
	}

	/********************************
	 * Init Caching VMs
	 ********************************/
	repo.InitCachingVms()

	/********************************
	 * Clear DNAT Rule in iptables
	 ********************************/
	utils.DeleteAllDnat()

	if !mcinflux.ConfigureInfluxDB() {
		fmt.Println("Failed to configure influxdb!")
		return false
	}

	// ConfigureMonitoring MonitorRoutine
	//if ! ConfigureMonitoring() {
	//	fmt.Println("Failed to configure agent!")
	//	return false
	//}

	//ConfigureVmList()

	/********************************
	 * Config Cron
	 ********************************/
	kvm.ConfigCron()

	/********************************
	 * Config Create Vm FSM
	 ********************************/
	kvm.ConfigCreateVmFsm()

	/********************************
	 * Sync Resource with current info
	 ********************************/
	server := SyncRepoWithCurrentInfo()

	/********************************
	 * Monitoring Resource
	 ********************************/
	kvm.ConfigureLibvirtResource(server)

	/********************************
	 * Statistics
	 ********************************/
	kvm.ConfigureLibvirtStatstics()

	return true
}

func ApplyCronForSnapshot() {
	for _, vm := range repo.GlobalVmCache {
		if vm.SnapType == false {
			continue
		}
		// apply cron
		fmt.Println("ApplyCronForSnapshot: ", vm.Name)
		kvm.AddSnapshotByMcVm(&vm)
	}
}

func CheckBackup() {
	// Check Backup Directory
	cfg := config2.GetMcGlobalConfig()
	if _, err := os.Stat(cfg.VmBackupDir); err != nil {
		err = os.MkdirAll(cfg.VmBackupDir, 0755)
		if err != nil {
			return
		}
	}
	// Check KT Storage Account
	err := ktrest.CheckKtAccount()
	if err != nil {
		fmt.Println("** CheckKtAccount Error : ", err)
	}
	// Backup configuration : KT Storage
	if ktrest.ConfigurationForKtContainer() != nil {
		fmt.Printf("\n** KT Storage configuration is invalid.\n\n")
	}
	// Check NAS Info
	// CheckNasInfo()
}

