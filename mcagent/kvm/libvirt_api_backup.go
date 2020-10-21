package kvm

import (
	"fmt"
)

/*************************************************************************
 * Backup
 *************************************************************************/
func CloneVm() {

}

func DecreaseQcow2Image() {

}

func BackupVmImage(vmName string) {
	// Stop Vm
	dom, err := GetDomainByName(vmName)
	if err != nil {
		return
	}
	name, _ := dom.GetName()
	status, _, _ := dom.GetState()
	fmt.Println("dom:", name, status)
	//if status != libvirt.DOMAIN_SHUTDOWN && status != libvirt.DOMAIN_SHUTOFF {
	//	dom.Destroy()
	//}

	// Clone Vm (long time)
	//disk :=""
	//backupXml, err := dom.BackupGetXMLDesc(0)
	//fmt.Println("res:", backupXml)

	//dom.BlockCopy(disk)

	// Destroy Backup Vm

	// Undefine Backup Vm

	// Decrease qcow2 image size (long time)

	// Start Vm
}
