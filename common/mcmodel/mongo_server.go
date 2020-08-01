package mcmodel

type MgoVm struct {
	Idx         uint   `json:"idx"`
	McServerIds int    `json:"serverIdx"`
	CompanyIdx  int    `json:"cpIdx"`
	Name        string `json:"name"`
	Cpu         int    `json:"cpu"`
	Ram         int    `json:"ram"`
	Hdd         int    `json:"hdd"`
	OS          string `json:"os"`
	Image       string `json:"image"`
	Network     string `json:"network"`
	IpAddr      string `json:"ipAddr"`
}

type MgoServer struct {
	Idx uint `json:`
}
