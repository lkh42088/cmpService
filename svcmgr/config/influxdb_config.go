package config

import (
	"cmpService/common/lib"
	"errors"
	client "github.com/influxdata/influxdb1-client/v2"
)

var infClient client.Client

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

	if infClient == nil {
		if c = NewClient(); c == nil {
			return nil, errors.New("Fail to client create\n")
		}
		// stored client
		infClient = c
	} else {
		c = infClient
	}
	defer c.Close()

	q.Command = query
	q.Database = SvcmgrGlobalConfig.InfluxdbConfig.DBName
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
	if SvcmgrGlobalConfig.InfluxdbConfig.Address == "" {
		lib.LogWarn("Collector config is empty.\n")
		return nil
	}
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     SvcmgrGlobalConfig.InfluxdbConfig.Address,
		Username: SvcmgrGlobalConfig.InfluxdbConfig.Username,
		Password: SvcmgrGlobalConfig.InfluxdbConfig.Password,
	})
	if err != nil {
		lib.LogWarn("Client create error : %s\n", err)
		return nil
	}
	return c
}

