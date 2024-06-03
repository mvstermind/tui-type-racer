package main

// api for requests
// https://random-word-api.herokuapp.com/word?number=10

// keylistyemer lib
// https://github.com/eiannone/keyboard?tab=readme-ov-file
import (
	"github.com/mvstermind/terminal-typeracer/key"
)

func main() {
	word := "kurwa"

	key.KeyListen(word)

}
