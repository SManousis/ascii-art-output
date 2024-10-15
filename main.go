package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Check if the number of command-line arguments is correct
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] || EX: go run . --output=<fileName.txt> something standard")
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
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] || EX: go run . --output=<fileName.txt> something standard")
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
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] || EX: go run . --output=<fileName.txt> something standard")
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
	file, err := os.ReadFile(banner + ".txt")
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
		printAsciiArt(sepText, lines)
	} else {
		// If output file is specified, create or overwrite the file
		file, err := os.Create(outputfile)
		if err != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] || EX: go run . --output=<fileName.txt> something standard")
			return
		}
		// Generate the ASCII art and write it to the file
		printAsciiArtToFile(file, sepText, lines)
	}
}

func printAsciiArt(sentences []string, textFile []string) {
	// Loop through each sentence (split by '\n')
	for i, word := range sentences {
		if word == "" {
			// If the sentence is empty (i.e., there was a '\n'), print an empty line
			if i != 0 {
				fmt.Println()
			}
			continue
		}
		// Loop through each of the 8 lines of ASCII art for the characters
		for h := 1; h < 9; h++ {
			// Loop through each character in the word
			for j := 0; j < len(word); j++ {
				// Loop through each line in the banner text file
				for lineIndex, line := range textFile {
					// Calculate the position of the character in the banner file
					// Each character occupies 9 lines in the banner file (including an empty line)
					// (int(word[j]) - 32) gives the index of the character in the ASCII table starting from space (' ')
					// Multiply by 9 to account for the lines per character
					// Add 'h' to get the specific line of the character's ASCII art
					if lineIndex == (int(word[j])-32)*9+h {
						// Print the line corresponding to the current character and line number
						fmt.Print(line)
					}
				}
			}
			// Move to the next line after printing all characters for the current line 'h'
			fmt.Println()
		}
	}
}

func printAsciiArtToFile(newFile *os.File, sentences []string, textFile []string) {
	// Loop through each sentence (split by '\n')
	for i, word := range sentences {
		if word == "" {
			// If the sentence is empty, write an empty line to the file
			if i != 0 {
				_, err := newFile.WriteString("\n")
				if err != nil {
					log.Fatal(err)
				}
			}
			continue
		}
		// Loop through each of the 8 lines of ASCII art for the characters
		for h := 1; h < 9; h++ {
			// Loop through each character in the word
			for j := 0; j < len(word); j++ {
				// Loop through each line in the banner text file
				for l, line := range textFile {
					// Calculate the position of the character in the banner file
					if l == int(word[j]-32)*9+h {
						// Write the line corresponding to the current character and line number to the file
						_, err := newFile.WriteString(line)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
			// Write a newline character to move to the next line in the file
			_, err := newFile.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	// Write an extra newline at the end of the file
	_, err := newFile.WriteString("\n")
	if err != nil {
		log.Fatal(err)
	}
}
