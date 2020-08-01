package mcrest

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/mcmongo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func addVmHandler(c *gin.Context) {
	var msg mcmodel.MgoVm
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

func deleteVmHandler(c *gin.Context) {
	var msg mcmodel.MgoVm
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = mcmongo.McMongo.DeleteVm(msg.Idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, msg)
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
