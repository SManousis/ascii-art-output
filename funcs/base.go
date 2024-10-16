package funcs

import "fmt"

func PrintAsciiArt(sentences []string, textFile []string) {
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
