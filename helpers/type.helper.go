package helpers

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"

	customJsonUnmarshal "github.com/aivanov1988/go-custom-json-unmarshal"
)

func TypeListConverter[R any, T any](data []T) (*[]R, error) {
	var res []R
	for _, item := range data {
		itemRes, err := TypeConverter[R](item)
		if err != nil {
			return nil, err
		}
		res = append(res, *itemRes)
	}
	return &res, nil
}

func TypeConverter[R any](data any) (*R, error) {
	var res R
	dataStr, errMarshal := json.Marshal(data)
	if errMarshal != nil {
		return nil, errMarshal
	}
	resAsMap, errUnmarshal := customJsonUnmarshal.UnmarshalJSON(dataStr, res)
	mapstructure.Decode(resAsMap, &res)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}
	return &res, nil
}
