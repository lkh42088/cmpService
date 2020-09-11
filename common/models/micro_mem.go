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
	Time           time.Time   `json:"time"`
	AvailableBytes json.Number `json:"availableBytes"`
	Total          json.Number `json:"total"`
}
