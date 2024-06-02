package key

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func KeyListen(word string) {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		go func() {
			fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		}()

		if key == keyboard.KeyEsc {
			break
		}
	}
}
