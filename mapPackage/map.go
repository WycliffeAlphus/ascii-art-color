package mapPackage

import (
	"fmt"
	"os"
	"strings"
)

// func checkFile(){
// 	args := os.Args
// 	if len(args) != 2 {
// 		fmt.Fprint(os.Stderr, `Usage<>
// 		go run . (stringValue) -> go run . "Hello"
// 		`)
// 		os.Exit(1)
// 	} else {
// 		// input := args[1]
// 		file := "standard.txt"
// 		resultingMap := Map(file)
// 		fmt.Println(resultingMap)
// 	}
// }

var fileContentString []string

func ReadStandard(fileName string) []string {
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading the file")
		os.Exit(1)
	} else {
		fileContentString = strings.Split(string(fileContents), "\n")
	}
	return fileContentString
}

func Map(fileName string) map[rune][]string {
	InputContents := ReadStandard(fileName)
	ourMap := make(map[rune][]string)
	start := ' '
	i := 1
	for i < len(InputContents) {
		pattern := []string{}
		for counter := 0; counter < 8; counter++ {
			pattern = append(pattern, InputContents[i])
			i++
		}
		i++
		ourMap[start] = pattern
		start++
	}
	//fmt.Println(ourMap)
	return ourMap
}
