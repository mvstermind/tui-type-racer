package key

import (
	"math"
	"time"
)

func calculateWPM(word string, duration time.Duration) int {
	wordLen := len(word)
	seconds := duration.Seconds()
	wpm := (wordLen / int(seconds)) * 12
	return wpm
}

func accuracy(wlist []rune, missedChar int) float64 {
	num := math.Round((float64(len(wlist)-missedChar) / float64(len(wlist))) * 100)
	return num

}
