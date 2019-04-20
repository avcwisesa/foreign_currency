package model

import (
	"time"
)

type ExchangeRateSummary struct {
	From string `json:"from"`
	To string 	`json:"to"`
	Latest float64 `json:"latest"`
	Avg float64 `json:"avg"`
}

type ExchangeRateSummaryPayload struct {
	Date time.Time `json:"from"`
	Summary []ExchangeRateSummary `json:"summary"`
}