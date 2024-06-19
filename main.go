package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-color/printingasciipackage"

	"ascii-art-color/utils"
)

func main() {
	// Checks for presence of standard file exists
	fileInfo, err := os.Stat("standard.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Checks incase of any modification done on standard text
	// Incase of modification the byte size will be not eqaul to 6623
	// Triggering a re-download
	if fileInfo.Size() != 6623 {
		err := utils.DownloadStd()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Fprintln(os.Stderr, `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"`)
		return
	}
	if len(os.Args) == 2 {
		if strings.Contains(os.Args[1], "--color=") {
			fmt.Fprintln(os.Stderr, `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"`)
			return
		}
		ap := printingasciipackage.PrintingAscii(os.Args[1], "standard.txt", "\033[0m", "")
		fmt.Print(ap)
	}

	if len(os.Args) == 3 {

		color := utils.FindingColor(strings.ToLower(os.Args[1]))
		ap := printingasciipackage.PrintingAscii(os.Args[2], "standard.txt", color, os.Args[2])
		fmt.Print(ap)
	}
	if len(os.Args) == 4 {
		color := utils.FindingColor(strings.ToLower(os.Args[1]))
		ap := printingasciipackage.PrintingAscii(os.Args[3], "standard.txt", color, os.Args[2])
		fmt.Print(ap)
	}
}
