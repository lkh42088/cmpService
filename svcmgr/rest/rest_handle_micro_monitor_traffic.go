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

func GetVmInterfaceTrafficByMac(c *gin.Context) {
	mac := c.Param("mac")
	dbname := "interface"
	field := `"time","hostname","ifDescr","ifPhysAddress","ifInOctets","ifOutOctets"`
	where := fmt.Sprintf(`"ifPhysAddress" =~ /.*%s/ AND time > now() - %s`, mac, "1h")
	res := conf.GetMeasurementsWithCondition(dbname, field, where)
    //fmt.Println(res.Err)
	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("InfluxDB Response Error : No Data\n")
		c.JSON(http.StatusInternalServerError, res.Err)
		return
	}

	// Convert response data
	v := res.Results[0].Series[0].Values
	stat := make([]models.VmIfStat, len(v))
	var convTime time.Time
	for i, data := range v {
		// select time check
		convTime, _ = time.Parse(time.RFC3339, data[0].(string))

		// make struct
		stat[i].Time = convTime
		stat[i].IfPhysAddress = mac
		if err := MakeStructForStats(&stat[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}

	deltaStats := MakeDeltaValues(stat)
	//fmt.Printf("%+v\n", deltaStats)
	c.JSON(http.StatusOK, deltaStats)
}

func MakeStructForStats(s *models.VmIfStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}
	s.IfDescr = data[2].(string)
	s.Hostname = data[1].(string)
	s.IfInOctets, _ = data[4].(json.Number).Int64()
	s.IfOutOctets, _ = data[5].(json.Number).Int64()
	return nil
}

func MakeStructForWinStats(s *models.WinVmIfStat, data []interface{}) error {
	//fmt.Println("data @@@@ : ", data)
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			fmt.Println("data[i] : ", data[i])
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}

	//s.Hostname = "none"
	s.BytesReceivedPersec, _ = data[1].(json.Number).Float64()
	s.BytesSentPersec, _ = data[2].(json.Number).Float64()

	s.BytesReceivedPersec = s.BytesReceivedPersec * 5
	s.BytesSentPersec = s.BytesSentPersec * 5

	return nil
}

func MakeDeltaValues(s []models.VmIfStat) models.VmStatseRsponse {
	//delta := make([]VmIfStat, len(s))
	var delta models.VmIfStatistics
	var result models.VmIfStat
	var response models.VmStatseRsponse
	var unit models.Stats

	response.Stats[0].Id = "RX"
	response.Stats[1].Id = "TX"

	for i := 0; i < len(s); i++ {
		if i == 0 {
			continue
		}
		//fmt.Println(i, s[i])
		result.Time = s[i].Time
		result.Hostname = s[i].Hostname
		result.IfDescr = s[i].IfDescr
		result.IfPhysAddress = s[i].IfPhysAddress
		result.IfInOctets = s[i].IfInOctets - s[i-1].IfInOctets
		result.IfOutOctets = s[i].IfOutOctets - s[i-1].IfOutOctets
		delta.Stats = append(delta.Stats, result)

		// Make response data set
		//unit.Xaxis = result.Time.Format("03:04:05")
		unit.Xaxis = result.Time
		unit.Yaxis = result.IfInOctets
		response.Stats[0].Data = append(response.Stats[0].Data, unit)
		unit.Yaxis = result.IfOutOctets
		response.Stats[1].Data = append(response.Stats[1].Data, unit)
	}
	response.Hostname = result.IfDescr
	return response
}

func MakeDeltaWinValues(s []models.WinVmIfStat) models.VmStatseRsponse {
	//delta := make([]VmIfStat, len(s))
	var delta models.WinVmIfStatistics
	var result models.WinVmIfStat
	var response models.VmStatseRsponse
	var unit models.Stats

	response.Stats[0].Id = "RX"
	response.Stats[1].Id = "TX"

	for i := 0; i < len(s); i++ {
		if i == 0 {
			continue
		}
		//fmt.Println(i, s[i])
		result.Time = s[i].Time
		//result.Hostname = s[i].Hostname
		// 여기다 곱하기 해줘야 하는거 아냐?
		result.BytesReceivedPersec = s[i].BytesReceivedPersec
		result.BytesSentPersec = s[i].BytesSentPersec
		delta.Stats = append(delta.Stats, result)

		// Make response data set
		//unit.Xaxis = result.Time.Format("03:04:05")
		unit.Xaxis = result.Time
		unit.Yaxis = int64(result.BytesReceivedPersec * 5)
		response.Stats[0].Data = append(response.Stats[0].Data, unit)
		unit.Yaxis = int64(result.BytesSentPersec * 5)
		response.Stats[1].Data = append(response.Stats[1].Data, unit)
	}
	//response.Hostname = result.Hostname
	//fmt.Println("★★★★★★ response : ", response)

	return response
}
