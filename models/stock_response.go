package models

// swagger:model getStocks
type StockResponseGet struct {
	// ID del stock
	// Required: false
	// Default: a927ed98-e7b3-460b-91e3-c0c72bb8900c
	ID string `json:"id"`
		// crecimiento del stock
	// Required: false
	// Default: 12.78
	Growth float64 `json:"growth"`
	Stock
}