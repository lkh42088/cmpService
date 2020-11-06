package lib

const (
	McUrlPrefix                = "/api/v1"
	McUrlResource              = "/resource"
	McUrlServer                = "/server"
	McUrlMon                   = "/mon"
	McUrlRegisterServer        = McUrlServer + "/register"
	McUrlUnRegisterServer      = McUrlServer + "/unregister"
	McUrlVm                    = "/vms"
	McUrlCreateVm              = McUrlVm + "/create"
	McUrlDeleteVm              = McUrlVm + "/delete"
	McUrlApplyVmAction         = McUrlVm + "/action"
	McUrlUpdateVmStatus        = McUrlVm + "/status"
	McUrlGetVmById             = McUrlVm + "/:id"
	McUrlGetVmSnapshot         = "/snapshot"
	McUrlAddVmSnapshot         = McUrlGetVmSnapshot + "/add"
	McUrlUpdateVmSnapshot      = McUrlGetVmSnapshot + "/update"
	McUrlDeleteVmSnapshot      = McUrlGetVmSnapshot + "/delete"
	McUrlDeleteVmSnapshotList  = McUrlGetVmSnapshot + "/delete-entry-list"
	McUrlRecoveryVmSnapshot    = McUrlGetVmSnapshot + "/recovery"
	McUrlFilterIpAddress       = "/filter"
	McUrlAddFilterIpAddress    = McUrlFilterIpAddress + "/add"
	McUrlDeleteFilterIpAddress = McUrlFilterIpAddress + "/delete"
	McUrlNetwork               = "/networks"
	McUrlNetworkAdd            = McUrlNetwork + "/add"
	McUrlNetworkDelete         = McUrlNetwork + "/delete"
	McUrlPublicIp              = McUrlPrefix + "/myip"
	McUrlVmInterfaceTraffic    = McUrlMon + "/interface/traffic/:mac"
	McUrlMonServer             = McUrlMon + "/server"

	// Window system api
	McUrlWinPrefix          = "/win"
	McUrlHealthCheckFromWin = McUrlWinPrefix + "/health"
	McUrlWinSystemModifyConf = McUrlWinPrefix + "/modifyConf"
	McUrlWinAgentRestart 	= McUrlWinPrefix + "/restart"

	// KT Rest API : Storage
	KtUrlPrefix      = "/kt/storage"
	KtUrlStorageInfo = "/info/get/:id"
	// Backup
	McUrlGetVmBackup        = "/bakcup"
	McUrlDeleteVmBackupList = McUrlGetVmBackup + "/delete-entry-list"
	McUrlVmBackup           = "/bakcup"
	McUrlAddVmBackup        = McUrlVmBackup + "/add"
	McUrlDeleteVmBackup     = McUrlVmBackup + "/delete"
	McUrlRestoreVmBackup    = McUrlVmBackup + "/restore"
	McUrlUpdateVmBackup     = McUrlVmBackup + "/update"

	// System
	McUrlSystemModifyConf 	= "/modifyConf"
	McUrlAgentRestart 		= "/restart"

)
