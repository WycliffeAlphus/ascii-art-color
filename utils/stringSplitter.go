package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringSplitter(s string) []int {
	res := []int{}

	editedString := strings.TrimLeft(s, "rgb(")
	editedString = strings.TrimRight(editedString, ")")

	stringSplit := strings.Split(editedString, ",")
	if len(stringSplit) != 3 {
		fmt.Fprintln(os.Stderr, `Array values should be three`)
		os.Exit(0)
	}

	for _, ch := range stringSplit {
		number, err := strconv.Atoi(strings.TrimSpace(ch))
		res = append(res, number)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting string %v to integer\n", ch)
			os.Exit(0)
		}
	}

	return res
}
