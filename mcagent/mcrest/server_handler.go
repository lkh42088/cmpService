package mcrest

import (
	"cmpService/common/mcmodel"
	"cmpService/common/utils"
	"cmpService/mcagent/config"
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

func getResourceHandler(c *gin.Context) {
	var resource McResourceMsg
	resource.GlobalConfig = config.GetMcGlobalConfig()
	resource.DnatList = utils.GetDnatList()
	resource.CreateVmList = &kvm.CreateVmFsm.Vms
	resource.CacheVmList = &repo.GlobalVmCache
	resource.LibvirtVmList = kvm.LibvirtR.Old.Vms
	resource.CronVmList = &kvm.CronSnap.Vms
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
	config.WriteServerStatus(msg.SerialNumber, msg.CompanyName, msg.CompanyIdx, true)
	config.SetSerialNumber2GlobalConfig(msg.SerialNumber)

	/*********************
	 * Add Repo
	 *********************/
	repo.AddServer2Repo(&msg)

	//server, _ := GetMcServer()
	server := kvm.GetMcServerInfo()
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

	/*********************
	 * Delete Repo
	 *********************/
	repo.DeleteServer2Repo()

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
