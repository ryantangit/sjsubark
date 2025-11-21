import Garage from "@/components/Garage";
import { Accordion, Heading} from "@chakra-ui/react";
import { fetchLatestStatus } from "@/lib/datastore"


export default async function Home() {	
	const data = await fetchLatestStatus();
	return (
		<Accordion.Root width={"320px"} marginX={"auto"} marginY={"auto"}>
			{
				data.map((d)=> (
					<Accordion.Item key={d.Garage_name} value={d.Garage_name} >
					<Accordion.ItemTrigger>
						<Heading flex="1"> {d.Garage_name} </Heading>
						<Accordion.ItemIndicator/>
					</Accordion.ItemTrigger>
					<Accordion.ItemContent>
						<Garage name={d.Garage_name} fullness={d.Fullness} />
					</Accordion.ItemContent>
					</Accordion.Item>
				))	
			}
		</Accordion.Root>
	);
}
