package datastore

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)

type DataStore struct {
	db *sql.DB
}

func New(db *sql.DB) DataStore {
	return DataStore{db}
}

func (d *DataStore) HelloWorldHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var name string
		// Execute the query.
		row := d.db.QueryRow("SELECT myname FROM mytable")
		if err := row.Scan(&name); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		// Write it back to the client.
		fmt.Fprintf(w, "hi %s!\n", name)
	})
}

func HelloHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var name string
		// Execute the query.
		row := db.QueryRow("SELECT myname FROM mytable")
		if err := row.Scan(&name); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		// Write it back to the client.
		fmt.Fprintf(w, "hi %s!\n", name)
	})
}

func WithMetrics(l *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()
		next.ServeHTTP(w, r)
		l.Printf("%s %s took %s", r.Method, r.URL, time.Since(began))
	})
}
