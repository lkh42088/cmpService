package rest

import (
	"cmpService/common/mcmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeviceCount struct {
	Total			int		`json:"total"`
	Operate			int		`json:"operate"`
}

const TOP_USAGE_COUNT = 5

type DeviceTop5 struct {
	Cpu 		[TOP_USAGE_COUNT]mcmodel.CpuStat	`json:"cpu"`
	Mem 		[TOP_USAGE_COUNT]mcmodel.MemStat	`json:"mem"`
	Disk 		[TOP_USAGE_COUNT]mcmodel.DiskStat	`json:"disk"`
	Rx 			[TOP_USAGE_COUNT]mcmodel.VmIfStat	`json:"rx"`
	Tx 			[TOP_USAGE_COUNT]mcmodel.VmIfStat	`json:"tx"`
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
func GetVmTop5(c *gin.Context) {
	var result DeviceTop5
	//cpu := []mcmodel.CpuStat{}
	//mem := []mcmodel.MemStat{}
	//disk := []mcmodel.DiskStat{}
	//rx := []mcmodel.VmIfStat{}
	//tx := []mcmodel.VmIfStat{}

	//dbname := "interface"
	//field := `"time","hostname","ifDescr","ifPhysAddress","ifInOctets","ifOutOctets"`
	//where := fmt.Sprintf(`"time > now() - %s`, "1h")
	//res := conf.GetMeasurementsWithCondition(dbname, field, where)
	//fmt.Println(res.Err)
	//
	//stat := make([]mcmodel.VmIfStat, 1)
	//
	//if res.Results[0].Series == nil ||
	//	len(res.Results[0].Series[0].Values) == 0 {
	//	lib.LogWarn("InfluxDB Response Error : No Data\n")
	//	/*c.JSON(http.StatusInternalServerError, res.Err)
	//	return*/
	//	stat[0].IfDescr = ""
	//	stat[0].Hostname = ""
	//	stat[0].IfInOctets = 0
	//	stat[0].IfOutOctets = 0
	//} else {
	//	// Convert response data
	//	v := res.Results[0].Series[0].Values
	//	stat = make([]mcmodel.VmIfStat, len(v))
	//	var convTime time.Time
	//	for i, data := range v {
	//		// select time check
	//		convTime, _ = time.Parse(time.RFC3339, data[0].(string))
	//
	//		// make struct
	//		stat[i].Time = convTime
	//		if err := MakeStructForStats(&stat[i], data); err != nil {
	//			lib.LogWarn("Error : %s\n", err)
	//			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
	//			return
	//		}
	//	}
	//}

	//fmt.Printf("%+v\n", deltaStats)
	c.JSON(http.StatusOK, result)
}