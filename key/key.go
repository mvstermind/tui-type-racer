package key

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

/*
this has to run forever
compare pressed key with word[i]
then lipgloss will do some like fucking coloring or something
i dunno
*/
func KeyListen(word string) {
	// convert words into slice of runes
	word_list := []rune(word)

	// current character
	index := 0

	for {
		// get current character
		char, key, err := keyboard.GetSingleKey()

		if err != nil {
			fmt.Println(err)
		}

		if key == keyboard.KeyEsc {
			fmt.Println("Exiting...")
			break
		}

		// compare current character with key pressed
		if char == word_list[index] || word_list[index] == rune(keyboard.KeySpace) {
			// fmt.Println("Literka sie zgadza") // for debugs
			index++

			if index == len(word_list) {
				fmt.Println("Koniec listy")
				break
			}
		}
		// for debugs
		// fmt.Printf("you pressed %q\n", char)
	}
}
