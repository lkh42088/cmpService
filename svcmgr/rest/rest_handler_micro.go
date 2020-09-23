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
	cpName := c.Param("cpName")
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
	fmt.Println("GetMcServers 1. page:")
	page.String()
	servers, err := h.db.GetMcServersPage(page, cpName)

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
	msg.Dump()

	msg.CurrentStatus = "Ready"
	msg, err := h.db.AddMcVm(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	server, err := h.db.GetMcServerByServerIdx(uint(msg.McServerIdx))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	cpName := c.Param("cpName")
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
	vms, err := h.db.GetMcVmsPage(page, cpName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	fmt.Println("😡😡😡 😡😡😡 😡😡😡 😡😡😡 😡😡😡 😡😡😡 vms : ", vms)
	fmt.Println("😡😡😡 cpName : ", cpName)

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

func (h *Handler) GetVmWinInterface(c *gin.Context) {
	fmt.Println("GetMcVmsGraphs start!!")
	mac := c.Param("mac")
	currentStatus := c.Param("currentStatus")

	var graph mcmodel.McWinVmGraph

	vmInfo, err := h.db.GetMcVmByMac(mac)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	/*---------------------------------------------------------------------------------------------------CPU*/
	dbname := "win_cpu"
	field := `"time","Percent_Idle_Time"`
	//where := fmt.Sprintf(`host = 'win_vm'`) /*MAC 조회 필요*/
	where := fmt.Sprintf(`instance = '_Total' AND "mac_address" = '%s'`, mac)
	res := conf.GetMeasurementsWithConditionOrderLimit(dbname, field, where)

	winCpu := make([]mcmodel.WinCpuStat, 1)
	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 || currentStatus != "running"{
		//fmt.Println("")
		lib.LogWarn(mac, "win_cpu InfluxDB Response Error" +
			" : No Data\n")

		winCpu[0].PercentIdleTime = 0
		winCpu[0].Total = 0

		graph.Cpu = winCpu[0]
	} else {
		v := res.Results[0].Series[0].Values
		winCpu := make([]mcmodel.WinCpuStat, len(v))
		var cpuTime time.Time
		for i, data := range v {
			cpuTime, _ = time.Parse(time.RFC3339, data[0].(string))

			winCpu[i].Time = cpuTime
			if err := MakeStructForStatsWinCpu(&winCpu[i], data, mac); err != nil {
				lib.LogWarn("Error : %s\n", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
				return
			}
		}

		graph.Cpu = winCpu[0]
	}

	/*---------------------------------------------------------------------------------------------------MEM*/
	dbname = "win_mem"
	field = `"time","Available_Bytes"`
	//where = fmt.Sprintf(`host = 'win_vm'`) /*MAC 조회 필요*/
	where = fmt.Sprintf(`mac_address = '%s'`, mac)
	res = conf.GetMeasurementsWithConditionOrderLimit(dbname, field, where)

	winMem := make([]mcmodel.WinMemStat, 1)
	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 || currentStatus != "running"{
		lib.LogWarn("win_mem InfluxDB Response Error : No Data\n")
		//var winMem models.WinMemStat;
		//winMem.AvailableBytes = json.Number(0)
		//c.JSON(http.StatusInternalServerError, "No Data")
		//return
		winMem[0].AvailableBytes = 0
		winMem[0].Total = float64(vmInfo.Ram)
		graph.Mem = winMem[0]
	} else {
		v := res.Results[0].Series[0].Values
		winMem := make([]mcmodel.WinMemStat, len(v))
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
		winMem[0].Total = float64(vmInfo.Ram)
		graph.Mem = winMem[0]
	}

	/*--------------------------------------------------------------------------------------------------DISK*/
	dbname = "win_disk"
	field = `"time","Free_Megabytes"`
	//where = fmt.Sprintf(`host = 'win_vm'`) /*MAC 조회 필요*/
	where = fmt.Sprintf(`instance = 'C:' AND "mac_address" = '%s'`, mac)
	res = conf.GetMeasurementsWithConditionOrderLimit(dbname, field, where)

	winDisk := make([]mcmodel.WinDiskStat, 1)
	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 || currentStatus != "running"{
		lib.LogWarn("win_disk InfluxDB Response Error : No Data\n")
		//var winDisk models.WinDiskStat
		//winDisk.FreeMegabytes = json.Number(0)
		//c.JSON(http.StatusInternalServerError, "No Data")
		//return
		winDisk[0].FreeMegabytes = 0
		winDisk[0].Total = float64(vmInfo.Hdd)
		graph.Disk = winDisk[0]
	} else {
		v := res.Results[0].Series[0].Values
		winDisk := make([]mcmodel.WinDiskStat, len(v))
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
		winDisk[0].Total = float64(vmInfo.Hdd)
		graph.Disk = winDisk[0]
	}
	/*-----------------------------------------------------------------------------------------------TRAFFIC*/

	dbname = "win_net"
	field = `"time","Bytes_Received_persec","Bytes_Sent_persec"`
	//where = fmt.Sprintf(`"MAC" =~ /.*%s/ AND time > now() - %s`, mac, "1h")
	//where = fmt.Sprintf(`time > now() - %s`, "10m")
	//where = fmt.Sprintf(`host = 'win_vm' AND "mac_address" =~ /.*%s/`, mac)
	where = fmt.Sprintf(`"mac_address" = '%s' AND time > now() - %s`, mac, "10m")
	res = conf.GetMeasurementsWithCondition(dbname, field, where)
	winTraffic := make([]mcmodel.WinVmIfStat, 1)
	//fmt.Println(res.Err)
	if res.Results[0].Series == nil ||
		len(res.Results[0].Series[0].Values) == 0 || currentStatus != "running"{
		lib.LogWarn("win_net InfluxDB Response Error : No Data\n")
		//var winTraffic models.WinVmIfStat
		//winTraffic.BytesSentPersec = 0
		//c.JSON(http.StatusInternalServerError, res.Err)
		//return
		winTraffic[0].BytesSentPersec = float64(0)
		graph.Traffic = winTraffic
	} else {
		v := res.Results[0].Series[0].Values
		winTraffic = make([]mcmodel.WinVmIfStat, len(v))
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
		graph.Traffic = winTraffic
	}

	//deltaStats := MakeDeltaWinValues(winTraffic)
	fmt.Println("")
	fmt.Println(mac, "-------------------------------------------------------------------------------------------------------------------")
	fmt.Println("graph CPU : ", graph.Cpu)
	fmt.Println("graph Mem : ", graph.Mem)
	fmt.Println("graph Disk : ", graph.Disk)
	fmt.Println("graph Traffic : ", graph.Traffic)

	c.JSON(http.StatusOK, graph)
}

func (h *Handler) GetVmSnapshotConfig(c *gin.Context) {
	serverIdx, _ := strconv.Atoi(c.Param("serverIdx"))
	fmt.Println("GetVmSnapshotConfig")
	server, err := h.db.GetMcServerByServerIdx(uint(serverIdx))
	if err != nil {
		return
	}
	vms, err := config.SvcmgrGlobalConfig.Mariadb.GetMcVmsByServerIdx(int(server.Idx))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, vms)
}

func (h *Handler) AddVmSnapshot(c *gin.Context) {
	var msg messages.SnapshotConfigMsg
	c.Bind(&msg)
	fmt.Println("AddVmSnapshot:", msg)
	server, err := h.db.GetMcServerByServerIdx(msg.ServerIdx)
	if err != nil {
		return
	}
	mcapi.SendAddVmSnapshot(msg, server)
}

func (h *Handler) DeleteVmSnapshot(c *gin.Context) {
	var msg messages.SnapshotConfigMsg
	c.Bind(&msg)
	fmt.Println("DeleteVmSnapshot:", msg)
	server, err := h.db.GetMcServerByServerIdx(msg.ServerIdx)
	if err != nil {
		return
	}
	mcapi.SendDeleteVmSnapshot(msg, server)
}

func (h *Handler) UpdateVmSnapshot(c *gin.Context) {
	var msg messages.SnapshotConfigMsg
	c.Bind(&msg)
	fmt.Println("UpdateVmSnapshot:", msg)
	server, err := h.db.GetMcServerByServerIdx(msg.ServerIdx)
	if err != nil {
		return
	}
	mcapi.SendUpdateVmSnapshot(msg, server)
}

func (h *Handler) UpdateVmStatus(c *gin.Context) {
	var msg messages.VmStatusActionMsg
	c.Bind(&msg)
	fmt.Println("UpdateVmSnapshot:", msg)
	server, err := h.db.GetMcServerByServerIdx(msg.ServerIdx)
	if err != nil {
		return
	}
	mcapi.SendUpdateVmStatus(msg, server)
}
