package mcmodel

import "encoding/xml"

const (
	McVmStatusCopyImage="copy vm image"
	McVmStatusCreateVm="create vm instance"
	McVmStatusRunning="running"
	McVmStatusShutdown="shutdown"
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
	VmNumber      int    `json:"-"`        // VmNumber: 1
}

// flavor
type MgoImage struct {
	Id          uint   `json:"id"`
	McServerIdx int    `json:"serverIdx"`
	Variant     string `json:"variant"` // os : win10
	Name        string `json:"name"`    // image : windows10-250G
	Hdd         int    `json:"hdd"`
	Desc        string `json:"desc"`
}

type MgoNetwork struct {
	Id uint `json:"id"`
	McServerIdx int `json:"serverIdx"`
	Name string `json:"name"`
	Mode string `json:"mode"`
	Subnet string `json:"subnet"`
}

type XmlNetwork struct {
	XMLName xml.Name `xml:"network"`
	Name string `xml:"name"`
	Uuid string `xml:"uuid"`
	Forward XmlForward `xml:"forward"`
	Bridge XmlBridge `xml:"bridge"`
}

type XmlForward struct {
	Mode string `xml:"mode,attr"`
}

type XmlBridge struct {
	Name string `xml:"name,attr"`
	Stp string `xml:"stp,attr"`
	Delay string `xml:"delay,attr"`
}

type XmlIp struct {
	Address string `xml:"address,attr"`
	Netmask string `xml:"netmask,attr"`
}

type XmlDhcpRange struct {
	Start string `xml:"start,attr"`
	End string `xml:"end,attr"`
}

type MgoServer struct {
	McServerIdx int `json:"serverIdx"`
}
