package mcinflux

import (
	"cmpService/common/lib"
	config2 "cmpService/mcagent/config"
	"errors"
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
)

type InfluxAccessor struct {
	Url      string
	Username string
	Password string
	Database string
	Bp       client.BatchPoints
	Client   client.Client
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
		Precision: "ms",
		Database:  config.Database,
	})
	if err != nil {
		lib.LogWarn("InfluxDB NewInfluxCfg: Failed to get BatchPoints!!\n")
	} else {
		config.Bp = bp
	}

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.Url,
		Username: config.Username,
		Password: config.Password,
	})
	config.Client = c
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return &config
}

func NewClient() client.Client {
	var c client.Client
	if Influx.Url == "" {
		lib.LogWarn("Collector config is empty.\n")
		return nil
	}
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     Influx.Url,
		Username: Influx.Username,
		Password: Influx.Password,
	})
	if err != nil {
		lib.LogWarn("Client create error : %s\n", err)
		return nil
	}
	return c
}

func InfluxdbCreateDB(dbname string) error {
	query := "CREATE DATABASE " + dbname
	_, err := InfluxdbQuery(query)
	if err != nil {
		return errors.New("Fail to create database in Influxdb.\n")
	}
	return nil
}

func InfluxdbCheckDB(dbname string) error {
	query := "USE " + dbname
	_, err := InfluxdbQuery(query)
	if err != nil {
		return errors.New("Find to database in Influxdb.\n")
	}
	return nil
}

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

	q.Command = query
	q.Database = Influx.Database
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

func ConfigureInfluxDB() bool {
	var influxConfig *InfluxAccessor
	globalCfg := config2.GetGlobalConfig()
	if globalCfg.InfluxUser == "" || globalCfg.InfluxIp == "" || globalCfg.InfluxDb == "" {
		lib.LogWarn("Failed MongoDb configuration!\n")
		return false
	}
	path := "http://" + globalCfg.InfluxIp + ":8086"
	influxConfig = NewInfluxCfg(
		path,
		globalCfg.InfluxUser,
		globalCfg.InfluxPassword,
		globalCfg.InfluxDb)
	SetInflux(*influxConfig)
	return true
}
