package snmpapi

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"nubes/collector/db/influx"
	"nubes/collector/device"
)

type ID string

type SnmpDevice struct {
	Device device.Device
	Snmp      *g.GoSNMP

	Cpu       Cpu
	Memory    Memory
	System    System
	IfTable   IfTable
	IpTable   IpTable
	l4table   L4Table
}

//func NewSnmpDevice(id device.CodeID, addr string, community string) *SnmpDevice {
func NewSnmpDevice(device device.Device) *SnmpDevice {
	return &SnmpDevice{
		Device:    device,
		Snmp:      nil,
		Cpu:       Cpu{},
		Memory:    Memory{},
		System:    System{},
		IfTable:   IfTable{},
		IpTable:   IpTable{},
		l4table:   L4Table{},
	}
}

type SnmpDeviceTable struct {
	devices map[device.ID]SnmpDevice
	count   int
	store   influx.Config
}

//var ErrDeviceNotExist = errors.New("device does not exist")

func (s *SnmpDevice) String() {
	output :=fmt.Sprintf("Device %s", s.Device)
	n := len(output)
	for i:=0 ; i < n; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
	fmt.Println(output)
	s.System.String()
	s.Cpu.String()
	s.Memory.String()
	s.l4table.String()
	s.IfTable.String()
	s.IpTable.String()
}

func (s *SnmpDeviceTable) String() {
	for _, d := range s.devices {
		d.String()
	}
	fmt.Println("Total:", len(s.devices), s.count)
}

func NewSnmpDeviceTable() *SnmpDeviceTable {
	return &SnmpDeviceTable {
		map[device.ID]SnmpDevice{},
		0,
		influx.Config{},
	}
}

func (sd *SnmpDeviceTable) Get(id device.ID) (*SnmpDevice, error) {
	d, exists := sd.devices[id]
	if !exists {
		return &SnmpDevice{}, device.ErrDeviceNotExist
	}
	return &d, nil
}

func (sd *SnmpDeviceTable) Put(id device.ID, d SnmpDevice) error {
	if _, exists := sd.devices[id]; !exists {
		return device.ErrDeviceNotExist
	}
	sd.devices[id] = d
	return nil
}

func (sd *SnmpDeviceTable) Post(d SnmpDevice) (device.ID, error) {
	sd.count++
	fmt.Printf("dev id : %s", d.Device.GetIdString())
	sd.devices[d.Device.GetIdString()] = d
	return d.Device.GetIdString(), nil
}

func (sd *SnmpDeviceTable) Delete(id device.ID) error {
	if _, exists := sd.devices[id]; !exists {
		return device.ErrDeviceNotExist
	}
	delete(sd.devices, id)
	sd.count--
	return nil
}


