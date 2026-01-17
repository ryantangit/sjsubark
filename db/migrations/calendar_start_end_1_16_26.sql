ALTER TABLE calendar
	ADD column IF NOT EXISTS utc_start TIMESTAMP,
	ADD column IF NOT EXISTS utc_end TIMESTAMP;

UPDATE calendar
SET 
    utc_start = (calendar_date::text || ' 00:00:00')::timestamp AT TIME ZONE 'America/Los_Angeles' AT TIME ZONE 'UTC',
    utc_end   = (calendar_date::text || ' 23:59:59.999')::timestamp AT TIME ZONE 'America/Los_Angeles' AT TIME ZONE 'UTC';
