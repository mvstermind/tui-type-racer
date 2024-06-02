package key

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func KeyListen(word string) {
	/*
		this has to run forever
		compare pressed key with word[i]
		then lipgloss will do some like fucking coloring or something
		i dunno
	*/

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("you pressed %q\n", char)
	}
}
