package interfaces

import (
	"log"
	"os"
)

// printDebugf behaves like log.Printf only in the debug env
func printDebugf(format string, args ...interface{}) {
	if env := os.Getenv("GO_SERVER_DEBUG"); len(env) != 0 {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

type ApiErrorResponse struct {
	Status int `json:status`
	Message string `json:message`
}

