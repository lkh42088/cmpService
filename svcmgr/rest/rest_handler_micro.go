package rest

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"cmpService/common/models"
	"cmpService/common/websocketproxy"
	"cmpService/svcmgr/config"
	conf "cmpService/svcmgr/config"
	"cmpService/svcmgr/mcapi"
	"flag"
	"fmt"
	"github.com/evangwt/go-vncproxy"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
	server, _ := h.db.GetMcServerByServerIdx(msg.Idx)
	mcapi.SendMcRegisterServer(server)

	c.JSON(http.StatusOK, msg)
}

func (h *Handler) UpdateMcServerResource(c *gin.Context) {
	var msg mcmodel.McServerMsg
	c.Bind(&msg)

	server, err := h.db.GetMcServerBySerialNumber(msg.SerialNumber)
	if err != nil && server.Idx == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	mcapi.ApplyMcServerResource(msg, server)
	c.JSON(http.StatusOK, msg)
}

func DeleteMcImagesByServerIdx(idx int) {
	images, _ := config.SvcmgrGlobalConfig.Mariadb.GetMcImagesByServerIdx(idx)
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
	if err != nil {
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

	msg, err := h.db.UpdateMcVm(msg)
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
		vm, err := h.db.GetMcVmByIdx(uint(idx))
		if err != nil {
			continue
		}
		// send to mcagent
		server, err := h.db.GetMcServerByServerIdx(uint(vm.McServerIdx))
		if err != nil {
			fmt.Println("DeleteMcVm: error", err)
		}
		mcapi.SendDeleteVm(vm, server)

		h.db.DeleteMcVm(vm)
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

func GetWebsockProxy() {
	addr := "ws://192.168.0.89:15901"
	flagBackend := flag.String("backend", addr, "Backend URL for proxying")
	target, err := url.Parse(*flagBackend)
	if err != nil {
		fmt.Println("GetWebsockProxy: error", err)
	}
	err = http.ListenAndServe("192.168.0.89:5900", websocketproxy.NewProxy(target))
	if err != nil {
		fmt.Println("GetWebsockProxy: listen error", err)
	}
}

func (h *Handler) GetMcVmVnc(c *gin.Context) {
	target := c.Param("target")
	port := c.Param("port")

	addr := fmt.Sprintf("%s:%s", target, port)
	fmt.Println("GetMcVmVnc:", addr)

	GetWebsockProxy()
	//vncProxy := NewVNCProxy(addr)
	//wh := websocket.Handler(vncProxy.ServeWS)
	//wh.ServeHTTP(c.Writer, c.Request)

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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, images)
}

func (h *Handler) GetMcNetworksByServerIdx(c *gin.Context) {
	serverIdx, _ := strconv.Atoi(c.Param("serverIdx"))
	images, err := h.db.GetMcNetworksByServerIdx(serverIdx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, images)
}

func GetVmWinInterface(c *gin.Context) {
	fmt.Println("GetMcVmsGraphs start!!")
	mac := c.Param("mac")
	fmt.Println("mac : ", mac)
	/* 20200911 todo where mac  조건 추가하고 mac값이 없을 경우 처리 필요 */

	/*---------------------------------------------------------------------------------------------------CPU*/
	dbname := "win_cpu"
	field := `"time","Percent_Idle_Time"`
	where := fmt.Sprintf(`host = 'win_vm'`) /*MAC 조회 필요*/
	//where := fmt.Sprintf(`host = 'win_vm' AND "mac_address" =~ /.*%s/`, mac)
	res := conf.GetMeasurementsWithConditionOrderLimit(dbname, field, where)

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		/*빈값일때 처리 필요......*/
		lib.LogWarn("win_cpu InfluxDB Response Error : No Data\n")
		c.JSON(http.StatusInternalServerError, "No Data")
		return
	}


	v := res.Results[0].Series[0].Values
	winCpu := make([]models.WinCpuStat, len(v))
	var cpuTime time.Time
	for i, data := range v {
		cpuTime, _ = time.Parse(time.RFC3339, data[0].(string))

		winCpu[i].Time = cpuTime
		if err := MakeStructForStatsWinCpu(&winCpu[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}

	/*---------------------------------------------------------------------------------------------------MEM*/
	dbname = "win_mem"
	field = `"time","Available_Bytes"`
	where = fmt.Sprintf(`host = 'win_vm'`) /*MAC 조회 필요*/
	res = conf.GetMeasurementsWithConditionOrderLimit(dbname, field, where)

	if res.Results[0].Series == nil ||
		/*빈값일때 처리......*/
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("win_mem InfluxDB Response Error : No Data\n")
		c.JSON(http.StatusInternalServerError, "No Data")
		return
	}

	v = res.Results[0].Series[0].Values
	winMem := make([]models.WinMemStat, len(v))
	var memTime time.Time
	for i, data := range v {
		memTime, _ = time.Parse(time.RFC3339, data[0].(string))

		winMem[i].Time = memTime
		if err := MakeStructForStatsWinMem(&winMem[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}

	/*--------------------------------------------------------------------------------------------------DISK*/
	dbname = "win_disk"
	field = `"time","Free_Megabytes"`
	where = fmt.Sprintf(`host = 'win_vm'`) /*MAC 조회 필요*/
	res = conf.GetMeasurementsWithConditionOrderLimit(dbname, field, where)

	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("win_disk InfluxDB Response Error : No Data\n")
		c.JSON(http.StatusInternalServerError, "No Data")
		return
	}

	v = res.Results[0].Series[0].Values
	winDisk := make([]models.WinDiskStat, len(v))
	var diskTime time.Time
	for i, data := range v {
		diskTime, _ = time.Parse(time.RFC3339, data[0].(string))

		winDisk[i].Time = diskTime
		if err := MakeStructForStatsWinDisk(&winDisk[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}
	/*-----------------------------------------------------------------------------------------------TRAFFIC*/

	dbname = "win_net"
	field = `"time","Bytes_Received_persec","Bytes_Sent_persec"`
	//where = fmt.Sprintf(`"MAC" =~ /.*%s/ AND time > now() - %s`, mac, "1h")
	where = fmt.Sprintf(`time > now() - %s`, "10m")
	res = conf.GetMeasurementsWithCondition(dbname, field, where)
	//fmt.Println(res.Err)
	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 {
		lib.LogWarn("win_net InfluxDB Response Error : No Data\n")
		c.JSON(http.StatusInternalServerError, res.Err)
		return
	}

	//Bytes_Received_persec RX
	//Bytes_Sent_persec TX
	// Convert response data
	v = res.Results[0].Series[0].Values
	winTraffic := make([]models.WinVmIfStat, len(v))
	var convTime time.Time
	for i, data := range v {
		// select time check
		convTime, _ = time.Parse(time.RFC3339, data[0].(string))

		// make struct
		winTraffic[i].Time = convTime
		//winTraffic[i].IfPhysAddress = mac
		if err := MakeStructForWinStats(&winTraffic[i], data); err != nil {
			lib.LogWarn("Error : %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}

	//deltaStats := MakeDeltaWinValues(winTraffic)

	graph := mcmodel.McWinVmGraph{
		Cpu:     winCpu[0],
		Mem:     winMem[0],
		Disk:    winDisk[0],
		Traffic: winTraffic,
	}

	/*fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("graph : ", graph)
	fmt.Println("graph CPU : ", graph.Cpu)
	fmt.Println("graph Mem : ", graph.Mem)
	fmt.Println("graph Disk : ", graph.Disk)
	fmt.Println("graph Traffic : ", graph.Traffic)*/

	c.JSON(http.StatusOK, graph)
}
