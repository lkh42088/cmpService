package snmpapi

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"cmpService/collector/config"
	"cmpService/collector/mongo"
	"cmpService/common/lib"
	"runtime"
	"strings"
	"sync"
	"time"
)

type System struct {
	desc     string
	uptime   int64
	hostname string
}

var RegularCollectTime = 5 * time.Second
var snmpDevTbMutex sync.Mutex

func getGOMAXPROCS() int {

	// Default GOMAXPROCS : The number of Server Cores.
	return runtime.GOMAXPROCS(0)
}

func GetDevicesFromMongoDB(snmpdevtable *SnmpDeviceTable) {

	if mongo.Mongo.Session == nil {
		lib.LogWarn("Mongo session is NULL!\n")
		return
	}

	devices, _ := mongo.Mongo.GetAll()
	for _, dev := range devices {
		snmpdev := SnmpDevice{}
		snmpdev.Device = dev
		lib.LogWarnln("From MongoDB: ", dev)
		snmpdevtable.Post(snmpdev)
	}
	lib.LogWarn("Total: count %d, collectdevice slice %d\n", SnmpDevTb.Count, len(SnmpDevTb.Devices))
}

func ApplyChangeThings() {
	// Snmp Devices
}

func InitConfig() {
	// Snmpdev Table
	SetSnmpDevTb(NewSnmpDeviceTable())

	// InfluxDb
	config.ConfigureInfluxDB()

	// MongoDb
	config.ConfigureMongoDB()
	GetDevicesFromMongoDB(SnmpDevTb)
}

func Start(parentwg *sync.WaitGroup) {

	for {
		ApplyChangeThings()
		time.Sleep(RegularCollectTime)
		CollectSnmpInfo()
	}

	if parentwg != nil {
		parentwg.Done()
	}
}

func CollectSnmpInfo() {
	devNum := len(SnmpDevTb.Devices)
	if devNum == 0 {
		lib.LogInfoln("It does not exist collectdevice!")
		return
	}

	//common.LogInfo("deviceList: %d\n", len(SnmpDevTb.devices))

	var wg sync.WaitGroup

	// sync Add
	wg.Add(devNum)

	for _, device := range SnmpDevTb.Devices {
		if "" == device.Device.Ip {

			// sync Delete
			wg.Add(-1)

			lib.LogInfo("go func --> %s %s\n", device.Device.Ip, "skip!!")
			continue
		}
		go func(device SnmpDevice) {
			dev := &device

			// sync Done
			defer wg.Done()

			if dev.Snmp == nil {
				dev.InitDeviceSnmp()
			}
			lib.LogInfo("go func - SNMP Connect %s %d %s\n", dev.Device.Ip, dev.Device.Port, dev.Device.SnmpCommunity)
			err := dev.Snmp.Connect()
			defer dev.Snmp.Conn.Close()
			if err != nil {
				lib.LogWarn("[collectdevice %s, comm %s] snmp connect: %s\n",
					dev.Snmp.Target, dev.Snmp.Community, err)
			} else {
				lib.LogInfoln("SNMP Get start", dev.Device.Ip)
				getDeviceSnmpInfo(dev)
				snmpDevTbMutex.Lock()
				SnmpDevTb.Devices[dev.Device.GetIdString()] = *dev
				snmpDevTbMutex.Unlock()
			}
		}(device)
	}
	wg.Wait()

	//WriteMetric(SnmpDevTb)
	MakeBpFromSnmpStruct(SnmpDevTb)

	//SnmpDevTb.String()
	fmt.Printf("Total: count %d, collectdevice slice %d v0.9.5\n",
		SnmpDevTb.Count, len(SnmpDevTb.Devices))
}

func ProcessSnmpAllDevice(deviceList []SnmpDevice) {

	var wg sync.WaitGroup

	// sync Add
	wg.Add(len(deviceList))

	lib.LogInfo("deviceList: %d", len(deviceList))
	for _, device := range deviceList {
		if "" == device.Device.Ip {

			// sync Delete
			wg.Add(-1)

			lib.LogInfo("go func --> ip address is null! %s %s", device.Device.Ip, "skip!!")
			continue
		}
		go func(device SnmpDevice) {
			dev := &device
			lib.LogInfo("go func %s", dev.Device.Ip)

			// sync Done
			defer wg.Done()

			dev.InitDeviceSnmp()
			err := dev.Snmp.Connect()
			if err != nil {
				lib.LogWarn("%s %s : error %s --> skip\n",
					dev.Snmp.Target, dev.Snmp.Community, err)
			} else {
				defer dev.Snmp.Conn.Close()
				lib.LogInfo("%s %s\n", dev.Snmp.Target, dev.Snmp.Community)
				getDeviceSnmpInfo(dev)
			}
		}(device)
	}
	wg.Wait()
}

func (s *SnmpDevice) InitDeviceSnmp() {
	s.Snmp = &g.GoSNMP{
		Target:             s.Device.Ip,
		Port:               161,
		Transport:          "udp",
		Community:          s.Device.SnmpCommunity,
		Version:            g.Version2c,
		Timeout:            time.Duration(2) * time.Second,
		Retries:            1,
		ExponentialTimeout: true,
		MaxOids:            g.MaxOids,
	}
}

func getDeviceSnmpInfo(s *SnmpDevice) {
	// Get System
	err := s.getSystemFromSnmp()
	if err != nil {
		return
	}

	// Get Cpu
	s.getCpu()

	// Get Memory
	s.getMemory()

	// Get Interface Table
	s.getIfTable()

	// Get IP Table
	s.GetIpTable()

	// Get L4 Port
	s.getL4Port()

	// Get IP Routing Table
	s.GetIpRouteTable()

	// Print collectdevice information
	//s.String()
}

func (s *SnmpDevice) getIfTable() {
	s.IfTable.ifNumber = s.getIfNumber()
	if s.IfTable.ifNumber < 1 {
		lib.LogInfo("collectdevice %s : ifNumber %d --> skip!\n", s.Device.Ip, s.IfTable.ifNumber)
		return
	}

	// Make interface entry
	s.IfTable.ifEntry = make([]IfEntry, s.IfTable.ifNumber)

	// Get interface entry information through SNMP Protocol
	for i := TypeOidIfGetBulkBegin; i <= TypeOidIfGetBulkEnd; i++ {
		s.getIfEntryFromSnmp(OidType(i))
	}
}

func (s *SnmpDevice) getSystemFromSnmp() error {
	oids := []string{StrOidSysDescr, StrOidSysUptime, StrOidSysHostname}
	result, err := s.Snmp.Get(oids)
	if err != nil {
		lib.LogWarn("getSystemFromSnmpGet() : %v\n", err)
		return err
	}

	system := &s.System
	for i, variable := range result.Variables {
		lib.LogInfo("getSystemFromSnmp: [collectdevice %s, community %s] %d: oid: %s ",
			s.Device.Ip, s.Device.SnmpCommunity, i, variable.Name)

		switch variable.Type {
		case g.OctetString:
			if strings.Contains(variable.Name, StrOidSysDescr) {
				system.desc = string(variable.Value.([]byte))
				lib.LogInfo("System desc(%s)\n", system.desc)
			} else if strings.Contains(variable.Name, StrOidSysHostname) {
				system.hostname = string(variable.Value.([]byte))
				lib.LogInfo("System hostname(%s)\n", system.hostname)
			} else {
				lib.LogInfo("*** string: %s\n", string(variable.Value.([]byte)))
			}
		default:
			if strings.Contains(variable.Name, StrOidSysUptime) {
				system.uptime = g.ToBigInt(variable.Value).Int64()
				lib.LogInfo("System uptime(%d)\n", system.uptime)
			} else {
				number := g.ToBigInt(variable.Value).Int64()
				lib.LogInfo("*** if number: %d\n", number)
			}
		}
	}
	return nil
}

func (s *System) String() {
	fmt.Printf(" [System]\n")
	fmt.Printf("  - desc: %s\n", s.desc)
	fmt.Printf("  - hostname: %s\n", s.hostname)
	fmt.Printf("  - uptime: %d\n", s.uptime)
}

func convertByte2StringMac(mac []byte) string {
	if len(mac) < 1 {
		return ""
	}
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
}

func string2mac(str string) string {
	var arr [6]byte
	if len(str) < 0 {
		return ""
	}
	copy(arr[:], str)
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		arr[0], arr[1], arr[2], arr[3], arr[4], arr[5])
}
