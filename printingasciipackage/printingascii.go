package printingasciipackage

import("strings"
"fmt"
"os"
"ascii-art/mapPackage")

// Reads input text,gets the pattern convert it to ascii art

func PrintingAscii(text, patternFile string) string {
	res:=""
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
		if text[i] > 127{
			fmt.Println("error: only Ascii Characters above 127 are not supported")
			os.Exit(0)
		}
	}
	lines := strings.Split(text, "\\n")
	asciiMap := mapPackage.AsciiMapping(patternFile)

	count := 0
	for _, word := range lines {
		if word == "" {
			count++
			if count < len(lines) {
				res += "\n"
			}
		} else {
			for n := 0; n < 8; n++ {
				for _, ch := range word {
					res += asciiMap[ch][n]
				}
				res += "\n"
			}

		}

	}
	return res
}