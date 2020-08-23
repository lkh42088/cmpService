package rest

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"cmpService/common/models"
	"cmpService/svcmgr/config"
	"cmpService/svcmgr/mcapi"
	"fmt"
	"github.com/evangwt/go-vncproxy"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
)

func (h *Handler) AddMcServer(c *gin.Context) {
	var msg mcmodel.McServer
	c.Bind(&msg)

	fmt.Printf("Add McServer : %v\n", msg)
	msg, err := h.db.AddMcServer(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Send to mc server
	server, _:= h.db.GetMcServerByServerIdx(msg.Idx)
	mcapi.SendMcRegisterServer(server)

	c.JSON(http.StatusOK, msg)
}

func DeleteMcImagesByServerIdx(idx int) {
	images , _ := config.SvcmgrGlobalConfig.Mariadb.GetMcImagesByServerIdx(idx)
	fmt.Println("DeleteMcImage: images ", images)
	for _, img := range images {
		fmt.Println("img ", img)
		config.SvcmgrGlobalConfig.Mariadb.DeleteMcImage(img)
	}
}

func DeleteMcNetworksByServerIdx(idx int) {
	networks, _ := config.SvcmgrGlobalConfig.Mariadb.GetMcNetworksByServerIdx(idx)
	fmt.Println("DeleteMcNetwork: networks ", networks)
	for _, net := range networks {
		fmt.Println("net ", net)
		config.SvcmgrGlobalConfig.Mariadb.DeleteMcNetwork(net)
	}
}

func (h *Handler) DeleteMcServer(c *gin.Context) {
	var msg messages.DeleteDataMessage
	c.Bind(&msg)
	fmt.Println("UnRegister Message: ", msg)
	for _, idx := range msg.IdxList {
		serverdetail, _ := config.SvcmgrGlobalConfig.Mariadb.GetMcServerByServerIdx(uint(idx))
		server := serverdetail.McServer
		fmt.Println("delete server : ", server)
		// Send to mc server
		mcapi.SendMcUnRegisterServer(server)
		// Dao: Network
		DeleteMcNetworksByServerIdx(idx)
		// Dao: Image
		DeleteMcImagesByServerIdx(idx)
		h.db.DeleteMcServer(server)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "created successfully"})
}

func (h *Handler) GetMcServers(c *gin.Context) {
	rowsPerPage, err := strconv.Atoi(c.Param("rows"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	orderBy := c.Param("orderby")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	order := c.Param("order")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	page := models.Pagination{
		TotalCount:  0,
		RowsPerPage: rowsPerPage,
		Offset:      offset,
		OrderBy:     orderBy,
		Order:       order,
	}
	fmt.Println("1. page:")
	page.String()
	servers, err := h.db.GetMcServersPage(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, servers)
}

func (h *Handler) GetMcServersByCpIdx(c *gin.Context) {
	cpIdx, _ := strconv.Atoi(c.Param("cpIdx"))
	servers, err := h.db.GetMcServersByCpIdx(cpIdx)
	if err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, servers)
}

func (h *Handler) AddMcVm(c *gin.Context) {
	var msg mcmodel.McVm
	c.Bind(&msg)

	fmt.Printf("Add McVm : %v\n", msg)

	msg.CurrentStatus = "Ready"
	msg, err := h.db.AddMcVm(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	server, err := h.db.GetMcServerByServerIdx(uint(msg.McServerIdx))
	// send to mcagent
	mcapi.SendAddVm(msg, server)

	c.JSON(http.StatusOK, msg)
}

// From Micro Cloud Server
func (h *Handler) UpdateMcVmFromMc(c *gin.Context) {
	var msg mcmodel.McVm
	c.Bind(&msg)

	fmt.Printf("Add McVm : %v\n", msg)

	msg, err := h.db.UpdateMcVmFromMc(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, msg)
}

func (h *Handler) DeleteMcVm(c *gin.Context) {
	var msg messages.DeleteDataMessage
	c.Bind(&msg)
	fmt.Println("UnRegister Message: ", msg)
	for _, idx := range msg.IdxList {
		var vm mcmodel.McVm
		vm.Idx = uint(idx)
		vm, err := h.db.DeleteMcVm(vm)
		if err != nil {
			continue
		}
		// send to mcagent
		server, err := h.db.GetMcServerByServerIdx(uint(vm.McServerIdx))
		mcapi.SendDeleteVm(vm, server)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "created successfully"})
}

func NewVNCProxy(targetAddr string) *vncproxy.Proxy {
	return vncproxy.New(&vncproxy.Config{
		LogLevel: vncproxy.DebugLevel,
		TokenHandler: func(r *http.Request) (addr string, err error) {
			// validate token and get forward vnc addr
			// ...
			addr = "192.168.0.73:5900"
			//addr = target
			return
		},
	})
}

func (h *Handler) GetMcVmVnc(c *gin.Context) {
	target := c.Param("target")
	port := c.Param("port")

	addr := fmt.Sprintf("%s:%s", target, port)
	fmt.Println("GetMcVmVnc:", addr)
	vncProxy := NewVNCProxy(addr)

	wh := websocket.Handler(vncProxy.ServeWS)
	wh.ServeHTTP(c.Writer, c.Request)
}

func (h *Handler) GetMcVms(c *gin.Context) {
	rowsPerPage, err := strconv.Atoi(c.Param("rows"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	orderBy := c.Param("orderby")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	order := c.Param("order")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	page := models.Pagination{
		TotalCount:  0,
		RowsPerPage: rowsPerPage,
		Offset:      offset,
		OrderBy:     orderBy,
		Order:       order,
	}
	fmt.Println("1. page:")
	page.String()
	vms, err := h.db.GetMcVmsPage(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, vms)
}

func (h *Handler) GetMcImages(c *gin.Context) {
	rowsPerPage, err := strconv.Atoi(c.Param("rows"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	orderBy := c.Param("orderby")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	order := c.Param("order")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	page := models.Pagination{
		TotalCount:  0,
		RowsPerPage: rowsPerPage,
		Offset:      offset,
		OrderBy:     orderBy,
		Order:       order,
	}
	fmt.Println("1. page:")
	page.String()
	images, err := h.db.GetMcImagesPage(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, images)
}

func (h *Handler) AddMcNetwork(c *gin.Context) {
	var msg mcmodel.McNetworks
	c.Bind(&msg)
	fmt.Println("AddMcNetwork:", msg)
	server, err := h.db.GetMcServerByServerIdx(uint(msg.McServerIdx))
	if err != nil {
		fmt.Println("AddMcNetwork: failed to get server - ", err)
		return
	}
	mcapi.SendAddNetwork(msg, server)
	c.JSON(http.StatusOK, msg)
}

func (h *Handler) DeleteMcNetwork(c *gin.Context) {
	var msg mcmodel.McNetworks
	c.Bind(&msg)
	fmt.Println("DeleteMcNetwork:", msg)
	server, err := h.db.GetMcServerByServerIdx(uint(msg.McServerIdx))
	if err != nil {
		fmt.Println("AddMcNetwork: failed to get server - ", err)
		return
	}
	mcapi.SendDeleteNetwork(msg, server)
	c.JSON(http.StatusOK, msg)
}

func (h *Handler) GetMcNetworks(c *gin.Context) {
	rowsPerPage, err := strconv.Atoi(c.Param("rows"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	orderBy := c.Param("orderby")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	order := c.Param("order")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	page := models.Pagination{
		TotalCount:  0,
		RowsPerPage: rowsPerPage,
		Offset:      offset,
		OrderBy:     orderBy,
		Order:       order,
	}
	fmt.Println("1. page:")
	page.String()
	networks, err := h.db.GetMcNetworksPage(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, networks)
}

func (h *Handler) GetMcImagesByServerIdx(c *gin.Context) {
	serverIdx, _ := strconv.Atoi(c.Param("serverIdx"))
	images, err := h.db.GetMcImagesByServerIdx(serverIdx)
	if err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, images)
}

func (h *Handler) GetMcNetworksByServerIdx(c *gin.Context) {
	serverIdx, _ := strconv.Atoi(c.Param("serverIdx"))
	images, err := h.db.GetMcNetworksByServerIdx(serverIdx)
	if err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, images)
}
