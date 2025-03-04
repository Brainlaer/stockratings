package utils

// successResponse representa una respuesta exitosa.
//
// swagger:response successResponse
type successResponse struct {
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}
// acceptedResponse representa una respuesta aceptada.
//
// swagger:response acceptedResponse
type acceptedResponse struct {
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}