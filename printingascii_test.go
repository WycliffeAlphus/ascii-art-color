// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"os"
// 	"reflect"
// 	"testing"
// )

// func TestAsciiMapping(t *testing.T) {
// 	//setting up data to mimic command line arguments
// 	args := "Hello"
// 	//redirect the standard output to a buffer
// 	var buf bytes.Buffer
// 	oldStdout := os.Stdout
// 	os.Stdout = &buf
// 	defer func() { os.Stdout = oldStdout }()

// 	file := "standard.txt"
// 	PrintingAscii(args, file)

// 	out := buf.String()

// 	expected := []string{
// 		` _              _   _          `,
// 		`| |            | | | |         `,
// 		`| |__     ___  | | | |   ___   `,
// 		`|  _ \   / _ \ | | | |  / _ \  `,
// 		`| | | | |  __/ | | | | | (_) | `,
// 		`|_| |_|  \___| |_| |_|  \___/  `,
// 		"                               ",
// 		"                               ",
// 	}
// 	if !reflect.DeepEqual(out, expected) {
// 		fmt.Println("Test passed unsuccessfully")
// 	} else {
// 		fmt.Println("Test passed successfully")
// 	}
// }

package main

import "testing"

func TestAsciiMapping(t *testing.T) {
	var tests = []struct{
		name string
		input string
		want string
	}{
		{"Test One","Hello",` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
								
								`},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			ans := PrintingAsciiTest(tt.input,"standard.txt")
			if ans != tt.want{
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}