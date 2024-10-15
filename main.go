package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if the number of command-line arguments is correct
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("ERROR! Usage: go run . [STRING] [BANNER] || Example: go run . \"test\" standard")
		return
	}

	// The text to be converted to ASCII art is the first argument
	text := string(os.Args[1])

	// Set the default banner to "standard" if not provided
	banner := "standard"
	if len(os.Args) == 3 {
		// Convert the banner name to lowercase to ensure consistency
		banner = strings.ToLower(os.Args[2])
	}

	// Split the input text by '\n' to handle newlines in the input
	sepText := strings.Split(text, "\\n")

	// Read the banner file (e.g., standard.txt, shadow.txt)
	file, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("Banner not found")
		return
	}

	// Split the contents of the banner file into lines
	str := string(file)
	str = strings.Replace(str, "\r\n", "\n", -1)
	lines := strings.Split(str, "\n")

	// Call the function to print the ASCII art representation
	printAsciiArt(sepText, lines)
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
			for i := 0; i < len(word); i++ {
				// Loop through each line in the banner text file
				for lineIndex, line := range textFile {
					// Calculate the position of the character in the banner file
					// Each character occupies 9 lines in the banner file (including an empty line)
					// (int(word[i]) - 32) gives the index of the character in the ASCII table starting from space (' ')
					// Multiply by 9 to account for the lines per character
					// Add 'h' to get the specific line of the character's ASCII art
					if lineIndex == (int(word[i])-32)*9+h {
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
