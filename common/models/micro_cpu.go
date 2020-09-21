package models

import (
	"time"
)

type CpuStat struct {
	Time      time.Time `json:"time"`
	Cpu       string    `json:"cpu"`
	UsageIdle float64   `json:"usage_idle"`
	Err       string    `json:"err"`
}

type WinCpuStat struct {
	Time            time.Time `json:"time"`
	PercentIdleTime float64   `json:"percentIdleTime"`
	Total           float64   `json:"total"`
}
