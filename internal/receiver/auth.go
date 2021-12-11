package receiver

import (
	"os"
)

func getAPIKey() string {
	return os.Getenv("API_KEY")
}
