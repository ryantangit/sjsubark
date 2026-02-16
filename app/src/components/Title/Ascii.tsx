export default function AsciiArt() {

	const art = `
	BARK BARK PARKING FULL?
        	\\ 
			/ \\ / \\_    
     _____  (  @  #\\___ 
    /     \\ /         0| 
   |       |/   (_____/  
    \\_____//     /      
 ____|___|/_____/________
 ________________________
	`
	
 	return (
		<div className="w-full overflow-hidden">
		<pre className="font-mono text-xs leading-none text-blue-500">
			{art}
		</pre>
		</div>
	)
}

