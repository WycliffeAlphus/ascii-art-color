package main

import (
	"fmt"
	"os"
)

// FindSubString function finds and highlights a substring within a main string.
func FindSubString(mainString, subString string) string {
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
		return msg
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

	return resString
}

func main() {
	// mainString := "A God morning is A Good God glory Morning Good"
	// subString := "oo"
	mainString := os.Args[1]
	subString := os.Args[2]
	fmt.Println(FindSubString(mainString, subString))
}