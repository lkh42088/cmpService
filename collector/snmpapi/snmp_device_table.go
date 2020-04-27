package snmpapi

import (
	"fmt"
	"nubes/collector/collectdevice"
	"nubes/common/lib"
)

type SnmpDeviceTable struct {
	Devices map[collectdevice.ID]SnmpDevice
	Count   int
}

var SnmpDevTb *SnmpDeviceTable

func (s *SnmpDeviceTable) String() {
	for _, d := range s.Devices {
		d.String()
	}
	fmt.Println("Total:", len(s.Devices), s.Count)
}

func NewSnmpDeviceTable() *SnmpDeviceTable {
	return &SnmpDeviceTable{
		map[collectdevice.ID]SnmpDevice{},
		0,
	}
}

func SetSnmpDevTb(t *SnmpDeviceTable) {
	SnmpDevTb = t
	lib.LogWarnln("SetSnmpDevTb:", *SnmpDevTb)
}

func (sd *SnmpDeviceTable) Get(id collectdevice.ID) (*SnmpDevice, error) {
	d, exists := sd.Devices[id]
	if !exists {
		return &SnmpDevice{}, collectdevice.ErrDeviceNotExist
	}
	return &d, nil
}

func (sd *SnmpDeviceTable) Put(id collectdevice.ID, d SnmpDevice) error {
	if _, exists := sd.Devices[id]; !exists {
		return collectdevice.ErrDeviceNotExist
	}
	sd.Devices[id] = d
	return nil
}

func (sd *SnmpDeviceTable) Post(d SnmpDevice) (collectdevice.ID, error) {
	sd.Count++
	lib.Debug("dev id : %s\n", d.Device.GetIdString())
	sd.Devices[d.Device.GetIdString()] = d
	return d.Device.GetIdString(), nil
}

func (sd *SnmpDeviceTable) Delete(id collectdevice.ID) error {
	if _, exists := sd.Devices[id]; !exists {
		return collectdevice.ErrDeviceNotExist
	}
	delete(sd.Devices, id)
	sd.Count--
	return nil
}

func (sd *SnmpDeviceTable) DeleteAll() error {
	return nil
}
