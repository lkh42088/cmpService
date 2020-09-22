package agent

import (
	"cmpService/common/utils"
	config2 "cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/mcinflux"
	"cmpService/mcagent/mcrest"
	"cmpService/mcagent/repo"
	"cmpService/svcmgr/config"
	"fmt"
	"sync"
)

func Start (config string) {
	var wg sync.WaitGroup

	if !config2.ApplyGlobalConfig(config) {
		return
	}

	if ! configure() {
		fmt.Println("Fatal: Failed configuration!")
		return
	}

	SetSysInfo()

	wg.Add(4)

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

	// BareMetal system info
	SendSysInfo()

	wg.Wait()
}

func configure() bool {

	/********************************
	 * Configure Mariadb
	 ********************************/
	cfg := config2.GetMcGlobalConfig()
	db, err := config.SetMariaDB(cfg.MariaUser, cfg.MariaPassword, cfg.MariaDb,
		cfg.MariaIp, 3306)
	if err != nil {
		fmt.Println("Main: ERROR - ", err)
		return false
	}
	config2.SetDbOrm(db)

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
