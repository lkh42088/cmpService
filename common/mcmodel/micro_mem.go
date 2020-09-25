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

type MemStatForRank struct {
	Time             time.Time `json:"time"`
	SN 		    	 string	   `json:"serial_number"`
	Available        float64   `json:"available"`
	AvailablePercent float64   `json:"available_percent"`
	Total            float64   `json:"total"`
}

type WinMemStat struct {
	Time           time.Time `json:"time"`
	AvailableBytes float64   `json:"availableBytes"`
	Total          float64   `json:"total"`
}
