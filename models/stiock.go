package models

import "time"

type Stock struct {
	// Ticker del stock
	// Required: true
	// Default: BSBR
	Ticker string `json:"ticker"`
	// Target_from del stock
	// Required: true
	// Minimum: 1
	// Default: 420
	Target_from float64 `json:"target_from"`
	// Target_to del stock
	// Required: true
	// Minimum: 1
	// Default: 470
	Target_to float64 `json:"target_to"`
	// Company del stock
	// Required: true
	// Default: Banco Santander (Brasil)
	Company string `json:"company"`
	// Action del stock
	// Required: true
	// Default: upgraded by
	Action string `json:"action"`
	// Brokerage del stock
	// Required: true
	// Default: The Goldman Sachs Group
	Brokerage string `json:"brokerage"`
	// Rating_from del stock
	// Required: true
	// Default: Sell
	Rating_from string `json:"rating_from"`
	// Rating_to del stock
	// Required: true
	// Default: Neutral
	Rating_to string `json:"rating_to"`
	// Time del stock
	// Required: false
	// Default: 2025-01-13T00:30:05.813548892Z
	Time *time.Time `json:"time"`
}
