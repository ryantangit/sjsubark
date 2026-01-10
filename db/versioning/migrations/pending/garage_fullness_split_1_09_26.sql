CREATE TABLE IF NOT EXISTS calendar (
	calendar_date date,
	is_holiday boolean,
	is_weekend boolean,
);

CREATE TABLE IF NOT EXISTS garage_info (
	garage_id SERIAL PRIMARY KEY,
	garage_name VARCHAR(50) NOT NULL,
	address VARCHAR(100) NOT NULL,
);
