package models

// swagger:parameters createStock
type StockRatingRequest struct {
	// Envoltura del cuerpo de la solicitud
	// in:body
	Body StockRatingCreate `json:"body"`
}

// swagger:parameters updateStock
type StockRatingUpdateRequest struct{
	// ID del stock
	// in: path
	ID string `json:"id" path:"id"`
	// Envoltura del cuerpo de la solicitud
	// in:body
	Body *StockRatingCreate `json:"body"`
}

// swagger:model createStock
type StockRatingCreate struct {
	// Ticker del stock
	// Required: true
	// Default: BSBR
	Ticker string `json:"ticker"`
	// Target_from del stock
	// Required: true
	// Minimum: 1
	// Default: 420
	Target_from int64 `json:"target_from"`
	// Target_to del stock
	// Required: true
	// Minimum: 1
	// Default: 470
	Target_to int64 `json:"target_to"`
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
	// Required: true
	// Default: 2025-01-13T00:30:05.813548892Z
	Time string `json:"time"`
}

// swagger:model createStock
type StockRatingGet struct {
	// ID del stock
	// Required: false
	// Default: a927ed98-e7b3-460b-91e3-c0c72bb8900c
	ID string `json:"id"`
	StockRatingCreate
}

// swagger:parameters id
type StockRatingId struct {
	// ID del stock
	// in: path
	ID string `json:"id" path:"id"`
}


func NewStockRatingCreate(stock *StockRatingCreate, id string) *StockRatingGet {
	return &StockRatingGet{
		ID:          id,
		StockRatingCreate: *stock,
	}
}
