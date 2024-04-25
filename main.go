package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	patternFile := "standard.txt"
	testfile, err := os.ReadFile(patternFile)
	if err != nil {
		fmt.Println("Error reading the file")
	}

	stringTestfile := string(testfile)
	splitted := strings.Split(stringTestfile, "\n")
	asciimapping := make(map[rune][]string)
	startAscii := 32
	for i := 1; i < len(splitted); {
		arrayString := []string{}
		for j := 0; j < 8; j++ {
			arrayString = append(arrayString, splitted[i])
			i++
		}
		i++
		asciimapping[rune(startAscii)] = arrayString
		startAscii++
	}
	args := os.Args[1]

	if args == "\\n" {
		fmt.Println()
		return
	} else if args == "" {
		return
	}
	newArgs := strings.Split(args, "\\n")

	count := 0
	for _, word := range newArgs {
		if word == "" {
			count++
			if count < len(newArgs) {
				fmt.Println()
			}
		} else {
			for n := 0; n < 8; n++ {

				for _, ch := range word {

					fmt.Print(asciimapping[ch][n])

				}
				fmt.Println()
			}
		}

	}
}
