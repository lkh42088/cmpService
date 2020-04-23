package snmpapi

type OidType int

/* 1. OID Type Definition */
const (
	_ = iota
	/* System */
	TypeOidSysDescr
	TypeOidSysUptime
	TypeOidSysHostname

	/* Memory */
	TypeOidTotalMemory
	TypeOidAvailMemory
	TypeOidTotalSwap
	TypeOidAvailSwap
	TypeOidTotalSharedMem
	TypeOidTotalBufferMem
	TypeOidTotalCacheMem /* 10 */

	/* Cpu Load */
	TypeOidCpuIdle
	TypeOidCpuMin1Av
	TypeOidCpuMin5Av
	TypeOidCpuMin10Av

	/* Interface */
	TypeOidIfNumber // The number of Interface
	TypeOidIfIndex  // Interface GetBulk begin
	TypeOidIfDescr
	TypeOidIfType
	TypeOidIfMtu
	TypeOidIfSpeed /* 20 */
	TypeOidIfPhysAddr
	TypeOidIfAdminStatus
	TypeOidIfOperStatus
	TypeOidIfLastChange
	TypeOidIfInOctets
	TypeOidIfInUcastPkts
	TypeOidIfInNUcastPkts
	TypeOidIfInDiscards
	TypeOidIfInErrors
	TypeOidIfInUnknownProto /* 30 */
	TypeOidIfOutOctets
	TypeOidIfOutUcastPkts
	TypeOidIfOutNUcastPkts
	TypeOidIfOutDiscards
	TypeOidIfOutErrors
	TypeOidIfOutQLen
	TypeOidIfSpecific
	TypeOidIfName // Interface GetBulk end

	/* L4 Port */
	TypeOidTcpConnState /* 40 */
	TypeOidUdpPort

	/* IP Address Entry */
	TypeOidIpAddr
	TypeOidIpIfIndex
	TypeOidIpMask

	/* IP Route Entry */
	TypeOidIpRouteDest
	TypeOidIpRouteIfIndex
	TypeOidIpRouteMetric1
	TypeOidIpRouteMetric2
	TypeOidIpRouteMetric3
	TypeOidIpRouteMetric4 /* 50 */
	TypeOidIpRouteNextHop
	TypeOidIpRouteType
	TypeOidIpRouteProto
	TypeOidIpRouteAge
	TypeOidIpRouteMask
	TypeOidIpRouteMetric5
	TypeOidIpRouteInfo

	/* OID MAX */
	typeOidMAX
)

const TypeOidIfGetBulkBegin int = TypeOidIfIndex
const TypeOidIfGetBulkEnd int = TypeOidIfName

/* 2. OID String Definition */
const (
	/* System */
	StrOidSysDescr    = ".1.3.6.1.2.1.1.1.0" // OCTET STRING (0~255)
	StrOidSysUptime   = ".1.3.6.1.2.1.1.3.0" // TIMETICKS
	StrOidSysHostname = ".1.3.6.1.2.1.1.5.0" // OCTET STRING (0~255)

	/* System */
	StrOidTotalMemory    = ".1.3.6.1.4.1.2021.4.5.0"  // INTEGER
	StrOidAvailMemory    = ".1.3.6.1.4.1.2021.4.6.0"  // INTEGER
	StrOidTotalSwap      = ".1.3.6.1.4.1.2021.4.3.0"  // INTEGER
	StrOidAvailSwap      = ".1.3.6.1.4.1.2021.4.4.0"  // INTEGER
	StrOidTotalSharedMem = ".1.3.6.1.4.1.2021.4.13.0" // INTEGER
	StrOidTotalBufferMem = ".1.3.6.1.4.1.2021.4.14.0" // INTEGER
	StrOidTotalCacheMem  = ".1.3.6.1.4.1.2021.4.15.0" // INTEGER

	/* Cpu Load */
	StrOidCpuIdle    = ".1.3.6.1.4.1.2021.11.11.0"  // INTEGER
	StrOidCpuMin1Av  = ".1.3.6.1.4.1.2021.10.1.3.1" // OCTET STRING
	StrOidCpuMin5Av  = ".1.3.6.1.4.1.2021.10.1.3.2" // OCTET STRING
	StrOidCpuMin10Av = ".1.3.6.1.4.1.2021.10.1.3.3" // OCTET STRING

	/* Interface */
	StrOidIfNumber         = ".1.3.6.1.2.1.2.1.0"
	StrOidIfIndex          = ".1.3.6.1.2.1.2.2.1.1"
	StrOidIfDescr          = ".1.3.6.1.2.1.2.2.1.2"
	StrOidIfType           = ".1.3.6.1.2.1.2.2.1.3"
	StrOidIfMtu            = ".1.3.6.1.2.1.2.2.1.4"
	StrOidIfSpeed          = ".1.3.6.1.2.1.2.2.1.5"
	StrOidIfPhysAddr       = ".1.3.6.1.2.1.2.2.1.6"
	StrOidIfAdminStatus    = ".1.3.6.1.2.1.2.2.1.7"
	StrOidIfOperStatus     = ".1.3.6.1.2.1.2.2.1.8"
	StrOidIfLastChange     = ".1.3.6.1.2.1.2.2.1.9"
	StrOidIfInOctets       = ".1.3.6.1.2.1.2.2.1.10"
	StrOidIfInUcastPkts    = ".1.3.6.1.2.1.2.2.1.11"
	StrOidIfInNUcastPkts   = ".1.3.6.1.2.1.2.2.1.12"
	StrOidIfInDiscards     = ".1.3.6.1.2.1.2.2.1.13"
	StrOidIfInErrors       = ".1.3.6.1.2.1.2.2.1.14"
	StrOidIfInUnknownProto = ".1.3.6.1.2.1.2.2.1.15"
	StrOidIfOutOctets      = ".1.3.6.1.2.1.2.2.1.16"
	StrOidIfOutUcastPkts   = ".1.3.6.1.2.1.2.2.1.17"
	StrOidIfOutNUcastPkts  = ".1.3.6.1.2.1.2.2.1.18"
	StrOidIfOutDiscards    = ".1.3.6.1.2.1.2.2.1.19"
	StrOidIfOutErrors      = ".1.3.6.1.2.1.2.2.1.20"
	StrOidIfOutQLen        = ".1.3.6.1.2.1.2.2.1.21"
	StrOidIfSpecific       = ".1.3.6.1.2.1.2.2.1.22"

	/* IP Route Entry */
	StrOidIpRouteDest    = ".1.3.6.1.2.1.4.21.1.1"  // IPADDRESS
	StrOidIpRouteIfIndex = ".1.3.6.1.2.1.4.21.1.2"  // INTEGER
	StrOidIpRouteMetric1 = ".1.3.6.1.2.1.4.21.1.3"  // INTEGER
	StrOidIpRouteMetric2 = ".1.3.6.1.2.1.4.21.1.4"  // INTEGER
	StrOidIpRouteMetric3 = ".1.3.6.1.2.1.4.21.1.5"  // INTEGER
	StrOidIpRouteMetric4 = ".1.3.6.1.2.1.4.21.1.6"  // INTEGER
	StrOidIpRouteNextHop = ".1.3.6.1.2.1.4.21.1.7"  // IPADDRESS
	StrOidIpRouteType    = ".1.3.6.1.2.1.4.21.1.8"  // INTEGER
	StrOidIpRouteProto   = ".1.3.6.1.2.1.4.21.1.9"  // INTEGER
	StrOidIpRouteAge     = ".1.3.6.1.2.1.4.21.1.10" // INTEGER
	StrOidIpRouteMask    = ".1.3.6.1.2.1.4.21.1.11" // IPADDRESS
	StrOidIpRouteMetric5 = ".1.3.6.1.2.1.4.21.1.12" // INTEGER
	StrOidIpRouteInfo    = ".1.3.6.1.2.1.4.21.1.13" // OBJECT IDENTIFIER

	StrOidIfName = ".1.3.6.1.2.1.31.1.1.1.1" // OCTET STRING or DISPLAYSTRING

	/* L4 Port */
	StrOidTcpConnState = ".1.3.6.1.2.1.6.13.1.1" // INTEGER {close(1), listen(2), synSent(3), synReceived(4), finWait1(6), finWait2(7), closeWait(8), lastAck(9), closing(10), timeWait(11) }
	StrOidUdpPort      = ".1.3.6.1.2.1.7.5.1.2"  // INTEGER

	/* IP Address Entry */
	StrOidIpAddr    = ".1.3.6.1.2.1.4.20.1.1" // IPADDRESS
	StrOidIpIfIndex = ".1.3.6.1.2.1.4.20.1.2" // INTEGER
	StrOidIpMask    = ".1.3.6.1.2.1.4.20.1.3" // IPADDRESS
)

var (
	/* 3. MAP { OID Type : OID String } */
	oidMap = map[OidType]string{
		/* System */
		TypeOidSysDescr:    StrOidSysDescr,
		TypeOidSysUptime:   StrOidSysUptime,
		TypeOidSysHostname: StrOidSysHostname,

		/* Memory */
		TypeOidTotalMemory:    StrOidTotalMemory,
		TypeOidAvailMemory:    StrOidAvailMemory,
		TypeOidTotalSwap:      StrOidTotalSwap,
		TypeOidAvailSwap:      StrOidAvailSwap,
		TypeOidTotalSharedMem: StrOidTotalSharedMem,
		TypeOidTotalBufferMem: StrOidTotalBufferMem,
		TypeOidTotalCacheMem:  StrOidTotalCacheMem,

		/* Cpu Load */
		TypeOidCpuIdle:    StrOidCpuIdle,
		TypeOidCpuMin1Av:  StrOidCpuMin1Av,
		TypeOidCpuMin5Av:  StrOidCpuMin5Av,
		TypeOidCpuMin10Av: StrOidCpuMin10Av,

		/* Interface */
		TypeOidIfNumber:         StrOidIfNumber,
		TypeOidIfIndex:          StrOidIfIndex,
		TypeOidIfDescr:          StrOidIfDescr,
		TypeOidIfType:           StrOidIfType,
		TypeOidIfMtu:            StrOidIfMtu,
		TypeOidIfSpeed:          StrOidIfSpeed,
		TypeOidIfPhysAddr:       StrOidIfPhysAddr,
		TypeOidIfAdminStatus:    StrOidIfAdminStatus,
		TypeOidIfOperStatus:     StrOidIfOperStatus,
		TypeOidIfLastChange:     StrOidIfLastChange,
		TypeOidIfInOctets:       StrOidIfInOctets,
		TypeOidIfInUcastPkts:    StrOidIfInUcastPkts,
		TypeOidIfInNUcastPkts:   StrOidIfInNUcastPkts,
		TypeOidIfInDiscards:     StrOidIfInDiscards,
		TypeOidIfInErrors:       StrOidIfInErrors,
		TypeOidIfInUnknownProto: StrOidIfInUnknownProto,
		TypeOidIfOutOctets:      StrOidIfOutOctets,
		TypeOidIfOutUcastPkts:   StrOidIfOutUcastPkts,
		TypeOidIfOutNUcastPkts:  StrOidIfOutNUcastPkts,
		TypeOidIfOutDiscards:    StrOidIfOutDiscards,
		TypeOidIfOutErrors:      StrOidIfOutErrors,
		TypeOidIfOutQLen:        StrOidIfOutQLen,
		TypeOidIfSpecific:       StrOidIfSpecific,
		TypeOidIfName:           StrOidIfName,

		/* L4 Port */
		TypeOidTcpConnState: StrOidTcpConnState,
		TypeOidUdpPort:      StrOidUdpPort,

		/* IP Address Entry */
		TypeOidIpAddr:    StrOidIpAddr,
		TypeOidIpIfIndex: StrOidIpIfIndex,
		TypeOidIpMask:    StrOidIpMask,

		/* IP Route	Table */
		TypeOidIpRouteDest:    StrOidIpRouteDest,
		TypeOidIpRouteIfIndex: StrOidIpRouteIfIndex,
		TypeOidIpRouteMetric1: StrOidIpRouteMetric1,
		TypeOidIpRouteMetric2: StrOidIpRouteMetric2,
		TypeOidIpRouteMetric3: StrOidIpRouteMetric3,
		TypeOidIpRouteMetric4: StrOidIpRouteMetric4,
		TypeOidIpRouteNextHop: StrOidIpRouteNextHop,
		TypeOidIpRouteType:    StrOidIpRouteType,
		TypeOidIpRouteProto:   StrOidIpRouteProto,
		TypeOidIpRouteAge:     StrOidIpRouteAge,
		TypeOidIpRouteMask:    StrOidIpRouteMask,
		TypeOidIpRouteMetric5: StrOidIpRouteMetric5,
		TypeOidIpRouteInfo:    StrOidIpRouteInfo,
	}

	/* 4. MAP { OID Type : OID Description } */
	oidDescMap = map[OidType]string{
		/* System */
		TypeOidSysDescr:    "SysDescr",
		TypeOidSysUptime:   "SysUptime",
		TypeOidSysHostname: "SysHostname",

		/* Memory */
		TypeOidTotalMemory:    "TotalMemory",
		TypeOidAvailMemory:    "AvailMemory",
		TypeOidTotalSwap:      "TotalSwap",
		TypeOidAvailSwap:      "AvailSwap",
		TypeOidTotalSharedMem: "TotalSharedMem",
		TypeOidTotalBufferMem: "TotalBufferMem",
		TypeOidTotalCacheMem:  "TotalCacheMem",

		/* Cpu Load */
		TypeOidCpuIdle:    "CpuIdle",
		TypeOidCpuMin1Av:  "CpuMin1Av",
		TypeOidCpuMin5Av:  "CpuMin5Av",
		TypeOidCpuMin10Av: "CpuMin10Av",

		/* Interface */
		TypeOidIfNumber:         "IfNumber",
		TypeOidIfIndex:          "IfIndex",
		TypeOidIfDescr:          "IfDescr",
		TypeOidIfType:           "IfType",
		TypeOidIfMtu:            "IfMtu",
		TypeOidIfSpeed:          "IfSpeed",
		TypeOidIfPhysAddr:       "IfPhysAddress",
		TypeOidIfAdminStatus:    "IfAdminStatus",
		TypeOidIfOperStatus:     "IfOperStatus",
		TypeOidIfLastChange:     "IfLastChange",
		TypeOidIfInOctets:       "IfInOctets",
		TypeOidIfInUcastPkts:    "IfInUcastPkts",
		TypeOidIfInNUcastPkts:   "IfInNUcastPkts",
		TypeOidIfInDiscards:     "IfInDiscards",
		TypeOidIfInErrors:       "IfInErrors",
		TypeOidIfInUnknownProto: "IfInUnknownProtos",
		TypeOidIfOutOctets:      "IfOutOctets",
		TypeOidIfOutUcastPkts:   "IfOutUcastPkts",
		TypeOidIfOutNUcastPkts:  "IfOutNUcastPkts",
		TypeOidIfOutDiscards:    "IfOutDiscards",
		TypeOidIfOutErrors:      "IfOutErrors",
		TypeOidIfOutQLen:        "IfOutQLen",
		TypeOidIfSpecific:       "IfSpecific",
		TypeOidIfName:           "IfName",

		/* L4 Port */
		TypeOidTcpConnState: "TcpConnState",
		TypeOidUdpPort:      "TcpUdpPort",

		/* IP Address Entry */
		TypeOidIpAddr:    "IpAddr",
		TypeOidIpIfIndex: "IpIfIndex",
		TypeOidIpMask:    "IpMask",
	}
)
