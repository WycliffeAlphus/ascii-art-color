package utils

/*This function is used to return ANSI escape codes for different colors. 
ANSI escape codes are sequences of characters used to control the formatting,
color, and other output options on text terminals.*/
func color(color string) string {
	mapColor := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magneta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"orange":  "\033[38;5;214m",
	}

	return mapColor[color]
}
