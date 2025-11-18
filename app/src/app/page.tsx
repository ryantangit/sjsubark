import { fetchLatestStatus } from "@/lib/datastore"
export default async function Home() {	
	const data = await fetchLatestStatus();
	console.log(data);
	return (
		<div> 
			Hello
		</div>
	);
}
