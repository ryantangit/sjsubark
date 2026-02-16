"use client"

import { useState, useEffect } from "react"
import { ChakraProvider, defaultSystem } from "@chakra-ui/react"
import {
  ColorModeProvider,
  type ColorModeProviderProps,
} from "./color-mode"

export function Provider(props: ColorModeProviderProps) {
	
	// Annoyingly throws up a hydration error and this useeffect/usestate fixes it
	//https://stackoverflow.com/questions/79116172/hydration-failed-problem-with-nextjs-v15-and-chakraui-when-using-components
	const [hydrated, setHydrated] = useState(false);
	useEffect(() => {
        setHydrated(true)
    }, [])
	if (!hydrated) {
		return null;
	}
  return (
    <ChakraProvider value={defaultSystem}>
      <ColorModeProvider {...props} />
    </ChakraProvider>
  )
}
