package main

import "time"

type WeatherMetric struct {
	Time        time.Time `json:"time"` // json tag to specify the key name in JSON
	Location    string    `json:"location"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
}

type WeatherStats struct {
	Bucket      time.Time `json:"bucket"`
	AvgTemp     float64   `json:"avg_temp"`
	MaxTemp     float64   `json:"max_temp"`
	SampleCount int       `json:"sample_count"`
}
