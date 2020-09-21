package mcrest

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/kvm"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func addNetworkHandler(c *gin.Context) {
	var msg mcmodel.McNetworks
	c.ShouldBindJSON(&msg)
	kvm.CreateNetworkByMgoNetwork(msg)
	c.JSON(http.StatusOK, msg)
}

func deleteNetworkHandler(c *gin.Context) {
	var msg mcmodel.McNetworks
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

