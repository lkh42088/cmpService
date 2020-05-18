package main

import (
	"cmpService/dbmigrator/config"
	"flag"
)

func main() {
	configFile := flag.String("file", "/etc/dbmigrator/dbmigrator.conf",
		"Input configuration file")
	flag.Parse()
	config.SetConfig(*configFile)
}

