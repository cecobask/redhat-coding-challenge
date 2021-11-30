package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

// Database ...
type Database struct {
	Conn *pgx.Conn
}

const (
	host = "database"
	port = 5432
)

// Initialize connects to the postgres database
func Initialize(user, password, dbname string) (Database, error) {
	database := Database{}
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	postgresConn, err := pgx.Connect(context.Background(), dataSourceName)
	if err != nil {
		return database, fmt.Errorf("Unable to connect to database: %v", err)
	}
	database.Conn = postgresConn
	log.Println("Database connection established")
	return database, nil
}
