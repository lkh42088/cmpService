package snmpapi

import (
	"nubes/collector/config"
	"testing"
)

func TestConvertSnmpData(t *testing.T) {
	config.SetConfig("/etc/collector/collector.conf")
	InitConfig()
	GetDevicesFromMongoDB(SnmpDevTb)
	CollectSnmpInfo()
	WriteMetricFromStruct(SnmpDevTb)
}