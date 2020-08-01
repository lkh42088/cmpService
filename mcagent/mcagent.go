package main

import (
	"cmpService/mcagent/agent"
	"flag"
)

func main() {
	config := flag.String("file", "mcagent.conf",
		"Input configuration file")
	flag.Parse()
	agent.Start(*config)
}

