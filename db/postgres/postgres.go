package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

const (
	host = "database"
	port = 5432
)

// Initialize connects to the postgres database
func Initialize(user, password, dbname string) (*pgx.Conn, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	postgresConn, err := pgx.Connect(context.Background(), dataSourceName)
	if err != nil {
		return postgresConn, fmt.Errorf("Unable to connect to database: %v", err)
	}
	log.Println("Database connection established")
	return postgresConn, nil
}
