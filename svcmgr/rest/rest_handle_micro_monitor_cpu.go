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
	dbname := "cpu"
	field := `"time","cpu","usage_idle"`
	where := fmt.Sprintf(`cpu = 'cpu-total'`)
	res := conf.GetMeasurementsWithConditionOrderLimit(dbname, field, where)

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("InfluxDB Response Error : No Data\n")
		c.JSON(http.StatusInternalServerError, "No Data")
		return
	}

	// Convert response data
	v := res.Results[0].Series[0].Values
	stat := make([]models.CpuStat, len(v))
	var convTime time.Time
	for i, data := range v {
		// select time check
		convTime, _ = time.Parse(time.RFC3339, data[0].(string))

		// make struct
		stat[i].Time = convTime
		if err := MakeStructForStatsCpu(&stat[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}

	//fmt.Printf("%+v\n", deltaStats)
	c.JSON(http.StatusOK, stat)
}

func MakeStructForStatsCpu(s *models.CpuStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}

	s.Cpu = data[1].(string)
	s.UsageIdle = data[2].(json.Number)
	return nil
}
