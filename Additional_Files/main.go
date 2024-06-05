package main

import (
	"fmt"
	"os"
	"strings"

	"additionalfiles/mapPackage"
)

// Reads input text,gets the pattern convert it to ascii art

func PrintingAscii(text, patternFile string) string {
	text = strings.ReplaceAll(text, "\n", "\\n")
	res := ""

	for i := 0; i < len(text); {
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'a' {
			fmt.Fprintf(os.Stderr, "error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'b' {
			l := len(text) - 2
			if i == 0 {
				text = text[i+2:]
			} else if i == l {
				text = text[:l]
			} else {
				text = text[:i-1] + text[i+2:]
				i = 0
			}
			continue
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 't' {
			text = text[:i] + "   " + text[i+2:]
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'v' {

			fmt.Fprintf(os.Stderr, "error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'f' {
			fmt.Fprintf(os.Stderr, "error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'r' {
			fmt.Fprintf(os.Stderr, "error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] > 127 {
			fmt.Fprintln(os.Stderr, "error: Only ascii characters are allowed")
			os.Exit(1)
		}
		i++
	}

	lines := strings.Split(text, "\\n")
	asciiMap := mapPackage.AsciiMapping(patternFile)
	alphaArray := findSubStringIndex(text, "ow")

	count := 0
	for wordIndex, word := range lines { // case of multiple newlines
		if word == "" {
			count++
			if count < len(lines) {
				res += "\n"
			}
		} else {
			for n := 0; n < 8; n++ {
				for runeIndex, ch := range word {
					if wordIndex > 0 {
						checkIndex := 0
						for x := 0; x < wordIndex; x++ {
							checkIndex += len(lines[x]) + 1
						}
						if indicesInArr(alphaArray, checkIndex+runeIndex+wordIndex) {
							res += "\033[33m" + asciiMap[ch][n] + "\033[0m"
							continue
						}
						res += asciiMap[ch][n]
					} else {
						if indicesInArr(alphaArray, runeIndex) {
							res += "\033[33m" + asciiMap[ch][n] + "\033[0m"
							continue
						}
						res += asciiMap[ch][n]
					}
				}
				res += "\n"
			}
		}
	}
	return res
}

func findSubStringIndex(mainString, subString string) []int {
	indices := []int{}
	for i := 0; i <= len(mainString)-len(subString); i++ {
		fmt.Println(mainString[i : i+len(subString)])
		if mainString[i:i+len(subString)] == subString {
			for j := i; j < i+len(subString); j++ {
				indices = append(indices, j)
			}
		}
	}
	fmt.Println(indices)
	return indices
}

func indicesInArr(arr []int, num int) bool {
	for _, x := range arr {
		if x == num {
			return true
		}
	}
	return false
}

func main() {
	fmt.Print(PrintingAscii("hellow \n from owalla\n no wahalla\n nowala", "standard.txt"))
}
