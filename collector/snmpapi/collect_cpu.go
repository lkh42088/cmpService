package snmpapi

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"nubes/common/lib"
	"strings"
)

type Cpu struct {
	Idle int
	min1av string
	min5av string
	min10av string
}

func (d *SnmpDevice) getCpu() error {
	var oids []string
	for i := TypeOidCpuIdle; i <= TypeOidCpuMin10Av; i ++ {
		oids = append(oids, oidMap[OidType(i)])
	}
	result, err := d.Snmp.Get(oids)
	if err != nil {
		lib.LogWarn("getCpu() : %v\n", err)
		return err
	}

	num := TypeOidCpuIdle
	for i, variable := range result.Variables {
		lib.LogInfo("[%s:%s] %d: oid: %s ",
			d.Device.Ip, d.Device.SnmpCommunity, i, variable.Name)

		if ! strings.Contains(variable.Name, oidMap[OidType(num)]) {
			lib.LogInfo(" - unmatch oid %s (%s) --> skip!\n",
				oidMap[OidType(num)], oidDescMap[OidType(num)])
			num++
			continue
		}
		num++
		switch variable.Type {
		case g.OctetString:
			lib.LogInfo("string: %s\n", string(variable.Value.([]byte)))
			(&d.Cpu).insertCpu(variable.Name, string(variable.Value.([]byte)))
		default:
			if variable.Name == StrOidCpuIdle {
				d.Cpu.Idle = int(g.ToBigInt(variable.Value).Int64())
				lib.LogInfo("%d\n", d.Cpu.Idle)
			} else {
				number := g.ToBigInt(variable.Value).Int64()
				lib.LogInfo("unknown oid (%s) %d\n", variable.Name, number)
			}
		}
	}
	return nil
}

func (m *Cpu) insertCpu(oid string, value string) {
	switch oid {
	case StrOidCpuMin1Av:
		m.min1av = value
	case StrOidCpuMin5Av:
		m.min5av = value
	case StrOidCpuMin10Av:
		m.min10av = value
	default:
		lib.LogWarn("unknown oid(%s) %s\n", oid, value)
	}
}

func (m *Cpu) String() {
	fmt.Println(" [CPU]")
	fmt.Printf("  - %10d : CPU Idle\n", m.Idle)
	fmt.Printf("  - %10s : CPU 1 minite average\n", m.min1av)
	fmt.Printf("  - %10s : CPU 5 minite average\n", m.min5av)
	fmt.Printf("  - %10s : CPU 10 minite average\n", m.min10av)
}

