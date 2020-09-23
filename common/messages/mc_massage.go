package messages

type VmActionMsg struct {
	VmIdx    int `json:"idx"`
	VmAction int `json:"vmAction"`
}

type McVmActionMsg struct {
	VmName   string `json:"vmName"`
	VmAction int    `json:"vmAction"`
}

type SnapshotConfigMsg struct {
	ServerIdx uint   `json:"serverIdx"`
	VmName    string `json:"vmName"`
	Type      string `json:"type"`
	Days      string `json:"days"`
	Hours     string `json:"hours"`
	Minutes   string `json:"minutes"`
}

type SnapshotEntry struct {
	VmName   string `json:"vmName"`
	SnapName string `json:"snapName"`
}

type SnapshotEntryMsg struct {
	CompanyIdx   string           `json:"companyIdx"`
	SerialNumber string           `json:"serialNumber"`
	Entry        *[]SnapshotEntry `json:"snapEntry"`
}

type VmStatusActionMsg struct {
	ServerIdx uint   `json:"serverIdx"`
	VmName    string `json:"vmName"`
	Status    string `json:"status"` // start, stop, reset
}
