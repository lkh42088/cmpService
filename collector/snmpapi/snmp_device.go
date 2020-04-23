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
	Snmp   *g.GoSNMP

	Cpu        Cpu
	Memory     Memory
	System     System
	IfTable    IfTable
	IpTable    IpTable
	l4table    L4Table
	RouteTable IpRouteTable
}

//func NewSnmpDevice(id device.CodeID, addr string, community string) *SnmpDevice {
func NewSnmpDevice(device device.Device) *SnmpDevice {
	return &SnmpDevice{
		Device:     device,
		Snmp:       nil,
		Cpu:        Cpu{},
		Memory:     Memory{},
		System:     System{},
		IfTable:    IfTable{},
		IpTable:    IpTable{},
		l4table:    L4Table{},
		RouteTable: IpRouteTable{},
	}
}

type SnmpDeviceTable struct {
	Devices map[device.ID]SnmpDevice
	Count   int
	Store   influx.Config
}

//var ErrDeviceNotExist = errors.New("device does not exist")

func (s *SnmpDevice) String() {
	output := fmt.Sprintf("Device %s", s.Device)
	n := len(output)
	for i := 0; i < n; i++ {
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
	s.RouteTable.String()
}

func (s *SnmpDeviceTable) String() {
	for _, d := range s.Devices {
		d.String()
	}
	fmt.Println("Total:", len(s.Devices), s.Count)
}

func NewSnmpDeviceTable() *SnmpDeviceTable {
	return &SnmpDeviceTable{
		map[device.ID]SnmpDevice{},
		0,
		influx.Config{},
	}
}

func (sd *SnmpDeviceTable) Get(id device.ID) (*SnmpDevice, error) {
	d, exists := sd.Devices[id]
	if !exists {
		return &SnmpDevice{}, device.ErrDeviceNotExist
	}
	return &d, nil
}

func (sd *SnmpDeviceTable) Put(id device.ID, d SnmpDevice) error {
	if _, exists := sd.Devices[id]; !exists {
		return device.ErrDeviceNotExist
	}
	sd.Devices[id] = d
	return nil
}

func (sd *SnmpDeviceTable) Post(d SnmpDevice) (device.ID, error) {
	sd.Count++
	fmt.Printf("dev id : %s", d.Device.GetIdString())
	sd.Devices[d.Device.GetIdString()] = d
	return d.Device.GetIdString(), nil
}

func (sd *SnmpDeviceTable) Delete(id device.ID) error {
	if _, exists := sd.Devices[id]; !exists {
		return device.ErrDeviceNotExist
	}
	delete(sd.Devices, id)
	sd.Count--
	return nil
}
