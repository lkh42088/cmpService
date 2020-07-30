package mcrest

import (
	"cmpService/mcagent/mcmodel"
	"cmpService/mcagent/mcmongo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addVmHandler(c *gin.Context) {
	var msg mcmodel.VmEntry
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert VM to Mongodb
	_, err = mcmongo.McMongo.AddVm(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, msg)
}

func getVmByIdHandler(c *gin.Context) {
	id := c.Param("id")

	// Get VMs from Mongodb
	vm, err := mcmongo.McMongo.GetVmById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, vm)
}
