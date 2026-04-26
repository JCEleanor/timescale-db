# Implementation Plan: Go & TimescaleDB Weather Tracker

## Overview
This plan outlines the creation of a Go-based REST API that ingests weather data (temperature/humidity) and stores it in TimescaleDB, utilizing time-series specific features.

## 1. Formatted Steps

### Phase 1: Environment & Database Setup
- [x] **Adjust Docker Setup**: Update the `docker run` command to use a `./data` subdirectory for persistence.
- [x] **Initialize Database Schema**:
    - Create a standard PostgreSQL table for `weather_metrics`.
    - Transform it into a **Hypertable** partitioned by time using `SELECT create_hypertable('weather_metrics', 'time');`.
- [x] **Go Project Initialization**: Initialize the module and install the `pgx` driver.

### Phase 2: Database Connectivity in Go
- [x] **Connection Pooling**: Implement a connection pool using `pgxpool` to safely handle concurrent requests.
- [x] **Health Check**: Create a simple function to verify the database connection on startup.

### Phase 3: Server Implementation
- [x] **REST API Setup**: Use the standard `net/http` package (or a router like `chi`) to create endpoints.
- [x] **Data Ingestion (POST /metrics)**:
    - Define a `Metric` struct.
    - Implement a handler to parse JSON and insert data into the `weather_metrics` table.
- [x] **Querying Data (GET /stats)**:
    - Implement a handler that uses TimescaleDB's `time_bucket` function to return 1-hour averages of temperature.

### Phase 4: Validation & Learning
- [ ] **Manual Testing**: Use `curl` or Postman to send sample data.
- [ ] **Analysis**: Connect via `psql` to verify data distribution across chunks.

### Phase 5: TimescaleDB Optimization & Lifecycle
- [ ] **Continuous Aggregates**: Replace the raw SQL query in `GET /stats` with a Continuous Aggregate for real-time performance.
- [ ] **Compression Policy**: Enable compression on the `weather_metrics` (or `records`) hypertable to reduce storage footprint by up to 90%.
- [ ] **Data Retention**: Implement a retention policy to automatically drop chunks older than 30 days.

### Phase 6: Robust Protocol Implementation (The "Pro" Ingest)
- [ ] **Schema Migration**: Implement the `streams` and `records` schema from `knowledge/protocol.md` to support idempotent upserts and HWM (High Water Mark).
- [ ] **Endpoint Evolution**:
    - `POST /v1/hello`: Initialize stream session and anchors.
    - `POST /v1/batch`: Implement batch processing with `ON CONFLICT DO NOTHING` for exactly-once semantics.
    - `GET /v1/cursor`: Allow devices to resume ingestion after disconnection.

### Phase 7: Go Concurrency & Simulation
- [ ] **Simulator CLI**: Create a separate Go sub-command/binary that simulates 100+ "weather stations" sending concurrent batches.
- [ ] **Graceful Shutdown**: Implement signal handling in the server to flush connection pools and stop background workers cleanly.
- [ ] **Middleware**: Add a custom logging and recovery middleware to the HTTP server.

---

## 2. Key Learning Points

### Golang
- **Structs and JSON**: Handling data serialization/deserialization.
- **Concurrency**: Using Goroutines and Channels for simulators; managing `sync.WaitGroup` for graceful shutdowns.
- **Error Handling**: The idiomatic "check if err != nil" pattern and custom error types.
- **Advanced API Design**: Functional options for server config, middleware patterns, and context propagation.

### TimescaleDB
- **Hypertables**: Understanding automatic partitioning by time.
- **Continuous Aggregates**: How to speed up dashboards using materialized views that update automatically.
- **Compression**: The columnar storage format and how it interacts with time-series data.
- **Idempotency**: Using `INSERT ... ON CONFLICT` to handle network retries without data duplication.

## 3. Key Takeaways
- Bridging statically typed Go with time-series relational power.
- Building systems that are "resilient by design" using the HWM protocol.
- Optimizing for both write-heavy ingest and read-heavy analytics simultaneously.
