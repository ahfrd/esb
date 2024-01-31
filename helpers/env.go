package helpers

import (
	"log"

	"github.com/joho/godotenv"
	// log "micro-auth/src/helpers"
)

// Env is a struct for env helpers
type Env struct{}

// StartingCheck is a function for checking a env ready to running
func (o Env) StartingCheck() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
