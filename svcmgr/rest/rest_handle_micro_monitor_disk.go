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

type DiskStat struct {
	Time        time.Time   `json:"time"`
	Device      string      `json:"device"`
	Fstype      string      `json:"fstype"`
	Path        string      `json:"path"`
	Total       json.Number `json:"total"`
	Used        json.Number `json:"used"`
	UsedPercent json.Number `json:"used_percent"`
}

func (h *Handler) GetMonitorDisk(c *gin.Context) {
	dbname := "disk"
	field := `"time","device","fstype","path","total","used","used_percent"`
	where := fmt.Sprintf(`path = '/'`)
	res := GetMeasurementsWithConditionDisk(dbname, field, where)

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("InfluxDB Response Error : No Data\n")
		c.JSON(http.StatusInternalServerError, "No Data")
		return
	}

	// Convert response data
	v := res.Results[0].Series[0].Values
	stat := make([]DiskStat, len(v))
	var convTime time.Time
	for i, data := range v {
		// select time check
		convTime, _ = time.Parse(time.RFC3339, data[0].(string))

		// make struct
		stat[i].Time = convTime
		if err := MakeStructForStatsStructDisk(&stat[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}

	c.JSON(http.StatusOK, stat)
}

// To create new influxDB client
func GetMeasurementsWithConditionDisk(collector string, field string, where string) *client.Response {
	query := "SELECT " + field + " FROM " + collector
	query += " WHERE " + where
	query += " ORDER BY time DESC LIMIT 1"
	fmt.Printf("Query: %s\n", query) // Need to debuggig
	res, err := InfluxdbQueryDisk(query)
	if err != nil {
		return nil
	}
	return res
}

func MakeStructForStatsStructDisk(s *DiskStat, data []interface{}) error {
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

func InfluxdbQueryDisk(query string) (*client.Response, error) {
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
