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
	SvcmgrApiMicroServer                    = SvcmgrApiMicro + "/servers"
	SvcmgrApiMicroServerPaging              = SvcmgrApiMicro + "/servers-paging"
	SvcmgrApiMicroServerRegister            = SvcmgrApiMicroServer + "/register"
	SvcmgrApiMicroServerUnRegister          = SvcmgrApiMicroServer + "/unregister"
	SvcmgrApiMicroServerSearchCompany       = SvcmgrApiMicroServer + "/search-company"
	SvcmgrApiMicroServerResource            = SvcmgrApiMicroServer + "/resource"
	SvcmgrApiMicroServerRegularMsg          = SvcmgrApiMicroServer + "/regular-msg"
	SvcmgrApiMicroSystemInfo                = SvcmgrApiMicro + "/sysinfo"
	SvcmgrApiMicroVm                        = SvcmgrApiMicro + "/vms" // /v1/micro/vms
	SvcmgrApiMicroVmRegister                = SvcmgrApiMicroVm + "/register"
	SvcmgrApiMicroVmUnRegister              = SvcmgrApiMicroVm + "/unregister"
	SvcmgrApiMicroVmUpdate                  = SvcmgrApiMicroVm + "/update"
	SvcmgrApiMicroVmUpdateFromMc            = SvcmgrApiMicroVm + "/update-from-mc"
	SvcmgrApiMicroVmUpdateFromMcSnapshot    = SvcmgrApiMicroVm + "/update-from-mc/snapshot"
	SvcmgrApiMicroVmVnc                     = SvcmgrApiMicroVm + "/vnc"
	SvcmgrApiMicroVmPaging                  = SvcmgrApiMicro + "/vms-paging"
	SvcmgrApiMicroVmAction                  = SvcmgrApiMicroVm + "/action"
	SvcmgrApiMicroVmSnapshotPaging          = SvcmgrApiMicro + "/snapshot-paging"
	SvcmgrApiMicroVmSnapshotConfig          = SvcmgrApiMicroVm + "/snapshot-config"
	SvcmgrApiMicroVmAddSnapshot             = SvcmgrApiMicroVm + "/snapshot/add"
	SvcmgrApiMicroVmUpdateSnapshot          = SvcmgrApiMicroVm + "/snapshot/update"
	SvcmgrApiMicroVmDeleteSnapshot          = SvcmgrApiMicroVm + "/snapshot/delete"
	SvcmgrApiMicroVmDeleteSnapshotEntryList = SvcmgrApiMicroVm + "/snapshot/delete-entry-list"
	/*SvcmgrApiMicroVmSnapshotCount           = SvcmgrApiMicroVm + "/snapshot/count"*/
	SvcmgrApiMicroVmRecoverySnapshot    = SvcmgrApiMicroVm + "/snapshot/recovery"
	SvcmgrApiMicroVmStatus              = SvcmgrApiMicroVm + "/status"
	SvcmgrApiMicroVmGraph               = SvcmgrApiMicro + "/vms-graph"
	SvcmgrApiMicroMcAgentNotifySnapshot = SvcmgrApiMicroVm + "/mcagent/snapshot/notify" // snapshot from mcagent
	SvcmgrApiMicroImage                 = SvcmgrApiMicro + "/images"                    // /v1/micro/images
	SvcmgrApiMicroImagePaging           = SvcmgrApiMicro + "/images-paging"
	SvcmgrApiMicroNetwork               = SvcmgrApiMicro + "/networks" // /v1/micro/networks
	SvcmgrApiMicroNetworkRegister       = SvcmgrApiMicroNetwork + "/register"
	SvcmgrApiMicroNetworkUnRegister     = SvcmgrApiMicroNetwork + "/unregister"
	SvcmgrApiMicroNetworkPaging         = SvcmgrApiMicro + "/networks-paging"
	SvcmgrApiMicroVmMonitor             = SvcmgrApiMicro + "/monitor"
	SvcmgrApiMicroVmStats               = SvcmgrApiMicroVmMonitor + "/stats"
	SvcmgrApiMicroVmInfo                = SvcmgrApiMicroVmMonitor + "/info"
	SvcmgrApiMicroMonitorCPU            = SvcmgrApiMicro + "/monitor/cpu"
	SvcmgrApiMicroDashboard             = SvcmgrApiMicro + "/dashboard" // /v1/micro/dashboard
	SvcmgrApiMicroTotalCount            = SvcmgrApiMicroDashboard + "/total-cnt"
	SvcmgrApiMicroVmCount               = SvcmgrApiMicroDashboard + "/vmcnt"
	SvcmgrApiMicroServerRank            = SvcmgrApiMicroDashboard + "/rank"
	SvcmgrApiMicroSnapshotCount         = SvcmgrApiMicroDashboard + "/snapshotcnt"
	SvcmgrApiMicroVmCheckUser           = SvcmgrApiMicroVm + "/check/user"
	SvcmgrApiMicroKtAuthUrl             = SvcmgrApiMicro + "/auth/url"

	// /v1/micro/vms/check/user
)
