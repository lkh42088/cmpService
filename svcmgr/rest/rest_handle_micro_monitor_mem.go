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

type MemStat struct {
	Time             time.Time   `json:"time"`
	Available        json.Number `json:"available"`
	AvailablePercent json.Number `json:"available_percent"`
}

func (h *Handler) GetMonitorMem(c *gin.Context) {
	dbname := "mem"
	field := `"time","available","available_percent"`
	/*where := fmt.Sprintf(`cpu = 'cpu-total'`)*/
	res := GetMeasurementsWithConditionMem(dbname, field)

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("InfluxDB Response Error : No Data\n")
		c.JSON(http.StatusInternalServerError, "No Data")
		return
	}

	// Convert response data
	v := res.Results[0].Series[0].Values
	stat := make([]MemStat, len(v))
	var convTime time.Time
	for i, data := range v {
		// select time check
		convTime, _ = time.Parse(time.RFC3339, data[0].(string))

		// make struct
		stat[i].Time = convTime
		if err := MakeStructForStatsStructMem(&stat[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}

	c.JSON(http.StatusOK, stat)
}

// To create new influxDB client
func GetMeasurementsWithConditionMem(collector string, field string) *client.Response {
	query := "SELECT " + field + " FROM " + collector
	/*query += " WHERE " + where*/
	query += " ORDER BY time DESC LIMIT 1"
	fmt.Printf("Query: %s\n", query) // Need to debuggig
	res, err := InfluxdbQueryMem(query)
	if err != nil {
		return nil
	}
	return res
}

func MakeStructForStatsStructMem(s *MemStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}

	s.Available = data[1].(json.Number)
	s.AvailablePercent = data[2].(json.Number)
	return nil
}

func InfluxdbQueryMem(query string) (*client.Response, error) {
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
