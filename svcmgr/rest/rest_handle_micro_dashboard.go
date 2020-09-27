package rest

import (
	"cmpService/common/mcmodel"
	conf "cmpService/svcmgr/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/influxdata/influxdb1-client/v2"
	"net/http"
	"sort"
	"time"
)

type DeviceCount struct {
	Total			int		`json:"total"`
	Operate			int		`json:"operate"`
}

const TOP_USAGE_COUNT = 5

type DeviceRank struct {
	Cpu 		[]mcmodel.CpuStatForRank	`json:"cpu"`
	Mem 		[]mcmodel.MemStatForRank	`json:"mem"`
	Disk 		[]mcmodel.DiskStatForRank	`json:"disk"`
	Traffic		[]mcmodel.VmIfStatForRank	`json:"traffic"`
}

// FOR MARIA-DB
func (h *Handler) GetServerTotalCount(c *gin.Context) {
	var deviceCount	DeviceCount
	count, operCount, err := h.db.GetServerTotalCount()
	deviceCount.Total = count
	deviceCount.Operate = operCount
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, deviceCount)
}

func (h *Handler) GetVmTotalCount(c *gin.Context) {
	var deviceCount DeviceCount
	count, operCount, err := h.db.GetVmTotalCount()
	deviceCount.Total = count
	deviceCount.Operate = operCount
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, deviceCount)
}

// FOR INFLUX-DB
func GetServerRank(c *gin.Context) {
	var result DeviceRank

	// GET SERVER CPU
	dbname := "cpu"
	field := `"time","serial_number","cpu","usage_idle"`
	query := fmt.Sprintf(`WHERE cpu = 'cpu-total' group by mac_address order by time desc limit 10`)
	resCpu := conf.GetMeasurementsWithQuery(dbname, field, query)
	result.Cpu = MakeAvgCpuData(*resCpu)

	// GET SERVER MEM
	dbname = "mem"
	field = `"time","serial_number","available","available_percent","total"`
	query = fmt.Sprintf(`group by mac_address order by time desc limit 10`)
	resMem := conf.GetMeasurementsWithQuery(dbname, field, query)
	result.Mem = MakeAvgMemData(*resMem)

	// GET SERVER DISK
	dbname = "disk"
	field = `"time","serial_number","device","total","used","used_percent"`
	query = fmt.Sprintf(`WHERE path = '/' group by mac_address order by time desc limit 10`)
	resDisk := conf.GetMeasurementsWithQuery(dbname, field, query)
	result.Disk = MakeAvgDiskData(*resDisk)

	// GET SERVER RX/TX
	dbname = "interface"
	field = `"time","serial_number","ifPhysAddress","ifInOctets","ifOutOctets"`
	query = fmt.Sprintf(`group by mac_address order by time desc limit 10`)
	resRxTx := conf.GetMeasurementsWithQuery(dbname, field, query)
	result.Traffic = MakeAvgRxTxData(*resRxTx)

	fmt.Printf("RANK RESULT : %+v\n", result)
	//c.JSON(http.StatusOK, result)
}

func MakeAvgCpuData(res client.Response) []mcmodel.CpuStatForRank {
	var rank []mcmodel.CpuStatForRank
	var tmp mcmodel.CpuStatForRank

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		return rank
	} else {
		// Collect response data
		for _, group := range res.Results[0].Series {
			total := 0.0
			avg := 0.0
			v := group.Values
			stat := make([]mcmodel.CpuStatForRank, len(v))
			var convTime time.Time
			for i, data := range v {
				convTime, _ = time.Parse(time.RFC3339, data[0].(string))
				stat[i].Time = convTime
				stat[i].SN = data[1].(string)
				stat[i].Cpu = data[2].(string)
				stat[i].UsageIdle, _ = data[3].(json.Number).Float64()
			}

			// Calc avg
			for _, data := range stat {
				total += data.UsageIdle
			}
			avg = total / float64(len(stat))
			tmp.Time = stat[0].Time
			tmp.SN = stat[0].SN
			tmp.Cpu = stat[0].Cpu
			tmp.UsageIdle = stat[0].UsageIdle
			tmp.Avg = avg
			rank = append(rank, tmp)
		}

		// Sorting
		sort.Slice(rank, func(i, j int) bool {
			return rank[i].Avg > rank[j].Avg
		})
		if len(rank) >= TOP_USAGE_COUNT {
			rank = rank[:5]
		}
	}
	return rank
}


func MakeAvgMemData(res client.Response) []mcmodel.MemStatForRank {
	var rank []mcmodel.MemStatForRank
	var tmp mcmodel.MemStatForRank

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		return rank
	} else {
		// Collect response data
		for _, group := range res.Results[0].Series {
			total := 0.0
			avg := 0.0
			v := group.Values
			stat := make([]mcmodel.MemStatForRank, len(v))
			var convTime time.Time
			for i, data := range v {
				convTime, _ = time.Parse(time.RFC3339, data[0].(string))
				stat[i].Time = convTime
				stat[i].SN = data[1].(string)
				stat[i].Available, _ = data[2].(json.Number).Float64()
				stat[i].AvailablePercent, _ = data[3].(json.Number).Float64()
				stat[i].Total, _ = data[4].(json.Number).Float64()
			}

			// Calc avg
			for _, data := range stat {
				total += data.Available
			}
			avg = total / float64(len(stat))
			tmp.Time = stat[0].Time
			tmp.SN = stat[0].SN
			tmp.Available = stat[0].Available
			tmp.AvailablePercent = stat[0].AvailablePercent
			tmp.Total = stat[0].Total
			tmp.Avg = avg
			rank = append(rank, tmp)
		}

		// Sorting
		sort.Slice(rank, func(i, j int) bool {
			return rank[i].Avg > rank[j].Avg
		})
		if len(rank) >= TOP_USAGE_COUNT {
			rank = rank[:5]
		}
	}
	return rank
}

func MakeAvgDiskData(res client.Response) []mcmodel.DiskStatForRank {
	var rank []mcmodel.DiskStatForRank
	var tmp mcmodel.DiskStatForRank

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		return rank
	} else {
		// Collect response data
		for _, group := range res.Results[0].Series {
			total := 0.0
			avg := 0.0
			v := group.Values
			stat := make([]mcmodel.DiskStatForRank, len(v))
			var convTime time.Time
			for i, data := range v {
				convTime, _ = time.Parse(time.RFC3339, data[0].(string))
				stat[i].Time = convTime
				stat[i].SN = data[1].(string)
				stat[i].Device = data[2].(string)
				stat[i].Total, _ = data[3].(json.Number).Float64()
				stat[i].Used, _ = data[4].(json.Number).Float64()
				stat[i].UsedPercent, _ = data[5].(json.Number).Float64()
			}

			// Calc avg
			for _, data := range stat {
				total += data.Used
			}
			avg = total / float64(len(stat))
			tmp.Time = stat[0].Time
			tmp.SN = stat[0].SN
			tmp.Device = stat[0].Device
			tmp.Total = stat[0].Total
			tmp.Used = stat[0].Used
			tmp.UsedPercent = stat[0].UsedPercent
			tmp.Avg = avg
			rank = append(rank, tmp)
		}

		// Sorting
		sort.Slice(rank, func(i, j int) bool {
			return rank[i].Used > rank[j].Used
		})
		if len(rank) >= TOP_USAGE_COUNT {
			rank = rank[:5]
		}
	}
	return rank
}

func MakeAvgRxTxData(res client.Response) []mcmodel.VmIfStatForRank {
	var rank []mcmodel.VmIfStatForRank
	var tmp mcmodel.VmIfStatForRank

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		return rank
	} else {
		// Collect response data
		for _, group := range res.Results[0].Series {
			var total int64
			var avg int64
			v := group.Values
			stat := make([]mcmodel.VmIfStatForRank, len(v))
			var convTime time.Time
			for i, data := range v {
				if data[2] == nil {	// nil mac-address
					continue
				}
				convTime, _ = time.Parse(time.RFC3339, data[0].(string))
				stat[i].Time = convTime
				stat[i].SN = data[1].(string)
				stat[i].IfPhysAddress = data[2].(string)
				stat[i].IfInOctets, _ = data[3].(json.Number).Int64()
				stat[i].IfOutOctets, _ = data[4].(json.Number).Int64()
			}

			// Calc avg
			lastIdx := len(stat) - 1
			lastValue := stat[lastIdx].IfInOctets + stat[lastIdx].IfOutOctets
			startValue := stat[0].IfInOctets + stat[0].IfOutOctets
			total = lastValue - startValue
			avg = total / int64(lastIdx + 1)

			tmp.Time = stat[lastIdx].Time
			tmp.SN = stat[lastIdx].SN
			tmp.IfPhysAddress = stat[lastIdx].IfPhysAddress
			tmp.IfInOctets = stat[lastIdx].IfInOctets
			tmp.IfOutOctets = stat[lastIdx].IfOutOctets
			tmp.Avg = avg
			rank = append(rank, tmp)
		}

		// Sorting
		sort.Slice(rank, func(i, j int) bool {
			return rank[i].Avg > rank[j].Avg
		})

		if len(rank) >= TOP_USAGE_COUNT {
			rank = rank[:5]
		}
	}
	return rank
}

