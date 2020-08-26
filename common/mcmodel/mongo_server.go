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
	Desc          string `json:"desc"`
	OS            string `json:"os"`       // OS: windows10
	Image         string `json:"image"`    // Image: windows10-250
	Filename      string `json:"filename"` // Filename: windows10-250-1.qcow2
	FullName      string `json:"fullName"`
	Network       string `json:"network"`
	IpAddr        string `json:"ipAddr"`
	Mac           string `json:"mac"`
	ConfigStatus  string `json:"configStatus"`
	CurrentStatus string `json:"currentStatus"`
	VmNumber      int    `json:"-"` // VmNumber: 1
	RemoteAddr    string `json:"remoteAddr"`
	IsCreated     bool   `json:"isCreated"`
	IsProcess     bool   `json:"isProcess"`
}

func (v *MgoVm) Dump() string {
	pretty, _ := json.MarshalIndent(v, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func DumpVmList(list []MgoVm) {
	pretty, _ := json.MarshalIndent(list, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("VM List: %d\n", len(list))
	fmt.Printf("%s\n", string(pretty))
}

func DumpNetworkList(list []MgoNetwork) {
	pretty, _ := json.MarshalIndent(list, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("Network List: %d\n", len(list))
	fmt.Printf("%s\n", string(pretty))
}

func DumpImageList(list []MgoImage) {
	pretty, _ := json.MarshalIndent(list, "", "  ")
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Printf("image List: %d\n", len(list))
	fmt.Printf("%s\n", string(pretty))
}

// flavor
type MgoImage struct {
	Id       uint   `json:"id"`
	Variant  string `json:"variant"` // os : win10
	Name     string `json:"name"`    // image : windows10-250G
	Hdd      int    `json:"hdd"`
	Desc     string `json:"desc"`
	FullName string `json:"fullName"`
}

func (n *MgoImage) Dump() string {
	pretty, _ := json.MarshalIndent(n, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

type MgoNetwork struct {
	Id        uint             `json:"id"`
	Uuid      string           `json:"uuid"`
	Name      string           `json:"name"`
	Bridge    string           `json:"bridge"`
	Mode      string           `json:"mode"`
	Mac       string           `json:"mac"`
	DhcpStart string           `json:"dhcpStart"`
	DhcpEnd   string           `json:"dhcpEnd"`
	Ip        string           `json:"ip"`
	Netmask   string           `json:"netmask"`
	Prefix    uint             `json:"prefix"`
	Host      []MgoNetworkHost `json:"host"`
}

type MgoNetworkHost struct {
	Id       uint   `json:"id"`
	Mac      string `json:"mac"`
	Ip       string `json:"ip"`
	Hostname string `json:"hostname"`
}

func (n *MgoNetwork) Dump() string {
	pretty, _ := json.MarshalIndent(n, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

type MgoServer struct {
	SerialNumber string        `json:"serialNumber"`
	Port         string        `json:"port"`
	Mac          string        `json:"mac"`
	Ip           string        `json:"ip"`
	PublicIp     string        `json:"publicIp"`
	Vms          *[]MgoVm      `json:"vms"`
	Networks     *[]MgoNetwork `json:"networks"`
	Images       *[]MgoImage   `json:"images"`
}

func (n *MgoServer) Dump() string {
	pretty, _ := json.MarshalIndent(n, "", "  ")

	fmt.Printf("%s\n", string(pretty))
	return string(pretty)
}

func (n *MgoServer) DumpSummary() {
	vmCount := 0
	netCount := 0
	imgCount := 0
	if n.Vms != nil {
		vmCount = len(*n.Vms)
	}
	if n.Networks != nil {
		netCount = len(*n.Networks)
	}
	if n.Images != nil {
		imgCount = len(*n.Images)
	}
	fmt.Printf("server(%s/%s, %s/%s), vm(%d), network(%d), image(%d)\n",
		n.Ip,
		n.PublicIp,
		n.Port,
		n.Mac,
		vmCount, netCount, imgCount)
}

func (s *MgoServer) Compare(n MgoServer) bool {
	isChanged := false
	return isChanged
}
