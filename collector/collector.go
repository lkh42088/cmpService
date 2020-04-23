package main

import (
	"fmt"
	"nubes/collector/rest"
	"nubes/collector/snmpapi"
	"sync"
)

func main() {
	collect()
	//simpleCollect()
}

func collect() {
	var wg sync.WaitGroup

	rest.FindConfig()
	var r = rest.ReadConf()

	fmt.Println("Start ++")

	wg.Add(2)

	go rest.RunAPI(&wg)

	go snmpapi.RegularCollect(&wg)

	fmt.Println("End --")

	wg.Wait()
}

func simpleCollect() {
	devices := make([]snmpapi.SnmpDevice, 2)

	devices[0].Device.Ip = "121.156.65.139"
	devices[0].Device.SnmpCommunity = "nubes"

	devices[1].Device.Ip = "192.168.122.11"
	devices[1].Device.SnmpCommunity = "nubes"

	snmpapi.ProcessSnmpAllDevice(devices)
}
