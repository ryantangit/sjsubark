import Garage from "@/components/Garage";
import { fetchLatestStatus } from "@/lib/datastore"
export default async function Home() {	
	const data = await fetchLatestStatus();
	console.log(data);
	return (
		<div> 
			{
				data.map((d)=> (
					<Garage key={d.Garage_name} name={d.Garage_name} fullness={d.Fullness} />
				))	
			}
		</div>
	);
}
