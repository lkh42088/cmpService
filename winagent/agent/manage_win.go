package agent

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
	"io"
	"os"
	"os/exec"
	"strings"
)

//type NetInterface struct {
//	Name    string   `json:name`
//	Address []string `json:address`
//}
//
//type SysInfo struct {
//	Idx             int    `gorm:"primary_key;column:idx;not null;auto_increment;comment:'INDEX'" json:"idx"`
//	Hostname        string `gorm:"column:hostname;not null;comment:'SERVER 이름'" json:"hostname"`
//	OS              string `gorm:"column:os;comment:'OS 명'" json:"os"`
//	Uptime          uint64 `gorm:"column:uptime;comment:'UPTIME'" json:"uptime"`
//	BootTime        uint64 `gorm:"column:boottime;comment:'SERVER BOOTTIME'" json:"bootTime"`
//	CpuCore         int    `gorm:"column:cpu_core;comment:'CPU Core 개수'" json:"cpuCore"`
//	CpuModel        string `gorm:"column:cpu_model;comment:'CPU 모델명'" json:"cpuModel"`
//	Platform        string `gorm:"column:platform;comment:'Platform'" json:"platform"`
//	PlatformVersion string `gorm:"column:platform_version;comment:'Platform 버전'" json:"platformVersion"`
//	KernelArch      string `gorm:"column:kernel_arch;comment:'KERNEL 아키텍처'" json:"kernelArch"`
//	KernelVersion   string `gorm:"column:kernel_version;comment:'KERNEL 버전'" json:"kernelVersion"`
//	IP              string `gorm:"column:ip;not null;comment:'SERVER IP'" json:"ip"`
//	IfName          string `gorm:"column:if_name;comment:'Interface Name'" json:"ifName"`
//	IfMac           string `gorm:"column:if_mac;comment:'Interface MAC'" json:"ifMac"`
//	MemTotal        int64  `gorm:"column:mem_total;comment:'MEMORY 용량'" json:"mem"`
//	DiskTotal       int64  `gorm:"column:disk_total;comment:'HDD DISK 용량'" json:"disk"`
//}
//
//type SysInfoDetail struct {
//	Host        host.InfoStat          `json:hostname`
//	Temperature []host.TemperatureStat `json:temperature`
//	Platform    string                 `json:platform`
//	CPU         []cpu.InfoStat         `json:cpu`
//	RAM         uint64                 `json:ram`
//	DiskUsage   uint64                 `json:diskUsage`
//	DiskPart    []disk.PartitionStat   `json:disk`
//	Net         []net.InterfaceStat    `json:interface`
//}

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
	globalSysInfo = *info

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
	return globalSysInfo
}

func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func InsertMacInTelegrafConf(mac string) bool {
	orgin_file := "c:\\Program files\\Telegraf\\telegraf.conf"
	fd, err := os.Open(orgin_file)
	if err != nil {
		fmt.Println("InsertMacInTelegrafConf: error", err)
		return false
	}
	defer fd.Close()

	backup_file := orgin_file +".cron"
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

func RestartTelegraf () {
	args := []string{
		"stop",
		"telegraf",
	}
	args2 := []string{
		"start",
		"telegraf",
	}

	cmd := exec.Command("net", args...)
	cmd1 := exec.Command("net", args2...)
	out, _ := cmd.Output()
	fmt.Println(out)
	out, _ = cmd1.Output()
	fmt.Println(out)
}


//func commandWindowsApp(command string, path string) {
//	verb := "runas"
//	exe, _ := os.Executable()
//	cwd, _ := os.Getwd()
//	args := strings.Join(os.Args[1:], " ")
//
//	verbPtr, _ := syscall.UTF16PtrFromString(verb)
//	exePtr, _ := syscall.UTF16PtrFromString(exe)
//	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
//	argPtr, _ := syscall.UTF16PtrFromString(args)
//
//	var showCmd int32 = windows.SW_HIDE
//
//	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
//	if err != nil {
//		fmt.Println(err)
//	}
//}