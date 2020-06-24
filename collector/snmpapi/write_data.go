package snmpapi

import (
	"cmpService/collector/influx"
	"cmpService/common/lib"
	client "github.com/influxdata/influxdb1-client/v2"
	"sync"
	"time"
)

func WriteMetric(s *SnmpDeviceTable) {
	for _, dev := range s.Devices {
		tags := map[string]string{
			"serverip": dev.Device.Ip,
		}
		fields := map[string]interface{}{
			"cpu-idle": dev.Cpu.Idle,
			"cpu-1m":   dev.Cpu.min1av,
			"cpu-5m":   dev.Cpu.min5av,
			"cpu-10m":  dev.Cpu.min10av,
		}

		eventTime := time.Now().Add(time.Second * -20)
		point, err := client.NewPoint(
			"cpu",
			tags,
			fields,
			eventTime.Add(time.Second*10),
		)
		if err != nil {
			lib.LogWarn("Error: %s\n", err)
			return
		}
		influx.Influx.Bp.AddPoint(point)
	}

	err := influx.Influx.Client.Write(influx.Influx.Bp)
	if err != nil {
		lib.LogWarn("InfluxDb Write Error: %s\n", err)
	}
}

func MakeBpFromSnmpStruct(s *SnmpDeviceTable) {
	for i, dev := range s.Devices {
		// IfTable
		MakeBpForIfTable(i, &dev)
		// IpTable
		MakeBpForIpTable(i, &dev)
		// Cpu
		MakeBpForCpu(i, &dev)
	}
}

func WriteMetricInfluxDB(parentwg *sync.WaitGroup) {
	// Store data
	for {
		time.Sleep(5 * time.Second)
		err := influx.Influx.Client.Write(influx.Influx.Bp)
		if err != nil {
			lib.LogWarn("InfluxDb Write Error: %s\n", err)
		}
	}

	if parentwg != nil {
		parentwg.Done()
	}
}
