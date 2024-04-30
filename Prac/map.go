package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileRead := "standard.txt"
	readStrings := os.Args[1]

	if readStrings == "\\n" {
		fmt.Println()
		return
	} else if readStrings == "" {
		return
	} else {
		splitString := strings.Split(readStrings, "\\n")
		ReadStrings(splitString, fileRead)

	}
}

func Map(patternFile string) map[rune][]string {
	ourMap := make(map[rune][]string)
	file, err := os.ReadFile(patternFile)
	if err != nil {
		fmt.Println("File read unsuccesfully")
		os.Exit(1)
	}
	fileContent := strings.Split(string(file), "\n")
	startAscii := ' '
	i := 1
	for i < len(fileContent) {
		pattern := []string{}
		for counter := 0; counter < 8; counter++ {
			pattern = append(pattern, fileContent[i])
			i++
		} //counter increment
		i++
		ourMap[startAscii] = pattern
		startAscii++
	}
	return ourMap

}

func ReadStrings(stringSlice []string, fileRead string) {
	mappedResult := Map(fileRead)
	count := 0
	for _, stringValue := range stringSlice {
		if stringValue == "" {
			count++
			if count < len(stringSlice) {
				fmt.Println()
			}
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range stringValue {
					fmt.Print(mappedResult[char][i])
				}
				fmt.Println()
			}
		}
	}

}
