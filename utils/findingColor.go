package utils

import (
	"fmt"
	"os"
	"strings"
)

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
	} else if strings.Contains(colorWanted,"#"){
		hexInt:=HexToDec(colorWanted)
		ansiNum:=RGBtoAnsi(hexInt[0],hexInt[1],hexInt[2])
        color:=ansiColor(ansiNum)
		return color

	}else {
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
