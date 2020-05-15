package snmpapi

import (
	"fmt"
	"testing"
)

func TestDevice01(t *testing.T) {
	devices := make([]SnmpDevice, 2)

	devices[0].Device.Ip = "121.156.65.139"
	devices[0].Device.SnmpCommunity = "cmpService"

	devices[1].Device.Ip = "192.168.56.10"
	devices[1].Device.SnmpCommunity = "cmpService"

	ProcessSnmpAllDevice(devices)
}

func TestDevice02(t *testing.T) {
	devices := make([]SnmpDevice, 1)

	devices[0].Device.Ip = "121.156.65.139"
	devices[0].Device.SnmpCommunity = "cmpService"

	ProcessSnmpAllDevice(devices)
}

func TestDevice03(t *testing.T) {
	devices := make([]SnmpDevice, 1)

	devices[0].Device.Ip = "192.168.56.10"
	devices[0].Device.SnmpCommunity = "cmpService"

	fmt.Printf("GOMAXPROCS: %d\n", getGOMAXPROCS())

	ProcessSnmpAllDevice(devices)
}
