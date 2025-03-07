package utils



// Response representa una respuesta del servidor.
//
// swagger:response Response
type Response struct{
	Body ResponseBody`json:"body"`
}

type ResponseBody struct{
	Status string `json:"status"`
	Data any `json:"data,omitempty"`
	Error ResponseError`json:"error,omitempty"`
}

type ResponseError struct{
	Code string `json:"code,omitempty"`
	Details string `json:"details,omitempty"`
}