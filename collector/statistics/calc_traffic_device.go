package statistics

import (
	"encoding/json"
	"errors"
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"nubes/collector/collectdevice"
	"nubes/collector/influx"
	"nubes/collector/snmpapi"
	"nubes/common/lib"
	"sync"
	"time"
)

// Do not modify : changed influxdb table, mismatched convert form
type IfStat struct {
	time 			time.Time
	id 				string		`influx:ID`
	ip				string		`influx:IP`
	ifIndex			int64		`influx:ifIndex`
	ifDescr			string		`influx:ifDescr`
	ifHCInOctets	int64		`influx:ifHCInOctets`
	ifHCOutOctets	int64		`influx:ifHCOutOctets`
}

var lastSelectTime sync.Map

const StatCollectTime = 10
const periodTime = "5m30s"
const iftable = "iftable"
const trafficField = `"time","id","ip","ifindex","descr","hc-in-octets","hc-out-octets"`

func GetdataAtInflux(id string) *client.Response {
	dbname := iftable
	field := trafficField
	// 5m30s To collect more data
	where := fmt.Sprintf(`"id"='%s' AND time > now() - %s`, id, periodTime)
	res := influx.GetMeasurementsWithCondition(dbname, field, where)
	//lib.LogWarn("data count : %d\n", len(res.Results[0].Series[0].Values))
	return res
}

func MakeStructForStatistics(s *IfStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}
	s.id				= data[1].(string)
	s.ip 				= data[2].(string)
	s.ifIndex, _ 		= data[3].(json.Number).Int64()
	s.ifDescr 			= data[4].(string)
	s.ifHCInOctets, _ 	= data[5].(json.Number).Int64()
	s.ifHCOutOctets, _ 	= data[6].(json.Number).Int64()
	return nil
}

func ConvertTrafficData(res *client.Response) []IfStat {
	// Check response data
	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("InfluxDB Response Error : No Data\n")
		return nil
	}

	// Convert response data
	v := res.Results[0].Series[0].Values
	stat := make([]IfStat, len(v))
	var timenano time.Time
	id := v[0][1].(string)

	for i, data := range v {
		// select time check
		timenano, _ = time.Parse(time.RFC3339, data[0].(string))
		if tmp, ok := lastSelectTime.Load(id); ok != true {
			if tmp.(time.Time).Sub(timenano) > 0 {
				continue
			}
		}

		// make struct
		stat[i].time = timenano
		if err := MakeStructForStatistics(&stat[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			return nil
		}

		lastSelectTime.Store(id, timenano)
	}
	//lib.LogWarn("Last data : %d\n", stat[len(stat)-1].ifHCOutOctets)
	return stat
}

func CalcTrafficPer5Min(stat []IfStat) (IfStat, error) {
	if stat == nil {
		return IfStat{}, errors.New("CalcTrafficPer5Min() Error : No data\n")
	}

	dummyid := 0
	for i := 0; i < len(stat); i++{
		if stat[i].ip != "" {
			dummyid = i
			break
		}
	}
	//lib.LogWarn("dummyid = %d\n", dummyid)

	last := len(stat) - 1
	rxAvg := (stat[last].ifHCInOctets - stat[dummyid].ifHCInOctets) / StatCollectTime
	txAvg := (stat[last].ifHCOutOctets - stat[dummyid].ifHCOutOctets) / StatCollectTime
	if rxAvg < 0 || txAvg < 0 {
		return IfStat{}, errors.New("Data is overflow. Need to check.")
	}

	 return IfStat{
		id: stat[last].id,
		ip:	stat[last].ip,
		ifIndex: stat[last].ifIndex,
		ifDescr: stat[last].ifDescr,
		ifHCInOctets: rxAvg,
		ifHCOutOctets: txAvg,
	}, nil
}

func StoreAvgData(stat IfStat) error {
	name := "Average5min"
	tags := snmpapi.MakeTagForInfluxDB(collectdevice.ID(stat.id), stat.ip)
	//fields := snmpapi.MakeFieldForInfluxDB(stat)
	fields := map[string]interface{}{
		"ifindex"	: stat.ifIndex,
		"ifdescr"	: stat.ifDescr,
		"rxavg5min"	: stat.ifHCInOctets,
		"txavg5min"	: stat.ifHCOutOctets,
	}
	if err := snmpapi.AddBpToInflux(name, tags, fields); err != nil {
		return fmt.Errorf("StoreAvgData() Error: %s\n", err)
	}
	return nil
}

func CollectStatistics(id collectdevice.ID) {
	res := GetdataAtInflux(string(id))
	stat :=	ConvertTrafficData(res)
	if stat == nil {
		lib.LogWarn("ConvertTrafficData() is failed.\n")
		return
	}
	avg, err := CalcTrafficPer5Min(stat)
	if err != nil {
		lib.LogWarn("CalcTrafficPer5Min() Error : %s\n", err)
		return
	}
	err = StoreAvgData(avg)
	if err != nil {
		lib.LogWarn("StatisticsTraffic() Error : %s\n", err)
		return
	}
}

func ActiveStatistics() {
	devNum := len(snmpapi.SnmpDevTb.Devices)
	if devNum == 0 {
		lib.LogInfoln("It does not exist collectdevice!")
		return
	}

	var wg sync.WaitGroup
	wg.Add(devNum)

	for _, device := range snmpapi.SnmpDevTb.Devices {
		if "" == device.Device.Ip {
			wg.Add(-1)
			lib.LogInfo("go func --> %s %s\n", device.Device.Ip, "skip!!")
			continue
		}
		fmt.Println(device.Device.Ip)
		go func(device snmpapi.SnmpDevice) {
			defer wg.Done()
			CollectStatistics(device.Device.Id)
		}(device)
	}

	wg.Wait()
	//fmt.Printf("%s\n", influx.Influx.Bp.Points())
}

func NewTimeStruct() {
	table := snmpapi.SnmpDevTb
	for _, device := range table.Devices {
		lastSelectTime.Store(string(device.Device.Id), time.Now())
	}
}

func Start(parentwg *sync.WaitGroup) {
	// LastSelectTime Init
	NewTimeStruct()

	for {
		time.Sleep(StatCollectTime * time.Second)
		ActiveStatistics()
	}

	if parentwg != nil {
		parentwg.Done()
	}
}
