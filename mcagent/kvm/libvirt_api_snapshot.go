package kvm

import (
	"fmt"
	"github.com/libvirt/libvirt-go"
)

/******************************************************************************
 * Snapshot
 ******************************************************************************/
func CreateSnapshot(name, desc string) (snap *libvirt.DomainSnapshot, err error) {
	dom, err := GetDomainByName(name)
	if err != nil {
		return nil, err
	}

	snap, err = dom.CreateSnapshotXML(fmt.Sprintf(`
		<domainsnapshot>
			<description>%s</description>
		</domainsnapshot>
		`, desc),
		0)
	snapName, _ := snap.GetName()
	fmt.Println("snap name:", snapName)
	return snap, err
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

func ApplySnapshot(domName, snapName string) error {
	dom, err := GetDomainByName(domName)
	if err != nil {
		return nil
	}
	snap, err := dom.SnapshotLookupByName(snapName, 0)
	err = snap.RevertToSnapshot(0)
	return err
}