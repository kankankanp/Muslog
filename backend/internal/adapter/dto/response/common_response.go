package response

type CommonResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
