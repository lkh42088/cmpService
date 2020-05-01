package main

import (
	"flag"
	"fmt"
	"nubes/collector/config"
	"nubes/collector/rest"
	"nubes/collector/snmpapi"
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
	wg.Add(2)

	// Start restapi server
	go rest.Start(&wg)

	// Start snmp collection
	go snmpapi.Start(&wg)

	fmt.Println("End --")
	wg.Wait()
}

