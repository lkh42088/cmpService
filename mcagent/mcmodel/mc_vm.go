package mcmodel

type VmEntry struct {
	Idx   int    `json:"idx"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Cpu   string `json:"cpu"`
	Ram   string `json:"ram"`
	Hdd   string `json:"hdd"`
}
