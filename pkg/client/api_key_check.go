package client

import (
	"log"
	"os"
)

func GetApiKey() string {
	if _, exists := os.LookupEnv("API_KEY"); !exists {
		log.Fatal("API_KEY not set! Cancelling search...\n")
	}
	return os.Getenv("API_KEY")
}
