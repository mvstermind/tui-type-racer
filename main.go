package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sentence := "this is the test sentence"

	fmt.Printf("%s\n", sentence)
	//	lenSen := len(sentence)

	b := bufio.NewReader(os.Stdin)
	input, err := b.ReadString('\n')
	if err != nil {
		return
	}
	re, err := compInputs(sentence, input)

	fmt.Println(re)
}

func compInputs(o string, u string) (string, error) {

	var oirignal int
	for _, v := range o {

	}
	return "", nil

}
