package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cecobask/redhat-coding-challenge/handler"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	handler := handler.NewHandler()
	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%s", os.Getenv("SERVER_PORT")),
		Handler: handler,
	}
	log.Println("Starting http server at", server.Addr)
	log.Fatal(server.ListenAndServe())
}
