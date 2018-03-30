package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/data-love/authrii/cors"
	"github.com/data-love/authrii/datastore"
	"github.com/data-love/authrii/logging"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=postgres sslmode=disable host=db"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	store := datastore.New(db)
	http.Handle("/health", http.HandlerFunc(HealthCheckHandler))
	corsMiddleware := cors.Handler(cors.Options{})
	loggingMiddleware := logging.Handler(os.Stdout)
	http.Handle("/hello", datastore.HelloHandler(db))
	http.Handle("/hello_again", loggingMiddleware(corsMiddleware(datastore.HelloHandler(db))))
	http.Handle("/hello_world", loggingMiddleware(corsMiddleware(store.HelloWorldHandler())))
	http.Handle("/", loggingMiddleware(corsMiddleware(IndexHandler())))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
