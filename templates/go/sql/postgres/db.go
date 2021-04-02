package postgres

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	log "github.com/sirupsen/logrus"

	// Driver for postgres
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/lib/pq"
)

// Database variables that should be treated as constant.
var (
	HOST     = os.Getenv("POSTGRES_HOST")
	USER     = os.Getenv("POSTGRES_USER")
	PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	DATABASE = os.Getenv("POSTGRES_DB")
	DB_TYPE  = os.Getenv("POSTGRES_DB_TYPE")
	SSLMODE  = "disable"
)

// DB : struct that encapsulates the sql database instance
type DB struct {
	*sqlx.DB
}

// Open : opens up a database connection using the postgres driver
func Open() (*DB, error) {
	log.Info(HOST)
	connStr := fmt.Sprintf("user=%s password=%s host=%s database=%s sslmode=%s", USER, PASSWORD, HOST, DATABASE, SSLMODE)
	log.Info("Opening database connection ...")
	db, err := sqlx.Connect(DB_TYPE, connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Error(err)
	}

	log.Info("Connection successful!")

	return &DB{db}, nil
}
