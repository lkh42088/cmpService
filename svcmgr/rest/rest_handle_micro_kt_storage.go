package rest

import (
	"cmpService/common/messages"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) UpdateKtAuthUrl(c *gin.Context) {
	var msg messages.KtAuthUrl
	c.Bind(&msg)
	fmt.Println("UpdateAuthUrl:", msg)
	server, err := h.db.UpdateKtAuthUrl(msg.Ip, msg.AuthUrl)
	if err != nil {
		return
	}
	fmt.Printf("UpdateAuthUrl: %+v\n", server)

	c.JSON(http.StatusOK, server)
}
