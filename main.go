package main

import "fmt"

func main() {

	s := []string{"test ", "word ", "less go "}

	// get single character from array
	for _, v := range s {
		for i := 0; i < len(v); i++ {
			fmt.Printf("%c\n", v[i])
		}

	}
}
