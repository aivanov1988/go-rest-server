package helpers

type errorResponse struct {
	Message string
}

func PrepareErrorResponse(err error) errorResponse {
	var res errorResponse
	res.Message = err.Error()
	return res
}
