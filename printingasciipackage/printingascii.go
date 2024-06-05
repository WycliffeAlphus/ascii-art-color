package printingasciipackage
import (
	"fmt"
	"os"
	"strings"
	"ascii-art-color/mapPackage"
	"ascii-art-color/utils"
)
// Reads input text,gets the pattern convert it to ascii art
	text = strings.ReplaceAll(text, "\n", "\\n")
	res := ""
	for i := 0; i < len(text); {
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'a' {
			fmt.Fprintf(os.Stderr, "error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'b' {
			l := len(text) - 2
			if i == 0 {
				text = text[i+2:]
			} else if i == l {
				text = text[:l]
			} else {
				text = text[:i-1] + text[i+2:]
				i = 0
			}
			continue
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 't' {
			text = text[:i] + "   " + text[i+2:]
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'v' {

			fmt.Fprintf(os.Stderr, "error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'f' {
			fmt.Fprintf(os.Stderr, "error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'r' {
			fmt.Fprintf(os.Stderr, "error: Special character %v%v is not supported \n", string(text[i]), string(text[i+1]))
			os.Exit(1)
		}
		if i+1 < len(text) && text[i] > 127 {
			fmt.Fprintln(os.Stderr, "error: Only ascii characters are allowed")
			os.Exit(1)
		}
		i++
	}
	lines := strings.Split(text, "\\n")
	asciiMap := mapPackage.AsciiMapping(patternFile)

	
	if !strings.Contains(text,letters){
		count := 0
	for _, word := range lines { // case of multiple newlines
		if word == "" {
			count++
			if count < len(lines) {
				res += "\n"
			}
		} else {
			for n := 0; n < 8; n++ {
				for _, ch := range word {
					if len(letters) == 0{
						res += color + asciiMap[ch][n] + "\033[0m"
					} else {
						if isInside(string(ch), letters){
							res +=  color + asciiMap[ch][n] + "\033[0m" 
							continue
						}
						res += asciiMap[ch][n]
						}
				}
				res += "\n"
			}
		}
		return res
	} 
	}else{
			alphaArray:=colorfeature.FindSubStringIndex(text,letters)
			count:=0
			for wordIndex, word := range lines { // case of multiple newlines
				if word == "" {
					count++
					if count < len(lines) {
						res += "\n"
					}
				} else {
					for n := 0; n < 8; n++ {
						for runeIndex, ch := range word {
		                    if wordIndex > 0{
								checkIndex := 0
								for x := 0; x < wordIndex; x++{
									checkIndex += len(lines[x]) + 1
								}
								if  colorfeature.IndicesInArr(alphaArray, checkIndex+runeIndex+wordIndex){
									res += color + asciiMap[ch][n] + "\033[0m"
									continue
								}
								res += asciiMap[ch][n]
							} else {
								if  colorfeature.IndicesInArr(alphaArray, runeIndex){
									res += color + asciiMap[ch][n] + "\033[0m"
									continue
								}
								res += asciiMap[ch][n]
							}
		
							
						}
						res += "\n"
					}
				}
			}
			
		}
		return res
	}
		
	func isInside(alpha, s1 string) bool {
		for _, ch := range s1{
			if alpha == string(ch){
				return true
			}
		}
		return false
	}


	