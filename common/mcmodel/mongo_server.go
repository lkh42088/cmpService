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
	Idx            uint   `json:"idx"`
	McServerIdx    int    `json:"serverIdx"`
	CompanyIdx     int    `json:"cpIdx"`
	Name           string `json:"name"`
	Cpu            int    `json:"cpu"`
	Ram            int    `json:"ram"`
	Hdd            int    `json:"hdd"`
	Desc           string `json:"desc"`
	OS             string `json:"os"`       // OS: windows10
	Image          string `json:"image"`    // Image: windows10-250
	Filename       string `json:"filename"` // Filename: windows10-250-1.qcow2
	VmIndex        int    `json:"vmIndex"`  // VmIndex: 1
	FullPath       string `json:"fullPath"`
	Network        string `json:"network"`
	IpAddr         string `json:"ipAddr"`
	IsChangeIpAddr bool   `json:"-"`
	Mac            string `json:"mac"`
	ConfigStatus   string `json:"configStatus"`
	CurrentStatus  string `json:"currentStatus"`
	RemoteAddr     string `json:"remoteAddr"`
	IsCreated      bool   `json:"isCreated"`
	IsProcess      bool   `json:"isProcess"`
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
	if n.SerialNumber == "" {
		fmt.Printf("SN(none), ")
	} else {
		fmt.Printf("SN(%s), ", n.SerialNumber)
	}
	fmt.Printf("server(%s/%s, %s/%s), vm(%d), network(%d), image(%d)\n",
		n.Ip,
		n.PublicIp,
		n.Port,
		n.Mac,
		vmCount, netCount, imgCount)
}

func (o *MgoVm) Compare(n *MgoVm) bool {
	if o.Name != n.Name {
		return true
	}
	if o.Cpu != n.Cpu {
		return true
	}
	if o.Ram != n.Ram {
		return true
	}
	if o.Hdd != n.Hdd {
		return true
	}
	if o.Desc != n.Desc {
		return true
	}
	if o.OS != n.OS {
		return true
	}
	if o.Image != n.Image {
		return true
	}
	if o.Filename != n.Filename {
		return true
	}
	if o.VmIndex != n.VmIndex {
		return true
	}
	if o.FullPath != n.FullPath {
		return true
	}
	if o.IpAddr != n.IpAddr {
		n.IsChangeIpAddr = true
		return true
	}
	if o.Mac != n.Mac {
		return true
	}
	if o.CurrentStatus != n.CurrentStatus {
		return true
	}
	if o.RemoteAddr != n.RemoteAddr {
		return true
	}
	return false
}

func (v MgoNetwork) Compare(n MgoNetwork) bool {
	if v.Name != n.Name {
		return true
	}
	if v.Bridge != n.Bridge {
		return true
	}
	if v.Mode != n.Mode {
		return true
	}
	if v.Mac != n.Mac {
		return true
	}
	if v.DhcpStart != n.DhcpStart {
		return true
	}
	if v.DhcpEnd != n.DhcpEnd {
		return true
	}
	if v.Ip != n.Ip {
		return true
	}
	if v.Netmask != n.Netmask {
		return true
	}
	if v.Prefix != n.Prefix {
		return true
	}
	return false
}

func (v MgoImage) Compare(n MgoImage) bool {
	if v.Name != n.Name {
		return true
	}
	if v.Variant != n.Variant {
		return true
	}
	if v.Hdd != n.Hdd {
		return true
	}
	if v.Desc != n.Desc {
		return true
	}
	if v.FullName != n.FullName {
		return true
	}
	return false
}

func (o *MgoServer) Compare(n *MgoServer) bool {
	isChanged := false

	if o.Vms != nil {
		if n.Vms == nil {
			isChanged = true
		} else if len(*(o.Vms)) != len(*(n.Vms)) {
			isChanged = true
		} else {
			for _, obj1 := range *o.Vms {
				obj2 := LookupMgoVm(n.Vms, obj1)
				if obj2 == nil {
					isChanged = true
				} else {
					res := obj1.Compare(obj2)
					if res == true {
						isChanged = true
					}
				}
			}
		}
	} else {
		if n.Vms != nil {
			isChanged = true
		}
	}

	if o.Networks != nil {
		if n.Networks == nil {
			isChanged = true
		} else if len(*(o.Networks)) != len(*(n.Networks)) {
			isChanged = true
		} else {
			for _, obj1 := range *o.Networks {
				obj2 := LookupMgoNetwork(n.Networks, obj1)
				if obj2 == nil {
					isChanged = true
				} else {
					res := obj1.Compare(*obj2)
					if res == true {
						isChanged = true
					}
				}
			}
		}
	} else {
		if n.Networks != nil {
			isChanged = true
		}
	}

	if o.Images != nil {
		if n.Images == nil {
			isChanged = true
	 	} else if len(*(o.Images)) != len(*(n.Images)) {
			isChanged = true
		} else {
			for _, obj1 := range *o.Images {
				obj2 := LookupMgoImage(n.Images, obj1)
				if obj2 == nil {
					isChanged = true
				} else {
					res := obj1.Compare(*obj2)
					if res == true {
						isChanged = true
					}
				}
			}
		}
	} else {
		if n.Images != nil {
			isChanged = true
		}
	}

	return isChanged
}

//----------------------------------------------------------------------

func (v McVm) Compare(n McVm) bool {
	if v.Name != n.Name {
		return true
	}
	if v.Cpu != n.Cpu {
		return true
	}
	if v.Ram != n.Ram {
		return true
	}
	if v.Hdd != n.Hdd {
		return true
	}
	if v.Desc != n.Desc {
		return true
	}
	if v.OS != n.OS {
		return true
	}
	if v.Image != n.Image {
		return true
	}
	if v.Filename != n.Filename {
		return true
	}
	if v.VmIndex != n.VmIndex {
		return true
	}
	if v.FullPath != n.FullPath {
		return true
	}
	if v.IpAddr != n.IpAddr {
		return true
	}
	if v.Mac != n.Mac {
		return true
	}
	if v.CurrentStatus != n.CurrentStatus {
		return true
	}
	if v.RemoteAddr != n.RemoteAddr {
		return true
	}
	return false
}

func (v McNetworks) Compare(n McNetworks) bool {
	if v.Name != n.Name {
		return true
	}
	if v.Bridge != n.Bridge {
		return true
	}
	if v.Mode != n.Mode {
		return true
	}
	if v.Mac != n.Mac {
		return true
	}
	if v.DhcpStart != n.DhcpStart {
		return true
	}
	if v.DhcpEnd != n.DhcpEnd {
		return true
	}
	if v.Ip != n.Ip {
		return true
	}
	if v.Netmask != n.Netmask {
		return true
	}
	if v.Prefix != n.Prefix {
		return true
	}
	return false
}

func (v McImages) Compare(n McImages) bool {
	if v.Name != n.Name {
		return true
	}
	if v.Variant != n.Variant {
		return true
	}
	if v.Hdd != n.Hdd {
		return true
	}
	if v.Desc != n.Desc {
		return true
	}
	if v.FullName != n.FullName {
		return true
	}
	return false
}

func (s *McServerMsg) Compare(n *McServerMsg) bool {
	isChanged := false

	if s.Vms != nil {
		if n.Vms == nil {
			return true
		}
		if len(*(s.Vms)) != len(*(n.Vms)) {
			return true
		}
		for _, obj1 := range *s.Vms {
			obj2 := LookupVm(n.Vms, obj1)
			if obj2 == nil {
				return true
			}
			res := obj1.Compare(*obj2)
			if res == true {
				return res
			}
		}
	} else {
		if n.Vms != nil {
			return true
		}
	}

	if s.Networks != nil {
		if n.Networks == nil {
			return true
		}
		if len(*(s.Networks)) != len(*(n.Networks)) {
			return true
		}
		for _, obj1 := range *s.Networks {
			obj2 := LookupNetwork(n.Networks, obj1)
			if obj2 == nil {
				return true
			}
			res := obj1.Compare(*obj2)
			if res == true {
				return res
			}
		}
	} else {
		if n.Networks != nil {
			return true
		}
	}

	if s.Images != nil {
		if n.Images == nil {
			return true
		}
		if len(*(s.Images)) != len(*(n.Images)) {
			return true
		}
		for _, obj1 := range *s.Images {
			obj2 := LookupImage(n.Images, obj1)
			if obj2 == nil {
				return true
			}
			res := obj1.Compare(*obj2)
			if res == true {
				return res
			}
		}
	} else {
		if n.Images != nil {
			return true
		}
	}

	return isChanged
}
