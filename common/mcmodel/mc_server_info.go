package mcmodel

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
	"time"
)

type NetInterface struct {
	Name    string   `json:name`
	Address []string `json:address`
}

type SysInfo struct {
	Idx             int    `gorm:"primary_key;column:idx;not null;unsigned;auto_increment;comment:'INDEX'" json:"idx"`
	Hostname        string `gorm:"type:varchar(50);column:hostname;not null;comment:'SYSTEM 이름'" json:"hostname"`
	OS              string `gorm:"type:varchar(50);column:os;comment:'OS 명'" json:"os"`
	Uptime          uint64 `gorm:"column:uptime;not null;unsigned;comment:'UPTIME'" json:"uptime"`
	BootTime        uint64 `gorm:"column:boottime;not null;unsigned;comment:'SYSTEM BOOT TIME'" json:"bootTime"`
	CpuCore         int    `gorm:"column:cpu_core;not null;unsigned;comment:'CPU Core 개수'" json:"cpuCore"`
	CpuModel        string `gorm:"type:varchar(50);column:cpu_model;comment:'CPU 모델명'" json:"cpuModel"`
	Platform        string `gorm:"type:varchar(50);column:platform;comment:'Platform'" json:"platform"`
	PlatformVersion string `gorm:"type:varchar(50);column:platform_version;comment:'Platform 버전'" json:"platformVersion"`
	KernelArch      string `gorm:"type:varchar(50);column:kernel_arch;comment:'KERNEL 아키텍처'" json:"kernelArch"`
	KernelVersion   string `gorm:"type:varchar(50);column:kernel_version;comment:'KERNEL 버전'" json:"kernelVersion"`
	IP              string `gorm:"type:varchar(15);column:ip;not null;comment:'SYSTEM IP'" json:"ip"`
	IfName          string `gorm:"type:varchar(50);column:if_name;not null;comment:'Interface Name'" json:"ifName"`
	IfMac           string `gorm:"type:varchar(50);column:if_mac;not null;comment:'Interface MAC'" json:"ifMac"`
	MemTotal        int64  `gorm:"column:mem_total;not null;unsigned;comment:'MEMORY 용량'" json:"mem"`
	DiskTotal       int64  `gorm:"column:disk_total;not null;unsigned;comment:'HDD DISK 용량'" json:"disk"`
	UpdateTime		time.Time  `gorm:"column:update_time;not null;comment:'UPDATE TIME'" json:"updateTime"`
}

func (s *SysInfo) Dump() string {
	pretty, _ := json.MarshalIndent(s, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func (SysInfo) TableName() string {
	return "sysinfo_tb"
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



