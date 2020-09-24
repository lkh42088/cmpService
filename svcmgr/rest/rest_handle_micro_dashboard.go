package rest

import (
	"cmpService/common/mcmodel"
	conf "cmpService/svcmgr/config"
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
	Cpu 		[TOP_USAGE_COUNT]mcmodel.CpuStatForRank		`json:"cpu"`
	Mem 		[TOP_USAGE_COUNT]mcmodel.MemStatForRank		`json:"mem"`
	Disk 		[TOP_USAGE_COUNT]mcmodel.DiskStatForRank	`json:"disk"`
	Traffic		[TOP_USAGE_COUNT]mcmodel.VmIfStatForRank	`json:"traffic"`
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
	result.Cpu = MakeAvgCpuData(*resMem)

	// GET SERVER DISK
	dbname = "disk"
	field = `"time","serial_number","device","total","used","used_percent"`
	query = fmt.Sprintf(`WHERE path = '/' group by mac_address order by time desc limit 10`)
	resDisk := conf.GetMeasurementsWithQuery(dbname, field, query)
	result.Cpu = MakeAvgCpuData(*resDisk)

	// GET SERVER RX/TX
	dbname = "interface"
	field = `"time","serial_number","ifPhysAddress","ifInOctets","ifOutOctets"`
	query = fmt.Sprintf(`group by mac_address order by time desc limit 10`)
	resRxTx := conf.GetMeasurementsWithQuery(dbname, field, query)
	result.Traffic = MakeAvgRxTxData(*resRxTx)

	c.JSON(http.StatusOK, result)
}

func MakeAvgCpuData(res client.Response) [TOP_USAGE_COUNT]mcmodel.CpuStatForRank {
	var rank [TOP_USAGE_COUNT]mcmodel.CpuStatForRank
	var store []mcmodel.CpuStatForRank
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
				stat[i].UsageIdle = data[3].(float64)
			}

			// Calc avg
			for _, data := range stat {
				total += data.UsageIdle
			}
			avg = total / float64(len(stat))
			tmp.Time = stat[0].Time
			tmp.SN = stat[0].SN
			tmp.Cpu = stat[0].Cpu
			tmp.UsageIdle = avg
			store = append(store, tmp)
		}

		// Sorting
		sort.Slice(store, func(i, j int) bool {
			return store[i].UsageIdle > store[j].UsageIdle
		})

		for i := 0; i < TOP_USAGE_COUNT; i++ {
			rank[i] =  store[i]
		}
	}
	return rank
}

func MakeAvgRxTxData(res client.Response) [TOP_USAGE_COUNT]mcmodel.VmIfStatForRank {
	var rank [TOP_USAGE_COUNT]mcmodel.VmIfStatForRank
	var store []mcmodel.VmIfStatForRank
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
				convTime, _ = time.Parse(time.RFC3339, data[0].(string))
				stat[i].Time = convTime
				stat[i].SN = data[1].(string)
				stat[i].IfPhysAddress = data[2].(string)
				stat[i].IfInOctets = data[3].(int64)
				stat[i].IfOutOctets = data[4].(int64)
			}

			// Calc avg
			for _, data := range stat {
				total += data.IfInOctets + data.IfOutOctets
			}
			avg = total / int64(len(stat))
			tmp.Time = stat[0].Time
			tmp.SN = stat[0].SN
			tmp.IfPhysAddress = stat[0].IfPhysAddress
			tmp.IfInOctets = stat[0].IfInOctets
			tmp.IfOutOctets = stat[0].IfOutOctets
			tmp.Avg = avg
			store = append(store, tmp)
		}

		// Sorting
		sort.Slice(store, func(i, j int) bool {
			return store[i].Avg > store[j].Avg
		})

		for i := 0; i < TOP_USAGE_COUNT; i++ {
			rank[i] =  store[i]
		}
	}
	return rank
}
