package main

import (
	"flag"
	"fmt"
	"nubes/collector/conf"
	"nubes/collector/rest"
	"nubes/collector/snmpapi"
	"sync"
)


func main() {
	configFile := flag.String("file", "/etc/collector/collector.conf",
		"Input configuration file")
	flag.Parse()
	collect(*configFile)
	//simpleCollect()
}

func collect(configFile string) {
	var wg sync.WaitGroup

	// Process configuration information
	conf.ProcessConfig(configFile)

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
