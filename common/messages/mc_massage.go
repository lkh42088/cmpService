package messages

type SnapshotConfigMsg struct {
	VmName  string `json:"vmName"`
	Type    string `json:"type"`
	Days    string `json:"days"`
	Hours   string `json:"hours"`
	Minutes string `json:"minutes"`
}

type SnapshotEntry struct {
	VmName       string `json:"vmName"`
	SnapName     string `json:"snapName"`
}

type SnapshotEntryMsg struct {
	CompanyIdx   string `json:"companyIdx"`
	SerialNumber string `json:"serialNumber"`
	Entry *[]SnapshotEntry `json:"snapEntry"`
}
