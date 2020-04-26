package influx

import (
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	"nubes/common/lib"
)

type InfluxAccessor struct {
	Url string
	Username string
	Password string
	Database string
	Bp client.BatchPoints
	Client client.Client
}

var Influx InfluxAccessor

func SetInflux(c InfluxAccessor) {
	Influx = c
	lib.LogWarnln("Set InfluxDb:", Influx)
}

func NewInfluxCfg(url string, user string, passwd string, db string) *InfluxAccessor {
	config := InfluxAccessor{
		Url:      url,
		Username: user,
		Password: passwd,
		Database: db,
	}
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Precision:        "s",
		Database:         config.Database,
	})
	if err != nil {
		lib.LogWarn("InfluxDB NewInfluxCfg: Failed to get BatchPoints!!\n")
	} else {
		config.Bp = bp
	}

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:               config.Url,
		Username:           config.Username,
		Password:           config.Password,
	})
	config.Client = c
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return &config
}

