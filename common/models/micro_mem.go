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
