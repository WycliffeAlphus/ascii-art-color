package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, `
USAGE:
	go run . <text> [OPTION]
		-st			standard.txt
		-standard	standard.txt
		-sh			shadow.txt
		-shadow		shadow.txt
		-t			thinkertoy.txt
		default		standard.txt

SEE ALSO:
	README.md
		`)
		os.Exit(1)
	}
	args := os.Args[1]

	// flag := os.Args[2]

	flag := ""

	if args == "\\n" {
		fmt.Println()
		return
	} else if args == "" {
		return
	}

	if len(os.Args) == 2 {
		flag = "-st"
	}

	if len(os.Args) == 3 {
		flag = os.Args[2]
		if flag == "-sh" || flag == "-shadow" || flag == "-t" || flag == "-st" || flag == "-standard" {
			flag = os.Args[2]
		} else {
			fmt.Println("error: Expected flags are -sh, -shadow, -t, -st, -standard. You used ", flag)
			os.Exit(1)
		}
	}

	patternFile := ""

	switch flag {
	case "-sh":
		patternFile = "shadow.txt"
	case "-shadow":
		patternFile = "shadow.txt"
	case "-t":
		patternFile = "thinkertoy.txt"
	case "-thinkertoy.txt":
		patternFile = "thinkertoy.txt"
	case "-st":
		patternFile = "standard.txt"
	case "-standard":
		patternFile = "standard.txt"
	default:
		patternFile = "standard.txt"

	}

	PrintingAscii(args, patternFile)

	// fmt.Print(PrintingAsciiTest(args, patternFile))
}

// PrintingAscii given a banner file and some ASCII text to print, prints the graphics of the ASCII text
func PrintingAscii(text, patternFile string) {
	for i := range text {
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'a' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'b' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 't' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'v' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'f' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'r' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
	}
	lines := strings.Split(text, "\\n")
	asciiMap := AsciiMapping(patternFile)

	count := 0
	for _, word := range lines {
		if word == "" {
			count++
			if count < len(lines) {
				fmt.Println()
			}
		} else {
			for n := 0; n < 8; n++ {
				for _, ch := range word {
					fmt.Print(asciiMap[ch][n])
				}
				fmt.Println()
			}
		}

	}

}

func PrintingAsciiTest(text, patternFile string) []string {
	res := ""
	ap := []string{}
	for i := range text {
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'a' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'b' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 't' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'v' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'f' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'r' {
			fmt.Printf("error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
	}
	lines := strings.Split(text, "\\n")
	asciiMap := AsciiMapping(patternFile)

	count := 0
	for _, word := range lines {
		if word == "" {
			count++
			if count < len(lines) {
				//fmt.Println()
				res += "\n"
			}
		} else {
			for n := 0; n < 8; n++ {
				for _, ch := range word {
					//fmt.Print(asciiMap[ch][n])
					res += asciiMap[ch][n]

				}
				// ap = append(ap, res)
				//fmt.Println()
				res += "\n"
			}
			ap = append(ap, res)
		}

	}
	return ap
}

// AsciiMapping given a banner file, reads all graphics representations of the ASCII characters and
// returns a map of the ASCII character to the graphics representations of the ASCII character
func AsciiMapping(patternFile string) map[rune][]string {
	var splitted []string
	if patternFile == "thinkertoy.txt" {
		testfile, err := os.ReadFile(patternFile)
		if len(testfile) == 0{
			fmt.Println("error:",patternFile,"is empty")
			os.Exit(0)
		}
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		splitted = strings.Split(string(testfile), "\r\n")
	} else {
		testfile, err := os.ReadFile(patternFile)
		if len(testfile) == 0{
			fmt.Println("error:",patternFile,"is empty")
			os.Exit(0)
		}
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		splitted = strings.Split(string(testfile), "\n")
	}

	// testfile, err := os.ReadFile(patternFile)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	// splitted := strings.Split(string(testfile), "\n")

	asciiMapping := make(map[rune][]string)
	startAscii := ' '
	for i := 1; i < len(splitted); {
		arrayString := []string{}
		for j := 0; j < 8; j++ {
			arrayString = append(arrayString, splitted[i])
			i++
		}
		i++
		asciiMapping[startAscii] = arrayString
		startAscii++
	}

	return asciiMapping
}
