package lib

const (
	SvcmgrApiPrefix  = "/v1"
	SvcmgrApiCode    = SvcmgrApiPrefix + "/codes"
	SvcmgrApiSubCode = SvcmgrApiPrefix + "/subcodes"
	SvcmgrApiDevice  = SvcmgrApiPrefix + "/device"
	SvcmgrApiDevices = SvcmgrApiPrefix + "/devices"
	SvcmgrApiLogin   = SvcmgrApiPrefix + "/auth"
	SvcmgrApiCompany = SvcmgrApiPrefix + "/companies"
	SvcmgrApiUser    = SvcmgrApiPrefix + "/users"
	SvcmgrApiMicro   = SvcmgrApiPrefix + "/micro"
)

// About Micro Cloud
const (
	// /v1/micro/servers
	SvcmgrApiMicroServer              = SvcmgrApiMicro + "/servers"
	SvcmgrApiMicroServerPaging        = SvcmgrApiMicro + "/servers-paging"
	SvcmgrApiMicroServerRegister      = SvcmgrApiMicroServer + "/register"
	SvcmgrApiMicroServerUnRegister    = SvcmgrApiMicroServer + "/register"
	SvcmgrApiMicroServerSearchCompany = SvcmgrApiMicroServer + "/search-company"
	// /v1/micro/vms
	SvcmgrApiMicroVm                  = SvcmgrApiMicro + "/vms"
	SvcmgrApiMicroVmRegister          = SvcmgrApiMicroVm + "/register"
	SvcmgrApiMicroVmUnRegister        = SvcmgrApiMicroVm + "/unregister"
	SvcmgrApiMicroVmUpdate            = SvcmgrApiMicroVm + "/update"
	SvcmgrApiMicroVmUpdateFromMc      = SvcmgrApiMicroVm + "/update-from-mc"
	SvcmgrApiMicroVmPaging            = SvcmgrApiMicro + "/vms-paging"
)
