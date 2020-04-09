package snmpapi

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"nubes/collector/lib"
	"strings"
)

type IpTable struct {
	IpList []IpAddrEntry
}

type IpAddrEntry struct {
	IpAddr string
	NetMask string
	IfIndex int
}

func (ip *IpAddrEntry) Dump() {
	fmt.Printf("  - IP %s Netmask %s ifindex %d\n", ip.IpAddr, ip.NetMask, ip.IfIndex)
}

func (t *IpTable) String() {
	fmt.Printf(" [IP Table]\n")
	for _, e := range t.IpList {
		e.Dump()
	}
}

func (t *IpTable)insertIpEntry2IfIndex(addr string, ifindex int) {
	for i, ent := range t.IpList {
		if ent.IpAddr == addr {
			t.IpList[i].IfIndex = ifindex
			lib.LogInfo("%d find it! --> ipaddr\n", i)
			return
		}
	}
	ent := IpAddrEntry{}
	ent.IpAddr = addr
	ent.IfIndex = ifindex
	t.IpList = append(t.IpList, ent)
	lib.LogInfo("new! --> ifindex\n")
}

func (t *IpTable)insertIpEntry2Ip(addr string) {
	for i, ent := range t.IpList {
		if ent.IpAddr == addr {
			lib.LogInfo("%d find it! --> ip\n", i)
			return
		}
	}
	ent := IpAddrEntry{}
	ent.IpAddr = addr
	t.IpList = append(t.IpList, ent)
	lib.LogInfo("new! --> ip\n")
}

func (t *IpTable)insertIpEntry2Netmask(addr string, netmask string) {
	for i, ent := range t.IpList {
		if ent.IpAddr == addr {
			t.IpList[i].NetMask = netmask
			lib.LogInfo("%d find it! --> netmask\n", i)
			return
		}
	}
	ent := IpAddrEntry{}
	ent.IpAddr = addr
	ent.NetMask = netmask
	t.IpList = append(t.IpList, ent)
	lib.LogInfo("new! --> netmask\n")
}

func (d *SnmpDevice) GetIpEntry(oid OidType) func() (ipEntry *IpAddrEntry, ret int) {
	var oidstr string = oidMap[oid]
	var num int = 0

	return func() (ipEntry *IpAddrEntry, ret int) {
		oids := []string{ oidstr, }
		result, err := d.Snmp.GetNext(oids)
		if err != nil {
			lib.LogWarn("getCpu() : %v\n", err)
		}

		num++
		for _, variable := range result.Variables {
			lib.LogInfo("[%s:%s] oid: %s ",
				d.Device.Ip, d.Device.SnmpCommunity, variable.Name)

			oidstr = variable.Name
			if ! strings.Contains(variable.Name, oidMap[oid]) {
				lib.LogInfo(" - unmatch oid %s --> skip!\n",
					oidstr)
				break
			}

			recvOidSlice := strings.Split(variable.Name, ".")
			orgOidSlice := strings.Split(oidMap[oid], ".")
			tmpSlice := recvOidSlice[len(orgOidSlice):]
			lib.LogInfo("%d\n", len(tmpSlice))
			if len(tmpSlice) < 4 {
				break
			}
			strip := fmt.Sprintf("%s.%s.%s.%s",
				tmpSlice[0], tmpSlice[1], tmpSlice[2], tmpSlice[3])

			switch variable.Type {
			case g.OctetString:
				lib.LogInfo("string: %s\n", string(variable.Value.([]byte)))
			case g.IPAddress:
				value := fmt.Sprint(variable.Value)
				if oid == TypeOidIpAddr {
					d.IpTable.insertIpEntry2Ip(value)
					lib.LogInfo("ip address %s\n", value)
				} else if oid == TypeOidIpMask {
					d.IpTable.insertIpEntry2Netmask(strip, value)
					lib.LogInfo("netmask %s\n", value)
				} else {
					lib.LogWarn("unknown type: value %s\n", value)
				}
				return ipEntry, 0
			default:
				if oid == TypeOidIpIfIndex {
					d.IpTable.insertIpEntry2IfIndex(strip, int(g.ToBigInt(variable.Value).Int64()))
					lib.LogInfo("ifindex %d\n", int(g.ToBigInt(variable.Value).Int64()))
				} else {
					lib.LogInfo("unknown type %d\n", int(g.ToBigInt(variable.Value).Int64()))
				}
				return ipEntry, 0
			}
		}
		return ipEntry, -1
	}
}

func (d *SnmpDevice) GetIpTable() {

	for i := TypeOidIpAddr; i <= TypeOidIpMask; i++ {
		getIpEntry := d.GetIpEntry(OidType(i))
		for k := 0; k < 100; k++ {
			_, ret := getIpEntry()
			if ret < 0 {
				break
			}
		}
	}
}


