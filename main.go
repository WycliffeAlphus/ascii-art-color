package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-color/printingasciipackage"

	"ascii-art-color/utils"
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

		color := utils.FindingColor(strings.ToLower(os.Args[1]))
		ap := printingasciipackage.PrintingAscii(os.Args[2], "standard.txt", color, "")
		fmt.Print(ap)
	}
	if len(os.Args) == 4 {
		color := utils.FindingColor(strings.ToLower(os.Args[1]))
		ap := printingasciipackage.PrintingAscii(os.Args[3], "standard.txt", color, os.Args[2])
		fmt.Print(ap)
	}
}
