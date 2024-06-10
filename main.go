package main

import (
	"fmt"
	"os"
	"strconv"
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
		if strings.Contains(os.Args[1], "--color=") {
			fmt.Fprintln(os.Stderr, `Usage: go run . [OPTION] [STRING]
EX: go run . --color=<color> <letters to be colored> "something"`)
			return
		}
		ap := printingasciipackage.PrintingAscii(os.Args[1], "standard.txt", "\033[0m", "")
		fmt.Print(ap)
	}

	if len(os.Args) == 3 {

		color := FindingColor(strings.ToLower(os.Args[1]))
		ap := printingasciipackage.PrintingAscii(os.Args[2], "standard.txt", color, "")
		fmt.Print(ap)
	}
	if len(os.Args) == 4 {
		color := FindingColor(strings.ToLower(os.Args[1]))
		ap := printingasciipackage.PrintingAscii(os.Args[3], "standard.txt", color, os.Args[2])
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
		"orange":  "\033[38;5;214m",
	}

	return mapColor[color]
}

func ansiColor(asciiNum int) string {
	return "\033[38;5;" + strconv.Itoa(asciiNum) + "m "
}

func FindingColor(s string) string {
	colorWanted := ""
	for i, ch := range s {
		if ch == '=' {
			colorWanted = s[i+1:]
			break
		}
	}

	if strings.Contains(colorWanted, "rgb") {
		rgbarray := stringSplitter(colorWanted)
		for _, ch := range rgbarray {
			if ch > 255 {
				fmt.Println("Invalid RGB value;", ch)
				os.Exit(0)
			}
		}
		ansiNum := RGBtoAnsi(rgbarray[0], rgbarray[1], rgbarray[2])
		color := ansiColor(ansiNum)
		return color
	} else {
		color := color(colorWanted)
		if color == "" {
			fmt.Println(`Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"
				`)
			os.Exit(0)
		}
		return color
	}
}

func RGBtoAnsi(R, G, B int) int {
	r := R / 51
	g := G / 51
	b := B / 51

	colorNum := 16 + 36*r + 6*g + b

	return colorNum
}

func stringSplitter(s string) []int {
	res := []int{}

	editedString := strings.TrimLeft(s, "rgb(")
	editedString = strings.TrimRight(editedString, ")")

	stringSplit := strings.Split(editedString, ",")

	for _, ch := range stringSplit {
		number, err := strconv.Atoi(strings.TrimSpace(ch))
		res = append(res, number)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting string %v to integer\n", ch)
			os.Exit(0)
		}
	}

	return res
}
