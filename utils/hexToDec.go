package utils

import (
	"fmt"
	"os"
	"strconv"
)

func HexToDec(s string) []int {
	for i, ch := range s {
		if ch == '#' && i != 0 {
			fmt.Fprintf(os.Stderr, "%v is not a valid hex color code\n", s)
			os.Exit(0)
		}
	
	}
	newString := s[1:]
	if len(newString) != 6 {
		
		fmt.Fprintf(os.Stderr, "%v is not a valid hex color> code\n", s)
		os.Exit(0)
	}
	var hexArray []string
	var word string
	var count int
	for _, char := range newString {
		count++
		word += string(char)
		if count%2 == 0 {
			hexArray = append(hexArray, word)
			word = ""

		}

	}
	var hexIntSlice []int
	for _, ch := range hexArray {
		dec, err := strconv.ParseInt(ch, 16, 64)
		if err != nil{
			fmt.Fprintf(os.Stderr, "%v is not a valid hex color code\n", s)
			os.Exit(0)
		}
		hexIntSlice = append(hexIntSlice, int(dec))
	}
	fmt.Println(hexIntSlice)
	return hexIntSlice
}
