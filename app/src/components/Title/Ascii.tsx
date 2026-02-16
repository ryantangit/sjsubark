export default function AsciiArt() {

	const art = `
		  BARK BARK PARKING FULL?
        	\\ 
			/ \\ / \\_      _________
     _____  (  @  #\\___   \\\\_______\\\\
    /     \\ /         0|   \\\\       \\\\
   |       |/   (_____/     \\\\       \\\\
    \\_____//     /           \\\\_______\\\\
 ____|___|/_____/_____________\\\\_______\\\\
 ________________________________________
	`
	
 	return (
		<div className="w-full overflow-hidden">
		<pre className="font-mono text-xs leading-none text-blue-500">
			{art}
		</pre>
		</div>
	)
}

