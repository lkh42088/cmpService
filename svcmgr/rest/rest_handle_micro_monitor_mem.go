package rest

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	conf "cmpService/svcmgr/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetVmInterfaceMem(c *gin.Context) {
	mac := c.Param("mac")
	fmt.Println("-------------------------------------")
	fmt.Println("memory mac : ", mac)
	dbname := "mem"
	field := `"time","available","available_percent", total`
	where := fmt.Sprintf(`mac_address = '%s'`, mac)
	res := conf.GetMeasurementsWithConditionOrderLimit(dbname, field, where)

	mem := make([]mcmodel.MemStat, 1)

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		fmt.Println("")
		lib.LogWarn("MEM InfluxDB Response Error : No Data\n")

		mem[0].Err = "nodata"
		mem[0].Available = 0
		mem[0].AvailablePercent = 0
		mem[0].Total = 0
		/*c.JSON(http.StatusInternalServerError, "No Data")
		return*/
	} else {
		// Convert response data
		v := res.Results[0].Series[0].Values
		mem = make([]mcmodel.MemStat, len(v))
		var convTime time.Time
		for i, data := range v {
			// select time check
			convTime, _ = time.Parse(time.RFC3339, data[0].(string))

			// make struct
			mem[i].Time = convTime
			if err := MakeStructForStatsMem(&mem[i], data); err != nil {
				lib.LogWarn("Error : %s\n", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
				return
			}
		}
	}

	//fmt.Printf("%+v\n", deltaStats)
	c.JSON(http.StatusOK, mem)
}

func MakeStructForStatsMem(s *mcmodel.MemStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}

	s.Err = "indata"
	s.Available, _ = data[1].(json.Number).Float64()
	s.AvailablePercent, _ = data[2].(json.Number).Float64()
	s.Total, _ = data[3].(json.Number).Float64()
	return nil
}

func MakeStructForStatsWinMem(s *mcmodel.WinMemStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}

	s.AvailableBytes, _ = data[1].(json.Number).Float64()
	return nil
}
