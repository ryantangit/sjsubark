export interface Garage { 
	garage_id: number;
	name: string;	
	utc_timestamp: string;
	fullness: number;
}

export interface Prediction {
	name: string;
	increments: number;
	forecast: number;
}

