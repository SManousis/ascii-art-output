package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/funcs"
)

func main() {
	// Check if the number of command-line arguments is correct
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> \"something\" standard")
		return
	}

	var text string       // The input text to be converted to ASCII art
	var output string     // The output flag (if any)
	banner := "standard"  // Default banner is 'standard'
	var outputfile string // The name of the output file (if any)
	hasOutput := false    // Flag to indicate if output to file is requested

	// Handle the case when there are 4 arguments: output flag, text, and banner
	if len(os.Args) == 4 {
		output = os.Args[1]
		text = os.Args[2]
		banner = strings.ToLower(os.Args[3])
		hasOutput = true

		// Validate the output flag format
		if !strings.HasPrefix(output, "--output=") || !strings.HasSuffix(output, ".txt") {
			fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> \"something\" standard")
			return
		}
		// Extract the output file name from the flag
		outputfile = output[9:] // Skips the "--output=" part
	}

	// Handle the case when there are 3 arguments
	if len(os.Args) == 3 {
		output = os.Args[1]
		if strings.HasPrefix(output, "--output=") && strings.HasSuffix(output, ".txt") {
			// If the first argument is the output flag
			text = os.Args[2]
			outputfile = output[9:] // Skips the "--output=" part
			hasOutput = true
		} else if strings.HasPrefix(output, "--output=") && !strings.HasSuffix(output, ".txt") {
			// If the output flag is missing the '.txt' extension
			fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> \"something\" standard")
			return
		} else {
			// If no output flag, treat arguments as text and banner
			text = os.Args[1]
			banner = strings.ToLower(os.Args[2])
		}
	}

	// Handle the case when there are 2 arguments
	if len(os.Args) == 2 {
		text = os.Args[1] // Only the text is provided; use default banner
	}

	// Split the input text by '\n' to handle newlines in the input
	sepText := strings.Split(text, "\\n")

	// Read the banner file (e.g., 'standard.txt', 'shadow.txt')
	file, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Banner not found")
		return
	}

	// Convert the file content to a string and normalize line endings
	str := string(file)
	str = strings.Replace(str, "\r\n", "\n", -1) // Replace Windows line endings with Unix ones

	// Split the banner file content into lines
	lines := strings.Split(str, "\n")

	if !hasOutput {
		// If no output file specified, print the ASCII art to the console
		funcs.PrintAsciiArt(sepText, lines)
	} else {
		// If output file is specified, create or overwrite the file
		file, err := os.Create(outputfile)
		if err != nil {
			fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> \"something\" standard")
			return
		}
		// Generate the ASCII art and write it to the file
		funcs.PrintAsciiArtToFile(file, sepText, lines)
	}
}
