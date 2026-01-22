export default function AsciiArt() {

	const art = `
				 / WOOF WOOF I WONDER IF THE SCHOOL PARKING LOT WILL BE FULL AGAIN?
                / 
			/ \\ / \\_        _________
     _____  (  @  #\\___     \\\\_______\\\\
    /     \\ /         0|     \\\\       \\\\
   |       |/   (_____/       \\\\       \\\\
    \\_____//     /             \\\\_______\\\\
 ____|___|/_____/_______________\\\\_______\\\\____
 _______________________________________________\\
	`
	
 	return (
		<pre className="font-mono text-xs leading-none text-blue-500 overflow-x-auto">
			{art}
		</pre>
	)
}

