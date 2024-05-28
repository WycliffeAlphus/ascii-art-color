package colorString

import "strings"

func colorStr(text, color, letters string) string { // The text is the string to be colored, color is the specified color
	// letters are the characters to be colored if specified

	reset := "\033[0m" // for resetting the applied color after use

	if letters == "" {
		return color + text + reset // If letter is not specified
	}

	coloredS := ""

	for _, char := range text {
		if strings.ContainsRune(letters, char) {
			coloredS += color + string(char) + reset
		} else {
			coloredS += string(char)
		}
	}
	return coloredS
}
