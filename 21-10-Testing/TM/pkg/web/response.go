package web

type ApiResponse struct {
	Code  int     `json:"code"`
	Data  any     `json:"data,omitempty"`
	Error *string `json:"error,omitempty"`
}

func New(code int, data interface{}, err error) ApiResponse {
	var errorMessage *string

	if err != nil {
		message := err.Error()
		errorMessage = &message
	}

	return ApiResponse{
		Code:  code,
		Data:  data,
		Error: errorMessage,
	}
}
