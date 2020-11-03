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
	_, err := repo.AddAccessFilter2Db(msg)
	if err != nil {
		fmt.Println("addMcFilterAddress error: ", err)
	}
	c.JSON(http.StatusOK, msg)
}

func deleteMcFilterAddress(c *gin.Context) {
	var msg mcmodel.McFilterRule
	c.ShouldBindJSON(&msg)
	fmt.Println("deleteMcFilterAddress:")
	msg.Dump()
	msg.Idx = 0
	rule, err := repo.GetAccessFilter2Db(msg)
	if err != nil {
		fmt.Println("deleteMcFilterAddress error1: ", err)
		return
	}
	_, err = repo.DeleteAccessFilter2Db(rule)
	if err != nil {
		fmt.Println("deleteMcFilterAddress error2: ", err)
	}
	c.JSON(http.StatusOK, msg)
}
