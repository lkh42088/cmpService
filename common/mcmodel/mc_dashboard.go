package mcmodel

type DeviceCount struct {
	Total			int		`json:"total"`
	Operate			int		`json:"operate"`
	Vm				int		`json:"vm"`
}

type DevicePlatform struct {
	Count 			int		`gorm:"column:count" json:"count"`
	ModelName		string	`gorm:"column:cpu_model" json:"modelName"`
}

type DeviceOsInfo struct {
	Count 			int 	`gorm:"column:count" json:"count"`
	OS 				string	`gorm:"column:vm_os" json:"os"`
}

type DeviceInfoForAdmin struct {
	Count		[]DeviceCount		`json:"count"`
	Platform    []DevicePlatform	`json:"platform"`
	OsInfo		[]DeviceOsInfo		`json:"osInfo"`
}

const TOP_USAGE_COUNT = 5

type DeviceRank struct {
	Cpu     []CpuStatForRank  `json:"cpu"`
	Mem     []MemStatForRank  `json:"mem"`
	Disk    []DiskStatForRank `json:"disk"`
	Traffic []VmIfStatForRank `json:"traffic"`
}
