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
	mac := c.Param("mac")
	dbname := "disk"
	field := `"time","device","fstype","path","total","used","used_percent"`
	//where := fmt.Sprintf(`path = '/'`)
	where := fmt.Sprintf(`path = '/' AND mac_address = '%s'`, mac)
	res := config.GetMeasurementsWithConditionOrderLimit(dbname, field, where)

	disk := make([]models.DiskStat, 1)

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		fmt.Println("")
		lib.LogWarn("DISK InfluxDB Response Error : No Data\n")

		disk[0].Err = "nodata"
		disk[0].Device = ""
		disk[0].Fstype = ""
		disk[0].Path = "/"
		disk[0].Total = 0
		disk[0].Used = 0
		disk[0].UsedPercent = 100
		/*c.JSON(http.StatusInternalServerError, "No Data")
		return*/
	} else {
		// Convert response data
		v := res.Results[0].Series[0].Values
		disk := make([]models.DiskStat, len(v))
		var convTime time.Time
		for i, data := range v {
			// select time check
			convTime, _ = time.Parse(time.RFC3339, data[0].(string))

			// make struct
			disk[i].Time = convTime
			if err := MakeStructForStatsDisk(&disk[i], data); err != nil {
				lib.LogWarn("Error : %s\n", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
				return
			}
		}
	}

	//fmt.Printf("%+v\n", deltaStats)
	c.JSON(http.StatusOK, disk)
}

func MakeStructForStatsDisk(s *models.DiskStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}

	s.Err = ""
	s.Device = data[1].(string)
	s.Fstype = data[2].(string)
	s.Path = data[3].(string)
	s.Total, _ = data[4].(json.Number).Float64()
	s.Used, _ = data[5].(json.Number).Float64()
	s.UsedPercent, _ = data[6].(json.Number).Float64()
	return nil
}

func MakeStructForStatsWinDisk(s *models.WinDiskStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}

	s.FreeMegabytes, _ = data[1].(json.Number).Float64()
	return nil
}
