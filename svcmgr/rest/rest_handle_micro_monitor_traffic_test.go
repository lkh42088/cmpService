package rest

import (
	"cmpService/common/mcmodel"
	conf "cmpService/svcmgr/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func BenchmarkGetVmInterfaceTrafficByMac(b *testing.B) {
	b.N = 10
	conf.SvcmgrConfigPath = "../etc/svcmgr.lkh.conf"
	conf.SetInfluxDB()

	var c *gin.Context
	mac := "52:54:00:01:b5:b7"
	dbname := "interface"
	field := `"time","hostname","ifDescr","ifPhysAddress","ifInOctets","ifOutOctets"`
	where := fmt.Sprintf(`"ifPhysAddress"='%s' AND time > now() - %s`, mac, "1h")
	res := conf.GetMeasurementsWithCondition(dbname, field, where)

	if (res != nil && res.Results[0].Series == nil) ||
		len(res.Results[0].Series[0].Values) == 0 {
		c.JSON(http.StatusInternalServerError, "No Data")
		return
	}

	// Convert response data
	v := res.Results[0].Series[0].Values
	stat := make([]mcmodel.VmIfStat, len(v))
	var convTime time.Time
	for i, data := range v {
		// select time check
		convTime, _ = time.Parse(time.RFC3339, data[0].(string))

		// make struct
		stat[i].Time = convTime
		stat[i].IfPhysAddress = mac
		if err := MakeStructForStats(&stat[i], data); err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}
	}

	resStat:= MakeDeltaValues(stat)
	fmt.Println(resStat)
}

func TestGetServerTop5(t *testing.T) {
	conf.SvcmgrConfigPath = "../etc/svcmgr.lkh.conf"
	conf.SetInfluxDB()
	var c *gin.Context
	GetServerRank(c)
}
