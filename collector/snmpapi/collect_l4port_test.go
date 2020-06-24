package snmpapi_test

import (
	"fmt"
	"log"
	"testing"
)
import "cmpService/collector/snmpapi"

func TestTcp(t *testing.T) {
	s := snmpapi.SnmpDevice{}
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
	//s.SubTree(snmpapi.TypeOidTcpConnState)
	var param snmpapi.L4TcpPort
	getNextL4PortGen := s.GetL4TcpPort(snmpapi.TypeOidTcpConnState)
	for i := 0; i < 10000; i++ {
		param = getNextL4PortGen()
		if param.Port < 0 {
			break
		}
		fmt.Printf("ip: %s, tcp: %d, status %d\n",
			param.IpAddr, param.Port, param.ConnStatus)
	}
}

func TestUdp(t *testing.T) {
	s := snmpapi.SnmpDevice{}
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
	//s.SubTree(snmpapi.TypeOidTcpConnState)
	var param snmpapi.L4UdpPort
	getNextL4PortGen := s.GetL4UdpPort(snmpapi.TypeOidUdpPort)
	for i := 0; i < 10000; i++ {
		param = getNextL4PortGen()
		if param.Port < 0 {
			break
		}
		fmt.Printf("ip: %s, udp: %d \n",
			param.IpAddr, param.Port)
	}
}

func TestTcpSlice(t *testing.T) {
	s := snmpapi.SnmpDevice{}
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
	//s.SubTree(snmpapi.TypeOidTcpConnState)
	var param snmpapi.L4TcpPort
	var tcpList []snmpapi.L4TcpPort
	//tcpList = make([]snmpapi.L4TcpPort, 1)
	getNextL4PortGen := s.GetL4TcpPort(snmpapi.TypeOidTcpConnState)
	for i := 0; i < 10000; i++ {
		param = getNextL4PortGen()
		if param.Port < 0 {
			break
		}
		fmt.Printf("ip: %s, tcp: %d, status %d\n",
			param.IpAddr, param.Port, param.ConnStatus)
		tcpList = append(tcpList, param)
	}
	fmt.Printf("-----------------------------\n")
	fmt.Printf("TCP List\n")
	fmt.Println("Count: ", len(tcpList))
	fmt.Println(tcpList)
	fmt.Printf("-----------------------------\n")
}

func TestUdpSlice(t *testing.T) {
	s := snmpapi.SnmpDevice{}
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
	//s.SubTree(snmpapi.TypeOidTcpConnState)
	var param snmpapi.L4UdpPort
	var udpList []snmpapi.L4UdpPort
	getNextL4PortGen := s.GetL4UdpPort(snmpapi.TypeOidUdpPort)
	for i := 0; i < 10000; i++ {
		param = getNextL4PortGen()
		if param.Port < 0 {
			break
		}
		fmt.Printf("ip: %s, udp: %d \n",
			param.IpAddr, param.Port)
		udpList = append(udpList, param)
	}

	fmt.Printf("-----------------------------\n")
	fmt.Printf("UDP List\n")
	fmt.Println("Count: ", len(udpList))
	fmt.Println(udpList)
	fmt.Printf("-----------------------------\n")
}
