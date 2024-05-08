package main

// import "golang.org/x/text/unicode/rangetable"

func main() {
	s1 := "Hello"

	for i:= range s1{
		if i == 2{
			s1[2] = "d"
		}
	}
}