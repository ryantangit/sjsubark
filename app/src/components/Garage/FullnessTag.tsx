import { Tag } from "@chakra-ui/react";

interface FullnessTagProps {
	fullness: number;
}

export default function FullnessTag(props: FullnessTagProps) {
	let colorPalette = "green";	
	if (props.fullness <= 25.0) {
		colorPalette = "green";
	} else if (props.fullness <= 50.0) {
		colorPalette = "yellow";
	} else if (props.fullness <= 75.0) {
		colorPalette = "orange"
	} else {
		colorPalette = "red";
	}
	return (
		<Tag.Root colorPalette={colorPalette} width={"100%"} height={"100%"} justifyContent={"space-evenly"}>
			<Tag.Label> 
				{props.fullness}%
			</Tag.Label>
		</Tag.Root>
	)
}
