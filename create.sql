-- Create the standard table
CREATE TABLE IF NOT EXISTS weather_metrics (
time        TIMESTAMPTZ       NOT NULL,
location    TEXT              NOT NULL,
temperature DOUBLE PRECISION  NULL,
humidity    DOUBLE PRECISION  NULL
);

-- Transform into a hypertable partitioned by the 'time' column
SELECT create_hypertable('weather_metrics', 'time');

INSERT INTO weather_metrics (time, location, temperature, humidity) VALUES 
(NOW() - INTERVAL '1 hour', 'New York', 72.1, 60.5),
(NOW() - INTERVAL '2 hours', 'New York', 71.0, 62.1),
(NOW() - INTERVAL '3 hours', 'New York', 69.8, 65.0);

