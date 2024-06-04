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
	alphaArray := findSubStringIndex(text, "l")


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

					if wordIndex > 0{
						checkIndex := 0
						for x := 0; x < wordIndex; x++{
							checkIndex += len(lines[x]) + 1
						}
						if indicesInArr(alphaArray, checkIndex+runeIndex+wordIndex){
							res += "\033[34m" + asciiMap[ch][n] + "\033[0m"
							continue
						}
						res += asciiMap[ch][n]
					} else {
						if indicesInArr(alphaArray, runeIndex){
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


func isSubstring(s1, s2 string) bool {

	count := 0
	for i:=0; i < len(s1);{
		for j:=0; j < len(s2);{
			if s1[i] == s2[j]{
				count ++
				j++	
			}
			i ++
			if i == len(s1){
				break
			}
		}
	}
	if count == len(s1) && count == len(s2){
		return true
	}
	return false
}


func findSubStringIndex(mainString, subString string) []int {
	indices := []int{}
	for i:=0; i<len(mainString);i++{
		if  i+len(subString) <= len(mainString) && isSubstring(subString, mainString[i:i+len(subString)]){
			for j:=i ; j < i+len(subString);j++{
				
				indices= append(indices, j)
			}
		}
	}
	return indices
}

func indicesInArr(arr []int, num int) bool {
	for _,x := range arr{
		
			if x == num{
				return true
			
		}
	}
	return false
}


func main() {
	fmt.Print(PrintingAscii("yellow\nblue", "standard.txt"))
}