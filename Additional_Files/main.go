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
	// text = FindSubString(text, "el")
	alphaArray := findArr(text, "ow")
	fmt.Println(alphaArray)
	lines := strings.Split(text, "\\n")
	asciiMap := mapPackage.AsciiMapping(patternFile)

	count := 0
	// alphaCount := 0 
	for j, word := range lines { // case of multiple newlines
		if word == "" {
			count++
			if count < len(lines) {
				res += "\n"
			}
		} else {
			for n := 0; n < 8; n++ {
				for i, ch := range word {

					if j > 0{
						checkIndex := 0
						for x := 0; x < j; x++{
							checkIndex += len(lines[x]) + 1
						}
						if intIsInArr(alphaArray, checkIndex+i+j){
							res += "\033[34m" + asciiMap[ch][n] + "\033[0m"
							continue
						}
						res += asciiMap[ch][n]
					} else {
						if intIsInArr(alphaArray, i){
							res += "\033[34m" + asciiMap[ch][n] + "\033[0m"
							continue
						}
	
						// currentIndex = len(lines[j])
						res += asciiMap[ch][n]
					}

					
				}
				res += "\n"
			}
		}
	}
	return res
}


func isDuplicate(s1, s2 string) bool {
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


func findArr(mainString, subString string) []int {
	res := []int{}
	for i:=0; i<len(mainString);i++{
		fmt.Println("Here")
		if  i+len(subString) <= len(mainString) && isDuplicate(subString, mainString[i:i+len(subString)]){
			for j:=i ; j < i+len(subString);j++{
				
				res = append(res, j)
			}
		}
	}
	fmt.Println(res)
	return res
}




func FindSubString(mainString, subString string) [][]int {
	s1 := []string{}

	resString := ""

	intArr := []int{}

	intArrArr := [][]int{}

	// Convert mainString to a slice of strings
	for i := 0; i < len(mainString); i++ {
		s1 = append(s1, string(mainString[i]))
	}

	// Iterate through the main string
	for i := 0; i < len(s1); {
		// Iterate through the substring
		for j := 0; j < len(subString); {
			// Check if characters match
			if string(subString[j]) == s1[i] {
				intArr = append(intArr, i)

				currIndex := intArr[len(intArr)-1]
				prevIndex := 0
				if len(intArr) < 2 {
					prevIndex = i - 1
				} else {
					prevIndex = intArr[len(intArr)-2]
				}

				indexDiff := currIndex - prevIndex

				j++

				// Reset if characters are not adjacent
				if indexDiff > 1 && len(subString) > 1 {
					j = 0
					intArr = []int{}
				}

			}

			// Store the index array if the substring is found
			if j == len(subString) {
				intArrArr = append(intArrArr, intArr)
			}

			i++

			// Break if end of main string is reached
			if i == len(s1) {
				break
			}
		}
	}

	// Return message if substring not found
	if len(intArrArr) < 1 {
		msg := "$" + subString + "$ is not present in $" + mainString + "$"
		fmt.Println(msg)
		os.Exit(0)
		// return msg
	}

	// Highlight the substring in the main string
	for _, ch := range intArrArr {
		for i := 0; i < len(ch); i++ {
			s1[ch[i]] = "\033[34m" + s1[ch[i]] + "\033[0m"
		}
	}

	// Construct the resulting string
	for i := 0; i < len(s1); i++ {
		resString += s1[i]
	}

	// return resString
	return intArrArr
}



func intIsInArr(arr []int, num int) bool {
	for _,x := range arr{
		
			if x == num{
				return true
			
		}
	}
	return false
}


func main() {
	fmt.Print(PrintingAscii("Hellow \nworld \n owmy \nmyow", "standard.txt"))
}