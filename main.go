package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-color/printingasciipackage"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Fprintln(os.Stderr, `Usage: go run . [OPTION] [STRING]
EX: go run . --color=<color> <letters to be colored> "something"`)
		return
	}
	if len(os.Args) == 2 {
		// fmt.Println("Here")
		if strings.Contains(os.Args[1], "--color=") {
			fmt.Fprintln(os.Stderr, `Usage: go run . [OPTION] [STRING]
EX: go run . --color=<color> <letters to be colored> "something"`)
			return
		}

		color := FindingColor("=red")
		ap := printingasciipackage.PrintingAscii(os.Args[1], "standard.txt", color, "") // normal pattern
		fmt.Print(ap)
	}

	if len(os.Args) == 3 {
		color := FindingColor(strings.ToLower(os.Args[1]))
		ap := printingasciipackage.PrintingAscii(os.Args[2], "standard.txt", color, os.Args[2]) // color without specific letters
		fmt.Print(ap)
	}
	if len(os.Args) == 4 {
		color := FindingColor(strings.ToLower(os.Args[1]))
		ap := printingasciipackage.PrintingAscii(os.Args[3], "standard.txt", color, os.Args[2]) // color with specified letters
		fmt.Print(ap)
	}
}

func color(color string) string {
	mapColor := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magneta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}
	// fmt.Println(mapColor[color])
	return mapColor[color]
}

func FindingColor(s string) string {
	colorWanted := ""
	for i, ch := range s { // finds color user wants
		if ch == '=' {
			colorWanted = s[i+1:]
			break
		}
	}
	
	// fmt.Println("Here")
	color := color(colorWanted)
	// fmt.Println(color, "Here")
	return color
}
