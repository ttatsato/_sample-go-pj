package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// printDebugf behaves like log.Printf only in the debug env
func printDebugf(format string, args ...interface{}) {
	if env := os.Getenv("GO_SERVER_DEBUG"); len(env) != 0 {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

func convertMapToJsonString(src interface{}) string {
	bytes, err := json.Marshal(src)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		return ""
	}
	return string(bytes)
}