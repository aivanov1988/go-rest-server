package helpers

import (
	"encoding/json"
	"fmt"
)

func JsonStringify(a any) string {
	res, err := json.Marshal(a)
	if err != nil {
		return fmt.Sprint(err)
	}
	return string(res)
}
