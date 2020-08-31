package kvm

import (
	"encoding/json"
	"fmt"
	"github.com/libvirt/libvirt-go"
	"sync"
	"time"
)

type LibvirtStatistics struct {
	Interval int
}

var LibvirtS *LibvirtStatistics

func NewLibvirtStatistics(interval int) *LibvirtStatistics {
	return &LibvirtStatistics{
		Interval: interval,
	}
}

func SetLibvirtStatistics(s *LibvirtStatistics) {
	LibvirtS = s
}

func ConfigureLibvirtStatstics() {
	s := NewLibvirtStatistics(10)
	SetLibvirtStatistics(s)
}

func (s *LibvirtStatistics) Start(parentwg *sync.WaitGroup) {
	loop := 1
	for {
		s.Run()
		time.Sleep(time.Duration(s.Interval * int(time.Second)))
		fmt.Printf("LibvirtStatistics Loop %d -----------------\n", loop)
		loop += 1
	}
	parentwg.Done()
}

func (s *LibvirtStatistics) Run() {

	doms, err := GetDomainListAll()
	if err != nil {
		return
	}
	for _, dom := range doms {
		name, _ := dom.GetName()
		fmt.Println(">>>> Statistics:", name)
		vcpumax, _ := dom.GetMaxVcpus()
		fmt.Println(" cpuinfo: max", vcpumax)
		//vcpuinfo, _ := dom.GetVcpus()
		//DumpCpuInfoSimple(&vcpuinfo)
		//vcpustat, _ := dom.GetCPUStats(0, 0,0)
		//DumpCpuState(&vcpustat)
	}
}

func DumpCpuState(c *[]libvirt.DomainCPUStats) string {
	pretty, _ := json.MarshalIndent(c, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func DumpCpuInfoSimple(c *[]libvirt.DomainVcpuInfo) {
	for _, entry := range *c {
		fmt.Println("Number:", entry.Number, "State:", entry.State,
			"CpuTime:", entry.CpuTime, "Cpu:", entry.Cpu, "CpuMap:", entry.CpuMap)
	}
}

func DumpCpuInfo(c *[]libvirt.DomainVcpuInfo) string {
	pretty, _ := json.MarshalIndent(c, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func SystemCpu() {
	conn, err := GetQemuConnect()
	if err != nil {
		return
	}
	cpumap, cpunum, err := conn.GetCPUMap(0)
	if err != nil {
		return
	}
	fmt.Println("cpunum:", cpunum)
	fmt.Println("cpumap:", cpumap)
}