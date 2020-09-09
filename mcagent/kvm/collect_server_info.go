package kvm

import (
	"encoding/json"
	"fmt"
	"github.com/google/gopacket/routing"
	"github.com/pkg/errors"
	cpu "github.com/shirou/gopsutil/cpu"
	disk "github.com/shirou/gopsutil/disk"
	host "github.com/shirou/gopsutil/host"
	mem "github.com/shirou/gopsutil/mem"
	net "github.com/shirou/gopsutil/net"
	network "net"
	"strings"
	"syscall"
)

type NetInterface struct {
	Name    string   `json:name`
	Address []string `json:address`
}

type SysInfo struct {
	Idx             int    `gorm:"primary_key;column:idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
	Hostname        string `gorm:"column:hostname;not null;comment:'SERVER 이름'" json:"hostname"`
	OS              string `gorm:"column:os;comment:'OS 명'" json:"os"`
	Uptime          uint64 `gorm:"column:uptime;comment:'UPTIME'" json:"uptime"`
	BootTime        uint64 `gorm:"column:boottime;comment:'SERVER BOOTTIME'" json:"bootTime"`
	CpuCore         int    `gorm:"column:cpu_core;comment:'CPU Core 개수'" json:"cpuCore"`
	CpuModel        string `gorm:"column:cpu_model;comment:'CPU 모델명'" json:"cpuModel"`
	Platform        string `gorm:"column:platform;comment:'Platform'" json:"platform"`
	PlatformVersion string `gorm:"column:platform_version;comment:'Platform 버전'" json:"platformVersion"`
	KernelArch      string `gorm:"column:kernel_arch;comment:'KERNEL 아키텍처'" json:"kernelArch"`
	KernelVersion   string `gorm:"column:kernel_version;comment:'KERNEL 버전'" json:"kernelVersion"`
	IP              string `gorm:"column:ip;not null;comment:'SERVER IP'" json:"ip"`
	IfName          string `gorm:"column:if_name;comment:'Interface Name'" json:"ifName"`
	IfMac           string `gorm:"column:if_mac;comment:'Interface MAC'" json:"ifMac"`
	MemTotal        int64  `gorm:"column:mem_total;comment:'MEMORY 용량'" json:"mem"`
	DiskTotal       int64  `gorm:"column:disk_total;comment:'HDD DISK 용량'" json:"disk"`
}

type SysInfoDetail struct {
	Host        host.InfoStat          `json:hostname`
	Temperature []host.TemperatureStat `json:temperature`
	Platform    string                 `json:platform`
	CPU         []cpu.InfoStat         `json:cpu`
	RAM         uint64                 `json:ram`
	DiskUsage   uint64                 `json:diskUsage`
	DiskPart    []disk.PartitionStat   `json:disk`
	Net         []net.InterfaceStat    `json:interface`
}

func GetSysInfoLegacy() {
	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)

	if err == nil {
		fmt.Printf("sysinfo : %+v\n", info)
	}
}

func GetEnvVar() {
	env := syscall.Environ()
	for i := range env {
		fmt.Println(env[i])
	}
}

func GetSysInfo() SysInfo {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")
	netIf, _ := net.Interfaces()

	info := new(SysInfo)
	info.Hostname = hostStat.Hostname
	info.OS = hostStat.OS
	info.Uptime = hostStat.Uptime
	info.BootTime = hostStat.BootTime
	info.CpuCore = len(cpuStat)
	info.CpuModel = cpuStat[0].ModelName
	info.Platform = hostStat.Platform
	info.PlatformVersion = hostStat.PlatformVersion
	info.KernelArch = hostStat.KernelArch
	info.KernelVersion = hostStat.KernelVersion
	tmpIP, _ := FindRouteInterface("0.0.0.0")
	info.IP = tmpIP.String()
	for _, v := range netIf {
		for _, ip := range v.Addrs {
			if info.IP == strings.Split(ip.Addr, "/")[0] {
				info.IfName = v.Name
				info.IfMac = v.HardwareAddr
			}
		}
	}
	info.MemTotal = int64(vmStat.Total)
	info.DiskTotal = int64(diskStat.Total)

	pretty, _ := json.MarshalIndent(info, "", "  ")
	fmt.Printf("%s\n", string(pretty))

	return *info
}

func GetSysInfoAll() SysInfoDetail {
	hostStat, _ := host.Info()
	temper, _ := host.SensorsTemperatures()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")
	diskPart, _ := disk.Partitions(true)
	netIf, _ := net.Interfaces()

	info := new(SysInfoDetail)
	info.Host = *hostStat
	info.Platform = hostStat.Platform
	info.RAM = vmStat.Total / 1024 / 1024
	info.DiskUsage = diskStat.Total / 1024 / 1024
	for _, v := range cpuStat {
		info.CPU = append(info.CPU, v)
	}
	for _, v := range temper {
		info.Temperature = append(info.Temperature, v)
	}
	for _, v := range diskPart {
		info.DiskPart = append(info.DiskPart, v)
	}
	for _, v := range netIf {
		info.Net = append(info.Net, v)
	}

	pretty, _ := json.MarshalIndent(info, "", "  ")
	fmt.Printf("%s\n", string(pretty))

	return *info
}

func FindRouteInterface(addr string) (network.IP, error) {
	var ip network.IP
	if ip = network.ParseIP(addr); ip == nil {
		return nil, fmt.Errorf("error as non-ip target %s\n", addr)
	}

	router, err := routing.New()
	if err != nil {
		return nil, errors.Wrap(err, "error while creating routing object")
	}

	//_, gatewayIP, preferredSrc, err := router.Route(ip)
	_, _, preferredSrc, err := router.Route(ip)
	if err != nil {
		return nil, errors.Wrapf(err, "error routing to ip: %s", addr)
	}

	//fmt.Printf("\ngatewayIP: %v preferredSrc: %v\n\n", gatewayIP, preferredSrc)
	return preferredSrc, nil
}
