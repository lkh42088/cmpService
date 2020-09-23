package rest

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	conf "cmpService/svcmgr/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/influxdata/influxdb1-client/v2"
	"math"
	"net/http"
	"time"
)

const CPU_FIELD = `"time", "Percent_Idle_Time"`
const MEM_FIELD = `"time", "Available_Bytes"`
const DISK_FIELD = `"time", "Free_Megabytes"`
const RX_FIELD = `"time","Bytes_Received_persec"`
const TX_FIELD = `"time","Bytes_Sent_persec"`
const CPU_INSTANCE = "_Total"
const DISK_INSTANCE = "C:"
const EMPTY_INSTANCE = ""
const INTERVAL = "10m"

func ChangeWhereString(instance string, mac string) string {
	where := fmt.Sprintf(`"instance" =~ /.*%s/ AND "mac_address" =~ /.*%s/ AND time > now() - %s`, instance, mac, INTERVAL)
	return where
}

func GetVmInfoStats(c *gin.Context) {
	var result mcmodel.VmInfoStatsResponse
	mac := c.Param("mac")

	// GET CPU
	dbname := "win_cpu"
	where := ChangeWhereString(CPU_INSTANCE, mac)
	resCpu := conf.GetMeasurementsWithCondition(dbname, CPU_FIELD, where)
	result.VmCpu = MakeDataSet("CPU", resCpu)

	// GET MEM
	dbname = "win_mem"
	where = ChangeWhereString(EMPTY_INSTANCE, mac)
	resMem := conf.GetMeasurementsWithCondition(dbname, MEM_FIELD, where)
	result.VmMem = MakeDataSet("MEM", resMem)

	// GET DISK
	dbname = "win_disk"
	where = ChangeWhereString(DISK_INSTANCE, mac)
	resDisk := conf.GetMeasurementsWithCondition(dbname, DISK_FIELD, where)
	result.VmDisk = MakeDataSet("DISK", resDisk)

	// GET RX/TX
	dbname = "win_net"
	where = ChangeWhereString(EMPTY_INSTANCE, mac)
	resRx := conf.GetMeasurementsWithCondition(dbname, RX_FIELD, where)
	result.VmRx = MakeDataSet("RX", resRx)
	resTx := conf.GetMeasurementsWithCondition(dbname, TX_FIELD, where)
	result.VmTx = MakeDataSet("TX", resTx)

	//fmt.Printf("%+v\n", result)
	c.JSON(http.StatusOK, result)
}

func MakeDataSet(dataType string, res *client.Response) mcmodel.VmStatsSet {
	var response mcmodel.VmStatsSet
	response.Id = dataType

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("InfluxDB Response Error : No Data\n")
		return response
	}

	v := res.Results[0].Series[0].Values
	var convTime time.Time
	for _, data := range v {
		// select time check
		convTime, _ = time.Parse(time.RFC3339, data[0].(string))
		val, _ := data[1].(json.Number).Float64()
		data := mcmodel.Stats{
			Xaxis: convTime,
			Yaxis: int64(math.Round(val)),
		}
		response.Data = append(response.Data, data)
	}

	//fmt.Printf("MakeDataSet: %+v\n", response)
	return response
}
