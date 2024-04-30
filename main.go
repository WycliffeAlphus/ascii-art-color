package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, `
Usage:
	ascii-art <text>

Display ASCII graphics art.
		`)
		os.Exit(1)
	}
	args := os.Args[1]

	if args == "\\n" {
		fmt.Println()
		return
	} else if args == "" {
		return
	}

	patternFile := "standard.txt"
	PrintingAscii(args, patternFile)
}

// PrintingAscii given a banner file and some ASCII text to print, prints the graphics of the ASCII text
func PrintingAscii(text, patternFile string) {
	lines := strings.Split(text, "\\n")
	asciiMap := AsciiMapping(patternFile)

	count := 0
	for _, word := range lines {
		if word == "" {
			count++
			if count < len(lines) {
				fmt.Println()
			}
		} else {
			for n := 0; n < 8; n++ {
				for _, ch := range word {
					fmt.Print(asciiMap[ch][n])
				}
				fmt.Println()
			}
		}

	}

}

// AsciiMapping given a banner file, reads all graphics representations of the ASCII characters and
// returns a map of the ASCII character to the graphics representations of the ASCII character
func AsciiMapping(patternFile string) map[rune][]string {
	testfile, err := os.ReadFile(patternFile)
	if err != nil {
		fmt.Println("Error reading the file")
		os.Exit(1)
	}

	splitted := strings.Split(string(testfile), "\n")

	asciiMapping := make(map[rune][]string)
	startAscii := ' '
	for i := 1; i < len(splitted); {
		arrayString := []string{}
		for j := 0; j < 8; j++ {
			arrayString = append(arrayString, splitted[i])
			i++
		}
		i++
		asciiMapping[startAscii] = arrayString
		startAscii++
	}

	return asciiMapping
}
