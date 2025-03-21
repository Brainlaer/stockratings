package utils



// Response representa una respuesta del servidor.
//
// swagger:response Response
type Response struct{
	Status string `json:"status"`
	Success ResponseSuccess `json:"success,omitempty"`
	Error ResponseError`json:"error,omitempty"`
}

type ResponseError struct{
	Code string `json:"code,omitempty"`
	Details string `json:"details,omitempty"`
}

type ResponseSuccess struct{
	Data any `json:"data,omitempty"`
	Meta any `json:"meta,omitempty"`
}