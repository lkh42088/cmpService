package models

import (
	"encoding/json"
	"time"
)

type MemStat struct {
	Time             time.Time   `json:"time"`
	Available        json.Number `json:"available"`
	AvailablePercent json.Number `json:"available_percent"`
	Total            json.Number `json:"total"`
}

type WinMemStat struct {
	Time           time.Time `json:"time"`
	AvailableBytes float64   `json:"availableBytes"`
	Total          float64   `json:"total"`
}
