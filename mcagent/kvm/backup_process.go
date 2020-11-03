package kvm

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"cmpService/mcagent/ktrest"
	"cmpService/mcagent/nas"
	"cmpService/mcagent/repo"
	"cmpService/mcagent/svcmgrapi"
	"errors"
	"fmt"
	"github.com/libvirt/libvirt-go"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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
	_, _ = cmd.Output()
	//fmt.Println("output", string(output))
}

func DeleteAllFile(path string, filenames []string) {
	if len(filenames) == 0 {
		return
	}
	fmt.Println("# FILE LIST : ", filenames)
	for i, filename := range filenames {
		DeleteFile(path + "/" + filename)
		fmt.Printf("%d 번째 파일이 삭제되었습니다.(%s)\n", i + 1, filename)
	}
	return
}

func DeleteDirectory(directoryPath string) {
	args := []string{
		"-rf",
		directoryPath,
	}

	binary := "rm"
	cmd := exec.Command(binary, args...)
	_, _ = cmd.Output()
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
	backupFile := config.GetMcGlobalConfig().VmBackupDir + "/" + backupVmName + ".qcow2"

	// Stop Vm
	dom, err := GetDomainByName(vmName)
	if err != nil {
		fmt.Printf("BackupVmImage (%s) error 0: %s", vmName, err)
		return "", 0
	}
	name, _ := dom.GetName()
	status, _, _ := dom.GetState()
	fmt.Println("dom:", name, status, backupFile)
	if status != libvirt.DOMAIN_SHUTDOWN && status != libvirt.DOMAIN_SHUTOFF {
		dom.Destroy()
	}

	// Clone Vm (long time)
	CloneVm(vmName, backupVmName, backupFile)
	fmt.Println("Finish clone....")

	// Start Vm
	err = dom.Create()
	if err != nil {
		fmt.Println("[BACKUP][Create] Error : ", err)
	}

	// Destroy/Undefine Backup Vm
	backupDom, err := GetDomainByName(backupVmName)
	if err != nil {
		fmt.Printf("BackupVmImage (%s) error 1: %s", vmName, err)
		return "", 0
	}
	err = backupDom.Undefine()
	if err != nil {
		fmt.Println("[BACKUP][Undefine] Error : ", err)
	}

	// Decrease qcow2 image size (long time)
	decreaseImage := backupFile + ".decrease"
	DecreaseQcow2Image(backupFile, decreaseImage)

	// Delete temp file
	DeleteFile(backupFile)

	// return cronsch file
	file, err := os.Stat(decreaseImage)
	fmt.Println("# Backup : ", decreaseImage, err)
	return decreaseImage, int(file.Size())
}

//func SafeBackup(name string) (entry *mcmodel.McVmSnapshot, snap *libvirt.DomainSnapshot, err error) {
func SafeBackup(vmName, backupName, desc string) {
	path := config.GetMcGlobalConfig().VmBackupDir
	/*****************
	* Make Bakcup entry
	*****************/
	backupFilePath, size := BackupVmImage(vmName)
	backupFile := strings.TrimPrefix(backupFilePath, path + "/")
	// Get container name or Create container
	ktrest.ConfigurationForKtContainer()

	/*****************
	* Upload cronsch file to KT Cloud Storage or NAS
	*****************/
	server, filenames, err := McVmBackup(vmName, backupFile, desc)
	if err != nil {
		fmt.Printf("! SafeBackup() : McVmBackup() Err - %s\n", err)
		return
	}

	/*****************
	* Make Backup message
	*****************/
	entry, svcmgrRestAddr := MakeBackupMsg(vmName, backupFile, desc, size, *server)

	_, err = repo.StoreVmBackup2Db(*entry)
	if err != nil {
		fmt.Printf("! StoreVmBackup2Db is failed.(%s)\n", err)
	} else {
		//fmt.Printf("Backup Entry : %+v\n", v)
		/*****************************
		 * Notify to svcmgr
		 *****************************/
		fmt.Println("Send to svcmgr... ", svcmgrRestAddr)
		svcmgrapi.SendMcVmBackup2Svcmgr(*entry, svcmgrRestAddr, lib.SvcmgrApiMicroMcAgentNotifyBackup)
	}

	/*****************
	 * Delete backupFile
	*****************/
	DeleteAllFile(path, filenames)

	/*****************
	 * Sync VM list
	*****************/
	vms := GetMcServerInfo().Vms
	fmt.Println("# Backup Vms : ", *vms)
	svcmgrapi.SendUpdateVmList2Svcmgr(*vms, svcmgrRestAddr)
}

func MakeBackupMsg(vmName string, backupName string, desc string, size int, server mcmodel.McServerDetail) (*mcmodel.McVmBackup, string) {
	entry := GetBackupEntry(vmName, GetTimeWord(), desc)
	entry.BackupSize = size
	if server.UcloudAccessKey != "" {
		entry.Name = backupName
		entry.KtContainerName = ktrest.GlobalContainerName
		entry.KtAuthUrl = server.UcloudAuthUrl
	} else {
		entry.NasBackupName = backupName
	}
	entry.Command = "add"
	entry.Dump()
	cfg := config.GetMcGlobalConfig()
	svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)
	return entry, svcmgrRestAddr
}

func McVmBackup(vmName string, backupFile string, command string) (*mcmodel.McServerDetail, []string, error) {
	server := repo.GetMcServer()
	var filenames []string
	var err error
	vm, _ := repo.GetVmFromDbByName(vmName)
	if command != "By action command" {
		if vm.BackupType == false {
			fmt.Println("! BackupType is false.")
			return nil, filenames, errors.New("BackupType is false.\n")
		}
	}

	if server.UcloudAccessKey != "" {
		fmt.Println("KT Storage SafeBackup: ", backupFile)
		filenames, err = ktrest.DivisionVmBackupFile(backupFile)
		if err != nil {
			fmt.Printf("\n! SafeBackup(FileDivision) Error : %s\n\n", err)
			return nil, nil, err
		}

		// delete backup decrease file
		DeleteFile(config.GetMcGlobalConfig().VmBackupDir + "/" + backupFile)

		tmp := time.Now()
		for i, filename := range filenames {
			err = ktrest.PutDynamicLargeObjects(ktrest.GlobalContainerName, backupFile, filename)
			if err != nil {
				//  오랜 시간 업로드 동작으로 인한 사용자 인증 해제 시 RETRY
				ktrest.PostAuthTokens()
				err = ktrest.PutDynamicLargeObjects(ktrest.GlobalContainerName, backupFile, filename)
				if err != nil {
					fmt.Printf("\n! SafeBackup(PutDLO) Error : %s\n\n", err)
					return nil, nil, err
				}
			}
			time.Sleep(1000)	// KT api restrictions : 1 sec term per 1 api
			fmt.Printf("%d.%s %d 초 소요되었습니다. \n", i+1, filename, int(time.Now().Sub(tmp).Seconds()))
			tmp = time.Now()
		}
		err = ktrest.PutDLOManifest(ktrest.GlobalContainerName, backupFile)
		if err != nil {
			fmt.Printf("\n! SafeBackup(PutManifest) Error : %s\n\n", err)
			return nil, nil, err
		}
	} else {
		fmt.Println("NAS SafeBackup: ", backupFile)
		src :=  config.GetMcGlobalConfig().VmBackupDir + "/" + backupFile
		dst := os.Getenv("HOME") + "/nas/backup/" + backupFile
		fmt.Println("NAS PATH: ", src, dst)
		MoveFile(src, dst)
		filenames = append(filenames, backupFile)
	}
	return server, filenames, nil
}

func GetBackupEntry(vmName, backupName, desc string) (*mcmodel.McVmBackup) {
	var backup mcmodel.McVmBackup
	backup.VmName = vmName
	backup.Desc = desc
	backup.McServerSn = repo.GetMcServer().SerialNumber
	backup.CompanyIdx = repo.GetMcServer().CompanyIdx
	backup.McServerIdx = int(repo.GetMcServer().Idx)

	arr := strings.Split(backupName, "-")
	backup.Year, _ = strconv.Atoi(arr[0])
	backup.Month = GetMonthStr2Num(arr[1])
	backup.Day, _ = strconv.Atoi(arr[2])
	backup.Hour, _ = strconv.Atoi(arr[3])
	backup.Minute, _ = strconv.Atoi(arr[4])
	backup.Second, _ = strconv.Atoi(arr[5])
	return &backup
}

// Restore Backup
func RebootingByBackupFile(src string, dst string, backup mcmodel.McVmBackup, vm mcmodel.McVm) {
	// old vm stop
	dom, err := GetDomainByName(backup.VmName)
	if err != nil {
		fmt.Printf("BackupVmImage (%s) error 0: %s", backup.VmName, err)
	} else {
		name, _ := dom.GetName()
		status, _, _ := dom.GetState()
		fmt.Println("dom: ", name, status)
		if status != libvirt.DOMAIN_SHUTDOWN && status != libvirt.DOMAIN_SHUTOFF {
			dom.Destroy()
		}
	}

	// new file move
	fmt.Println("# Movefile : ", src, dst)
	MoveFile(src, dst)

	// delete temp file & directory
	currentPath, _ := os.Getwd()
	DeleteFile(currentPath + "/" + backup.Name)
	DeleteDirectory(currentPath + "/opt")

	// new vm start
	fmt.Println("# New VM Start!")
	//CreateVmInstance(vm)
	StartVm(vm)
}

func RebootingByBackupFileWithNas(src string, dst string, backup mcmodel.McVmBackup, vm mcmodel.McVm) {
	// old vm stop
	dom, err := GetDomainByName(backup.VmName)
	if err != nil {
		fmt.Printf("BackupVmImage (%s) error 0: %s", backup.VmName, err)
	} else {
		name, _ := dom.GetName()
		status, _, _ := dom.GetState()
		fmt.Println("dom: ", name, status)
		if status != libvirt.DOMAIN_SHUTDOWN && status != libvirt.DOMAIN_SHUTOFF {
			dom.Destroy()
		}
	}

	// new file move
	fmt.Println("# Movefile : ", src, dst)
	MoveFile(src, dst)

	// new vm start
	fmt.Println("# New VM Start!")
	//CreateVmInstance(vm)
	StartVm(vm)
}

func MoveFile(src string, dst string) {
	args := []string{
		src,
		dst,
	}

	binary := "mv"
	cmd := exec.Command(binary, args...)
	_, _ = cmd.Output()
}

func CheckBackup() {
	// Check Backup Directory
	cfg := config.GetMcGlobalConfig()
	if _, err := os.Stat(cfg.VmBackupDir); err != nil {
		err = os.MkdirAll(cfg.VmBackupDir, 0755)
		if err != nil {
			return
		}
	}
	// Check KT Storage Account
	err := ktrest.CheckKtAccount()
	if err != nil {
		fmt.Println("** CheckKtAccount Error : ", err)
	}
	// Backup configuration : KT Storage
	if ktrest.ConfigurationForKtContainer() != nil {
		fmt.Printf("\n** KT Storage configuration is invalid.\n\n")
	}
	// Check NAS Info
	CheckNasInfo()
}

func CheckNasInfo() {
	// Check Nas DB info
	server := repo.GetMcServer()
	if server.NasUrl == "" {
		fmt.Println("# NAS path info is empty.")
		return
	}
	nas.MountNasDirectory(server.NasUrl)
	if _, err := os.Stat(os.Getenv("HOME") + "/" + "nas/backup"); err != nil {
		fmt.Println("!! NAS directory is not mounted.\n")
	}
}

