CREATE TEMP TABLE campus_closed (
	closed_campus_date DATE		
);

COPY campus_closed FROM '/campus_closed.txt';

UPDATE calendar
SET is_campus_closed = TRUE
WHERE calendar_date IN (SELECT closed_campus_date from campus_closed);

DROP campus_closed;

