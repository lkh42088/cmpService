package snmpapi

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"cmpService/collector/collectdevice"
	"cmpService/collector/influx"
	"cmpService/common/lib"
	"reflect"
	"strings"
	"sync"
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
	var mutex = &sync.Mutex{}

	// Add batch point
	eventTime := time.Now().Add(time.Second * -20)
	point, err := client.NewPoint(
		strings.ToLower(name),
		tags,
		fields,
		eventTime.Add(time.Second * 10),
	)
	if err != nil {
		return fmt.Errorf("Error: %s\n", err)
	}

	mutex.Lock()
	influx.Influx.Bp.AddPoint(point)
	mutex.Unlock()

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
		//fields := MakeFieldForInfluxDB(dev.IfTable.ifEntry[j])
		fields := map[string]interface{}{
			"ifindex" 		: dev.IfTable.ifEntry[j].ifIndex,
			"descr" 		: dev.IfTable.ifEntry[j].ifDescr,
			"type" 			: dev.IfTable.ifEntry[j].ifType,
			"mtu" 			: dev.IfTable.ifEntry[j].ifMTU,
			"speed" 		: dev.IfTable.ifEntry[j].ifSpeed,
			"physaddress" 	: dev.IfTable.ifEntry[j].ifPhysAddress,
			"adminstatus" 	: dev.IfTable.ifEntry[j].ifAdminStatus,
			"operstatus" 	: dev.IfTable.ifEntry[j].ifOperStatus,
			"lastchange" 	: dev.IfTable.ifEntry[j].ifLastChange,
			"in-octets" 	: dev.IfTable.ifEntry[j].ifInOctets,
			"in-ucastpkts" 	: dev.IfTable.ifEntry[j].ifInUcastPkts,
			"in-n-ucastpkts" : dev.IfTable.ifEntry[j].ifInNUcastPkts,
			"in-discards" 	: dev.IfTable.ifEntry[j].ifInDiscards,
			"in-errors" 	: dev.IfTable.ifEntry[j].ifInErrors,
			"out-octets" 	: dev.IfTable.ifEntry[j].ifOutOctets,
			"out-ucastpkts" : dev.IfTable.ifEntry[j].ifOutUcastPkts,
			"out-n-ucastpkts" : dev.IfTable.ifEntry[j].ifOutNUcastPkts,
			"out-discards" 	: dev.IfTable.ifEntry[j].ifOutDiscards,
			"out-errors" 	: dev.IfTable.ifEntry[j].ifOutErrors,
			"out-qlen" 		: dev.IfTable.ifEntry[j].ifOutQLen,
			"ifspecific" 	: dev.IfTable.ifEntry[j].ifSpecific,
			"ifname" 		: dev.IfTable.ifEntry[j].ifName,
			"hc-in-octets" 	: dev.IfTable.ifEntry[j].ifHCInOctets,
			"hc-out-octets" : dev.IfTable.ifEntry[j].ifHCOutOctets,
		}

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
		//fields := MakeFieldForInfluxDB(dev.IpTable.IpList[j])
		fields := map[string]interface{}{
			"ipaddr"	: dev.IpTable.IpList[j].IpAddr,
			"ifindex"	: dev.IpTable.IpList[j].IfIndex,
			"newmask"	: dev.IpTable.IpList[j].NetMask,
		}

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
	//fields := MakeFieldForInfluxDB(dev.Cpu)
	fields := map[string]interface{}{
		"idle" 		: dev.Cpu.Idle,
		"min1av" 	: dev.Cpu.min1av,
		"min5av" 	: dev.Cpu.min5av,
		"min10av" 	: dev.Cpu.min10av,
	}

	// Add batch point
	if AddBpToInflux(name, tags, fields) != nil {
		lib.LogWarn("Failed to store CPU info.")
	}
}

