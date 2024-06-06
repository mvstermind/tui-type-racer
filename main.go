package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/mvstermind/terminal-typeracer/key"
	"github.com/mvstermind/terminal-typeracer/wordlist"
)

func main() {

	wordSlice := generateList()
	words := sliceToString(wordSlice)

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
	wordChannel := make(chan string)
	var wg sync.WaitGroup

	// use keyword argument to genreate len of list
	osarg := os.Args[1]
	o, err := strconv.Atoi(osarg)
	if err != nil {
		fmt.Println("error: ", err)
		return nil
	}
	// this is so it will generate 20 words in the list
	totalAmount := o

	// start totalAmount goroutines
	for i := 0; i < totalAmount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			wordSeed := randomSeed()
			word, err := wordlist.Get(wordSeed)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}
			wordChannel <- word
		}()
	}

	// close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(wordChannel)
	}()

	// collect results from the channel
	var wordList []string
	for word := range wordChannel {
		wordList = append(wordList, word)
	}

	return wordList
}

func sliceToString(s []string) string {

	w := strings.Join(s, " ")
	w = strings.TrimLeft(w, "[")
	w = strings.TrimRight(w, "]")
	return w

}
