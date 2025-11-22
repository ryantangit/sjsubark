import { Accordion, Flex, Heading} from "@chakra-ui/react";
import { fetchLatestStatus } from "@/lib/datastore"
import GarageTitle from "@/components/Garage";

export default async function Home() {	
	const data = await fetchLatestStatus();

	return (
		<Flex alignContent={"center"} justifyContent={"center"} height={"500px"}>
		<Accordion.Root collapsible width={"750px"} marginX={"auto"} marginY={"auto"}>
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
		</Flex>
	);

}
