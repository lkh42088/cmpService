package snmpapi

import (
	"fmt"
	"log"
	"testing"
)

func TestMemory(t *testing.T) {
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

	s.getMemory()
	s.Memory.String()
}
