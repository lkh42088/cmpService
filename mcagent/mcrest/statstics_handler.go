package mcrest

import (
	"cmpService/common/lib"
	"cmpService/mcagent/mcinflux"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Get VM Interface traffic
type VmIfStat struct {
	time          	time.Time
	hostname 		string
	ifDescr       	string
	ifPhysAddress 	string
	ifInOctets  	int64
	ifOutOctets		int64
}

func GetVmInterfaceTrafficByMac(c *gin.Context) {
	mac := c.Param("mac")
	dbname := "interface"
	field := `"time","hostname","ifDescr","ifPhysAddress","ifInOctets","ifOutOctets"`
	where := fmt.Sprintf(`"ifPhysAddress"='%s' AND time > now() - %s`, mac, "3h")
	res := mcinflux.GetMeasurementsWithCondition(dbname, field, where)

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("InfluxDB Response Error : No Data\n")
		c.JSON(http.StatusInternalServerError, "No Data")
		return
	}

	// Convert response data
	v := res.Results[0].Series[0].Values
	stat := make([]VmIfStat, len(v))
	var convTime time.Time
	for i, data := range v {
		// select time check
		convTime, _ = time.Parse(time.RFC3339, data[0].(string))

		// make struct
		stat[i].time = convTime
		stat[i].ifPhysAddress = mac
		if err := MakeStructForStats(&stat[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}

	deltaStats := MakeDeltaValues(stat)

	fmt.Printf("VM STAT : %v\n", deltaStats)
	c.JSON(http.StatusOK, deltaStats)
}

func MakeStructForStats(s *VmIfStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}
	s.ifDescr = data[2].(string)
	s.hostname = data[1].(string)
	s.ifInOctets, _ = data[4].(json.Number).Int64()
	s.ifOutOctets, _ = data[5].(json.Number).Int64()
	return nil
}

func MakeDeltaValues(s []VmIfStat) []VmIfStat {
	var result []VmIfStat
	var idx = 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			result[idx].time = s[i].time
			result[idx].hostname = s[i].hostname
			result[idx].ifPhysAddress = s[i].ifPhysAddress
			result[idx].ifInOctets = s[i].ifInOctets - s[i-1].ifInOctets
			result[idx].ifOutOctets = s[i].ifOutOctets - s[i-1].ifOutOctets
		}
	}
	return result
}

