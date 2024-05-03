package mapPackage_test

import (
	"ascii-art/mapPackage"
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	filetext := "standard.txt"

	out := mapPackage.Map(filetext)
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
