package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"cmpService/mcagent/repo"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"github.com/libvirt/libvirt-go"
	"strconv"
	"strings"
)

/******************************************************************************
 * Snapshot
 ******************************************************************************/
func CreateSnapshot(name, snapName, desc string) (snap *libvirt.DomainSnapshot, err error) {
	dom, err := GetDomainByName(name)
	if err != nil {
		return nil, err
	}

	snap, err = dom.CreateSnapshotXML(fmt.Sprintf(`
		<domainsnapshot>
			<name>%s</name>
			<description>%s</description>
		</domainsnapshot>
		`, snapName, desc),
		0)
	if err != nil {
		fmt.Println("snap error: ", err)
		return nil, err
	}
	snapName, _ = snap.GetName()
	fmt.Println("snap name:", snapName)
	return snap, err
}

func SafeSnapshot(name, snapName, desc string) (entry *mcmodel.McVmSnapshot, snap *libvirt.DomainSnapshot, err error) {
	dom, err := GetDomainByName(name)
	if err != nil {
		return nil, nil, err
	}

	isChangeState := false
	state, _, _ := dom.GetState()
	if state == libvirt.DOMAIN_RUNNING {
		fmt.Println("Suspend")
		dom.Suspend()
		isChangeState = true
	}

	snap, err = dom.CreateSnapshotXML(fmt.Sprintf(`
		<domainsnapshot>
			<name>%s</name>
			<description>%s</description>
		</domainsnapshot>
		`, snapName, desc),
		0)

	if isChangeState == true {
		dom.Resume()
		fmt.Println("Resume")
	}

	if err != nil {
		fmt.Println("snap error: ", err)
		return nil, nil, err
	}

	snapName, _ = snap.GetName()
	fmt.Println("snap name:", snapName)
	/*****************
	* Make Snapshot entry
	*****************/
	entry = GetSnapEntry(name, snapName, desc)
	entry.Dump()
	cfg := config.GetMcGlobalConfig()
	svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)
	svcmgrapi.SendMcVmSnapshot2Svcmgr(*entry, svcmgrRestAddr)
	return entry, snap, err
}

func GetSnapEntry(vmName, snapName, desc string) (*mcmodel.McVmSnapshot) {
	var snap mcmodel.McVmSnapshot
	snap.VmName = vmName
	snap.Name = snapName
	snap.Desc = desc
	snap.ServerSn = repo.GetMcServer().SerialNumber
	snap.CompanyIdx = repo.GetMcServer().CompanyIdx

	arr := strings.Split(snapName, "-")
	snap.Year, _ = strconv.Atoi(arr[0])
	snap.Month = GetMonthStr2Num(arr[1])
	snap.Day, _ = strconv.Atoi(arr[2])
	snap.Hour, _ = strconv.Atoi(arr[3])
	snap.Minute, _ = strconv.Atoi(arr[4])
	snap.Second, _ = strconv.Atoi(arr[5])
	snap.Current = true
	return &snap
}

func GetMonthStr2Num(month string) int {
	switch month {
	case "Jan": return 1
	case "Feb": return 2
	case "Mar": return 3
	case "Apr": return 4
	case "May": return 5
	case "Jun": return 6
	case "Jul": return 7
	case "Aug": return 8
	case "Sep": return 9
	case "Oct": return 10
	case "Nov": return 11
	case "Dec": return 12
	}
	return 0
}

func GetSnapshotsListName(name string) (snaps[]string, err error) {
	dom, err := GetDomainByName(name)
	if err != nil {
		fmt.Println("GetSnapshotListName error:", err)
		return snaps, err
	}

	snaps, err = dom.SnapshotListNames(0)
	if err != nil {
		fmt.Println("GetSnapshotListName error:", err)
		return snaps, err
	}

	for index, snap := range snaps {
		fmt.Println("index ", index, " name:", snap)
	}
	return snaps, err
}

func GetAllSnapshots(name string) (snaps []libvirt.DomainSnapshot, err error) {
	dom, err := GetDomainByName(name)
	if err != nil {
		return nil, err
	}

	snaps, err = dom.ListAllSnapshots(0)
	for index, snap := range snaps {
		name, _ := snap.GetName()
		fmt.Println("index ", index, " name:", name)
		//desc, _ := snap.GetXMLDesc(0)
		//fmt.Println("        desc: ", desc)
	}
	return snaps, err
}

func DeleteAllSnapshot(name string) (snaps []libvirt.DomainSnapshot, err error) {
	dom, err := GetDomainByName(name)
	if err != nil {
		return nil, err
	}

	snaps, err = dom.ListAllSnapshots(0)
	for index, snap := range snaps {
		name, _ := snap.GetName()
		fmt.Println("index ", index, " name:", name)
		err = snap.Delete(0)
		if err != nil {
			fmt.Println(" error:", err)
		}
	}
	return snaps, err
}

func ApplySnapshot(domName, snapName string) error {
	dom, err := GetDomainByName(domName)
	if err != nil {
		return nil
	}
	snap, err := dom.SnapshotLookupByName(snapName, 0)
	err = snap.RevertToSnapshot(0)
	return err
}