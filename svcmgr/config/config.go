package config

import (
	client "github.com/influxdata/influxdb1-client"
	"nubes/common/config"
	"nubes/common/mariadblayer"
)

type global_config struct {
	MariadbConfig config.DBConfig
	Mariadb       mariadblayer.DBORM

	InfluxdbConfig config.DBConfig
	InfluxdbBp client.BatchPoints
	InfluxdbClient client.Client
	RestServer     string
}

var SvcmgrConfig = &global_config{}

