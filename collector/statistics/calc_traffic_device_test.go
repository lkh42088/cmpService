package statistics

import (
	"cmpService/collector/config"
	"cmpService/collector/influx"
	"cmpService/common/lib"
	"testing"
)
const homeid = "5eaeda991d41c85a06e98f1d"

func TestGetTraffic(t *testing.T) {
	config.CollectorConfigPath = "/etc/collector/collector.conf"
	config.ConfigureInfluxDB()
	res := GetdataAtInflux(homeid)
 	stat :=	ConvertTrafficData(res)
 	avg, _ := CalcTrafficPer5Min(stat)
 	StoreAvgData(avg)
	err := influx.Influx.Client.Write(influx.Influx.Bp)
	if err != nil {
		lib.LogWarn("InfluxDb Write Error: %s\n", err)
	}
}
