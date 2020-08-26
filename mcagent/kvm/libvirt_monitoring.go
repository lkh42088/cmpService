package kvm

import (
	"cmpService/common/mcmodel"
	"fmt"
	"sync"
	"time"
)

type LibvirtResource struct {
	Interval int
	Old *mcmodel.MgoServer
}

var LibvirtR *LibvirtResource

func NewLibvirtResource(interval int) *LibvirtResource {
	return &LibvirtResource{
		Interval: interval,
	}
}

func SetLibvirtResource(l *LibvirtResource) {
	LibvirtR = l
}

func ConfigureLibvirtResource() {
	l := NewLibvirtResource(10)
	SetLibvirtResource(l)
}

func (l *LibvirtResource) Start(parentwg *sync.WaitGroup) {
	loop := 1
	for {
		l.Run()
		time.Sleep(time.Duration(l.Interval * int(time.Second)))
		fmt.Printf("LibvirtResource Loop %d -----------------\n", loop)
		loop += 1
	}
}

func (l *LibvirtResource) Run() {
	server := GetMcServerInfo()
	server.DumpSummary()

	isChanged := false
	if l.Old != nil {
		isChanged = l.Old.Compare(server)
	} else {
		l.Old = &server
		isChanged = true
	}
	if isChanged {
		// Notify ...
		fmt.Printf("Changed! --> Notify...\n")
	}
}