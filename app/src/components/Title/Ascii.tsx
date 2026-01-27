export default function AsciiArt() {

	const art = `
		  BARK BARK PARKING FULL?
        	\\ 
			/ \\ / \\_       _________
     _____  (  @  #\\___    \\\\_______\\\\
    /     \\ /         0|    \\\\       \\\\
   |       |/   (_____/      \\\\       \\\\
    \\_____//     /            \\\\_______\\\\
 ____|___|/_____/______________\\\\_______\\\\____
 ______________________________________________\\
	`
	
 	return (
		<pre className="font-mono text-xs leading-none text-blue-500 overflow-x-auto">
			{art}
		</pre>
	)
}

