package lib

// From MC-Agent
const (
	WinUrlPrefix = "/api/v1/win"
	WinUrlHealth = "/health"
	WinUrlModifyConf = "/modifyConf"
	WinUrlAgentRestart = "/restart"
)

// To MC-Agent
const (
	ToMcUrlPrefix = "/api/v1/win"
	ToMcUrlHealth = ToMcUrlPrefix + "/health"
)
