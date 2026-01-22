"use client"
import { use } from "react"
import { Prediction } from "@/lib/types"
import { DataList, HStack } from "@chakra-ui/react"
import FullnessTag from "./FullnessTag"

interface GaragePredictionsProps {
	predictions: Promise<Prediction | null>[]
}

interface GaragePredictionProps {
	prediction: Promise<Prediction | null>
}

export default function GaragePredictions(props: GaragePredictionsProps) {
	return (
		<HStack>
		{
			props.predictions.map((prediction, index)=> (
				<GaragePrediction key={index} prediction={prediction}/> 
			))	
		}
		</HStack>
	)
}

function GaragePrediction(props: GaragePredictionProps){
	const prediction = use(props.prediction);
	if (prediction == null) return null;
	return (
		<DataList.Root>
		<DataList.Item >
		  <DataList.ItemLabel>{incrementToHourString(prediction.increments)}</DataList.ItemLabel>
		  <DataList.ItemValue><FullnessTag fullness={prediction.forecast}></FullnessTag></DataList.ItemValue>
		</DataList.Item>
		</DataList.Root>
	)
}

function incrementToHourString(increments: number) {
	const hour = increments / 6;	
	let plural = ""
	if (hour > 1) plural = "s";
	return "In " + hour + " Hour" + plural;
}
