package mcrest

import (
	"cmpService/collector/influx"
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/utils"
	"cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/mcmongo"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func checkValidation(msg mcmodel.MgoVm) bool {
	if msg.Idx == 0 {
		fmt.Printf("error: idx is zero!\n")
		return false
	}
	if msg.McServerIdx == 0 {
		fmt.Printf("error: serverIdx is zero!\n")
		return false
	}
	if msg.CompanyIdx == 0 {
		fmt.Printf("error: cpIdx is zero!\n")
		return false
	}
	if msg.Name == "" {
		fmt.Printf("error: name is nil!\n")
		return false
	}
	if msg.OS == "" {
		fmt.Printf("error: os is nil!\n")
		return false
	}
	if msg.Image == "" {
		fmt.Printf("error: image is nil!\n")
		return false
	}
	return true
}

func GetMgoServer() (mcmodel.MgoServer, error) {
	var server mcmodel.MgoServer
	networks, err := kvm.GetMgoNetworksFromXmlNetwork()
	if err != nil {
		return server, err
	}
	images := kvm.GetImages()
	cfg := config.GetGlobalConfig()
	server.Mac = cfg.ServerMac
	server.Port = cfg.ServerPort
	server.Ip = cfg.ServerIp
	if len(networks) > 0 {
		server.Networks = &networks
	}
	if len(images) > 0 {
		server.Images = &images
	}
	fmt.Println("server:", server)
	fmt.Println("networks:", server.Networks)
	fmt.Println("images:", server.Images)
	return server, err
}

func registerServerHandler(c *gin.Context) {
	var msg mcmodel.McServerDetail
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("registerServerHandler: %v\n", msg)
	config.WriteServerStatus(msg.SerialNumber, msg.CompanyName, msg.CompanyIdx)

	server, _ := GetMgoServer()
	c.JSON(http.StatusOK, server)
}

func unRegisterServerHandler(c *gin.Context) {
	var msg mcmodel.McServer
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("unRegisterServerHandler: %v\n", msg)
	config.DeleteServerStatus()
	c.JSON(http.StatusOK, msg)
}

func addVmHandler(c *gin.Context) {
	var msg mcmodel.MgoVm
	err := c.ShouldBindJSON(&msg)
	fmt.Printf("addVmHandler: %s\n", msg.Dump())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !checkValidation(msg) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid message"})
		return
	}

	// Insert VM to Mongodb
	_, err = mcmongo.McMongo.AddVm(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg.CurrentStatus = "Ready"

	fmt.Printf("addVmHandler: success - %v\n", msg)
	c.JSON(http.StatusOK, msg)

	// Update Vm
	msg.Filename = kvm.MakeFilename(&msg)
	msg.IsCreated = false
	msg.IsProcess = true

	cfg := config.GetGlobalConfig()
	_, err = mcmongo.McMongo.UpdateVmByInternal(&msg)

	filepath := cfg.VmInstanceDir+"/"+msg.Filename+".qcow2"
	if ! utils.IsExistFile(filepath) {
		//kvm.KvmR.Vms = append(kvm.KvmR.Vms, msg)
		kvm.KvmR.Vms[msg.Idx] = msg
	}
}

func deleteVmHandler(c *gin.Context) {
	var msg mcmodel.MgoVm
	err := c.ShouldBindJSON(&msg)
	fmt.Printf("deleteVmHandler: %v\n", msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	vm, err := mcmongo.McMongo.GetVmById(int(msg.Idx))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = mcmongo.McMongo.DeleteVm(int(msg.Idx))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("deleteVmHandler: success\n")
	c.JSON(http.StatusOK, msg)

	// 1. Delete Vm instance
	kvm.DeleteVm(vm)
	// 2. Delete Vm image
	kvm.DeleteVmInstance(vm)
}

func getVmByIdHandler(c *gin.Context) {
	idStr := c.Param("id")

	// Get VMs from Mongodb
	id, _ := strconv.Atoi(idStr)
	vm, err := mcmongo.McMongo.GetVmById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, vm)
}

func getVmAllHandler(c *gin.Context) {
	// Get VMs from Mongodb
	vm, err := mcmongo.McMongo.GetVmAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, vm)
}

func addNetworkHandler(c *gin.Context) {
	var msg mcmodel.MgoNetwork
	c.ShouldBindJSON(&msg)
	kvm.CreateNetworkByMgoNetwork(msg)
	c.JSON(http.StatusOK, msg)
}

func deleteNetworkHandler(c *gin.Context) {
	var msg mcmodel.MgoNetwork
	c.ShouldBindJSON(&msg)

	kvm.DeleteNetwork(msg.Name)
	c.JSON(http.StatusOK, msg)
}

// Search public ip
func GetClientIp(c *gin.Context) {
	// search public ip
	url := "https://domains.google.com/checkip"

	response, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println("response:", string(data))
	c.JSON(http.StatusOK, string(data))
}

// Get VM Interface traffic
type VmIfStat struct {
	time          	time.Time
	ifDescr       	string
	ifPhysAddress 	string
	ifInUcastPkts  	int64
	ifOutUcastPkts	int64
}
var lastSelectTime sync.Map

func GetVmInterfaceTrafficByMac(mac string) {
	//mac := c.Param("mac")
	dbname := "interface"
	field := `"time","ifDescr"","ifPhysAddress","ifInUcastPkts","ifOutUcastPkts"`
	where := fmt.Sprintf(`"ifPhysAddress"='%s' AND time > now() - %s`, mac, "3hour")
	res := influx.GetMeasurementsWithCondition(dbname, field, where)

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("InfluxDB Response Error : No Data\n")
		return
	}

	// Convert response data
	v := res.Results[0].Series[0].Values
	stat := make([]VmIfStat, len(v))
	var timenano time.Time
	id := v[0][1].(string)
	for i, data := range v {
		// select time check
		timenano, _ = time.Parse(time.RFC3339, data[0].(string))
		if tmp, ok := lastSelectTime.Load(id); ok != true {
			if tmp.(time.Time).Sub(timenano) > 0 {
				continue
			}
		}

		// make struct
		stat[i].time = timenano
		stat[i].ifPhysAddress = mac
		if err := MakeStructForStats(&stat[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			return
		}

		lastSelectTime.Store(id, timenano)
	}

}

func MakeStructForStats(s *VmIfStat, data []interface{}) error {
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			return fmt.Errorf("Data interface is nil.(%d)\n", i)
		}
	}
	s.ifDescr = data[1].(string)
	s.ifInUcastPkts, _ = data[2].(json.Number).Int64()
	s.ifOutUcastPkts, _ = data[3].(json.Number).Int64()
	return nil
}
