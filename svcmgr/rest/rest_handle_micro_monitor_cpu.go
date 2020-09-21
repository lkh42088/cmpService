package rest

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	conf "cmpService/svcmgr/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetVmInterfaceCpu(c *gin.Context) {
	mac := c.Param("mac")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("mac : ", mac)
	fmt.Println("")
	dbname := "cpu"
	field := `"time","cpu","usage_idle"`
	where := fmt.Sprintf(`cpu = 'cpu-total' AND mac_address = '%s'`, mac)
	//where := fmt.Sprintf(`"ifPhysAddress" =~ /.*%s/ AND time > now() - %s`, mac, "1h")
	res := conf.GetMeasurementsWithConditionOrderLimit(dbname, field, where)

	cpu := make([]models.CpuStat, 1)

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		fmt.Println("")
		lib.LogWarn("CPU InfluxDB Response Error : No Data\n")
		/*c.JSON(http.StatusInternalServerError, "No Data")
		return*/
		cpu[0].Err = "nodata"
		cpu[0].UsageIdle = 0
	} else {
		// Convert response data
		v := res.Results[0].Series[0].Values
		cpu := make([]models.CpuStat, len(v))
		var convTime time.Time
		for i, data := range v {
			// select time check
			convTime, _ = time.Parse(time.RFC3339, data[0].(string))

			// make struct
			cpu[i].Time = convTime
			if err := MakeStructForStatsCpu(&cpu[i], data); err != nil {
				lib.LogWarn("Error : %s\n", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
				return
			}
		}
	}

	//fmt.Printf("%+v\n", deltaStats)
	c.JSON(http.StatusOK, cpu)
}

func MakeStructForStatsCpu(s *models.CpuStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}

	s.Cpu = data[1].(string)
	s.Err = ""
	s.UsageIdle, _ = data[2].(json.Number).Float64()
	return nil
}

func MakeStructForStatsWinCpu(s *models.WinCpuStat, data []interface{}, mac string) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}

	//s.PercentIdleTime = data[1].(json.Number).(float64)
	s.PercentIdleTime, _ = data[1].(json.Number).Float64()

	/*fmt.Println("")
	fmt.Println(mac, "😡😡😡😡😡 CPU DATA 1 : ", s.PercentIdleTime)*/
	return nil
}
