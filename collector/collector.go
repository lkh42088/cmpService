package main

import (
	"cmpService/collector/config"
	"cmpService/collector/rest"
	"cmpService/collector/snmpapi"
	"cmpService/collector/statistics"
	"flag"
	"fmt"
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

	fmt.Println("StartByVirsh ++")
	wg.Add(4)

	// StartByVirsh restapi server
	go rest.Start(&wg)

	// StartByVirsh snmp collection
	go snmpapi.Start(&wg)

	// StartByVirsh statistics
	go statistics.Start(&wg)

	// Store to influxdb
	go snmpapi.WriteMetricInfluxDB(&wg)

	fmt.Println("End --")
	wg.Wait()
}
