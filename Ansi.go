package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	colorArg := args[0]
	subStringArg := args[1]
	fmt.Println(subStringArg)
	givenString := args[2]
	var colorWanted string
	for i, ch := range colorArg {
		if ch == '=' {
			colorWanted = colorArg[i+1:]
			break
		}
	}
	color(givenString, subStringArg, colorWanted)
}
func color(str, sub, color string) {
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
	
   

		for key, value := range mapColor {
			if key == color {
				for _,ch:=range str{
					for _,char:=range sub{
						if char==ch{
						modify:=func(r rune)rune{
							  if r==char{
								colored,s:=value,char
				                fmt.Println(colored)
								return s
							}
							return r
						}
						result:=strings.Map(modify,str)
					fmt.Println(result)
					}
					}
					
				}
			}
		}
		
	}


