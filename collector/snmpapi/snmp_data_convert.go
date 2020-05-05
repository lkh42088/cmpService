package snmpapi

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"nubes/collector/collectdevice"
	"nubes/collector/influx"
	"nubes/common/lib"
	"reflect"
	"time"
)

// Not use
func RemoveUnnecessaryTable(name string) bool {
	switch name {
	case "Device":
		return true
	case "Snmp":
		return true
	default:
		return false
	}
}

func MakeTagForInfluxDB(id collectdevice.ID, ip string) map[string]string {
	return map[string]string{
		"id": string(id),
		"ip": ip,
	}
}

func MakeFieldForInfluxDB(data interface{}) map[string]interface{} {
	elements := reflect.ValueOf(data)
	target := elements.Type()
	fields := map[string]interface{}{}
	for i := 0; i < target.NumField(); i++ {
		fields[target.Field(i).Name] = elements.Field(i)
	}
	return fields
}

func AddBpToInflux(name string,
	tags map[string]string, fields map[string]interface{}) error {

	// Add batch point
	eventTime := time.Now().Add(time.Second * -20)
	point, err := client.NewPoint(
		name,
		tags,
		fields,
		eventTime.Add(time.Second * 10),
	)
	if err != nil {
		return fmt.Errorf("Error: %s\n", err)
	}
	influx.Influx.Bp.AddPoint(point)
	return nil
}

// Add SNMP:IFTABLE Batch point
func MakeBpForIfTable(id collectdevice.ID, dev *SnmpDevice) {
	for j := 0; j < len(dev.IfTable.ifEntry); j++ {
		// remove loopback interface
		if dev.IfTable.ifEntry[j].ifName == "lo" {
			continue
		}

		// Table name, tags, fields
		name := reflect.TypeOf(dev.IfTable).Name()
		tags := MakeTagForInfluxDB(id, dev.Device.Ip)
		fields := MakeFieldForInfluxDB(dev.IfTable.ifEntry[j])

		// Add batch point
		if AddBpToInflux(name, tags, fields) != nil {
			lib.LogWarn("Failed to store IfTable info.")
		}
	}
}

// Add SNMP:IPTABLE Batch point
func MakeBpForIpTable(id collectdevice.ID, dev *SnmpDevice) {
	for j := 0; j < len(dev.IpTable.IpList); j++ {
		// Table name, tags, fields
		name := reflect.TypeOf(dev.IpTable).Name()
		tags := MakeTagForInfluxDB(id, dev.Device.Ip)
		fields := MakeFieldForInfluxDB(dev.IpTable.IpList[j])

		// Add batch point
		if AddBpToInflux(name, tags, fields) != nil {
			lib.LogWarn("Failed to store IpTable info.")
		}
	}
}

// Add SNMP:CPU Batch point
func MakeBpForCpu(id collectdevice.ID, dev *SnmpDevice) {
	// Table name, tags, fields
	name := reflect.TypeOf(dev.Cpu).Name()
	tags := MakeTagForInfluxDB(id, dev.Device.Ip)
	fields := MakeFieldForInfluxDB(dev.Cpu)

	// Add batch point
	if AddBpToInflux(name, tags, fields) != nil {
		lib.LogWarn("Failed to store CPU info.")
	}
}

