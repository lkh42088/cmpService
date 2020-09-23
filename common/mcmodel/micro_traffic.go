package mcmodel

import "time"

// Get VM Interface traffic
type VmIfStat struct {
	Time          time.Time `json:"time"`
	Hostname      string    `json:"hostname"`
	IfDescr       string    `json:"ifDescr"`
	IfPhysAddress string    `json:"ifPhysAddress"`
	IfInOctets    int64     `json:"ifInOctets"`
	IfOutOctets   int64     `json:"ifOutOctets"`
}

type WinVmIfStat struct {
	Time                time.Time `json:"time"`
	BytesReceivedPersec float64     `json:"bytesReceivedPersec"`
	BytesSentPersec     float64     `json:"bytesSentPersec"`
}

type VmIfStatistics struct {
	Stats []VmIfStat `json:"stats"`
}

type WinVmIfStatistics struct {
	Stats []WinVmIfStat `json:"stats"`
}

type Stats struct {
	Xaxis time.Time `json:"x"`
	Yaxis int64     `json:"y"`
}

type VmStatsSet struct {
	Id   string  `json:"id"`
	Data []Stats `json:"data"`
}

type VmStatsRsponse struct {
	Hostname string        `json:"hostname"`
	Stats    [2]VmStatsSet `json:"stats"`
}

type VmInfoStatsResponse struct {
	VmCpu  VmStatsSet `json:"cpu"`
	VmMem  VmStatsSet `json:"mem"`
	VmDisk VmStatsSet `json:"disk"`
	VmRx   VmStatsSet `json:"rx"`
	VmTx   VmStatsSet `json:"tx"`
}
