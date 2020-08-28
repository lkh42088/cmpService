package rest

import (
	"cmpService/common/lib"
	"cmpService/svcmgr/config"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/influxdata/influxdb1-client/v2"
	"net/http"
	"time"
)

// Get VM Interface traffic
type VmIfStat struct {
	Time          	time.Time	`json:"time"`
	Hostname 		string		`json:"hostname"`
	IfDescr       	string		`json:"ifDescr"`
	IfPhysAddress 	string		`json:"ifPhysAddress"`
	IfInOctets  	int64		`json:"ifInOctets"`
	IfOutOctets		int64		`json:"ifOutOctets"`
}

type VmIfStatistics struct {
	Stats 			[]VmIfStat	`json:"stats"`
}

type Stats struct {
	Xaxis 			string		`json:"x"`
	Yaxis			int64		`json:"y"`
}

type VmStatsSet struct {
	Id				string		`json:"id"`
	Data			[]Stats		`json:"data"`
}

type VmStatseRsponse struct {
	Hostname 		string			`json:"hostname"`
	Stats 			[2]VmStatsSet	`json:"stats"`
}

func GetVmInterfaceTrafficByMac(c *gin.Context) {
	mac := c.Param("mac")
	dbname := "interface"
	field := `"time","hostname","ifDescr","ifPhysAddress","ifInOctets","ifOutOctets"`
	where := fmt.Sprintf(`"ifPhysAddress"='%s' AND time > now() - %s`, mac, "1h")
	res := GetMeasurementsWithCondition(dbname, field, where)

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
		stat[i].Time = convTime
		stat[i].IfPhysAddress = mac
		if err := MakeStructForStats(&stat[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}

	deltaStats := MakeDeltaValues(stat)
	fmt.Printf("%+v\n", deltaStats)
	c.JSON(http.StatusOK, deltaStats)
}

func MakeStructForStats(s *VmIfStat, data []interface{}) error {
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

func MakeDeltaValues(s []VmIfStat) VmStatseRsponse {
	//delta := make([]VmIfStat, len(s))
	var delta VmIfStatistics
	var result VmIfStat
	var response VmStatseRsponse
	var unit Stats

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
		unit.Xaxis = result.Time.Format("01:02:03")
		unit.Yaxis = result.IfInOctets
		response.Stats[0].Data = append(response.Stats[0].Data, unit)
		unit.Yaxis = result.IfOutOctets
		response.Stats[1].Data = append(response.Stats[1].Data, unit)
	}
	response.Hostname = result.IfDescr
	return response
}


// To create new influxDB client
func GetMeasurementsWithCondition(collector string, field string, where string) *client.Response {
	query := "SELECT " + field + " FROM " + collector
	query += " WHERE " + where
	//fmt.Printf("Query: %s\n", query)	// Need to debuggig
	res, err := InfluxdbQuery(query)
	if err != nil {
		return nil
	}
	return res
}

func InfluxdbQuery(query string) (*client.Response, error) {
	if query == "" {
		return nil, errors.New("Invalid query message.\n")
	}

	var q client.Query
	var c client.Client

	if c = NewClient(); c == nil {
		return nil, errors.New("Fail to client create\n")
	}
	defer c.Close()

	q.Command = query
	q.Database = config.SvcmgrGlobalConfig.InfluxdbConfig.DBName
	var res *client.Response
	var err error
	if res, err = c.Query(q); err != nil {
		lib.LogWarn("Influxdb query error : %s\n", err)
		return nil, err
	}
	if res != nil {
		return res, nil
	}
	return nil, nil
}

func NewClient() client.Client {
	var c client.Client
	if config.SvcmgrGlobalConfig.InfluxdbConfig.Address == "" {
		lib.LogWarn("Collector config is empty.\n")
		return nil
	}
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.SvcmgrGlobalConfig.InfluxdbConfig.Address,
		Username: config.SvcmgrGlobalConfig.InfluxdbConfig.Username,
		Password: config.SvcmgrGlobalConfig.InfluxdbConfig.Password,
	})
	if err != nil {
		lib.LogWarn("Client create error : %s\n", err)
		return nil
	}
	return c
}