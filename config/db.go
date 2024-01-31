package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"os"

	log "github.com/sirupsen/logrus"
)

// Database is a
type Database struct{}

// ConnectDB is a
func (o Database) ConnectDB() (*sql.DB, error) {
	DB := os.Getenv("DB")
	URI := os.Getenv("DBURL")
	fmt.Println(DB)
	fmt.Println("......")
	db, err := sql.Open(DB, URI)

	if err != nil {
		log.Warnf("failed connection to DB : %v", err)
		return nil, fmt.Errorf("failed connection to DB : %v", err)
	}
	//

	return db, nil
}
