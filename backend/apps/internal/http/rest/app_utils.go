package rest

type errorResp struct {
	Error string `json:"error"`
}

func newErrorResp(err error) *errorResp {
	return &errorResp{
		Error: err.Error(),
	}
}

type resp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func newResp(message string, data interface{}) *resp {
	return &resp{
		Message: message,
		Data:    data,
	}
}
