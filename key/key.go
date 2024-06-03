package key

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/mvstermind/terminal-typeracer/timer"
)

func KeyListen(word string) {
	// mistakes counter
	mistake := 0

	// convert words into slice of runes
	wordList := []rune(word)

	// current character
	index := 0
	fmt.Println(NotTypedChar.Render(word))

	// start timer as soon as it's printed
	timer.StartTimer()

	// flag to track if a mistake has been printed
	mistakePrinted := false

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
		if char == wordList[index] || wordList[index] == rune(keyboard.KeySpace) {
			fmt.Printf("%s", ValidChar.Render(string(wordList[index])))

			// reset mistake flag if a correct character is typed
			mistakePrinted = false

			index++
			if index == len(wordList) {
				fmt.Printf("\n")

				m := fmt.Sprintf("Mistakes: %v", mistake)
				fmt.Println(Mistakes.Render(m))

				duration := timer.StopTimer()
				wpm := calculateWPM(word, duration)
				fmt.Println("WPM: ", wpm)
				break
			}

		} else {
			// print mistake only once
			if !mistakePrinted {
				fmt.Printf("%s", Mistakes.Render(string(wordList[index])))
				index++
				mistake++
				mistakePrinted = true
			}
		}
	}
}
