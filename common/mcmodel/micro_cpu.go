package mcmodel

import (
	"time"
)

type CpuStat struct {
	Time      time.Time `json:"time"`
	Cpu       string    `json:"cpu"`
	UsageIdle float64   `json:"usage_idle"`
	Err       string    `json:"err"`
}

type CpuStatForRank struct {
	Time      time.Time `json:"time"`
	SN 		  string	`json:"serial_number"`
	Cpu       string    `json:"cpu"`
	UsageIdle float64   `json:"usage_idle"`
}

type WinCpuStat struct {
	Time            time.Time `json:"time"`
	PercentIdleTime float64   `json:"percentIdleTime"`
	Total           float64   `json:"total"`
}
