package common

import (
	"bufio"
	"cmpService/common/mcmodel"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"os"
	"os/exec"
	"strings"
)

var GlobalSysInfo mcmodel.SysInfo

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
	for _, v := range netIf {
		if v.Name == "Ethernet" || v.Name == "이더넷" {
			info.IfName = v.Name
			info.IfMac = v.HardwareAddr
			info.IP = v.Addrs[0].Addr
		}
	}
	info.MemTotal = int64(vmStat.Total)
	info.DiskTotal = int64(diskStat.Total)

	pretty, _ := json.MarshalIndent(info, "", "  ")
	fmt.Printf("%s\n", string(pretty))
	GlobalSysInfo = *info

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

func CheckMySystem() bool {
	data := GetSysInfo()
	if data.IfMac == "" {
		return false
	}
	return true
}

func GetGlbalSysInfo() mcmodel.SysInfo {
	return GlobalSysInfo
}

func InsertMacInTelegrafConf(mac string) bool {
	orgin_file := "c:\\Program files\\Telegraf\\telegraf.conf"
	fd, err := os.Open(orgin_file)
	if err != nil {
		fmt.Println("InsertMacInTelegrafConf: error", err)
		return false
	}
	defer fd.Close()

	backup_file := orgin_file +".cronsch"
	backup_fd, err := os.Create(backup_file)
	if err != nil {
		fmt.Println("InsertMacInTelegrafConf: error", err)
		return false
	}
	defer backup_fd.Close()

	w := bufio.NewWriter(backup_fd)
	if err != nil {
		fmt.Println("InsertMacInTelegrafConf: error", err)
		return false
	}

	isFind := false
	global_tags_area := false
	findStr := "mac_address"
	reader := bufio.NewReader(fd)
	for {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix  || err != nil {
			fmt.Println(isPrefix, "error", err)
			break
		}
		lineStr := string(line)
		if isFind == true {
			w.WriteString(lineStr+"\n")
		} else if strings.Contains(lineStr, findStr) == true {
			w.WriteString(fmt.Sprintf("  %s = \"%s\"\n", findStr, mac))
			isFind = true
		} else if strings.Contains(lineStr, "[global_tags]") == true {
			w.WriteString(lineStr+"\n")
			global_tags_area = true
		} else if isFind == false && global_tags_area == true &&
			len(strings.Trim(lineStr, " ")) == 0 {
			w.WriteString(fmt.Sprintf("  %s = \"%s\"\n", findStr, mac))
			w.WriteString("\n")
			isFind = true
		} else {
			w.WriteString(lineStr+"\n")
		}
	}
	w.Flush()
	backup_fd.Sync()

	err = CopyFile(backup_file, orgin_file)
	if err != nil {
		return false
	}

	return true
}

func addFireWallRule(names, appnames, dirs, actions string) error {
	c := exec.Command("netsh", "advfirewall", "firewall", "add", "rule",
			"name=" + names,
			"dir=" + dirs,
			"action=" + actions,
			"program=" + appnames,
		)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
