package mcmodel

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
	variant     string `json:"variant"` // os : window10
	Name        string `json:"name"`    // image : window10-250
	Hdd         int    `json:"hdd"`
	Desc        string `json:"desc"`
}

type MgoServer struct {
	McServerIdx int `json:"serverIdx"`
}
