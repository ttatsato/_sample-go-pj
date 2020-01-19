package interfaces

import (
	"encoding/json"
	"fmt"
)

func convertMapToJsonString(src interface{}) string {
	bytes, err := json.Marshal(src)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		return ""
	}
	return string(bytes)
}