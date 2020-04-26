package snmpapi

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"nubes/common/lib"
	"strconv"
	"strings"
)

const L4PortMax int = 65535

type L4Table struct {
	tcp []L4TcpPort
	udp []L4UdpPort
}

type L4TcpPort struct {
	Port   int
	IpAddr string
	ConnStatus int
}

type L4UdpPort struct {
	Port   int
	IpAddr string
}

func (s *SnmpDevice) GetL4TcpPort(oid OidType) func() L4TcpPort {
	var oidstr string = oidMap[oid]
	var param L4TcpPort
	param = L4TcpPort{}

	return func() L4TcpPort {
		oids := []string{ oidstr, }
		result, err := s.Snmp.GetNext(oids)
		if err != nil {
			lib.LogWarn("GetL4TcpPort() : %v\n", err)
			param.Port = -1
			param.ConnStatus = -1
			param.IpAddr = ""
			return param
		}

		for _, variable := range result.Variables {
			lib.LogInfo("[%s:%s] oid: %s ",
				s.Device.Ip, s.Device.SnmpCommunity, variable.Name)

			if ! strings.Contains(variable.Name, oidstr) {
				 lib.LogInfo(" - unmatch oid %s --> skip!\n",
					oidstr)
				continue
			}
			switch variable.Type {
			case g.OctetString:
				lib.LogInfo("string: %s\n", string(variable.Value.([]byte)))
			default:
				param.ConnStatus = int(g.ToBigInt(variable.Value).Int64())
				getOid := variable.Name
				if strings.Contains(getOid, oidMap[oid]) {
					//common.LogInfo("%s, len %d\n", oidMap[oid], len(oidMap[oid]))
					lib.LogInfo("%s\n", getOid)
					byteoid := []byte(getOid)
					cutPrefixOid := byteoid[len(oidMap[oid]) + 1 /* . */:]
					//cutPrefixOid := strings.TrimLeft(getOid, oidMap[oid])
					//common.LogInfo("%s\n", cutPrefixOid)
					sliceValue := strings.Split(string(cutPrefixOid), ".")
					param.IpAddr = fmt.Sprintf("%s.%s.%s.%s",
						sliceValue[0], sliceValue[1], sliceValue[2], sliceValue[3])
					param.Port, _ = strconv.Atoi(sliceValue[4])
					oidstr = getOid
					return param
				}
			}
		}
		param.Port = -1
		param.ConnStatus = -1
		param.IpAddr = ""
		return param
	}
}

func (s *SnmpDevice) GetL4UdpPort(oid OidType) func() L4UdpPort {
	var oidstr string = oidMap[oid]
	var param L4UdpPort
	param = L4UdpPort{}

	return func() L4UdpPort {
		oids := []string{ oidstr, }
		result, err := s.Snmp.GetNext(oids)
		if err != nil {
			lib.LogWarn("GetL4UdpPort() : %v\n", err)
			param.Port = -1
			param.IpAddr = ""
			return param
		}

		for _, variable := range result.Variables {
			//common.LogInfo("SubTree [collectdevice %s, community %s] oid: %s ",
			//	s.Device.Ip, s.Device.SnmpCommunity, variable.Name)

			switch variable.Type {
			case g.OctetString:
				lib.LogInfo("string: %s\n", string(variable.Value.([]byte)))
			default:
				param.Port = int(g.ToBigInt(variable.Value).Int64())
				getOid := variable.Name
				if strings.Contains(getOid, oidMap[oid]) {
					//common.LogInfo("%s, len %d\n", oidMap[oid], len(oidMap[oid]))
					//common.LogInfo("%s\n", getOid)
					byteoid := []byte(getOid)
					cutPrefixOid := byteoid[len(oidMap[oid]) + 1 /* . */:]
					//cutPrefixOid := strings.TrimLeft(getOid, oidMap[oid])
					//common.LogInfo("%s\n", cutPrefixOid)
					sliceValue := strings.Split(string(cutPrefixOid), ".")
					param.IpAddr = fmt.Sprintf("%s.%s.%s.%s",
						sliceValue[0], sliceValue[1], sliceValue[2], sliceValue[3])
					param.Port, _ = strconv.Atoi(sliceValue[4])
					oidstr = getOid
					return param
				}
			}
		}
		param.Port = -1
		param.IpAddr = ""
		return param
	}
}

func (d *SnmpDevice) getL4Port() {
	/* TCP Port */
	var param L4TcpPort
	var tcpList []L4TcpPort
	getNextL4PortGen := d.GetL4TcpPort(TypeOidTcpConnState)
	for i := 0 ; i < L4PortMax; i++ {
		param = getNextL4PortGen()
		if param.Port < 0 {
			break
		}
		//fmt.Printf("ip: %s, tcp: %d, status %d\n", param.IpAddr, param.Port, param.ConnStatus)
		tcpList = append(tcpList, param)
	}
	d.l4table.tcp = tcpList

	/* UDP Port */
	var udp L4UdpPort
	var udpList []L4UdpPort
	getNextL4UdpPortGen := d.GetL4UdpPort(TypeOidUdpPort)
	for i := 0 ; i < L4PortMax; i++ {
		udp = getNextL4UdpPortGen()
		if udp.Port < 0 {
			break
		}
		//fmt.Printf("ip: %s, udp: %d", udp.IpAddr, udp.Port)
		udpList = append(udpList, udp)
		//d.l4table.udp = append(d.l4table.udp, udp)
	}
	d.l4table.udp = udpList
}

func (l *L4Table) String() {
	fmt.Printf(" [L4 Port]\n")
	fmt.Printf("  - TCP Count %d : ", len(l.tcp))
	fmt.Print(l.tcp)
	fmt.Printf("\n")
	fmt.Printf("  - UDP Count %d : ", len(l.udp))
	fmt.Print(l.udp)
	fmt.Printf("\n")
}