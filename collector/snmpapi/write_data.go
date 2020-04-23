package snmpapi

import (
	client "github.com/influxdata/influxdb1-client/v2"
	"nubes/collector/lib"
	"time"
)

func WriteMetric(s *SnmpDeviceTable) {
	for _, dev := range s.devices {
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
		s.store.Bp.AddPoint(point)
	}

	err := s.store.Client.Write(s.store.Bp)
	if err != nil {
		lib.LogWarn("Influxdb Write Error: %s\n", err)
	}
}
