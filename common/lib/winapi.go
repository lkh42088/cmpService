package lib

// From MC-Agent
const (
	WinUrlPrefix = "/api/v1/win"
	WinUrlHealth = WinUrlPrefix + "/health"
	WinUrlModifyConf = WinUrlPrefix + "/modifyConf"
	WinUrlAgentRestart = WinUrlPrefix + "/restart"
)

// To MC-Agent
const (
	ToMcUrlPrefix = "/api/v1/win"
	ToMcUrlHealth = ToMcUrlPrefix + "/health"
)
