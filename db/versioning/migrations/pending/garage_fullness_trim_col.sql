ALTER TABLE garage_fullness
	DROP column second, 
	DROP column minute, 
	DROP column hour, 
	DROP column day, 
	DROP column month, 
	DROP column year, 
	DROP column weekday, 
	DROP column is_weekend,
	DROP column is_campus_closed, 
	ADD column garage_id INTEGER,
	ADD constraint fk_garage_id 
		FOREIGN KEY (garage_id) 
		REFERENCES garage_info(garage_id) 
		ON DELETE RESTRICT;

-- Populate all the relevant garage_id matched with the names
UPDATE garage_fullness gf 
SET gf.garage_id = gi.garage_id
FROM garage_info gi 
WHERE gf.name = gi.garage_name;

