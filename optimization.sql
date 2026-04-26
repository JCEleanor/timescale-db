-- 1. Create a Continuous Aggregate for hourly stats
-- This replaces the manual GROUP BY in our Go code with a high-performance materialized view.
CREATE MATERIALIZED VIEW IF NOT EXISTS weather_stats_hourly
WITH (timescaledb.continuous) AS
SELECT 
    time_bucket('1 hour', time) AS bucket, 
    location,
    AVG(temperature) as avg_temp, 
    MAX(temperature) as max_temp,
    COUNT(*) as sample_count
FROM weather_metrics 
GROUP BY bucket, location;

-- 2. Set a Refresh Policy
-- Automatically refresh the view every 30 minutes to include new data.
SELECT add_continuous_aggregate_policy('weather_stats_hourly',
    start_offset => INTERVAL '3 hours',
    end_offset => INTERVAL '1 hour',
    schedule_interval => INTERVAL '30 minutes');

-- 3. Enable Compression
-- Compress data older than 7 days to save space.
ALTER TABLE weather_metrics SET (
  timescaledb.compress,
  timescaledb.compress_segmentby = 'location'
);

SELECT add_compression_policy('weather_metrics', INTERVAL '7 days');
