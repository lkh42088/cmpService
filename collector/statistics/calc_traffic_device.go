package statistics

import (
	"errors"
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"nubes/collector/collectdevice"
	"nubes/collector/influx"
	"nubes/collector/snmpapi"
	"nubes/common/lib"
	"strconv"
	"sync"
	"time"
)

// Do not modify : changed influxdb table, mismatched convert form
type IfStat struct {
	//time 			string		`influx:time`
	id 				string		`influx:ID`
	ip				string		`influx:IP`
	ifIndex			string		`influx:ifIndex`
	ifDescr			string		`influx:ifDescr`
	ifHCInOctets	uint64		`influx:ifHCInOctets`
	ifHCOutOctets	uint64		`influx:ifHCOutOctets`
}

var Avg []IfStat

const iftable = "IfTable"
const trafficField = `"ID","IP","ifIndex","ifDescr","ifHCInOctets","ifHCOutOctets"`
const StatCollecTime = 300

func GetdataAtInflux(id string) *client.Response {
	dbname := iftable
	field := trafficField
	where := fmt.Sprintf(`"ID"='%s' AND time >= now() - 5m`, id)
	res := influx.GetMeasurementsWithCondition(dbname, field, where)
	//lib.LogWarn("data count : %d\n", len(res.Results[0].Series[0].Values))
	return res
}

func ConvertTrafficData(res *client.Response) []IfStat {
	if res.Results[0].Series == nil {
		lib.LogWarn("InfluxDB Response Error : No Data\n")
		return nil
	}
	// Convert Response Data
	v := res.Results[0].Series[0].Values
	stat := make([]IfStat, len(v))
	for i, data := range v {
		//stat[i].time = data[0].(string)	//remove time data
		stat[i].id = data[1].(string)
		stat[i].ip = data[2].(string)
		stat[i].ifIndex = data[3].(string)
		stat[i].ifDescr = data[4].(string)
		num, _ := strconv.Atoi(data[5].(string))
		stat[i].ifHCInOctets = uint64(num)
		num, _ = strconv.Atoi(data[6].(string))
		stat[i].ifHCOutOctets = uint64(num)
	}
	//lib.LogWarn("Last data : %d\n", stat[len(stat)-1].ifHCOutOctets)
	return stat
}

func CalcTrafficPer5Min(stat []IfStat) (IfStat, error) {
	if stat == nil {
		return IfStat{}, errors.New("CalcTrafficPer5Min() Error : No data\n")
	}

	rxAvg := (stat[len(stat)-1].ifHCInOctets - stat[0].ifHCInOctets) / StatCollecTime
	txAvg := (stat[len(stat)-1].ifHCOutOctets - stat[0].ifHCOutOctets) / StatCollecTime
	if rxAvg < 0 || txAvg < 0 {
		return IfStat{}, errors.New("Data is overflow. Need to check.")
	}

	 return IfStat{
		id: stat[0].id,
		ip:	stat[0].ip,
		ifIndex: stat[0].ifIndex,
		ifDescr: stat[0].ifDescr,
		ifHCInOctets: rxAvg,
		ifHCOutOctets: txAvg,
	}, nil
}

func StoreAvgData(stat IfStat) error {
	name := "Average5min"
	tags := snmpapi.MakeTagForInfluxDB(collectdevice.ID(stat.id), stat.ip)
	fields := snmpapi.MakeFieldForInfluxDB(stat)
	if err := snmpapi.AddBpToInflux(name, tags, fields); err != nil {
		return fmt.Errorf("StoreAvgData() Error: %s\n", err)
	}
	return nil
}

func CollectStatistics(id collectdevice.ID) {
	res := GetdataAtInflux(string(id))
	stat :=	ConvertTrafficData(res)
	avg, err := CalcTrafficPer5Min(stat)
	if err != nil {
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
	//var wg sync.WaitGroup
	//wg.Add(devNum)
	for _, device := range snmpapi.SnmpDevTb.Devices {
		if "" == device.Device.Ip {
			//wg.Add(-1)
			lib.LogInfo("go func --> %s %s\n", device.Device.Ip, "skip!!")
			continue
		}
		go func(device snmpapi.SnmpDevice) {
			//defer wg.Done()
			CollectStatistics(device.Device.Id)
		}(device)
	}
	//wg.Wait()
	//fmt.Printf("%s\n", influx.Influx.Bp.Points())
}

func Start(parentwg *sync.WaitGroup) {
	for {
		time.Sleep(StatCollecTime * time.Second)
		ActiveStatistics()
	}

	if parentwg != nil {
		parentwg.Done()
	}
}
