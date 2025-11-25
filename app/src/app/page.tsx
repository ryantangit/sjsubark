import { Accordion, Flex, Badge, Stack} from "@chakra-ui/react";
import { fetchLatestStatus } from "@/lib/datastore"
import GarageTitle from "@/components/Garage/GarageTitle";
import Title from "@/components/Title/PageTitle";

export default async function Home() {	
	const data = await fetchLatestStatus();
	const timestamp = new Date(data[0].Utc_timestamp);
	return (
		<Flex alignContent={"center"} justifyContent={"center"} height={"500px"}>
		<Stack>
		<Title marginX={"auto"} marginY={"auto"}>	
			SJSU Bark
		</Title>
		<Accordion.Root collapsible width={"750px"} marginX={"auto"} marginY={"auto"}>
			<Badge> Updated: {timestamp.toString()} </Badge>
			{
				data.map((d)=> (
					<Accordion.Item key={d.Garage_name} value={d.Garage_name} >
					<Accordion.ItemTrigger>
						<GarageTitle name={d.Garage_name} fullness={d.Fullness} />
						<Accordion.ItemIndicator/>
					</Accordion.ItemTrigger>
					<Accordion.ItemContent>
					</Accordion.ItemContent>
					</Accordion.Item>
				))	
			}
		</Accordion.Root>
		</Stack>
		</Flex>
	);

}
