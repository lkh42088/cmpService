package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"cmpService/collector/collectdevice"
	config2 "cmpService/collector/config"
	"cmpService/collector/mongo"
	"cmpService/collector/snmpapi"
	"cmpService/common/config"
)

type HandlerInterface interface {
	// MONGO-DB
	apiDeviceGetAllHandler(c *gin.Context)
	apiDeviceGetHandler(c *gin.Context)
	apiDevicePostHandler(c *gin.Context)
	apiDeviceRemoveAllHandler(c *gin.Context)
	apiDeviceRemoveHandler(c *gin.Context)
	// CONFIG
	apiRestConfigHandler(c *gin.Context)
}

func apiDeviceGetAllHandler(c *gin.Context) {
	d, err := mongo.Mongo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, d)
	return
}

func apiDeviceGetHandler(c *gin.Context) {
	id := collectdevice.ID(c.Param("id"))
	d, err := mongo.Mongo.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, d)
	return
}

func apiDevicePostHandler(c *gin.Context) {
	devices, err := getDevices(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	for _, d := range devices {
		fmt.Println(d)
		id, err := mongo.Mongo.Post(&d)
		fmt.Println(id, err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}

		d.Id = id
		snmpdev := snmpapi.NewSnmpDevice(d)
		snmpapi.SnmpDevTb.Post(*snmpdev)
	}
	c.JSON(http.StatusOK, devices)
	return
}

func apiDeviceRemoveHandler(c *gin.Context) {
	id := collectdevice.ID(c.Param("id"))
	err := mongo.Mongo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, id)
}

func apiDeviceRemoveAllHandler(c *gin.Context) {
	_, err := mongo.Mongo.DeleteAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, "")
}

func apiConfInfluxdbPostHandler(c *gin.Context) {
	var cfg config.InfluxDbConfig
	err := c.ShouldBindJSON(&cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	config2.SetConfigInfluxdb(cfg)
	c.JSON(http.StatusOK, cfg)
}

func apiConfMongodbPostHandler(c *gin.Context) {
	var cfg config.MongoDbConfig
	err := c.ShouldBindJSON(&cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	config2.SetConfigMongodb(cfg)
	c.JSON(http.StatusOK, cfg)
}


