package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	conf "nubes/collector/conf"
	"nubes/collector/device"
	"sync"
)

const (
	apiPathPrefix = "/api/v1"
	idPattern = "/{id:[0-9a-f]+}"
	apiDevice = "/device"
	apiConfig = "/conf"
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

func getDevices(r *http.Request) ([]device.Device, error) {
	var result []device.Device
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
	ID     device.ID     `json:"id,omitempty"`
	Device []device.Device `json:"device"`
	Error  ResponseError `json:"error"`
}

type Response struct {
	ID     device.ID     `json:"id,omitempty"`
	Device device.Device `json:"device"`
	Error  ResponseError `json:"error"`
}

func RestAPIConfigure() {
	// MongoDB configure
	MongoDBConfigChange()
	// InfluxDB configure
	InfluxDBConfigChange()
}

func RunAPI(parentwg *sync.WaitGroup) {
	// Configure
	RestAPIConfigure()

	// Read REST api conf
	config := conf.ReadConfig()
	if config.Restip == "" || config.Restport == "" {
		fmt.Println("===== Need to REST server configuration. =====")
		return
	}
	address := config.Restip + ":" + config.Restport

	// Activate GIN
	router := gin.Default()

	rg := router.Group(apiPathPrefix)
	// GET
	rg.GET(apiDevice, apiDeviceGetAllHandler)
	rg.GET(apiDevice + "/:get", apiDeviceGetHandler)
	// POST
	rg.POST(apiDevice, apiDevicePostHandler)
	// DELETE
	rg.DELETE(apiDevice + "/all", apiDeviceRemoveAllHandler)
	//rg.DELETE(apiDevice + "/:del", apiDeviceRemoveHandler)

	// REST CONFIG CHANGE
	rg.POST(apiConfig + "/:key" + "/:conf", apiRestConfigHandler)

	router.Run(address)
	if parentwg != nil {
		parentwg.Done()
	}
}

