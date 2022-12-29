package helpers

import (
	"encoding/json"
)

type errorResponse struct {
	Message string
}

func PrepareEntityResponse[T any](data T) any {
	var res any
	dataStr, _ := JsonMarshalIgnoreTags(data)
	json.Unmarshal(dataStr, &res)
	return res
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
