package wordlist

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func Get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %v", err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error: %v", err)
		return "", err
	}

	shuffledWords, err := shuffleNFormat(body)
	if err != nil {
		fmt.Printf("error: %v", err)
		return "", err
	}

	return shuffledWords, nil
}

func shuffleNFormat(b []byte) (string, error) {
	var words []string
	err := json.Unmarshal(b, &words)
	if err != nil {
		return "", err
	}

	shuffle(words)

	formatted := strings.Join(words, " ")
	return formatted, nil
}

func shuffle(s []string) {
	// make sure it's "more randomized"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range s {
		j := r.Intn(len(s))
		s[i], s[j] = s[j], s[i]
	}
}

