package lib

const SVCMGR_AGENT	= 1
const MC_AGENT		= 2
const WIN_AGENT		= 3

type ConfVariable struct {
	AgentType int    `json:agentType`
	IpAddr    string `json:ipAddr`
	FieldName string `json:fieldName`
	Value     string `json:value`
}

