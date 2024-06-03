package key

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// use valid char style if letters matched letters typed
// use currentChar to indicate index position
// use notTypedChar to show how many charcters are left to type

func KeyListen(word string) {

	// convert words into slice of runes
	word_list := []rune(word)

	// current character
	index := 0
	fmt.Println(NotTypedChar.Render(word))

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
			fmt.Printf("%s", ValidChar.Render(string(word_list[index])))
			// fmt.Println("Literka sie zgadza") // for debugs
			index++

			if index == len(word_list) {
				fmt.Printf("\n")
				break
			}
		}
		// for debugs
		// fmt.Printf("you pressed %q\n", char)
	}
}
