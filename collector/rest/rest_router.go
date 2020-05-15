package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	conf "cmpService/collector/config"
	"cmpService/collector/collectdevice"
	"sync"
)

const (
	apiPathPrefix = "/api/v1"
	apiDevice = "/collectdevice"
	apiConfig = "/config"
)

var Router *gin.Engine

type ResponseError struct {
	Err error
}

func (err ResponseError) MarshalJSON() ([]byte, error) {
	if err.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", err.Err)), nil
}

func getDevices(r *http.Request) ([]collectdevice.ColletDevice, error) {
	var result []collectdevice.ColletDevice
	resp, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(resp))
	err := json.Unmarshal(resp, &result)
	if err != nil {
		fmt.Printf("getDevices() func fail.(err:%s)\n", err)
		return result, err
	}
	return result, nil
}

type Responses struct {
	ID     collectdevice.ID             `json:"id,omitempty"`
	Device []collectdevice.ColletDevice `json:"collectdevice"`
	Error  ResponseError                `json:"error"`
}

type Response struct {
	ID     collectdevice.ID           `json:"id,omitempty"`
	Device collectdevice.ColletDevice `json:"collectdevice"`
	Error  ResponseError              `json:"error"`
}

func Start(parentwg *sync.WaitGroup) {
	// Read REST api config
	config := conf.ReadConfig(conf.CollectorConfigPath)
	if config.RestServerIp == "" || config.RestServerPort == "" {
		fmt.Println("===== Need to REST server configuration. =====")
		return
	}
	address := config.RestServerIp + ":" + config.RestServerPort

	// Activate GIN
	router := gin.Default()

	rg := router.Group(apiPathPrefix)
	rg.GET(apiDevice, apiDeviceGetAllHandler)
	rg.GET(apiDevice + "/:id", apiDeviceGetHandler)
	rg.POST(apiDevice, apiDevicePostHandler)
	rg.DELETE(apiDevice + "", apiDeviceRemoveAllHandler)
	rg.DELETE(apiDevice + "/:id", apiDeviceRemoveHandler)

	// REST CONFIG CHANGE
	rg.POST(apiConfig + "influxdb", apiConfInfluxdbPostHandler)
	rg.POST(apiConfig + "mongodb", apiConfMongodbPostHandler)

	router.Run(address)
	if parentwg != nil {
		parentwg.Done()
	}
}

