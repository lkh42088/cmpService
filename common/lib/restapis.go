package lib

const (
	McUrlPrefix             = "/api/v1"
	McUrlResource           = "/resource"
	McUrlServer             = "/server"
	McUrlMon                = "/mon"
	McUrlRegisterServer     = McUrlServer + "/register"
	McUrlUnRegisterServer   = McUrlServer + "/unregister"
	McUrlVm                 = "/vms"
	McUrlCreateVm           = McUrlVm + "/create"
	McUrlDeleteVm           = McUrlVm + "/delete"
	McUrlUpdateVmStatus     = McUrlVm + "/status"
	McUrlGetVmById          = McUrlVm + "/:id"
	McUrlGetVmSnapshot      = "/snapshot"
	McUrlAddVmSnapshot      = McUrlGetVmSnapshot + "/add"
	McUrlUpdateVmSnapshot   = McUrlGetVmSnapshot + "/update"
	McUrlDeleteVmSnapshot   = McUrlGetVmSnapshot + "/delete"
	McUrlNetwork            = "/networks"
	McUrlNetworkAdd         = McUrlNetwork + "/add"
	McUrlNetworkDelete      = McUrlNetwork + "/delete"
	McUrlPublicIp           = McUrlPrefix + "/myip"
	McUrlVmInterfaceTraffic = McUrlMon + "/interface/traffic/:mac"
	McUrlMonServer          = McUrlMon + "/server"

	// Window system api
	McUrlWinPrefix			= "/win"
	McUrlHealthCheckFromWin = McUrlWinPrefix + "/health"
)
