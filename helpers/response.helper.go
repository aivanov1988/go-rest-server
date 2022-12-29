package helpers

import (
	"encoding/json"
)

type errorResponse struct {
	Message string
}

func PrepareListResponse[T any](data []T) []any {
	var res []any
	for _, item := range data {
		itemStr, _ := JsonMarshalIgnoreTags(item)
		var itemRes map[string]interface{}
		json.Unmarshal(itemStr, &itemRes)
		res = append(res, itemRes)
	}
	return res
}

func PrepareErrorResponse(err error) errorResponse {
	var res errorResponse
	res.Message = err.Error()
	return res
}
