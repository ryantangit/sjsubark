import { DATASTORE_URL } from "@/lib/constants";

export async function fetchLatestStatus(): Promise<any[]> {
	try {
		console.log(DATASTORE_URL);
		const response = await fetch(`${DATASTORE_URL}/latest`);
		if (!response.ok) {
			throw new Error("Fetch Latest Garage Status failed.");
		}
		const json = await response.json();
		return json;
	} catch (error) {
		console.error(error);
		return [];
	}
}
