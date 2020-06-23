package snmpapi

import (
	"fmt"
	"log"
	"testing"
)

func TestIpTable(t *testing.T) {
	s := SnmpDevice{}
	s.Device.Ip = "121.156.65.139"
	s.Device.SnmpCommunity = "cmpService"
	s.InitDeviceSnmp()
	err := s.Snmp.Connect()
	if err != nil {
		fmt.Printf("[server %s, comm %s]\n",
			s.Snmp.Target, s.Snmp.Community)
		log.Fatalf("Connect() err: %v", err)
	}
	defer s.Snmp.Conn.Close()

	// Get IP Table
	s.GetIpTable()
	s.IpTable.String()
}

func TestIpEntry(t *testing.T) {
	s := SnmpDevice{}
	s.Device.Ip = "121.156.65.139"
	s.Device.SnmpCommunity = "cmpService"
	s.InitDeviceSnmp()
	err := s.Snmp.Connect()
	var ret int
	if err != nil {
		fmt.Printf("[server %s, comm %s]\n",
			s.Snmp.Target, s.Snmp.Community)
		log.Fatalf("Connect() err: %v", err)
	}
	defer s.Snmp.Conn.Close()

	getip := s.GetIpEntry(TypeOidIpAddr)
	for i := 0; i < 100; i++ {
		_, ret = getip()
		if ret < 0 {
			break
		}
	}
	getmask := s.GetIpEntry(TypeOidIpMask)
	for i := 0; i < 100; i++ {
		_, ret = getmask()
		if ret < 0 {
			break
		}
	}
	getIpIfIndex := s.GetIpEntry(TypeOidIpIfIndex)
	for i := 0; i < 100; i++ {
		_, ret = getIpIfIndex()
		if ret < 0 {
			break
		}
	}
}

func TestIpAddr(t *testing.T) {
	s := SnmpDevice{}
	s.Device.Ip = "121.156.65.139"
	s.Device.SnmpCommunity = "cmpService"
	s.InitDeviceSnmp()
	err := s.Snmp.Connect()
	if err != nil {
		fmt.Printf("[server %s, comm %s]\n",
			s.Snmp.Target, s.Snmp.Community)
		log.Fatalf("Connect() err: %v", err)
	}
	defer s.Snmp.Conn.Close()

	var ret int
	getip := s.GetIpEntry(TypeOidIpAddr)
	for i := 0; i < 100; i++ {
		_, ret = getip()
		if ret < 0 {
			break
		}
	}
}

func TestIpIfIndex(t *testing.T) {
	s := SnmpDevice{}
	s.Device.Ip = "121.156.65.139"
	s.Device.SnmpCommunity = "cmpService"
	s.InitDeviceSnmp()
	err := s.Snmp.Connect()
	if err != nil {
		fmt.Printf("[server %s, comm %s]\n",
			s.Snmp.Target, s.Snmp.Community)
		log.Fatalf("Connect() err: %v", err)
	}
	defer s.Snmp.Conn.Close()

	var ret int
	getIpIfIndex := s.GetIpEntry(TypeOidIpIfIndex)
	for i := 0; i < 100; i++ {
		_, ret = getIpIfIndex()
		if ret < 0 {
			break
		}
	}
}

func TestIpMask(t *testing.T) {
	s := SnmpDevice{}
	s.Device.Ip = "121.156.65.139"
	s.Device.SnmpCommunity = "cmpService"
	s.InitDeviceSnmp()
	err := s.Snmp.Connect()
	if err != nil {
		fmt.Printf("[server %s, comm %s]\n",
			s.Snmp.Target, s.Snmp.Community)
		log.Fatalf("Connect() err: %v", err)
	}
	defer s.Snmp.Conn.Close()

	var ret int
	getmask := s.GetIpEntry(TypeOidIpMask)
	for i := 0; i < 100; i++ {
		_, ret = getmask()
		if ret < 0 {
			break
		}
	}
}
