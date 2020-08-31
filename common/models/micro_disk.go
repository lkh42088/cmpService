package models

import (
	"encoding/json"
	"time"
)

type DiskStat struct {
	Time        time.Time   `json:"time"`
	Device      string      `json:"device"`
	Fstype      string      `json:"fstype"`
	Path        string      `json:"path"`
	Total       json.Number `json:"total"`
	Used        json.Number `json:"used"`
	UsedPercent json.Number `json:"used_percent"`
}
