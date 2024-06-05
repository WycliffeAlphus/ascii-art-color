package printingasciipackage_test

import (
	"ascii-art-color/printingasciipackage"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestPrintingAscii(t *testing.T) {
	args := "hello" // string to be patterned and printed
	file := "standard.txt"
	out := printingasciipackage.PrintingAscii(args, file,"\033[31m", "")
	expected := []string{
		"\033[31m _      \033[0m\033[31m       \033[0m\033[31m _  \033[0m\033[31m _  \033[0m\033[31m        \033[0m",
		"\033[31m| |     \033[0m\033[31m       \033[0m\033[31m| | \033[0m\033[31m| | \033[0m\033[31m        \033[0m",
		"\033[31m| |__   \033[0m\033[31m  ___  \033[0m\033[31m| | \033[0m\033[31m| | \033[0m\033[31m  ___   \033[0m",
		"\033[31m|  _ \\  \033[0m\033[31m / _ \\ \033[0m\033[31m| | \033[0m\033[31m| | \033[0m\033[31m / _ \\  \033[0m",
		"\033[31m| | | | \033[0m\033[31m|  __/ \033[0m\033[31m| | \033[0m\033[31m| | \033[0m\033[31m| (_) | \033[0m",
		"\033[31m|_| |_| \033[0m\033[31m \\___| \033[0m\033[31m|_| \033[0m\033[31m|_| \033[0m\033[31m \\___/  \033[0m",
		"\033[31m        \033[0m\033[31m       \033[0m\033[31m    \033[0m\033[31m    \033[0m\033[31m        \033[0m",
		"\033[31m        \033[0m\033[31m       \033[0m\033[31m    \033[0m\033[31m    \033[0m\033[31m        \033[0m",
	}

	//PrintingAscii function returns a string of len 256
	//we have 8 lines ,so length of each line is (256/8)=32
	//we access from index 0 to 255
	splitOut := strings.Split(out, "\n")

	if !reflect.DeepEqual(splitOut[0], expected[0]) {
		fmt.Println("Test  line 1 failed")
		t.Errorf("got\n %v, want %v", splitOut[0], expected[0])
		return
	} else {
		fmt.Println(" line 1 Test passed successfully")
	}
	if !reflect.DeepEqual(splitOut[1], expected[1]) {
		fmt.Println("Test line 2 failed")
		t.Errorf("got\n %v, want %v", splitOut[1], expected[1])
		return
	} else {
		fmt.Println(" line 2 Test passed successfully")
	}
	if !reflect.DeepEqual(splitOut[2], expected[2]) {
		fmt.Println("Test line 3 failed")
		t.Errorf("got\n %v, want %v", splitOut[2], expected[2])
		return
	} else {
		fmt.Println(" line 3 Test passed successfully")
	}
	if !reflect.DeepEqual(splitOut[3], expected[3]) {
		fmt.Println("Test line 4 failed")
		t.Errorf("got\n %v, want %v", splitOut[3], expected[3])
		return
	} else {
		fmt.Println(" line 4 Test passed successfully")
	}
	if !reflect.DeepEqual(splitOut[4], expected[4]) {
		fmt.Println("Test line 5 failed")
		t.Errorf("got\n %v, want %v", splitOut[4], expected[4])
		return
	} else {
		fmt.Println(" line 5 Test passed successfully")
	}
	if !reflect.DeepEqual(splitOut[5], expected[5]) {
		fmt.Println("Test line 6 failed")
		t.Errorf("got\n %v, want %v", splitOut[5], expected[5])
		return
	} else {
		fmt.Println(" line 6 Test passed successfully")
	}
	if !reflect.DeepEqual(splitOut[6], expected[6]) {
		fmt.Println("Test line 7 failed")
		t.Errorf("got\n %v, want %v", splitOut[6], expected[6])
		return
	} else {
		fmt.Println(" line 7 Test passed successfully")
	}
	if !reflect.DeepEqual(splitOut[7], expected[7]) {
		fmt.Println("Test  line 8 failed")
		t.Errorf("got\n %v, want %v", splitOut[7], expected[7])
		return
	} else {
		fmt.Println(" line 8 Test passed successfully")
	}
}
