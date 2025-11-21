"use client"
import { Card } from "@chakra-ui/react";

interface GarageProps {
	name: string;
	fullness: number;
}

export default function Garage(props: GarageProps) {
	return (
	<Card.Root width="320px">
      <Card.Body gap="2">
        <Card.Title mt="2">{props.name}</Card.Title>
        <Card.Description>
		{props.fullness}
        </Card.Description>
      </Card.Body>
    </Card.Root>
	)
}
