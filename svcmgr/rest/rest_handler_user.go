package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetCustomersByName(c *gin.Context) {
	if h.db == nil {
		return
	}
	name := c.Param("name")
	customers, err := h.db.GetCustomersByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}