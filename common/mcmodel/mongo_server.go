package mcmodel

const (
	McVmStatusCopyImage = "copy vm image"
	McVmStatusCreateVm  = "create vm instance"
	McVmStatusRunning   = "running"
	McVmStatusShutdown  = "shutdown"
)

//
//type MgoVm struct {
//	Idx            uint   `json:"idx"`
//	McServerIdx    int    `json:"serverIdx"`
//	CompanyIdx     int    `json:"cpIdx"`
//	Name           string `json:"name"`
//	Cpu            int    `json:"cpu"`
//	Ram            int    `json:"ram"`
//	Hdd            int    `json:"hdd"`
//	Desc           string `json:"desc"`
//	OS             string `json:"os"`       // OS: windows10
//	Image          string `json:"image"`    // Image: windows10-250
//	Filename       string `json:"filename"` // Filename: windows10-250-1.qcow2
//	VmIndex        int    `json:"vmIndex"`  // VmIndex: 1
//	FullPath       string `json:"fullPath"`
//	Network        string `json:"network"`
//	IpAddr         string `json:"ipAddr"`
//	IsChangeIpAddr bool   `json:"-"`
//	VncPort        string `json:"vncPort"`
//	Mac            string `json:"mac"`
//	ConfigStatus   string `json:"configStatus"`
//	CurrentStatus  string `json:"currentStatus"`
//	RemoteAddr     string `json:"remoteAddr"`
//	IsCreated      bool   `json:"isCreated"`
//	IsProcess      bool   `json:"isProcess"`
//}
//
//func (v *MgoVm) Dump() string {
//	pretty, _ := json.MarshalIndent(v, "", "  ")
//	fmt.Printf("%s\n", string(pretty))
//	return string(pretty)
//}

// flavor
//type MgoImage struct {
//	Idx      uint   `json:"id"`
//	Variant  string `json:"variant"` // os : win10
//	Name     string `json:"name"`    // image : windows10-250G
//	Hdd      int    `json:"hdd"`
//	Desc     string `json:"desc"`
//	FullName string `json:"fullName"`
//}
//
//func (n *MgoImage) Dump() string {
//	pretty, _ := json.MarshalIndent(n, "", "  ")
//
//	fmt.Printf("%s\n", string(pretty))
//	return string(pretty)
//}
//
//type MgoNetwork struct {
//	Idx       uint             `json:"id"`
//	Uuid      string           `json:"uuid"`
//	Name      string           `json:"name"`
//	Bridge    string           `json:"bridge"`
//	Mode      string           `json:"mode"`
//	Mac       string           `json:"mac"`
//	DhcpStart string           `json:"dhcpStart"`
//	DhcpEnd   string           `json:"dhcpEnd"`
//	Ip        string           `json:"ip"`
//	Netmask   string           `json:"netmask"`
//	Prefix    uint             `json:"prefix"`
//	Host      []MgoNetworkHost `json:"host"`
//}
//
//type MgoNetworkHost struct {
//	Idx      uint   `json:"id"`
//	Mac      string `json:"mac"`
//	Ip       string `json:"ip"`
//	Hostname string `json:"hostname"`
//}
//
//func (n *MgoNetwork) Dump() string {
//	pretty, _ := json.MarshalIndent(n, "", "  ")
//
//	fmt.Printf("%s\n", string(pretty))
//	return string(pretty)
//}
//
//type MgoServer struct {
//	SerialNumber string        `json:"serialNumber"`
//	Port         string        `json:"port"`
//	Mac          string        `json:"mac"`
//	Ip           string        `json:"ip"`
//	PublicIp     string        `json:"publicIp"`
//	Vms          *[]McVm       `json:"vms"`
//	Networks     *[]MgoNetwork `json:"networks"`
//	Images       *[]MgoImage   `json:"images"`
//}
//
//func (n *MgoServer) Dump() string {
//	pretty, _ := json.MarshalIndent(n, "", "  ")
//
//	fmt.Printf("%s\n", string(pretty))
//	return string(pretty)
//}
//
//func (n *MgoServer) DumpSummary() {
//	vmCount := 0
//	netCount := 0
//	imgCount := 0
//	if n.Vms != nil {
//		vmCount = len(*n.Vms)
//	}
//	if n.Networks != nil {
//		netCount = len(*n.Networks)
//	}
//	if n.Images != nil {
//		imgCount = len(*n.Images)
//	}
//	if n.SerialNumber == "" {
//		fmt.Printf("SN(none), ")
//	} else {
//		fmt.Printf("SN(%s), ", n.SerialNumber)
//	}
//	fmt.Printf("server(%s/%s, %s/%s), vm(%d), network(%d), image(%d)\n",
//		n.Ip,
//		n.PublicIp,
//		n.Port,
//		n.Mac,
//		vmCount, netCount, imgCount)
//}

//func LookupMcVm(list *[]McVm, target McVm) *McVm {
//	if list == nil {
//		return nil
//	}
//	for _, obj := range *list {
//		if obj.Name == target.Name {
//			return &obj
//		}
//	}
//	return nil
//}
