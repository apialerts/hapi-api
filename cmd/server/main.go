package main

import (
	"hapi/internal/util/webserver"
	"log"
	"net/http"
	"os"
	"time"

	"hapi/internal/db"
	"hapi/internal/handlers"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /export", handlers.Export(database))
	mux.HandleFunc("GET /import/{code}", handlers.Import(database))
	mux.HandleFunc("GET /health", handlers.Health(database))
	mux.HandleFunc("GET /tasks/cleanup", handlers.CleanupExpired(database))

	notFoundHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		webserver.RespondHeader(w, http.StatusNotFound)
	})
	mux.Handle("/", notFoundHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("Listening on :" + port)
	log.Fatal(server.ListenAndServe())
}
