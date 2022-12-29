package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func getVariantStructValue(v reflect.Value, t reflect.Type) reflect.Value {
	sf := make([]reflect.StructField, 0)
	for i := 0; i < t.NumField(); i++ {
		sf = append(sf, t.Field(i))

		if t.Field(i).Tag.Get("json") != "" {
			sf[i].Tag = ``
		}
	}
	newType := reflect.StructOf(sf)
	return v.Convert(newType)
}

func JsonMarshalIgnoreTags(obj interface{}) ([]byte, error) {
	value := reflect.ValueOf(obj)
	t := value.Type()
	newValue := getVariantStructValue(value, t)
	return json.Marshal(newValue.Interface())
}

func JsonStringify(a any) string {
	res, err := json.Marshal(a)
	if err != nil {
		return fmt.Sprint(err)
	}
	return string(res)
}
