package stores

import (
	"database/sql"
	_ "github.com/lib/pq" // postgres golang driver
)

func CreateConnection() (*sql.DB, error) {
	// Open the connection
	db, err := sql.Open("postgres", "host=localhost port=5432 dbname=postgres user=postgres password=postgres sslmode=disable connect_timeout=10")

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
