INSERT INTO calendar (calendar_date, is_holiday, is_weekend)
SELECT 
    datum AS calendar_date,
    FALSE AS is_holiday, --default to false for now
	CASE 
        WHEN EXTRACT(DOW FROM datum) IN (0, 6) THEN TRUE 
        ELSE FALSE 
    END AS is_weekend
FROM generate_series(
    '2025-01-01'::DATE, 
    '2030-12-31'::DATE, 
    '1 day'::interval
) AS datum;
