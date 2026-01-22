"use client"
import { Heading, HStack } from "@chakra-ui/react";
import FullnessTag from "./FullnessTag";

interface GarageProps {
	name: string;
	fullness: number;
}

export default function GarageTitle(props: GarageProps) {
	return (
		<HStack width={"320px"} flex={1} justifyContent={"space-between"}>
			<Heading> {props.name} </Heading>	
			<FullnessTag fullness={props.fullness} />
			</HStack>
	)
}
