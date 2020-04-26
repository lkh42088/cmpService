package snmpapi

import (
	client "github.com/influxdata/influxdb1-client/v2"
	"nubes/collector/db/influx"
	"nubes/collector/lib"
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
		s.Store.Bp.AddPoint(point)
	}

	err := s.Store.Client.Write(s.Store.Bp)
	if err != nil {
		lib.LogWarn("Influxdb Write Error: %s\n", err)
		// Create database
		err := influx.InfluxdbCreateDB(s.Store.Database)
		if err != nil {
			lib.LogWarn("Error : $s\n", err)
			return
		}
	}
}
