package influx

import (
	"errors"
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	conf2 "nubes/collector/conf"
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

func NewClient() client.Client {
	var c client.Client
	conf := conf2.ReadConfig()
	if conf.Influxip == ""  {
		lib.LogWarn("Collector config is empty.")
		return nil
	}
	addr := "http://" + conf.Influxip + ":8086"

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: conf.InfluxUser,
		Password: "",
	})
	if err != nil {
		lib.LogWarn("Client create error : %s", err)
		return nil
	}
	return c
}

func InfluxdbCreateDB(dbname string) error {
	query := "CREATE DATABASE " + dbname
	err := InfluxdbQuery(query)
	if err != nil {
		return errors.New("Fail to create database in Influxdb.")
	}
	return nil
}

func InfluxdbCheckDB(dbname string) error {
	query := "USE " + dbname
	err := InfluxdbQuery(query)
	if err != nil {
		return errors.New("Find to database in Influxdb.")
	}
	return nil
}

func InfluxdbQuery(query string) error {
	if query == "" {
		return errors.New("Invalid query message.\n")
	}

	var q client.Query
	var c client.Client

	if c = NewClient(); c == nil {
		return errors.New("Fail to client create\n")
	}

	q.Command = query
	if _, err := c.Query(q); err != nil {
		lib.LogWarn("Influxdb query error : %s", err)
		return err
	}
	return nil
}
