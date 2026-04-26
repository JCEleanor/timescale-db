#!/bin/bash
# test_ingest.sh - verifies the Go + TimescaleDB integration

echo "--- Sending Sample Metrics ---"
curl -X POST http://localhost:8080/metrics \
     -H "Content-Type: application/json" \
     -d '{"time":"'$(date -u +"%Y-%m-%dT%H:%M:%SZ")'","location":"London","temperature":15.5,"humidity":70.2}'

echo -e "\n\n--- Fetching Stats ---"
curl -s http://localhost:8080/stats | jq .
