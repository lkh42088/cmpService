package kvm

import (
	"bufio"
	"bytes"
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"log"
	"net"
	"os/exec"
	"strings"
	"time"
)

func CreateXrdpNat(vm mcmodel.MgoVm) {
	// iptables -t nat -A PREROUTING -d 192.168.0.57 -p tcp --dport 10022 -j DNAT --to 10.0.0.197:22
	// iptables-save
}

func CreateNetwork(vm mcmodel.MgoVm) {
	// create virbr1.xml file
	// virsh define virbr1.xml
	// mapping a virtual bridge per physical interface
	args := []string{
		"start",
		vm.Name,
	}

	fmt.Println("args: ", args)

	binary := "virsh"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output", string(output))
}

func GetIpAddressOfVm(vm mcmodel.MgoVm) (ip, mac string, res int) {
	res = 0

	// virsh domifaddr NAME
	args := []string{
		"domifaddr",
		vm.Name,
	}

	fmt.Println("args: ", args)
	binary := "virsh"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	//fmt.Println("output", string(output))

	scanner := bufio.NewScanner(bytes.NewReader(output))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println("line num: ", len(lines))

	if len(lines) < 4 {
		return ip, mac, -1
	}

	var tags []string
	line := strings.TrimSpace(lines[2])
	tags = strings.Split(line, " ")

	var arr[]string
	for _, tag := range tags {
		if len(tag) > 0 {
			arr = append(arr, tag)
		}
	}

	for index, ar := range arr {
		fmt.Printf("%d: (%s), %d\n", index, ar, len(ar))
	}
	ip = arr[3]
	mac = arr[1]
	return ip, mac, res
}

func MakeFilename(vm *mcmodel.MgoVm) string {
	cfg := config.GetGlobalConfig()
	for index, num := range cfg.VmNumber {
		if num == 0 {
			config.SetGlobalConfigByVmNumber(uint(index), vm.Idx)
			//cfg.VmNumber[index] = vm.Idx
			vm.VmNumber = index
			return fmt.Sprintf("%s-%d", vm.Image, index)
		}
	}
	return ""
}

func DeleteFilename(vm mcmodel.MgoVm) {
	cfg := config.GetGlobalConfig()
	cfg.VmNumber[vm.VmNumber] = 0
}

func ConfigDNAT(vm *mcmodel.MgoVm) {
	cfg := config.GetGlobalConfig()
	//iptables -t nat -A PREROUTING -d 192.168.0.73 -p tcp --dport 13389 -j DNAT --to 10.0.0.159:3389
	dport:= fmt.Sprintf("%d", 10000+vm.VmNumber)
	ip := strings.Split(vm.IpAddr,"/")
	target := fmt.Sprintf("%s:3389", ip[0])
	vm.RemoteAddr = fmt.Sprintf("%s:%s", cfg.ServerIp, dport)
	args := []string{
		"-t",
		"nat",
		"-A",
		"PREROUTING",
		"-d",
		cfg.ServerIp,
		"-p",
		"tcp",
		"--dport",
		dport,
		"-j",
		"DNAT",
		"--to",
		target,
	}

	fmt.Println("args: ", args)

	binary := "iptables"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output", string(output))
}

func CopyVmInstance(vm *mcmodel.MgoVm) {
	cfg := config.GetGlobalConfig()
	org := fmt.Sprintf("%s/%s.qcow2", cfg.VmImageDir, vm.Image)
	target := fmt.Sprintf("%s/%s.qcow2", cfg.VmInstanceDir, vm.Filename)
	args := []string{
		org,
		target,
	}

	fmt.Println("args: ", args)

	binary := "cp"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output", string(output))
}

func CreateVmInstance(vm mcmodel.MgoVm) {
	cfg := config.GetGlobalConfig()
	if vm.Filename == "" {
		fmt.Printf("CreateVmInstance: %s failed to get filename!\n", vm.Name)
		return
	}
	diskPath := fmt.Sprintf("path=%s/%s.qcow2,format=qcow2,bus=virtio", cfg.VmInstanceDir,vm.Filename)
	RamStr := fmt.Sprintf("%d", vm.Ram)
	cpuStr := fmt.Sprintf("--vcpus=%d", vm.Cpu)
	netStr := fmt.Sprintf("network=%s,model=virtio", vm.Network)
	args := []string{
		"--connect=qemu:///system",
		"--virt-type",
		"kvm",
		"--name",
		vm.Name,
		"--memory",
		RamStr,
		cpuStr,
		"--cpu",
		"host-passthrough",
		"--os-type",
		"win10",
		"--import",
		"--disk",
		diskPath,
		"--network",
		//"network=default,model=virtio",
		netStr,
		"--noautoconsole",
	}

	fmt.Println("args:", args)

	binary := "virt-install"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output", string(output))
}

func StartVm(vm mcmodel.MgoVm) {
	// virsh shutdown NAME
	args := []string{
		"start",
		vm.Name,
	}

	fmt.Println("args: ", args)

	binary := "virsh"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output", string(output))
}

func ShutdownVm(vm mcmodel.MgoVm) {
	// virsh shutdown NAME
	args := []string{
		"shutdown",
		vm.Name,
	}

	fmt.Println("args: ", args)

	binary := "virsh"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output", string(output))
}

func UndefineVm(vm mcmodel.MgoVm) {
	// virsh undefine NAME
	args := []string{
		"undefine",
		vm.Name,
	}
	fmt.Println("args: ", args)
	binary := "virsh"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output", string(output))
}

func DestroyVm(vm mcmodel.MgoVm) {
	// virsh undefine NAME
	args := []string{
		"destroy",
		vm.Name,
	}
	fmt.Println("args: ", args)
	binary := "virsh"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output", string(output))
}

func StatusVm(vm mcmodel.MgoVm) string {
	// virsh undefine NAME
	args := []string{
		"list",
		"--all",
	}
	fmt.Println("args: ", args)
	binary := "virsh"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	//fmt.Println("output", string(output))

	scanner := bufio.NewScanner(bytes.NewReader(output))
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Contains(line, vm.Name) {
			fmt.Println("fine: ", line)
			break
		}
	}
	if len(line) < 1 {
		return ""
	}

	var tags []string
	line = strings.TrimSpace(line)
	tags = strings.Split(line, " ")

	var arr[]string
	for _, tag := range tags {
		if len(tag) > 0 {
			arr = append(arr, tag)
		}
	}

	//for index, ar := range arr {
	//	fmt.Printf("%d: (%s), %d\n", index, ar, len(ar))
	//}

	//fmt.Printf("status: ", arr[2])
	return arr[2]
}

func DeleteVm(vm mcmodel.MgoVm) {
	DestroyVm(vm)
	UndefineVm(vm)
}

func DeleteVmInstance(vm mcmodel.MgoVm) {
	DeleteFilename(vm)
	cfg := config.GetGlobalConfig()
	args := []string{
		cfg.VmInstanceDir+"/"+vm.Filename+".qcow2",
		"-f",
	}

	fmt.Println("args: ", args)

	binary := "rm"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output", string(output))
}

//func DeleteVmInstance(vm *mcmodel.McVm) {
//	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
//	if err != nil {
//		log.Fatalf("failed to dial libvirt: %v", err)
//	}
//
//	l := libvirt.New(c)
//	if err := l.Connect(); err != nil {
//		log.Fatalf("failed to connect: %v", err)
//	}
//
//	v, err := l.Version()
//	//v, err := l.ConnectGetLibVersion()
//	if err != nil {
//		log.Fatalf("failed to retrieve libvirt version: %v", err)
//	}
//	fmt.Println("Version:", v)
//
//	domains, err := l.Domains()
//	//domains, err, _ := l.ConnectListAllDomains()
//	if err != nil {
//		log.Fatalf("failed to retrieve domains: %v", err)
//	}
//
//}


//func CreateVmByLibvirt() {
//	var drive uint
//	drive = 0
//	conn, err := libvirt.NewConnect("qemu:///system")
//	if err != nil {
//		log.Fatalf("failed to connect to qemu")
//	}
//	defer conn.Close()
//
//	domcfg := &libvirtxml.Domain{
//		Type: "kvm", Name: "demo", Memory:
//		&libvirtxml.DomainMemory{Value: 4096, Unit: "MB", DumpCore: "on"},
//		VCPU: &libvirtxml.DomainVCPU{Value: 1},
//		CPU:  &libvirtxml.DomainCPU{Mode: "host-model"},
//		Devices: &libvirtxml.DomainDeviceList{
//			Disks: []libvirtxml.DomainDisk{
//				{
//					Source: &libvirtxml.DomainDiskSource{File:
//					&libvirtxml.DomainDiskSourceFile{File: "/home/dude/projects/go/src/gitlab.com/driftavalii/gkvmlibvirt/vm.qcow2"}},
//						Target:  &libvirtxml.DomainDiskTarget{Dev: "hda",
//						Bus: "ide"},
//						Alias:   &libvirtxml.DomainAlias{Name: "ide0-0-0"},
//						Address: &libvirtxml.DomainAddress{Drive:
//						&libvirtxml.DomainAddressDrive{Controller: &drive, Bus: &drive, Target:
//						&drive, Unit: &drive}}}}}},
//			xml, _ := domcfg.Marshal()
//			domain, _ := conn.DomainDefineXML(xml)
//			createDomain, _ := conn.DomainCreateXML(xml, 0)
//
//			fmt.Println(xml)
//			fmt.Println(domain)
//			fmt.Println(createDomain)
//		},
//	}
//}

func GetVmFromLibvirt() {
	// This dials libvirt on the local machine, but you can substitute the first
	// two parameters with "tcp", "<ip address>:<port>" to connect to libvirt on
	// a remote machine.
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	v, err := l.Version()
	//v, err := l.ConnectGetLibVersion()
	if err != nil {
		log.Fatalf("failed to retrieve libvirt version: %v", err)
	}
	fmt.Println("Version:", v)

	domains, err := l.Domains()
	//domains, err, _ := l.ConnectListAllDomains()
	if err != nil {
		log.Fatalf("failed to retrieve domains: %v", err)
	}

	fmt.Println("ID\tName\t\tUUID")
	fmt.Printf("--------------------------------------------------------\n")
	for _, d := range domains {
		fmt.Printf("%d\t%s\t%x\n", d.ID, d.Name, d.UUID)
	}

	if err := l.Disconnect(); err != nil {
		log.Fatalf("failed to disconnect: %v", err)
	}
}