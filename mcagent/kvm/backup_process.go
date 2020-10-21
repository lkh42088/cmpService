package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"cmpService/mcagent/repo"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"github.com/libvirt/libvirt-go"
	"os/exec"
	"strconv"
	"strings"
)

/*************************************************************************
 * Backup
 *************************************************************************/
func CloneVm(vmName, backupVmName, backupFile string) {
	//"virt-clone --connect qemu:///system  --original vm1 --name vm1-clone --file  /vm-images/vm1-clone.img"
	args := []string{
		"--connect",
		"qemu:///system",
		"--original",
		vmName,
		"--name",
		backupVmName,
		"--file",
		backupFile,
	}

	binary := "virt-clone"
	cmd := exec.Command(binary, args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("output error : ", err)
		fmt.Println("output", string(output))
	} else {
		fmt.Println("output", string(output))
	}
}

func DeleteFile(fileName string) {
	args := []string{
		"-f",
		fileName,
	}

	binary := "rm"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output", string(output))
}

func DecreaseQcow2Image(image, decreaseImage string) {
	//"qemu-img convert -c -O qcow2 backup2.qcow2 backup2_zero.qcow2"
	args := []string{
		"convert",
		"-c",
		"-O",
		"qcow2",
		image,
		decreaseImage,
	}

	binary := "qemu-img"
	cmd := exec.Command(binary, args...)
	output, _ := cmd.Output()
	fmt.Println("output", string(output))
}

func BackupVmImage(vmName string) string {
	backupVmName := vmName + "-cronsch"
	backupFile := "/opt/vm_instances/" + backupVmName + ".qcow2"

	// Stop Vm
	dom, err := GetDomainByName(vmName)
	if err != nil {
		fmt.Printf("BackupVmImage (%s) error 0: %s", vmName, err)
		return ""
	}
	name, _ := dom.GetName()
	status, _, _ := dom.GetState()
	fmt.Println("dom:", name, status)
	if status != libvirt.DOMAIN_SHUTDOWN && status != libvirt.DOMAIN_SHUTOFF {
		dom.Destroy()
	}

	// Clone Vm (long time)
	CloneVm(vmName, backupVmName, backupFile)
	fmt.Println("Finish clone....")

	// Start Vm
	dom.Create()

	// Destroy/Undefine Backup Vm
	backupDom, err := GetDomainByName(backupVmName)
	if err != nil {
		fmt.Printf("BackupVmImage (%s) error 1: %s", vmName, err)
		return ""
	}

	backupDom.Undefine()

	// Decrease qcow2 image size (long time)
	decreaseImage := backupFile + ".decrease"
	DecreaseQcow2Image(backupFile, decreaseImage)

	// Delete temp file
	DeleteFile(backupFile)

	// return cronsch file
	return decreaseImage
}

//func SafeBackup(name string) (entry *mcmodel.McVmSnapshot, snap *libvirt.DomainSnapshot, err error) {
func SafeBackup(name, snapName, desc string) {
	/*****************
	* Make Snapshot entry
	*****************/
	backupFile := BackupVmImage(name)

	/*****************
	* Upload cronsch file to KT Cloud Storage or NAS
	*****************/
	fmt.Println("SafeBackup:", backupFile)

	/*****************
	* Make Snapshot entry
	*****************/
	entry := GetBackupEntry(name, snapName, desc)
	entry.Command = "add"
	entry.Dump()
	cfg := config.GetMcGlobalConfig()
	svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)

	/*****************************
	 * Notify to svcmgr
	 *****************************/
	svcmgrapi.SendMcVmBackup2Svcmgr(*entry, svcmgrRestAddr)
}

func GetBackupEntry(vmName, snapName, desc string) (*mcmodel.McVmBackup) {
	var backup mcmodel.McVmBackup
	backup.VmName = vmName
	backup.Name = snapName
	backup.Desc = desc
	backup.ServerSn = repo.GetMcServer().SerialNumber
	backup.CompanyIdx = repo.GetMcServer().CompanyIdx

	arr := strings.Split(snapName, "-")
	backup.Year, _ = strconv.Atoi(arr[0])
	backup.Month = GetMonthStr2Num(arr[1])
	backup.Day, _ = strconv.Atoi(arr[2])
	backup.Hour, _ = strconv.Atoi(arr[3])
	backup.Minute, _ = strconv.Atoi(arr[4])
	backup.Second, _ = strconv.Atoi(arr[5])
	return &backup
}

func RecoveryBackup(vmName, backupImage string) {
	// vm stop
	dom, err := GetDomainByName(vmName)
	if err != nil {
		fmt.Printf("BackupVmImage (%s) error 0: %s", vmName, err)
		return
	}
	name, _ := dom.GetName()
	status, _, _ := dom.GetState()
	fmt.Println("dom:", name, status)
	if status != libvirt.DOMAIN_SHUTDOWN && status != libvirt.DOMAIN_SHUTOFF {
		dom.Destroy()
	}

	// delete vm snapshot
	DeleteAllSnapshot(vmName)
	//DeleteVm(vmName)
	// download cronsch file
	// change qcow2 file
}
