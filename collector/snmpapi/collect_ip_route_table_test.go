package snmpapi

import (
	"fmt"
	"log"
	"testing"
)

func TestIpRouteTable(t *testing.T) {
	s := SnmpDevice{}
	s.Device.Ip = "127.0.0.1"
	s.Device.SnmpCommunity = "nubes"

	// Init Snmp
	s.InitDeviceSnmp()

	err := s.Snmp.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer s.Snmp.Conn.Close()

	// Get IP Route Table
	s.GetIpRouteTable()
	s.RouteTable.String()
}

func TestIpRouteTableGetNext(t *testing.T) {
	s := SnmpDevice{}
	s.Device.Ip = "127.0.0.1"
	s.Device.SnmpCommunity = "nubes"
	oidstr := ".1.3.6.1.2.1.4.21.1.1"
	oid := []string{oidstr}

	// Init Snmp
	s.InitDeviceSnmp()

	err := s.Snmp.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer s.Snmp.Conn.Close()

	// Get IP Route Table
	for i := 0; i < 20; i++ {
		result, _ := s.Snmp.GetNext(oid)
		fmt.Println(result)
		for _, variable := range result.Variables {
			oidstr = variable.Name
			oid = []string{oidstr}
		}
	}
}

func BenchmarkIpRouteTable(b *testing.B) {
	s := SnmpDevice{}
	s.Device.Ip = "127.0.0.1"
	s.Device.SnmpCommunity = "nubes"

	// Init Snmp
	s.InitDeviceSnmp()

	err := s.Snmp.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer s.Snmp.Conn.Close()

	// Get IP Route Table
	s.GetIpRouteTable()
	s.RouteTable.String()
}
