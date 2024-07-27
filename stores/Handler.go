package stores

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // postgres golang driver
)

func CreateConnection() (*sql.DB, error) {
	db_host, env_err := os.LookupEnv("DB_HOST")
	if env_err {
		return nil, fmt.Errorf("could not find variable DB_HOST in env")
	}
	dataSourceName := fmt.Sprintf("host=%s port=5432 dbname=postgres user=postgres password=postgres sslmode=disable connect_timeout=10", db_host)
	// Open the connection
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// return the connection
	return db, nil
}
