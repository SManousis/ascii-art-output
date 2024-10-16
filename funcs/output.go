package funcs

import (
	"log"
	"os"
)

func PrintAsciiArtToFile(newFile *os.File, sentences []string, textFile []string) {
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
