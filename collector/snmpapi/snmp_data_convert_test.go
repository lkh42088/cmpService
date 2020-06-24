package snmpapi

import (
	"cmpService/collector/config"
	"sync"
	"testing"
)

func TestConvertSnmpData(t *testing.T) {
	config.SetConfig("/etc/collector/collector.conf")
	InitConfig()
	GetDevicesFromMongoDB(SnmpDevTb)
	CollectSnmpInfo()

	var wg *sync.WaitGroup
	WriteMetricInfluxDB(wg)
}
