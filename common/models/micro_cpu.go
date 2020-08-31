package models

import (
	"encoding/json"
	"time"
)

type CpuStat struct {
	Time      time.Time   `json:"time"`
	Cpu       string      `json:"cpu"`
	UsageIdle json.Number `json:"usage_idle"`
}
