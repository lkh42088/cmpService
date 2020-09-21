package models

import (
	"time"
)

type DiskStat struct {
	Time        time.Time `json:"time"`
	Device      string    `json:"device"`
	Fstype      string    `json:"fstype"`
	Path        string    `json:"path"`
	Total       float64   `json:"total"`
	Used        float64   `json:"used"`
	UsedPercent float64   `json:"used_percent"`
	Err         string    `json:"err"`
}

type WinDiskStat struct {
	Time          time.Time `json:"time"`
	FreeMegabytes float64   `json:"freeMegabytes"`
	Total         float64   `json:"total"`
}
