package kvm

import (
	"encoding/json"
	"fmt"
	"github.com/libvirt/libvirt-go"
	"math"
	"sync"
	"time"
)

type LibvirtStatistics struct {
	Interval int
}

var LibvirtS *LibvirtStatistics

type LibvirtCpuDelta struct {
	TotalCpu	uint64
	EachCpu		uint64
}
const NANO_SECONDS = 10000000000
const INTERVAL_SECONDS = 5

var testVal = make([]LibvirtCpuDelta, 5)

func NewLibvirtStatistics(interval int) *LibvirtStatistics {
	return &LibvirtStatistics{
		Interval: interval,
	}
}

func SetLibvirtStatistics(s *LibvirtStatistics) {
	LibvirtS = s
}

func ConfigureLibvirtStatstics() {
	s := NewLibvirtStatistics(INTERVAL_SECONDS)
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
		dom.SetVcpus(0)
	}
	for i, dom := range doms {
		name, _ := dom.GetName()
		//fmt.Println(">>>> Statistics:", name)

		vcpumax, _ := dom.GetMaxVcpus()
		//fmt.Println(" cpuinfo: max", vcpumax)

		//vcpuinfo, _ := dom.GetVcpus()
		//DumpCpuInfoSimple(&vcpuinfo)

		//vcpustat, _ := dom.GetCPUStats(-1, 0,0)
		//DumpCpuStats(&vcpustat)

		//jobStats, _ := dom.GetJobStats(1)
		//DumpDomainJobInfo(jobStats)

		//tmpStats, _ := dom.GetInfo()
		//DumpDomainInfo(tmpStats) // vm simple cpu usage

		vcpustat, _ := dom.GetCPUStats(-1, 0, 0)
		//DumpCpuStats(&vcpustat)
		CalcEachCPUUsage(vcpustat, i, name, vcpumax)

		/** memory */
		//memstat, _ := dom.MemoryStats(4, 0)
		//DumpDomainMemStats(memstat)

		//mempeek, _ := dom.MemoryPeek(0, 64000, 0)
		//DumpDomainMemPeek(mempeek)

		//memmax, _ := dom.GetMaxMemory()
		//DumpDomainMemMax(memmax)

		//memparam, _ := dom.GetMemoryParameters(0)
		//DumpDomainMemParam(memparam)

		//fsinfo, _ := dom.GetFSInfo(0)
		//DumpFSInfo(fsinfo)

	}
	//SystemCpu()
}

func DumpFSInfo(fs []libvirt.DomainFSInfo) string {
	pretty, _ := json.MarshalIndent(fs, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func DumpDomainMemParam(m *libvirt.DomainMemoryParameters) string {
	pretty, _ := json.MarshalIndent(m, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func DumpDomainMemMax(m uint64) string {
	pretty, _ := json.MarshalIndent(m, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func DumpDomainMemStats(m []libvirt.DomainMemoryStat) string {
	pretty, _ := json.MarshalIndent(m, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func DumpDomainJobInfo(c *libvirt.DomainJobInfo) string {
	pretty, _ := json.MarshalIndent(c, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func CalcEachCPUUsage (s []libvirt.DomainCPUStats, i int, name string, max uint) {
	for _, cpu := range(s) {
		deltaCpuTime := cpu.CpuTime - cpu.UserTime - cpu.SystemTime
		timeUnit := uint64(INTERVAL_SECONDS * NANO_SECONDS)
		result := float64(deltaCpuTime - testVal[i].TotalCpu) / float64(timeUnit)

		fmt.Printf("%s cpu (used : %.2f %%)\n", name, result * 100)
		testVal[i].TotalCpu = uint64(math.Abs(float64(deltaCpuTime)))
	}
}

func CalcCPUUsage (s []libvirt.DomainCPUStats, i int) {
	for _, cpu := range(s) {
		fmt.Println("CPU ", i, " USAGE: ", (float64(cpu.CpuTime-cpu.SystemTime-cpu.UserTime) - float64(testVal[i].TotalCpu))/(2*NANO_SECONDS), "%")
		testVal[i].TotalCpu = cpu.CpuTime-cpu.SystemTime-cpu.UserTime
	}
}

func DumpDomainInfo(c *libvirt.DomainInfo) string {
	pretty, _ := json.MarshalIndent(c, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func DumpCpuStats(c *[]libvirt.DomainCPUStats) string {
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