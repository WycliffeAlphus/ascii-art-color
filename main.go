package main

import (
	"fmt"
	"os"

	"ascii-art/printingasciipackage"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, `
USAGE:
	go run . <text> [OPTION]
		-sh			shadow.txt
		-t			thinkertoy.txt
		default		standard.txt

SEE ALSO:
	README.md
		`)
		os.Exit(1)
	}
	args := os.Args[1]

	flag := ""

	if args == "" {
		return
	}

	if len(os.Args) == 3 { // checks that correct flag is assigned as argument
		flag = os.Args[2]
		if flag == "-sh" || flag == "-t" {
			flag = os.Args[2]
		} else {
			fmt.Println(`error: Expected flags are -sh or -t
			 To read standard.txt,do not include any flag
			 You used `, flag)
			os.Exit(1)
		}
	}

	patternFile := ""

	switch flag {
	case "-sh":
		patternFile = "shadow.txt"
	case "-t":
		patternFile = "thinkertoy.txt"
	default:
		patternFile = "standard.txt"

	}

	ap := printingasciipackage.PrintingAscii(args, patternFile)
	fmt.Print(ap)
}
