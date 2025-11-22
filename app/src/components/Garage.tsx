"use client"
import { Card, Heading, HStack, Tag } from "@chakra-ui/react";

interface GarageProps {
	name: string;
	fullness: number;
}

export default function GarageTitle(props: GarageProps) {
	return (
		<HStack width={"320px"} flex={1} justifyContent={"space-between"}>
			<Heading> {props.name} </Heading>	
			<Tag.Root>
				<Tag.Label>
					{props.fullness}%
				</Tag.Label>
			</Tag.Root>
		</HStack>
	)
}
