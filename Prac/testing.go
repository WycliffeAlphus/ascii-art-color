package main

import (
	"fmt"
	"strings"
)

func main2() {

	s := strings.Split("Hello\n\nWorld", "\n")
	s1 := strings.Split("Hello\nWorld", "\n")
	fmt.Println(s)
	fmt.Println(s1[1])
	fmt.Println(s1)
}
