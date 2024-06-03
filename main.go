package main

// api for requests
// https://random-word-api.herokuapp.com/word?number=10

// keylistyemer lib
// https://github.com/eiannone/keyboard?tab=readme-ov-file
import (
	"fmt"

	"github.com/mvstermind/terminal-typeracer/key"
	"github.com/mvstermind/terminal-typeracer/wordlist"
)

func main() {
	word, err := wordlist.Get("https://random-word-api.herokuapp.com/word?number=10&length=5")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	key.KeyListen(word)

}
