import { Accordion, Box, Badge, Heading,Stack} from "@chakra-ui/react";
import { fetchGaragePrediction, fetchRecentGarage } from "@/lib/inference"
import GarageTitle from "@/components/Garage/GarageTitle";
import Title from "@/components/Title/PageTitle";
import GaragePredictions from "@/components/Garage/GaragePrediction";
import { Garage } from "@/lib/types";
import AsciiArt from "@/components/Title/Ascii";

export default async function Home() {	
	const data = await fetchRecentGarage();
	if (!data || data.length === 0) {
		return <div> Loading </div>
	}
	const timestamp = new Date(data[0].utc_timestamp+"Z").toLocaleString("en-US", {
  		timeZone: "America/Los_Angeles"
	});
	return (
		<Box width={{base: "100%", md: "750px", lg: "1000px"}} mx="auto" marginTop={"10px"}>
		<Stack alignContent={"center"} justifyContent={"center"} minHeight={"500px"}>
		<Title margin={"auto"} size={"7xl"}>	
			SJSU Bark
		</Title>
		<Heading margin={"auto"} size={"md"} opacity={0.7}>
			San Jose State Garage Occupancy Inference
		</Heading>
		<Box margin={"auto"}>
		<AsciiArt />
		</Box>
		<Accordion.Root multiple collapsible defaultValue={data.map(d=>(d.garage_id.toString()))} marginX={"auto"} marginY={"auto"} padding={5} size={"lg"}>
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
		</Box>
	);

}


// Fetches next hour, 3 hours and 6 hours as a promise to be resolved by React's use hook
function fetchFuturePredictions(garage: Garage) {
	let results = [
		fetchGaragePrediction(garage.garage_id, (1) * 6),
		fetchGaragePrediction(garage.garage_id, (3) * 6),
		fetchGaragePrediction(garage.garage_id, (6) * 6),
	];
	return results
}

