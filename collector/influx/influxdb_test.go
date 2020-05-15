package influx

import (
	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"cmpService/collector/config"
	"testing"
)

func TestCreateDB(t *testing.T) {
	config.CollectorConfigPath = "/home/lkh/go/src/cmpService/collector/collector.conf"
	config := config.ReadConfig(config.CollectorConfigPath)
	var q client.Query
	var c client.Client

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://127.0.0.1:8086",
		Username: "cmpService",
		Password: "",
	})
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	q.Command = "CREATE DATABASE " + config.InfluxDb
	c.Query(q)
}
