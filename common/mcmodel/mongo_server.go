package mcmodel

import (
	"encoding/json"
	"fmt"
)

const (
	McVmStatusCopyImage = "copy vm image"
	McVmStatusCreateVm  = "create vm instance"
	McVmStatusRunning   = "running"
	McVmStatusShutdown  = "shutdown"
)

type MgoVm struct {
	Idx           uint   `json:"idx"`
	McServerIdx   int    `json:"serverIdx"`
	CompanyIdx    int    `json:"cpIdx"`
	Name          string `json:"name"`
	Cpu           int    `json:"cpu"`
	Ram           int    `json:"ram"`
	Hdd           int    `json:"hdd"`
	OS            string `json:"os"`       // OS: windows10
	Image         string `json:"image"`    // Image: windows10-250
	Filename      string `json:"filename"` // Filename: windows10-250-1.qcow2
	Network       string `json:"network"`
	IpAddr        string `json:"ipAddr"`
	Mac           string `json:"mac"`
	ConfigStatus  string `json:"configStatus"`
	CurrentStatus string `json:"currentStatus"`
	VmNumber      int    `json:"-"` // VmNumber: 1
	IsCreated     bool   `json:"-"`
}

func (v *MgoVm) Dump() string {
	pretty, _:= json.MarshalIndent(v, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

// flavor
type MgoImage struct {
	Id      uint   `json:"id"`
	Variant string `json:"variant"` // os : win10
	Name    string `json:"name"`    // image : windows10-250G
	Hdd     int    `json:"hdd"`
	Desc    string `json:"desc"`
}

type MgoNetwork struct {
	Id      uint   `json:"id"`
	Uuid    string `json:"uuid"`
	Name    string `json:"name"`
	Bridge  string `json:"bridge"`
	Mode    string `json:"mode"`
	Ip      string `json:"ip"`
	Netmask string `json:"netmask"`
	Prefix  uint   `json:"prefix"`
}

type MgoServer struct {
	Port     string        `json:"port"`
	Mac      string        `json:"mac"`
	Ip       string        `json:"ip"`
	Networks *[]MgoNetwork `json:"networks"`
	Images   *[]MgoImage   `json:"images"`
}
