import { Heading } from "@chakra-ui/react"
import { ComponentPropsWithRef } from "react"

interface TitleProps extends ComponentPropsWithRef<typeof Heading>{
	children?: React.ReactNode
}

export default function Title({children, ...rest}: TitleProps) {
	return (
		<Heading size="6xl" {...rest}>
			{children}		
		</Heading>
	)
}
