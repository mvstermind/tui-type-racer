package main

// https://github.com/eiannone/keyboard?tab=readme-ov-file
import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/mvstermind/terminal-typeracer/key"
	getwords "github.com/mvstermind/terminal-typeracer/wordlist"
)

func main() {

	wordSlice := generateList()

	fmt.Println(wordSlice)

	words := strings.Join(wordSlice, " ")
	// do this when it's fucking got full list
	key.KeyListen(words)

}

func randomSeed() string {
	lenght := rand.Intn(5) + 3

	word := 5

	list_seed := fmt.Sprintf("https://random-word-api.herokuapp.com/word?number=%v&length=%v", word, lenght)
	return list_seed

}

func generateList() []string {

	var wordList []string

	// it's 4 cuz i want to have 20 words
	// random seed generates 5 words
	totalAmount := 4
	for i := 0; i < totalAmount; i++ {

		wordSeed := randomSeed()

		word, err := getwords.Get(wordSeed)
		if err != nil {
			fmt.Printf("error: %v", err)

		}
		wordList = append(wordList, word)
	}

	return wordList
}
