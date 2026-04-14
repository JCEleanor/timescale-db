# Implementation Plan: Go & TimescaleDB Weather Tracker

## Overview
This plan outlines the creation of a Go-based REST API that ingests weather data (temperature/humidity) and stores it in TimescaleDB, utilizing time-series specific features.

## 1. Formatted Steps

### Phase 1: Environment & Database Setup
- [ ] **Adjust Docker Setup**: Update the `docker run` command to use a `./data` subdirectory for persistence.
- [ ] **Initialize Database Schema**:
    - Create a standard PostgreSQL table for `weather_metrics`.
    - Transform it into a **Hypertable** partitioned by time using `SELECT create_hypertable('weather_metrics', 'time');`.
- [ ] **Go Project Initialization**: Initialize the module and install the `pgx` driver.

### Phase 2: Database Connectivity in Go
- [ ] **Connection Pooling**: Implement a connection pool using `pgxpool` to safely handle concurrent requests.
- [ ] **Health Check**: Create a simple function to verify the database connection on startup.

### Phase 3: Server Implementation
- [ ] **REST API Setup**: Use the standard `net/http` package (or a router like `chi`) to create endpoints.
- [ ] **Data Ingestion (POST /metrics)**:
    - Define a `Metric` struct.
    - Implement a handler to parse JSON and insert data into the `weather_metrics` table.
- [ ] **Querying Data (GET /stats)**:
    - Implement a handler that uses TimescaleDB's `time_bucket` function to return 1-hour averages of temperature.

### Phase 4: Validation & Learning
- [ ] **Manual Testing**: Use `curl` or Postman to send sample data.
- [ ] **Analysis**: Connect via `psql` to verify data distribution across chunks.

---

## 2. Key Learning Points

### Golang
- **Structs and JSON**: Handling data serialization/deserialization.
- **Concurrency**: How Go handles multiple HTTP requests and database connections.
- **Error Handling**: The idiomatic "check if err != nil" pattern in database operations.

### TimescaleDB
- **Hypertables**: Understanding how TimescaleDB automatically partitions data by time while treating it as a single table.
- **Chunks**: How data is physically stored in the background.
- **Time-Series Functions**: Using `time_bucket` for easy data aggregation over time intervals.

## 3. Key Takeaways
- How to bridge a statically typed language (Go) with a powerful time-series relational database.
- The performance benefits of using Hypertables over standard PostgreSQL tables for time-stamped data.
- Basic CRUD operations and time-series aggregations in a production-like environment.
