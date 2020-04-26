package influx

import (
	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"nubes/collector/conf"
	"testing"
)

func TestCreateDB(t *testing.T) {
	conf.ConfigPath = "/home/lkh/go/src/nubes/collector/collector.conf"
	config := conf.ReadConfig()
	var q client.Query
	var c client.Client

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://127.0.0.1:8086",
		Username: "nubes",
		Password: "",
	})
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	q.Command = "CREATE DATABASE " + config.Influxdb
	c.Query(q)
}
