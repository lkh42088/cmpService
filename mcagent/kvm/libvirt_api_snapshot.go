package kvm

import (
	"fmt"
	"github.com/libvirt/libvirt-go"
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

func SafeSnapshot(name, snapName, desc string) (snap *libvirt.DomainSnapshot, err error) {
	dom, err := GetDomainByName(name)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	snapName, _ = snap.GetName()
	fmt.Println("snap name:", snapName)
	return snap, err
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
		desc, _ := snap.GetXMLDesc(0)
		fmt.Println("        desc: ", desc)
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