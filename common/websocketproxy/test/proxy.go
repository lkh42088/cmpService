package main

import (
	"cmpService/common/websocketproxy"
	"flag"
	"log"
	"net/http"
	"net/url"
)

var (
	flagBackend = flag.String("backend", "", "Backend URL for proxying")
)

func main() {
	u, err := url.Parse(*flagBackend)
	if err != nil {
		log.Fatalln(err)
	}

	err = http.ListenAndServe("192.168.0.89:5900", websocketproxy.NewProxy(u))
	if err != nil {
		log.Fatalln(err)
	}
}