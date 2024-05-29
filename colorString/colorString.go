package colorString

import (
	"strings"
)

func colorStr(text, color, letters string) string {
	reset := "\033[0m"
	if letters == "" {
		return color + text + reset
	}
	var r []rune
	for _, char := range letters {
		r = append(r, char)
	}

	coloredS := ""

	if len(r) == 1 {
		for _, char := range text {
			if strings.ContainsRune(letters, char) {
				coloredS += color + string(char) + reset
			} else {
				coloredS += string(char)
			}
		}
	}
	i := 0
	for i < len(text) {
		if strings.HasPrefix(text[i:], letters) {
			coloredS += color + letters + reset
			i += len(letters)
		} else {
			coloredS += string(text[i])
			i++
		}
	}
	return coloredS
}
