package mcrest

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/repo"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getMcFilterAddress(c *gin.Context) {
	var msg []mcmodel.McFilterRule
	// todo
	c.JSON(http.StatusOK, msg)
}

func addMcFilterAddress(c *gin.Context) {
	var msg mcmodel.McFilterRule
	c.ShouldBindJSON(&msg)
	fmt.Println("addMcFilterAddress:")
	msg.Dump()
	msg.Idx = 0
	rule, err := repo.AddAccessFilter(msg)
	if err != nil {
		fmt.Println("addMcFilterAddress error: ", err)
	}
	c.JSON(http.StatusOK, rule)
}

func deleteMcFilterAddress(c *gin.Context) {
	var msg mcmodel.McFilterRule
	c.ShouldBindJSON(&msg)
	fmt.Println("deleteMcFilterAddress:")
	msg.Dump()
	msg.Idx = 0
	rule, err := repo.DeleteAccessFilter(msg)
	if err != nil {
		fmt.Println("deleteMcFilterAddress error: ", err)
	}
	c.JSON(http.StatusOK, rule)
}
