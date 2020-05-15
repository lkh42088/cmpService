package main

import (
	"flag"
	"fmt"
	"cmpService/collector/config"
	"cmpService/collector/rest"
	"cmpService/collector/snmpapi"
	"cmpService/collector/statistics"
	"sync"
)

func main() {
	configFile := flag.String("file", "/etc/collector/collector.conf",
		"Input configuration file")
	flag.Parse()
	collect(*configFile)
}

func collect(configPath string) {
	var wg sync.WaitGroup

	// Set CollectorConfig Path
	config.SetConfig(configPath)
	snmpapi.InitConfig()

	fmt.Println("Start ++")
	wg.Add(4)

	// Start restapi server
	go rest.Start(&wg)

	// Start snmp collection
	go snmpapi.Start(&wg)

	// Start statistics
	go statistics.Start(&wg)

	// Store to influxdb
	go snmpapi.WriteMetricInfluxDB(&wg)

	fmt.Println("End --")
	wg.Wait()
}

