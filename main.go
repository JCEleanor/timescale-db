package main // compiled into an executable program, not just a library.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// format: postgres://<user>:<password>@<host>:<port>/<dbname>
	conStr := "postgres://postgres:password@localhost:5432/postgres"

	config, err := pgxpool.ParseConfig(conStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse connection string: %v\n", err)
		os.Exit(1)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer pool.Close()

	http.HandleFunc("/metrics", handlePostMetrics(pool))
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	// err = pool.Ping(context.Background())
	// if err != nil {
	// 	log.Fatalf("Unable to connect to database %v\n", err)

	// }

	// fmt.Println("Successfully connected to TimescaleDB!")
}

// handlePostMetrics is a factory
// handlePostMetrics creates a handler function that has access to the database pool
func handlePostMetrics(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// only allow POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Decode the JSON body into our WeatherMetric struct
		var m WeatherMetric
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// insert
		query := `INSERT INTO weather_metrics (time, location, temperature, humidity) VALUES ($1, $2, $3, $4)`
		_, err = pool.Exec(context.Background(), query, m.Time, m.Location, m.Temperature, m.Humidity)
		if err != nil {
			http.Error(w, "Failed to insert data", http.StatusInternalServerError)
			log.Printf("Insert error: %v\n", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Metric recorded successfully for %s", m.Location)
	}
}
