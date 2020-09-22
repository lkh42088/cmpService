package agent

import (
	"cmpService/common/mcmodel"
	"cmpService/common/package/routing"
	"cmpService/mcagent/config"
	"cmpService/mcagent/svcmgrapi"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	network "net"
	"strings"
	"syscall"
	"time"
)

//func GetSysInfoLegacy() {
//	info := syscall.Sysinfo_t{}
//	err := syscall.Sysinfo(&info)
//
//	if err == nil {
//		fmt.Printf("sysinfo : %+v\n", info)
//	}
//}

func GetEnvVar() {
	env := syscall.Environ()
	for i := range env {
		fmt.Println(env[i])
	}
}

func GetSysInfo() mcmodel.SysInfo {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")
	netIf, _ := net.Interfaces()

	info := new(mcmodel.SysInfo)
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
	info.UpdateTime = time.Now()

	pretty, _ := json.MarshalIndent(info, "", "  ")
	fmt.Printf("%s\n", string(pretty))

	return *info
}

func GetSysInfoAll() mcmodel.SysInfoDetail {
	hostStat, _ := host.Info()
	temper, _ := host.SensorsTemperatures()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")
	diskPart, _ := disk.Partitions(true)
	netIf, _ := net.Interfaces()

	info := new(mcmodel.SysInfoDetail)
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

	router, err := routing.New()	// fix to bug : reference redmine document
	if err != nil {
		fmt.Println(err)
		//return nil, errors.Wrap(err, "error while creating routing object")
	}

	_, gatewayIP, preferredSrc, err := router.Route(ip)
	//_, _, preferredSrc, err := router.Route(ip)
	if err != nil {
		return nil, errors.Wrapf(err, "error routing to ip: %s", addr)
	}

	fmt.Printf("\ngatewayIP: %v preferredSrc: %v\n\n", gatewayIP, preferredSrc)
	return preferredSrc, nil
}

func SetSysInfo() {
	conf := config.GetMcGlobalConfig()
	fmt.Println(conf)
	info := GetSysInfo()
	if info.IfMac == "" {
		return
	}

	// set to global config
	config.SetGlobalConfigWithSysInfo(info)

	config.SetTelegraf("", info.IfMac)
	config.RestartTelegraf()
}

func SendSysInfo() {
	conf := config.GetMcGlobalConfig()

	url := conf.SvcmgrIp+ ":" + conf.SvcmgrPort
	svcmgrapi.SendSysInfoToSvcmgr(conf.SystemInfo, url)
}
