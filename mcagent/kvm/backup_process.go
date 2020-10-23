package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"cmpService/mcagent/ktrest"
	"cmpService/mcagent/repo"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"github.com/libvirt/libvirt-go"
	"os"
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

func BackupVmImage(vmName string) (string, int) {
	backupVmName := vmName + "-cronsch"
	backupFile := "/opt/vm_instances/" + backupVmName + ".qcow2"

	// Stop Vm
	dom, err := GetDomainByName(vmName)
	if err != nil {
		fmt.Printf("BackupVmImage (%s) error 0: %s", vmName, err)
		return "", 0
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
		return "", 0
	}

	backupDom.Undefine()

	// Decrease qcow2 image size (long time)
	decreaseImage := backupFile + ".decrease"
	DecreaseQcow2Image(backupFile, decreaseImage)

	// Delete temp file
	DeleteFile(backupFile)
	file, err := os.Stat(decreaseImage)

	// return cronsch file
	return decreaseImage, int(file.Size())
}

//func SafeBackup(name string) (entry *mcmodel.McVmSnapshot, snap *libvirt.DomainSnapshot, err error) {
func SafeBackup(name, backupName, desc string) {
	/*****************
	* Make Bakcup entry
	*****************/
	backupFile, size := BackupVmImage(name)
	// Get container name or Create container
	ktrest.ConfigurationForKtContainer()

	/*****************
	* Upload cronsch file to KT Cloud Storage or NAS
	*****************/
	server := repo.GetMcServer()
	vm, _ := repo.GetVmFromDbByName(name)
	if vm.BackupType == false {
		return
	}
	if server.UcloudAccessKey != "" {
		fmt.Println("SafeBackup:", backupFile)
		filenames, err := ktrest.DivisionVmBackupFile(backupFile)
		if err != nil {
			fmt.Printf("\n! SafeBackup(FileDivision) Error : %s\n\n", err)
			return
		}
		for i, filename := range filenames {
			err = ktrest.PutDynamicLargeObjects(ktrest.GlobalContainerName, backupFile, filename)
			if err != nil {
				fmt.Printf("\n! SafeBackup(PutDLO) Error : %s\n\n", err)
				return
			}
			fmt.Printf(" KT Storage backup file upload %d. %s\n", i+1, filename)
		}
		err = ktrest.PutDLOManifest(ktrest.GlobalContainerName, backupFile)
		if err != nil {
			fmt.Printf("\n! SafeBackup(PutManifest) Error : %s\n\n", err)
			return
		}
	} else {
		// NAS backup
	}

	/* Next, khlee delete backupFile */

	/*****************
	* Make Backup message
	*****************/
	entry := GetBackupEntry(name, backupName, desc)
	entry.BackupSize = size
	if server.UcloudAccessKey != "" {
		entry.Name = backupName
		entry.KtContainerName = ktrest.GlobalContainerName
	} else {
		entry.NasBackupName = backupName
	}
	entry.Command = "add"
	entry.Dump()
	cfg := config.GetMcGlobalConfig()
	svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)

	/*****************************
	 * Notify to svcmgr
	 *****************************/
	fmt.Println("Send to svcmgr... ", svcmgrRestAddr)
	svcmgrapi.SendMcVmBackup2Svcmgr(*entry, svcmgrRestAddr)
}

func GetBackupEntry(vmName, backupName, desc string) (*mcmodel.McVmBackup) {
	var backup mcmodel.McVmBackup
	backup.VmName = vmName
	backup.Desc = desc
	backup.ServerSn = repo.GetMcServer().SerialNumber
	backup.CompanyIdx = repo.GetMcServer().CompanyIdx

	//backup.NasBackupName = backupName
	//backup.KtContainerName = ktrest.GlobalContainerName
	//backup.Name = backupName
	//backup.BackupSize = size

	arr := strings.Split(backupName, "-")
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
