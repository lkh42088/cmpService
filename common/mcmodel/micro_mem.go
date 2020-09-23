package mcmodel

import (
	"time"
)

type MemStat struct {
	Time             time.Time `json:"time"`
	Available        float64   `json:"available"`
	AvailablePercent float64   `json:"available_percent"`
	Total            float64   `json:"total"`
	Err              string    `json:"err"`
}

type WinMemStat struct {
	Time           time.Time `json:"time"`
	AvailableBytes float64   `json:"availableBytes"`
	Total          float64   `json:"total"`
}
