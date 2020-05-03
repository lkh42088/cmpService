package snmpapi

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"log"
	"nubes/common/lib"
	"strings"
)

type IfEntry struct {
	ifIndex           int         // 1.3.6.1.2.1.2.2.1.1   INTEGER
	ifDescr           string      // 1.3.6.1.2.1.2.2.1.2   OCTET STRING 0..255
	ifType            int         // 1.3.6.1.2.1.2.2.1.3   INTEGER
	ifMTU             int         // 1.3.6.1.2.1.2.2.1.4   INTEGER
	ifSpeed           int         // 1.3.6.1.2.1.2.2.1.5   Gauge
	ifPhysAddress     string      // 1.3.6.1.2.1.2.2.1.6   OCTET STRING
	ifAdminStatus     int         // 1.3.6.1.2.1.2.2.1.7   INTEGER  { up(1), down(2), testing(3) }
	ifOperStatus      int         // 1.3.6.1.2.1.2.2.1.8   INTEGER  { up(1), down(2), testing(3) }
	ifLastChange      int         // 1.3.6.1.2.1.2.2.1.9   TIMETICKS (millisecond)

	ifInOctets        int64       // 1.3.6.1.2.1.2.2.1.10  COUNT (Counter32)
	ifInUcastPkts     int64       // 1.3.6.1.2.1.2.2.1.11  COUNT (Counter32)
	ifInNUcastPkts    int64       // 1.3.6.1.2.1.2.2.1.12  COUNT (Counter32)
	ifInDiscards      int64       // 1.3.6.1.2.1.2.2.1.13  COUNT (Counter32)
	ifInErrors        int64       // 1.3.6.1.2.1.2.2.1.14  COUNT (Counter32)
	ifInUnknownProtos int64       // 1.3.6.1.2.1.2.2.1.15  COUNT (Counter32)

	ifOutOctets       int64       // 1.3.6.1.2.1.2.2.1.16  COUNT (Counter32)
	ifOutUcastPkts    int64       // 1.3.6.1.2.1.2.2.1.17  COUNT (Counter32)
	ifOutNUcastPkts   int64       // 1.3.6.1.2.1.2.2.1.18  COUNT (Counter32)
	ifOutDiscards     int64       // 1.3.6.1.2.1.2.2.1.19  COUNT (Counter32)
	ifOutErrors       int64       // 1.3.6.1.2.1.2.2.1.20  COUNT (Counter32)

	ifOutQLen         int64       // 1.3.6.1.2.1.2.2.1.21  GAUGE
	ifSpecific        int64       // 1.3.6.1.2.1.2.2.1.22  OBJECT IDENTIFIER (OID)

	ifName            string	  // 1.3.6.1.2.1.31.1.1.1.1 DISPLAY STRING
	ifHCInOctets	  uint64	  // 1.3.6.1.2.1.31.1.1.1.6 COUNT (Counter64)
	ifHCOutOctets	  uint64	  // 1.3.6.1.2.1.31.1.1.1.10 COUNT (Counter64)
}

type IfTable struct {
	ifNumber int64
	ifEntry []IfEntry
}

func (d *SnmpDevice) getIfNumber() int64 {
	var number int64 = 0
	oids := []string{StrOidIfNumber, }
	result, err2 := d.Snmp.Get(oids)
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		lib.LogInfo("[%s:%s] %d: oid: %s ",
			d.Device.Ip, d.Device.SnmpCommunity, i, variable.Name)

		switch variable.Type {
		case g.OctetString:
			lib.LogInfo("string: %s\n", string(variable.Value.([]byte)))
		default:
			number = g.ToBigInt(variable.Value).Int64()
			lib.LogInfo("if number: %d\n", number)
		}
	}
	return number
}

func (s *SnmpDevice) getIfEntryFromSnmp(oid OidType) error {
	if s.IfTable.ifNumber < 1 {
		return nil
	}
	result, err := s.Snmp.GetBulk([]string{oidMap[oid],}, 0, uint8(s.IfTable.ifNumber))
	if err != nil {
		log.Fatalf("getIfEntryFromSnmp() : %v", err)
		return err
	}

	for i, variable := range result.Variables {
		lib.LogInfo("[%s:%s] %d: oid: %s ", s.Device.Ip, s.Device.SnmpCommunity, i, variable.Name)
		if ! strings.Contains(variable.Name, oidMap[oid]) {
			lib.LogInfo(" - unmatch oid %s (%s) --> skip!\n", oidMap[oid], oidDescMap[oid])
			continue
		}
		switch variable.Type {
		case g.OctetString:
			s.insertOctetString2IfEntry(i, oid, variable.Value.([]byte))
		default:
			s.insertIntegerIfEntry(i, oid, g.ToBigInt(variable.Value).Int64())
		}
	}
	return nil
}

func (s *SnmpDevice)insertOctetString2IfEntry(num int, oid OidType, raw []byte) {

	lib.LogInfo("%s, %s\n", oidDescMap[oid], string(raw))
	switch oid {
	// System
	case TypeOidSysDescr:
		s.System.desc = string(raw)
	case TypeOidSysHostname:
		s.System.hostname = string(raw)
	// Interface
	case TypeOidIfDescr:
		s.IfTable.ifEntry[num].ifDescr = string(raw)
	case TypeOidIfPhysAddr:
		s.IfTable.ifEntry[num].ifPhysAddress = convertByte2StringMac(raw)
	case TypeOidIfName:
		s.IfTable.ifEntry[num].ifName = string(raw)
	default:
		log.Fatalf("insertOctetString2IfEntry() err: oid %d, value %s\n", oid, string(raw))
	}
}

func (s *SnmpDevice)insertIntegerIfEntry(num int, oid OidType, value int64) {

	lib.LogInfo("%s, %d\n", oidDescMap[oid], value)
	switch oid {
	// System
	case TypeOidSysUptime:
		s.System.uptime = value
	// Interface
	case TypeOidIfIndex:
		s.IfTable.ifEntry[num].ifIndex = int(value)
	case TypeOidIfType:
		s.IfTable.ifEntry[num].ifType = int(value)
	case TypeOidIfSpeed:
		s.IfTable.ifEntry[num].ifSpeed = int(value)
	case TypeOidIfMtu:
		s.IfTable.ifEntry[num].ifMTU = int(value)
	case TypeOidIfAdminStatus:
		s.IfTable.ifEntry[num].ifAdminStatus = int(value)
	case TypeOidIfOperStatus:
		s.IfTable.ifEntry[num].ifOperStatus = int(value)
	case TypeOidIfLastChange:
		s.IfTable.ifEntry[num].ifLastChange = int(value)
	case TypeOidIfInOctets:
		s.IfTable.ifEntry[num].ifInOctets = value
	case TypeOidIfInUcastPkts:
		s.IfTable.ifEntry[num].ifInUcastPkts= value
	case TypeOidIfInNUcastPkts:
		s.IfTable.ifEntry[num].ifInNUcastPkts= value
	case TypeOidIfInDiscards:
		s.IfTable.ifEntry[num].ifInDiscards = value
	case TypeOidIfInErrors:
		s.IfTable.ifEntry[num].ifInErrors = value
	case TypeOidIfInUnknownProto:
		s.IfTable.ifEntry[num].ifInUnknownProtos = value
	case TypeOidIfOutOctets:
		s.IfTable.ifEntry[num].ifOutOctets = value
	case TypeOidIfOutUcastPkts:
		s.IfTable.ifEntry[num].ifOutUcastPkts = value
	case TypeOidIfOutNUcastPkts:
		s.IfTable.ifEntry[num].ifOutNUcastPkts = value
	case TypeOidIfOutDiscards:
		s.IfTable.ifEntry[num].ifOutDiscards = value
	case TypeOidIfOutErrors:
		s.IfTable.ifEntry[num].ifOutErrors = value
	case TypeOidIfOutQLen:
		s.IfTable.ifEntry[num].ifOutQLen = value
	case TypeOidIfSpecific:
		s.IfTable.ifEntry[num].ifSpecific = value
	case TypeOidIfHCInOctets:
		s.IfTable.ifEntry[num].ifHCInOctets = uint64(value)
	case TypeOidIfHCOutOctets:
		s.IfTable.ifEntry[num].ifHCOutOctets = uint64(value)
	default:
		log.Fatalf("insertIntegerIfEntry() err: oid %d, value %d\n", oid, value)
	}
}

func (s *SnmpDevice) getIfDescr(){
	oids := []string{StrOidIfDescr, }
	if s.IfTable.ifNumber < 1 {
		return
	}
	result, err2 := s.Snmp.GetBulk(oids, 0, uint8(s.IfTable.ifNumber))
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		lib.LogInfo("[collectdevice %s, community %s] %d: oid: %s ", s.Device.Ip, s.Device.SnmpCommunity, i, variable.Name)
		switch variable.Type {
		case g.OctetString:
			s.IfTable.ifEntry[i].ifDescr = string(variable.Value.([]byte))
			lib.LogInfo("string: %s\n",s.IfTable.ifEntry[i].ifDescr)
		default:
		}
	}
}

func (s *SnmpDevice) getIfIndex(){
	oids := []string{StrOidIfIndex, }
	if s.IfTable.ifNumber < 1 {
		return
	}
	result, err2 := s.Snmp.GetBulk(oids, 0, uint8(s.IfTable.ifNumber))
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		lib.LogInfo("[%s:%s] %d: oid: %s ",
			s.Device.Ip, s.Device.SnmpCommunity, i, variable.Name)

		switch variable.Type {
		case g.OctetString:
			lib.LogInfo("string: %s\n", string(variable.Value.([]byte)))
		default:
			ifindex := g.ToBigInt(variable.Value)
			s.IfTable.ifEntry[i].ifIndex = int(ifindex.Int64())
			lib.LogInfo("ifEntry number: %d\n", s.IfTable.ifEntry[i].ifIndex)
		}
	}
}

func (t *IfTable) String() {
	fmt.Println(" [Interface Table]")
	fmt.Printf("  - The number of Interface: %d\n", t.ifNumber)
	for i:=0 ; int(t.ifNumber) > i; i++ {
		ife := t.ifEntry[i]
		ife.DumpWithNum(i)
	}
}

func (ife *IfEntry) DumpWithNum(num int) {
	fmt.Printf("   [%d] IFINDEX[%d] NAME[%s] Desc(%s)\n",
		num, ife.ifIndex, ife.ifName, ife.ifDescr)
	fmt.Printf("         Type(%d), MTU(%d), Speed(%d), MAC(%s)\n",
		ife.ifType, ife.ifMTU, ife.ifSpeed, strings.ToUpper(string2mac(ife.ifPhysAddress)))
	fmt.Printf("         Status(Admin %d, Oper %d), LastChange(%d)\n",
		ife.ifAdminStatus, ife.ifOperStatus, ife.ifLastChange)

	fmt.Print("         IN  ")
	fmt.Printf("%d bytes, ", ife.ifInOctets)
	fmt.Printf("U %d pkts, ", ife.ifInUcastPkts)
	fmt.Printf("NU %d pkts, ", ife.ifInNUcastPkts)
	fmt.Printf("D %d pkts, ", ife.ifInDiscards)
	fmt.Printf("E %d pkts, ", ife.ifInErrors)
	fmt.Printf("HC %d bytes", ife.ifHCInOctets)
	fmt.Print("\n")

	fmt.Print("         OUT ")
	fmt.Printf("%d bytes, ", ife.ifOutOctets)
	fmt.Printf("U %d pkts, ", ife.ifOutUcastPkts)
	fmt.Printf("NU %d pkts, ", ife.ifOutNUcastPkts)
	fmt.Printf("D %d pkts, ", ife.ifOutDiscards)
	fmt.Printf("E %d pkts, ", ife.ifOutErrors)
	fmt.Printf("HC %d bytes", ife.ifHCOutOctets)
	fmt.Print("\n")

	fmt.Print("         OutQLen ")
	fmt.Printf("%d ", ife.ifOutOctets)
	fmt.Printf("Specific %d", ife.ifSpecific)
	fmt.Print("\n")
}

func (ife *IfEntry) Dump() {
	fmt.Printf("   IFINDEX[%d] NAME[%s] Desc(%s)\n",
		ife.ifIndex, ife.ifName, ife.ifDescr)
	fmt.Printf("         Type(%d), MTU(%d), Speed(%d), MAC(%s)\n",
		ife.ifType, ife.ifMTU, ife.ifSpeed, strings.ToUpper(string2mac(ife.ifPhysAddress)))
	fmt.Printf("         Status(Admin %d, Oper %d), LastChange(%d)\n",
		ife.ifAdminStatus, ife.ifOperStatus, ife.ifLastChange)

	fmt.Print("         IN  ")
	fmt.Printf("%d, ", ife.ifInOctets)
	fmt.Printf("U %d, ", ife.ifInUcastPkts)
	fmt.Printf("NU %d, ", ife.ifInNUcastPkts)
	fmt.Printf("D %d, ", ife.ifInDiscards)
	fmt.Printf("E %d, ", ife.ifInErrors)
	fmt.Printf("HC %d", ife.ifHCInOctets)
	fmt.Print("\n")

	fmt.Print("         OUT ")
	fmt.Printf("%d, ", ife.ifOutOctets)
	fmt.Printf("U %d, ", ife.ifOutUcastPkts)
	fmt.Printf("NU %d, ", ife.ifOutNUcastPkts)
	fmt.Printf("D %d, ", ife.ifOutDiscards)
	fmt.Printf("E %d, ", ife.ifOutErrors)
	fmt.Printf("HC %d", ife.ifHCOutOctets)
	fmt.Print("\n")

	fmt.Print("         OutQLen ")
	fmt.Printf("%d ", ife.ifOutOctets)
	fmt.Printf("Specific %d", ife.ifSpecific)
	fmt.Print("\n")
}
