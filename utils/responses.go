package utils



// Response representa una respuesta del servidor.
//
// swagger:response Response
type Response struct{
	Body struct {
		Status string `json:"status"`
		Data any `json:"data,omitempty"`
		Error *struct{
			Code string `json:"code"`
			Details string `json:"details"`
		}`json:"error,omitempty"`
	} `json:"body"`
}