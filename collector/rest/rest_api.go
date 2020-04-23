package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"io/ioutil"
	"log"
	"net/http"
	"nubes/collector/db/influx"
	"nubes/collector/device"
	"nubes/collector/lib"
	"nubes/collector/mongodao"
	"nubes/collector/snmpapi"
	"strings"
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

// Avoid import cycle (rest <-> mongodao)
var Mongo MongoUser
type MongoUser interface {
	Get(device.ID) (device.Device, error)
	Put(device.ID, device.Device) error
	Post(*device.Device) (device.ID, error)
	DeleteAll() (*mgo.ChangeInfo, error)
	Delete(device.ID) error
	GetAll() ([]device.Device, error)
}

func MongoDBConfigChange() {
	r := ReadConf()
	if r == nil {
		fmt.Println("NewmongoDB Readconf fail")
		Mongo = mongodao.New("127.0.0.1", "collector", "devices")
	} else {
		fmt.Printf("Mongo Config IP:%s DB:%s TABLE:%s\n",
			r["mongoip"], r["mongodb"], r["mongotable"])
		Mongo = mongodao.New(r["mongoip"], r["mongodb"], r["mongotable"])
	}
	return
}

func InfluxDBConfigChange() {
	var config *influx.Config
	r := ReadConf()
	if r == nil {
		config = influx.Init(
			"http://192.168.10.19:8086",
			"nubes",
			"",
			"snmp_nodes")
	} else {
		path := "http://" + r["influxip"] + ":8086"
		config = influx.Init(
			path,
			"nubes",
			//"",	// id
			"",
			r["influxdb"])
	}

	fmt.Println(config)
	snmpapi.InfluxConfigure(config)
}

func RestAPIServerRestart() {
	// For Rest API Restart
}

/// REST-GET
func apiDeviceGetAllHandler(c *gin.Context) {
	d, err := Mongo.GetAll()
	if d != nil {
		fmt.Println("devices:", d)
	}
	response := Responses{
		Device: d,
		Error:  ResponseError{err},
	}

	c.JSON(http.StatusOK, response)
	return
}

func apiDeviceGetHandler(c *gin.Context) {
	id := device.ID(c.Param("get"))
	d, err := Mongo.Get(id)
	response := Response{
		ID:	id,
		Device: d,
		Error:  ResponseError{err},
	}

	c.JSON(http.StatusOK, response)
	return
}

/// REST-POST
func apiDevicePostHandler(c *gin.Context) {
	devices, err := getDevices(c.Request)
	var response Response
	if err != nil {
		log.Println(err)
		return
	}
	for _, d := range devices {
		id, err := Mongo.Post(&d)
		response = Response{
			ID:     id,
			Device: d,
			Error:  ResponseError{},
		}
		if err != nil {
			log.Println(err)
			return
		}

		d.Id = id
		snmpdev := snmpapi.NewSnmpDevice(d)
		snmpapi.SnmpDevices.Post(*snmpdev)
	}
	c.JSON(http.StatusOK, response)
	return
}

// Not used
// Need to fix : not changing to GIN Module
func apiDevicePostJsonHandler(c *gin.Context) {
	//body, err := ioutil.ReadAll(&io.LimitedReader{r.Body, 1048657})
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	if err := c.Request.Body.Close(); err != nil {
		panic(err)
	}
	dev := device.Device{}
	if err :=json.Unmarshal(body, &dev); err != nil {
		json.NewEncoder(c.Writer).Encode(Response {
			"-1",
			dev,
			ResponseError{err},
		})
	}
	lib.LogInfo("apiDevicePostJsonHandler..")
	if dev.Ip == "" || dev.SnmpCommunity == "" {
		json.NewEncoder(c.Writer).Encode(Response {
			"-1",
			dev,
			ResponseError{err},
		})
	}
	if _, err := Mongo.Post(&dev); err != nil {
		json.NewEncoder(c.Writer).Encode(Response {
			"-1",
			dev,
			ResponseError{err},
		})
	}
	fmt.Println("Post: ", dev)
	snmpdev := snmpapi.NewSnmpDevice(dev)
	snmpapi.SnmpDevices.Post(*snmpdev)
	if err := json.NewEncoder(c.Writer).Encode(Response {
		dev.GetIdString(),
		dev, ResponseError{err},}); err != nil {
		panic(err)
	}
}


/// REST-DELETE
// Not used
func apiDeviceRemoveAllHandler(c *gin.Context) {
	var err error
	_, err = Mongo.DeleteAll()
	response := Response{
		Error:  ResponseError{err},
	}
	c.JSON(http.StatusOK, response)
}

// /url/all : all delete
// /url/id : specific id delete
func apiDeviceRemoveHandler(c *gin.Context) {
	id := device.ID(c.Param("del"))
	var err error
	if strings.ToUpper(string(id)) == "ALL" {
		_, err = Mongo.DeleteAll()
		id = ""
	} else {
		err = Mongo.Delete(id)
	}
	response := Response{
		ID:     id,
		Error:  ResponseError{err},
	}
	c.JSON(http.StatusOK, response)
}

// REST CONFIG change
func apiRestConfigHandler(c *gin.Context) {
	key := c.Param("key")
	config := c.Param("config")
	fmt.Printf("===== key:%s, config:%s =====\n", key, config)
	if WriteConf(key, config) == nil {
		// Apply config
		if strings.Contains(key, "mongo") {
			MongoDBConfigChange()
		} else if strings.Contains(key, "influx") {
			InfluxDBConfigChange()
		} else if strings.Contains(key, "rest") {
			RestAPIServerRestart()
		} else if strings.Contains(key, "svcmgr"){
			// SvcmgrReconnect()
		}
		c.JSON(http.StatusOK, gin.H{"message": "Success config change.\n"})
	} else {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message":"Config file cannot modify.\n"})
	}
	return
}