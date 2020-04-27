package config

import (
	"nubes/collector/influx"
	"nubes/collector/mongo"
	"nubes/common/lib"
)

func ConfigureMongoDB() {
	config := ReadConfig(CollectorConfigPath)
	if config.MongoIp == "" || config.MongoDb == "" || config.MongoCollection == "" {
		lib.LogWarn("Failed MongoDb configuration!\n")
		return
	}
	m := mongo.NewMongoAccessor(config.MongoIp, config.MongoDb, config.MongoCollection)
	mongo.SetMongo(m)
}

func ConfigureInfluxDB() {
	var influxConfig *influx.InfluxAccessor
	config := ReadConfig(CollectorConfigPath)
	if config.InfluxIp == "" || config.InfluxDb == "" {
		lib.LogWarn("Failed InfluxDb configuration!\n")
		return
	}
	path := "http://" + config.InfluxIp + ":8086"
	influxConfig = influx.NewInfluxCfg(
		path,
		config.InfluxUser,
		config.InfluxPassword,
		config.InfluxDb)
	influx.SetInflux(*influxConfig)
}
