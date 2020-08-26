package lib

const (
	McUrlPrefix             = "/api/v1"
	McUrlServer             = "/server"
	McUrlMon                = "/mon"
	McUrlRegisterServer     = McUrlServer + "/register"
	McUrlUnRegisterServer   = McUrlServer + "/unregister"
	McUrlVm                 = "/vms"
	McUrlCreateVm           = McUrlVm + "/create"
	McUrlDeleteVm           = McUrlVm + "/delete"
	McUrlGetVmById          = McUrlVm + "/:id"
	McUrlNetwork            = "/networks"
	McUrlNetworkAdd         = McUrlNetwork + "/add"
	McUrlNetworkDelete      = McUrlNetwork + "/delete"
	McUrlPublicIp           = McUrlPrefix + "/myip"
	McUrlVmInterfaceTraffic = McUrlMon + "/interface/traffic/:mac"
	McUrlMonServer          = McUrlMon + "/server"
)
