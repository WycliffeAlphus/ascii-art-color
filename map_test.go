package main

import (
	//"ascii-art/mappackage"
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	filetext := "standard.txt"

	out := AsciiMapping(filetext)
	expected := []string{
		"           ",
		`    /\     `,
		`   /  \    `,
		`  / /\ \   `,
		` / ____ \  `,
		`/_/    \_\ `,
		"           ",
		"           ",
	}
	if !reflect.DeepEqual(out['A'], expected) {
		fmt.Println("Test unsuccessful")
	} else {
		fmt.Println("Test passed successfully")
	}
}
