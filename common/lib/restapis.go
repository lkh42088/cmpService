package lib

const (
	McUrlPrefix               = "/api/v1"
	McUrlResource             = "/resource"
	McUrlServer               = "/server"
	McUrlMon                  = "/mon"
	McUrlRegisterServer       = McUrlServer + "/register"
	McUrlUnRegisterServer     = McUrlServer + "/unregister"
	McUrlVm                   = "/vms"
	McUrlCreateVm             = McUrlVm + "/create"
	McUrlDeleteVm             = McUrlVm + "/delete"
	McUrlApplyVmAction        = McUrlVm + "/action"
	McUrlUpdateVmStatus       = McUrlVm + "/status"
	McUrlGetVmById            = McUrlVm + "/:id"
	McUrlGetVmSnapshot        = "/snapshot"
	McUrlAddVmSnapshot        = McUrlGetVmSnapshot + "/add"
	McUrlUpdateVmSnapshot     = McUrlGetVmSnapshot + "/update"
	McUrlDeleteVmSnapshot     = McUrlGetVmSnapshot + "/delete"
	McUrlDeleteVmSnapshotList = McUrlGetVmSnapshot + "/delete-entry-list"
	McUrlRecoveryVmSnapshot     = McUrlGetVmSnapshot + "/recovery"
	McUrlNetwork              = "/networks"
	McUrlNetworkAdd           = McUrlNetwork + "/add"
	McUrlNetworkDelete        = McUrlNetwork + "/delete"
	McUrlPublicIp             = McUrlPrefix + "/myip"
	McUrlVmInterfaceTraffic   = McUrlMon + "/interface/traffic/:mac"
	McUrlMonServer            = McUrlMon + "/server"

	// Window system api
	McUrlWinPrefix          = "/win"
	McUrlHealthCheckFromWin = McUrlWinPrefix + "/health"

	// KT Rest API : Storage
	KtUrlPrefix				  = "/kt/storage"
	KtUrlStorageInfo		  = "/info/get/:id"
	// Backup
	McUrlGetVmBackup = "/bakcup"
	McUrlDeleteVmBackupList = McUrlGetVmBackup + "/delete-entry-list"
)
