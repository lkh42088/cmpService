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
	s := NewLibvirtStatistics(5)
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
		fmt.Println(">>>> Statistics:", name)

		//vcpumax, _ := dom.GetMaxVcpus()
		//fmt.Println(" cpuinfo: max", vcpumax)

		//vcpuinfo, _ := dom.GetVcpus()
		//DumpCpuInfoSimple(&vcpuinfo)

		//vcpustat, _ := dom.GetCPUStats(-1, 0,0)
		//DumpCpuStats(&vcpustat)

		//jobStats, _ := dom.GetJobStats(1)
		//DumpDomainJobInfo(jobStats)
		if i == 1 {
			//tmpStats, _ := dom.GetInfo()
			//DumpDomainInfo(tmpStats) // vm simple cpu usage
			vcpustat, _ := dom.GetCPUStats(-1, 0, 0)
			//DumpCpuStats(&vcpustat)
			//CalcTotalCPUUsage(*tmpStats, vcpustat, i)
			CalcEachCPUUsage(vcpustat, i)
		}
	}
	//SystemCpu()
}

func DumpDomainJobInfo(c *libvirt.DomainJobInfo) string {
	pretty, _ := json.MarshalIndent(c, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func CalcTotalCPUUsage (s libvirt.DomainInfo, cpu []libvirt.DomainCPUStats, i int) {
	//fmt.Println("CPU: ", i, ", Time: ", cpu.CpuTime / 1000000000.)
	if len(cpu) > 0 {
		fmt.Println("CPU ", i, " USAGE: ",
			//float64(s.CpuTime- testVal[i].TotalCpu - cpu[0].SystemTime - cpu[0].UserTime + testVal[i].EachCpu)*100/(NANO_SECONDS), "%")
		float64(s.CpuTime- testVal[i].TotalCpu)*100/(5*NANO_SECONDS), "%")
		testVal[i].TotalCpu = s.CpuTime
		testVal[i].EachCpu = cpu[0].SystemTime + cpu[0].UserTime
	}
}

func CalcEachCPUUsage (s []libvirt.DomainCPUStats, i int) {
	for _, cpu := range(s) {
		//fmt.Println("CPU: ", i, ", Time: ", cpu.CpuTime / 1000000000.)
		fmt.Println(float64(cpu.CpuTime - cpu.UserTime - cpu.SystemTime))
		fmt.Println(math.Abs(float64(cpu.CpuTime - cpu.UserTime - cpu.SystemTime)))
		fmt.Println(((math.Abs(float64(cpu.CpuTime - cpu.UserTime - cpu.SystemTime)) - float64(testVal[i].TotalCpu)))*100)
		fmt.Println("CPU ", i, " USAGE: ", ((math.Abs(float64(cpu.CpuTime - cpu.UserTime - cpu.SystemTime)) - float64(testVal[i].TotalCpu))*100*10000)/(5*NANO_SECONDS), "%")
		testVal[i].TotalCpu = uint64(math.Abs(float64(cpu.CpuTime - cpu.UserTime - cpu.SystemTime)))
	}
}

func CalcCPUUsage (s []libvirt.DomainCPUStats, i int) {
	for _, cpu := range(s) {
		//fmt.Println("CPU: ", i, ", Time: ", cpu.CpuTime / 1000000000.)
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