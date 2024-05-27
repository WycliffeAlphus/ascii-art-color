package main

import (
	"fmt"
	"os"
	// "strings"
)

func main() {
	args := os.Args[1:]
	colorArg := args[0]
	subStringArg := args[1]
	givenString := args[2]
	var colorWanted string
	for i, ch := range colorArg { //finds color user wants
		if ch == '=' {
			colorWanted = colorArg[i+1:]
			break
		}
	}

	color := color(colorWanted) // gets the ANSI code of the color
	for _, char := range givenString {
		toColor := false
		for _, ch :=  range subStringArg{
			if ch == char {
				toColor = true
			}
		}
		if toColor {
			fmt.Printf("%s%s" + "\033[0m", color,string(char)) //colors then resets
			continue
		} else {
			fmt.Printf("%s",string(char)) //prints non-colored runes
			continue
		}
	}
	fmt.Println()
}
func color(color string)  string{
	mapColor := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magneta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"test": "Hello world",
	}
	return mapColor[color]
}


