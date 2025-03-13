package models
// swagger:parameters createStock
type StockRequestCreate struct {
	// Envoltura del cuerpo de la solicitud
	// in:body
	Body Stock `json:"body"`
}

// swagger:parameters updateStock
type StockRequestUpdate struct{
	// ID del stock
	// in: path
	ID string `json:"id" path:"id"`
	// Envoltura del cuerpo de la solicitud
	// in:body
	Body *Stock `json:"body"`
}

// swagger:parameters deleteStock
type StockRequestDelete struct {
	// ID del stock
	// in: path
	ID string `json:"id" path:"id"`
}


// swagger:parameters getStock
type StockRequestGetOne struct {
	// ID del stock
	// in: path
	ID string `json:"id" path:"id"`
}