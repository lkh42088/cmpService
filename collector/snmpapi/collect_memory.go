package snmpapi

import (
	"cmpService/common/lib"
	"fmt"
	g "github.com/soniah/gosnmp"
)

type Memory struct {
	totalMemory    int
	availMemory    int
	totalSwap      int
	availSwap      int
	totalSharedMem int
	totalBufferMem int
	totalCacheMem  int
}

func (s *SnmpDevice) getMemory() error {
	var oids []string
	for i := TypeOidTotalMemory; i <= TypeOidTotalCacheMem; i++ {
		oids = append(oids, oidMap[OidType(i)])
	}
	result, err := s.Snmp.Get(oids)
	if err != nil {
		lib.LogWarn("Get() err: %v", err)
		return err
	}

	for i, variable := range result.Variables {
		lib.LogInfo("[%s:%s] %d: oid: %s ",
			s.Device.Ip, s.Device.SnmpCommunity, i, variable.Name)

		switch variable.Type {
		case g.OctetString:
			lib.LogInfo("string: %s\n", string(variable.Value.([]byte)))
		default:
			number := g.ToBigInt(variable.Value).Int64()
			lib.LogInfo("if number: %d\n", number)
			(&s.Memory).insertMemory(variable.Name, int(g.ToBigInt(variable.Value).Int64()))
		}
	}
	return nil
}

func (m *Memory) insertMemory(oid string, value int) {
	switch oid {
	case StrOidTotalMemory:
		m.totalMemory = value
	case StrOidAvailMemory:
		m.availMemory = value
	case StrOidTotalSwap:
		m.totalSwap = value
	case StrOidAvailSwap:
		m.availSwap = value
	case StrOidTotalSharedMem:
		m.totalSharedMem = value
	case StrOidTotalBufferMem:
		m.totalBufferMem = value
	case StrOidTotalCacheMem:
		m.totalCacheMem = value
	default:
		lib.LogWarn("unknown : %s %d\n", oid, value)
	}
}

func (m *Memory) String() {
	fmt.Println(" [Memory]")
	fmt.Printf("  - %10d : Total Memory \n", m.totalMemory)
	fmt.Printf("  - %10d : Available Memory \n", m.availMemory)
	fmt.Printf("  - %10d : Total Swap \n", m.totalSwap)
	fmt.Printf("  - %10d : Available Swap \n", m.availMemory)
	fmt.Printf("  - %10d : Total Buffer Memory \n", m.totalBufferMem)
	fmt.Printf("  - %10d : Total Cache Memory \n", m.totalCacheMem)
}
