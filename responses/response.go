package responses

type HttpResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"omitempty"`
}
