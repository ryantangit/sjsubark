ALTER TABLE garage_fullness
	DROP column IF EXISTS second, 
	DROP column IF EXISTS minute, 
	DROP column IF EXISTS hour, 
	DROP column IF EXISTS day, 
	DROP column IF EXISTS month, 
	DROP column IF EXISTS year, 
	DROP column IF EXISTS weekday, 
	DROP column IF EXISTS is_weekend,
	DROP column IF EXISTS is_campus_closed, 
	ADD column garage_id INTEGER,
	ADD constraint fk_garage_id 
		FOREIGN KEY (garage_id) 
		REFERENCES garage_info(garage_id) 
		ON DELETE RESTRICT;

-- Populate all the relevant garage_id matched with the names
UPDATE garage_fullness gf 
SET garage_id = gi.garage_id
FROM garage_info gi 
WHERE gf.name = gi.garage_name;

