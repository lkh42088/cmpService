package rest

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"cmpService/svcmgr/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetVmInterfaceDisk(c *gin.Context) {
	dbname := "disk"
	field := `"time","device","fstype","path","total","used","used_percent"`
	where := fmt.Sprintf(`path = '/'`)
	res := config.GetMeasurementsWithConditionOrderLimit(dbname, field, where)

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("InfluxDB Response Error : No Data\n")
		c.JSON(http.StatusInternalServerError, "No Data")
		return
	}

	// Convert response data
	v := res.Results[0].Series[0].Values
	stat := make([]models.DiskStat, len(v))
	var convTime time.Time
	for i, data := range v {
		// select time check
		convTime, _ = time.Parse(time.RFC3339, data[0].(string))

		// make struct
		stat[i].Time = convTime
		if err := MakeStructForStatsDisk(&stat[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}

	//fmt.Printf("%+v\n", deltaStats)
	c.JSON(http.StatusOK, stat)
}

func MakeStructForStatsDisk(s *models.DiskStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}

	s.Device = data[1].(string)
	s.Fstype = data[2].(string)
	s.Path = data[3].(string)
	s.Total = data[4].(json.Number)
	s.Used = data[5].(json.Number)
	s.UsedPercent = data[6].(json.Number)
	return nil
}
