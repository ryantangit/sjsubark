import { Accordion, Flex, Badge, Stack} from "@chakra-ui/react";
import { fetchGaragePrediction, fetchRecentGarage } from "@/lib/inference"
import GarageTitle from "@/components/Garage/GarageTitle";
import Title from "@/components/Title/PageTitle";
import GaragePredictions from "@/components/Garage/GaragePrediction";
import { Garage } from "@/lib/types";

export default async function Home() {	
	const data = await fetchRecentGarage();
	const timestamp = new Date(data[0].utc_timestamp+"Z").toLocaleString("en-US", {
  		timeZone: "America/Los_Angeles"
	});
	return (
		<Flex alignContent={"center"} justifyContent={"center"} minHeight={"500px"} paddingTop={40}>
		<Stack>
		<Title marginX={"auto"} marginY={"auto"}>	
			SJSU Bark
		</Title>
		<Accordion.Root multiple collapsible defaultValue={data.map(d=>(d.garage_id.toString()))} width={"750px"} marginX={"auto"} marginY={"auto"} padding={5}>
			<Badge> Updated: {timestamp.toString()} </Badge>
			{
				data.map((d)=> (
					<Accordion.Item key={d.garage_id} value={d.garage_id.toString()} >
					<Accordion.ItemTrigger>
						<GarageTitle name={d.name} fullness={d.fullness} />
						<Accordion.ItemIndicator/>
					</Accordion.ItemTrigger>
					<Accordion.ItemContent>
						<GaragePredictions key={d.garage_id} predictions={fetchFuturePredictions(d)}/>	
					</Accordion.ItemContent>
					</Accordion.Item>
				))	
			}
		</Accordion.Root>
		</Stack>
		</Flex>
	);

}

function fetchFuturePredictions(garage: Garage) {
	let results = [];
	//1 hr
	results.push(fetchGaragePrediction(garage.garage_id, (1) * 6));
	//3 hrs
	results.push(fetchGaragePrediction(garage.garage_id, (3) * 6));
	//6 hrs
	results.push(fetchGaragePrediction(garage.garage_id, (6) * 6));
	return results
}

