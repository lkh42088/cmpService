package snmpapi

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"cmpService/collector/collectdevice"
)

type ID string

type SnmpDevice struct {
	Device collectdevice.ColletDevice
	Snmp   *g.GoSNMP

	Cpu        Cpu
	Memory     Memory
	System     System
	IfTable    IfTable
	IpTable    IpTable
	l4table    L4Table
	RouteTable IpRouteTable
}

func NewSnmpDevice(device collectdevice.ColletDevice) *SnmpDevice {
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

func (s *SnmpDevice) String() {
	output := fmt.Sprintf("ColletDevice %s", s.Device)
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

