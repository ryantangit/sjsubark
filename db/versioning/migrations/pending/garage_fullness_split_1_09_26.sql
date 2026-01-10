CREATE TABLE IF NOT EXISTS calendar (
	calendar_date DATE,
	is_holiday BOOLEAN,
	is_weekend BOOLEAN
);

CREATE TABLE IF NOT EXISTS garage_info (
	garage_id SERIAL PRIMARY KEY,
	garage_name VARCHAR(50) NOT NULL,
	address VARCHAR(100) NOT NULL,
);
