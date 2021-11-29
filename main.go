package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cecobask/redhat-coding-challenge/db/postgres"
	"github.com/cecobask/redhat-coding-challenge/handler"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Initialize the Postgres database
	postgresConn, err := postgres.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer postgresConn.Close(context.Background())

	// Setup the server
	handler := handler.NewHandler(postgresConn)
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", os.Getenv("SERVER_PORT")),
		Handler: handler,
	}
	log.Println("Starting http server at", server.Addr)
	log.Fatal(server.ListenAndServe())
}
