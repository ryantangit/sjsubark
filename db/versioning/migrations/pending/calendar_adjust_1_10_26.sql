CREATE TEMP TABLE raw_holiday_data (content jsonb);
COPY raw_holiday_data FROM '/campus_close.json'

WITH base_json AS (
    -- If your file is a JSON array, we expand the root array first
    SELECT jsonb_array_elements(content) AS year_data
    FROM raw_holiday_data
),
parsed_holidays AS (
    SELECT 
        (year_data->>'year')::int AS holiday_year,
        jsonb_array_elements(year_data->'instances') AS inst
    FROM base_json
),
holiday_ranges AS (
    SELECT 
        make_date(holiday_year, (inst->>'startmonth')::int, (inst->>'startday')::int) AS start_date,
        make_date(holiday_year, (inst->>'endmonth')::int, (inst->>'endday')::int) AS end_date
    FROM parsed_holidays
)
UPDATE calendar c
SET is_holiday = TRUE
FROM holiday_ranges hr
WHERE c.calendar_date BETWEEN hr.start_date AND hr.end_date;
