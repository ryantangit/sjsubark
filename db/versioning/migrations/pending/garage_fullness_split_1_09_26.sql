CREATE TABLE calendar (
	calendar_date date,
	is_holiday boolean,
	is_weekend boolean,
);

CREATE TABLE garage_info (
	garage_id SERIAL PRIMARY KEY,
	garage_name VARCHAR(50) NOT NULL,
	address VARCHAR(100) NOT NULL,
);
