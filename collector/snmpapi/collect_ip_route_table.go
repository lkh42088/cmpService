package snmpapi

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"cmpService/common/lib"
	"strings"
)

type IpRouteTable struct {
	IpRouteList []IpRouteEntry
}

type IpRouteEntry struct {
	IpRouteDest    string
	IpRouteIfIndex int
	IpRouteMetric1 int
	IpRouteMetric2 int
	IpRouteMetric3 int
	IpRouteMetric4 int
	IpRouteNextHop string
	IpRouteType    int
	IpRouteProto   int
	IpRouteAge     int
	IpRouteMask    string
	IpRouteMetric5 int
	IpRouteInfo    string // Octet Identifier check
}

func (rt *IpRouteTable) Dump() {
	for i := 0; i < len(rt.IpRouteList); i++ {
		fmt.Printf(" - IP ROUTE TABLE [%d] - \n"+
			"   DEST: %s\n"+
			"   IFINDEX: %d\n"+
			"   METRIC1: %d\n"+
			"   METRIC2: %d\n"+
			"   METRIC3: %d\n"+
			"   METRIC4: %d\n"+
			"   NEXTHOP: %s\n"+
			"   TYPE: %d\n"+
			"   PROTO: %d\n"+
			"   AGE: %d\n"+
			"   MASK: %s\n"+
			"   METRIC5: %d\n"+
			"   IP ROUTE INFO: %s\n",
			i, rt.IpRouteList[i].IpRouteDest, rt.IpRouteList[i].IpRouteIfIndex,
			rt.IpRouteList[i].IpRouteMetric1, rt.IpRouteList[i].IpRouteMetric2,
			rt.IpRouteList[i].IpRouteMetric3, rt.IpRouteList[i].IpRouteMetric4,
			rt.IpRouteList[i].IpRouteNextHop, rt.IpRouteList[i].IpRouteType,
			rt.IpRouteList[i].IpRouteProto, rt.IpRouteList[i].IpRouteAge,
			rt.IpRouteList[i].IpRouteMask, rt.IpRouteList[i].IpRouteMetric5,
			rt.IpRouteList[i].IpRouteInfo)
	}
}

func (rt *IpRouteTable) String() {
	rt.Dump()
}

func (rt *IpRouteTable) insertIpRouteTable2Ip(addr string) {
	for i, ent := range rt.IpRouteList {
		if ent.IpRouteDest == addr {
			lib.LogInfo("%d find it! --> ipaddr\n", i)
			return
		}
	}
	table := IpRouteEntry{}
	table.IpRouteDest = addr
	rt.IpRouteList = append(rt.IpRouteList, table)
}

func (rt *IpRouteTable) GetIpRouteTable2Ip(addr string) int {
	if rt != nil {
		for i, ent := range rt.IpRouteList {
			if ent.IpRouteDest == addr {
				lib.LogInfo("%d find it! --> ipaddr\n", i)
				return i
			}
		}
	}
	return 0
}

func (d *SnmpDevice) GetIpRouteEntry(oid OidType) func() (t *IpRouteTable, ret int) {
	var oidstr string = oidMap[oid]
	var idx = 0

	return func() (rt *IpRouteTable, ret int) {
		oids := []string{oidstr}
		result, err := d.Snmp.GetNext(oids)
		if err != nil {
			lib.LogWarn("getIpRouteEntry() : %s\n", err.Error())
			return rt, -1
		}

		for _, variable := range result.Variables {
			lib.LogInfo("[%s:%s] oid: %s\n",
				d.Device.Ip, d.Device.SnmpCommunity, variable.Name)

			oidstr = variable.Name
			if !strings.Contains(oidstr, oidMap[oid]) {
				lib.LogInfo(" - Unmatch oid %s --> skip!\n", oidstr)
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

			idx = d.RouteTable.GetIpRouteTable2Ip(strip)

			switch oid {
			case TypeOidIpRouteDest:
				if IsSameVarType(variable.Type, g.IPAddress) {
					d.RouteTable.insertIpRouteTable2Ip(strip)
					return rt, 0
				}
			case TypeOidIpRouteIfIndex:
				if IsSameVarType(variable.Type, g.Integer) {
					d.RouteTable.IpRouteList[idx].IpRouteIfIndex =
						int(g.ToBigInt(variable.Value).Int64())
					return rt, 0
				}
			case TypeOidIpRouteMetric1:
				if IsSameVarType(variable.Type, g.Integer) {
					d.RouteTable.IpRouteList[idx].IpRouteMetric1 =
						int(g.ToBigInt(variable.Value).Int64())
					return rt, 0
				}
			case TypeOidIpRouteMetric2:
				if IsSameVarType(variable.Type, g.Integer) {
					d.RouteTable.IpRouteList[idx].IpRouteMetric2 =
						int(g.ToBigInt(variable.Value).Int64())
					return rt, 0
				}
			case TypeOidIpRouteMetric3:
				if IsSameVarType(variable.Type, g.Integer) {
					d.RouteTable.IpRouteList[idx].IpRouteMetric3 =
						int(g.ToBigInt(variable.Value).Int64())
					return rt, 0
				}
			case TypeOidIpRouteMetric4:
				if IsSameVarType(variable.Type, g.Integer) {
					d.RouteTable.IpRouteList[idx].IpRouteMetric4 =
						int(g.ToBigInt(variable.Value).Int64())
					return rt, 0
				}
			case TypeOidIpRouteNextHop:
				if IsSameVarType(variable.Type, g.IPAddress) {
					d.RouteTable.IpRouteList[idx].IpRouteNextHop =
						fmt.Sprint(variable.Value)
					return rt, 0
				}
			case TypeOidIpRouteType:
				if IsSameVarType(variable.Type, g.Integer) {
					d.RouteTable.IpRouteList[idx].IpRouteType =
						int(g.ToBigInt(variable.Value).Int64())
					return rt, 0
				}
			case TypeOidIpRouteProto:
				if IsSameVarType(variable.Type, g.Integer) {
					d.RouteTable.IpRouteList[idx].IpRouteProto =
						int(g.ToBigInt(variable.Value).Int64())
					return rt, 0
				}
			case TypeOidIpRouteAge:
				if IsSameVarType(variable.Type, g.Integer) {
					d.RouteTable.IpRouteList[idx].IpRouteAge =
						int(g.ToBigInt(variable.Value).Int64())
					return rt, 0
				}
			case TypeOidIpRouteMask:
				if IsSameVarType(variable.Type, g.IPAddress) {
					d.RouteTable.IpRouteList[idx].IpRouteMask =
						fmt.Sprint(variable.Value)
					return rt, 0
				}
			case TypeOidIpRouteMetric5:
				if IsSameVarType(variable.Type, g.Integer) {
					d.RouteTable.IpRouteList[idx].IpRouteMetric5 =
						int(g.ToBigInt(variable.Value).Int64())
					return rt, 0
				}
			case TypeOidIpRouteInfo:
				if IsSameVarType(variable.Type, g.ObjectIdentifier) {
					d.RouteTable.IpRouteList[idx].IpRouteInfo =
						fmt.Sprint(variable.Value)
					return rt, 0
				}
			default:
				lib.LogInfo("Unmatched Oid : %s", variable.Name)
				return rt, -1
			}
		}
		return rt, -1
	}
}

// true: match, false: unmatch
func IsSameVarType(a interface{}, b interface{}) bool {
	if a == nil || b == nil {
		lib.LogInfoln("Unmatch variable type.")
		return false
	}
	return a == b
}

func (d *SnmpDevice) GetIpRouteTable() {
	for i := TypeOidIpRouteDest; i <= TypeOidIpRouteInfo; i++ {
		getIpRouteEntry := d.GetIpRouteEntry(OidType(i))
		for {
			_, ret := getIpRouteEntry()
			if ret < 0 {
				break
			}
		}
	}
}
