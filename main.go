package main

import (
	"github.com/mvstermind/terminal-typeracer/key"
)

func main() {

	wordSlice := generateList()
	words := sliceToString(wordSlice)

	// do this when it's fucking got full list
	key.KeyListen(words)
}
