package key

import "time"

func calculateWPM(word string, duration time.Duration) int {
	wordLen := len(word)
	seconds := duration.Seconds()
	wpm := (wordLen / int(seconds)) * 12
	return wpm
}
