package mcrest

import (
	"cmpService/common/mcmodel"
	"cmpService/common/utils"
	"cmpService/mcagent/config"
	"cmpService/mcagent/ddns"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/repo"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMcServer() (mcmodel.McServerMsg, error) {
	var server mcmodel.McServerMsg
	networks, err := kvm.GetMgoNetworksFromXmlNetwork()
	if err != nil {
		return server, err
	}
	images := kvm.GetImages()
	cfg := config.GetMcGlobalConfig()
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

type McResourceMsg struct {
	GlobalConfig config.McAgentConfig
	DnatList *[]utils.DnatRule
	CreateVmList *map[uint]mcmodel.McVm
	CacheVmList *[]mcmodel.McVm
	LibvirtVmList *[]mcmodel.McVm
	CronVmList *[]kvm.SnapVm
}

func (n *McResourceMsg) Dump() string {
	pretty, _ := json.MarshalIndent(n, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

type ResourceMsg struct {
	Command string
}

func clearAllResource () {
	/********************************************
	 * Delete VM
	 ********************************************/
	// get Vm List Object
	vmList := repo.GetVmCacheObject()
	for _, vm := range vmList {
		deleteVm(vm.Name)
	}

	/********************************************
	 * Delete Server
	 ********************************************/
	deleteServer()
}

func resourceControlHandler(c *gin.Context) {
	var msg ResourceMsg
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		fmt.Println("resourceControlHandler: error")
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	fmt.Println("resourceControlHandler:", msg.Command)
	switch msg.Command  {
	case "clear":
		/****************
		 * Clear Resource
		 ****************/
		clearAllResource()
		c.JSON(http.StatusOK, "success")
		return
	default:
		c.JSON(http.StatusInternalServerError, "default")
	}
}

func getResourceHandler(c *gin.Context) {
	var resource McResourceMsg
	resource.GlobalConfig = config.GetMcGlobalConfig()
	resource.DnatList = utils.GetDnatList()
	resource.CreateVmList = &kvm.CreateVmFsm.Vms
	resource.CacheVmList = &repo.GlobalVmCache
	resource.LibvirtVmList = kvm.LibvirtR.Old.Vms
	resource.CronVmList = &kvm.CronSch.SnapVms
	c.JSON(http.StatusOK, resource)
}

func registerServerHandler(c *gin.Context) {
	var msg mcmodel.McServerDetail
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("registerServerHandler: %v\n", msg)

	// Insert to DB and create etc file.
	AddMcServer(msg, true)
	ddns.ApplyDdns(msg.McServer)

	//server, _ := GetMcServer()
	server := kvm.GetMcServerInfo()
	c.JSON(http.StatusOK, server)
}

func AddMcServer(msg mcmodel.McServerDetail, enable bool) {
	// Create etc file
	msg.Enable = enable
	config.WriteServerStatus(msg.SerialNumber, msg.CompanyName, msg.CompanyIdx, enable)
	config.SetSerialNumber2GlobalConfig(msg.SerialNumber)
	// Add repo (DB)
	repo.AddServer2Repo(&msg)
}

func deleteServer() {
	config.DeleteServerStatus()
	// Delete Repo
	repo.DeleteServer2Repo()
}

func unRegisterServerHandler(c *gin.Context) {
	var msg mcmodel.McServer
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("unRegisterServerHandler: %v\n", msg)

	/*********************
	 * Check Vm
	 *********************/
	list, err := repo.GetAllVmFromDb()
	if len(list) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "VMs exist!!"})
		return
	}

	/*********************
	 * Delete Server
	 *********************/
	deleteServer()

	c.JSON(http.StatusOK, msg)
}

func getServerHandler(c *gin.Context) {
	if kvm.LibvirtR == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "LibvirtR dose not exist!"})
	}
	server := kvm.LibvirtR.Old
	if server == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server dose not exist!"})
	}

	c.JSON(http.StatusOK, server)
}
