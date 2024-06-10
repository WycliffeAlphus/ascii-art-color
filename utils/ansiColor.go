package utils

import "strconv"

func ansiColor(asciiNum int) string {
	return "\033[38;5;" + strconv.Itoa(asciiNum) + "m "
}
