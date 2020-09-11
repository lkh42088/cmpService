package lib

const (
	SvcmgrApiPrefix      = "/v1"
	SvcmgrApiCode        = SvcmgrApiPrefix + "/codes"
	SvcmgrApiCodeSetting = SvcmgrApiPrefix + "/codesetting"
	SvcmgrApiSubCode     = SvcmgrApiPrefix + "/subcodes"
	SvcmgrApiDevice      = SvcmgrApiPrefix + "/device"
	SvcmgrApiDevices     = SvcmgrApiPrefix + "/devices"
	SvcmgrApiLogin       = SvcmgrApiPrefix + "/auth"
	SvcmgrApiCompany     = SvcmgrApiPrefix + "/companies"
	SvcmgrApiUser        = SvcmgrApiPrefix + "/users"
	SvcmgrApiMicro       = SvcmgrApiPrefix + "/micro"
)

// About Micro Cloud
const (
	// /v1/micro/servers
	SvcmgrApiMicroServer              = SvcmgrApiMicro + "/servers"
	SvcmgrApiMicroServerPaging        = SvcmgrApiMicro + "/servers-paging"
	SvcmgrApiMicroServerRegister      = SvcmgrApiMicroServer + "/register"
	SvcmgrApiMicroServerUnRegister    = SvcmgrApiMicroServer + "/unregister"
	SvcmgrApiMicroServerSearchCompany = SvcmgrApiMicroServer + "/search-company"
	SvcmgrApiMicroServerResource      = SvcmgrApiMicroServer + "/resource"
	// /v1/micro/vms
	/*<<<<<<< HEAD
		SvcmgrApiMicroVm             = SvcmgrApiMicro + "/vms"
		SvcmgrApiMicroVmRegister     = SvcmgrApiMicroVm + "/register"
		SvcmgrApiMicroVmUnRegister   = SvcmgrApiMicroVm + "/unregister"
		SvcmgrApiMicroVmUpdate       = SvcmgrApiMicroVm + "/update"
		SvcmgrApiMicroVmUpdateFromMc = SvcmgrApiMicroVm + "/update-from-mc"
		SvcmgrApiMicroVmVnc          = SvcmgrApiMicroVm + "/vnc"
		SvcmgrApiMicroVmPaging       = SvcmgrApiMicro + "/vms-paging"
		SvcmgrApiMicroVmGraph        = SvcmgrApiMicro + "/vms-graph"
	=======*/
	SvcmgrApiMicroVm               = SvcmgrApiMicro + "/vms"
	SvcmgrApiMicroVmRegister       = SvcmgrApiMicroVm + "/register"
	SvcmgrApiMicroVmUnRegister     = SvcmgrApiMicroVm + "/unregister"
	SvcmgrApiMicroVmUpdate         = SvcmgrApiMicroVm + "/update"
	SvcmgrApiMicroVmUpdateFromMc   = SvcmgrApiMicroVm + "/update-from-mc"
	SvcmgrApiMicroVmVnc            = SvcmgrApiMicroVm + "/vnc"
	SvcmgrApiMicroVmPaging         = SvcmgrApiMicro + "/vms-paging"
	SvcmgrApiMicroVmSnapshot       = SvcmgrApiMicroVm + "/snapshot"
	SvcmgrApiMicroVmSnapshotConfig = SvcmgrApiMicroVm + "/snapshot-config"
	SvcmgrApiMicroVmAddSnapshot    = SvcmgrApiMicroVm + "/snapshot/add"
	SvcmgrApiMicroVmUpdateSnapshot = SvcmgrApiMicroVm + "/snapshot/update"
	SvcmgrApiMicroVmDeleteSnapshot = SvcmgrApiMicroVm + "/snapshot/delete"
	SvcmgrApiMicroVmStatus         = SvcmgrApiMicroVm + "/status"
	SvcmgrApiMicroVmGraph          = SvcmgrApiMicro + "/vms-graph"
	//>>>>>>> d05604a9bd6a003bc0ebf44748c287c5f85fd51e
	// /v1/micro/images
	SvcmgrApiMicroImage       = SvcmgrApiMicro + "/images"
	SvcmgrApiMicroImagePaging = SvcmgrApiMicro + "/images-paging"
	// /v1/micro/networks
	SvcmgrApiMicroNetwork           = SvcmgrApiMicro + "/networks"
	SvcmgrApiMicroNetworkRegister   = SvcmgrApiMicroNetwork + "/register"
	SvcmgrApiMicroNetworkUnRegister = SvcmgrApiMicroNetwork + "/unregister"
	SvcmgrApiMicroNetworkPaging     = SvcmgrApiMicro + "/networks-paging"
	SvcmgrApiMicroVmMonitor         = SvcmgrApiMicro + "/monitor"
	SvcmgrApiMicroVmStats           = SvcmgrApiMicroVmMonitor + "/stats"
	SvcmgrApiMicroMonitorCPU        = SvcmgrApiMicro + "/monitor/cpu"
)
