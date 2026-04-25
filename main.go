package main // compiled into an executable program, not just a library.

import (
	"context"
	"fmt"
	"log"
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

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to connect to database %v\n", err)

	}

	fmt.Println("Successfully connected to TimescaleDB!")
}
