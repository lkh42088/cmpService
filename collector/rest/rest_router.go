package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"nubes/collector/device"
	"nubes/collector/lib"
	"nubes/collector/mongodao"
	"nubes/collector/snmpapi"
	"sync"
)

const (
	apiPathPrefix = "/api/v1"
	idPattern = "/{id:[0-9a-f]+}"
	apiDevice ="/device"
)

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
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	encodedDevices, ok := r.PostForm["device"]
	if !ok {
		return nil, errors.New("device parameter expected")
	}
	for _, encodedDevice := range encodedDevices {
		var d device.Device
		if err := json.Unmarshal([]byte(encodedDevice), &d); err != nil {
			return nil, err
		}
		result = append(result, d)
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

func RestRouter(parentwg *sync.WaitGroup) {

	r := mux.NewRouter()
	s := r.PathPrefix(apiPathPrefix).Subrouter()
	s.HandleFunc(apiDevice, apiDeviceGetAllHandler).Methods("GET")
	s.HandleFunc(apiDevice + idPattern, apiDeviceGetHandler).Methods("GET")
	//s.HandleFunc(apiDevice, apiDevicePostHandler).Methods("POST")
	s.HandleFunc(apiDevice , apiDevicePostJsonHandler).Methods("POST")
	s.HandleFunc(apiDevice + idPattern, apiDeviceRemoveHandler).Methods("DELETE")
	s.HandleFunc(apiDevice + "/all", apiDeviceRemoveAllHandler).Methods("DELETE")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8884", nil))

	if parentwg != nil {
		parentwg.Done()
	}
}

func apiDeviceGetAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all device")
	d, err := mongodao.Mongo.GetAll()
	if d != nil {
		fmt.Println("devices:", d)
	}
	err = json.NewEncoder(w).Encode(Responses{
		Device: d,
		Error:  ResponseError{err},
	})
}

func apiDeviceGetHandler(w http.ResponseWriter, r *http.Request) {
	id := device.ID(mux.Vars(r)["id"])
	fmt.Println("Get device id:", id)
	d, err := mongodao.Mongo.Get(id)
	err = json.NewEncoder(w).Encode(Response{
		ID:     id,
		Device: d,
		Error:  ResponseError{err},
	})
}

func apiDevicePostHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := getDevices(r)
	if err != nil {
		log.Println(err)
		return
	}
	for _, d := range devices {
		id, err := mongodao.Mongo.Post(&d)
		err = json.NewEncoder(w).Encode(Response{
			ID:     id,
			Device: d,
			Error:  ResponseError{},
		})
		if err != nil {
			log.Println(err)
			return
		}
		lib.LogInfo("apiDevicePostHandler..")
		d.Id = id
		snmpdev := snmpapi.NewSnmpDevice(d)
		snmpapi.SnmpDevices.Post(*snmpdev)
	}
}

func apiDevicePostJsonHandler(w http.ResponseWriter, r *http.Request) {
	//body, err := ioutil.ReadAll(&io.LimitedReader{r.Body, 1048657})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	dev := device.Device{}
	if err :=json.Unmarshal(body, &dev); err != nil {
		json.NewEncoder(w).Encode(Response {
			"-1",
			dev,
			ResponseError{err},
		})
	}
	lib.LogInfo("apiDevicePostJsonHandler..")
	if dev.Ip == "" || dev.SnmpCommunity == "" {
		json.NewEncoder(w).Encode(Response {
			"-1",
			dev,
			ResponseError{err},
		})
	}
	if _, err := mongodao.Mongo.Post(&dev); err != nil {
		json.NewEncoder(w).Encode(Response {
			"-1",
			dev,
			ResponseError{err},
		})
	}
	fmt.Println("Post: ", dev)
	snmpdev := snmpapi.NewSnmpDevice(dev)
	snmpapi.SnmpDevices.Post(*snmpdev)
	if err := json.NewEncoder(w).Encode(Response {
		dev.GetIdString(),
		dev, ResponseError{err},}); err != nil {
		panic(err)
	}
}

func apiDeviceRemoveAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REST: remove all device")
	var err error
	_, err = mongodao.Mongo.DeleteAll()
	err = json.NewEncoder(w).Encode(Response{
		Error:  ResponseError{err},
	})
}

func apiDeviceRemoveHandler(w http.ResponseWriter, r *http.Request) {
	id := device.ID(mux.Vars(r)["id"])
	fmt.Println("REST: remove device", id)
	var err error
	if id == "ALL" {
		_, err = mongodao.Mongo.DeleteAll()
	} else {
		err = mongodao.Mongo.Delete(id)
	}
	err = json.NewEncoder(w).Encode(Response{
		ID:     id,
		Error:  ResponseError{err},
	})
}

