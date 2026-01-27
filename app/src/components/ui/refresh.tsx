"use client"

import { useState } from "react"
import { Button } from "@chakra-ui/react";

export default function RefeshButton() { 
	const [loading, setLoading] = useState(false)
	return ( 
	<Button size={"xs"} loading={loading} onClick={()=> {setLoading(true); location.reload(); setLoading(false)}}>
		Refresh
	</Button>
	)
}
