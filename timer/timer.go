package timer

import (
	"time"
)

var startTime time.Time

func StartTimer() {
	startTime = time.Now()
}

func StopTimer() time.Duration {
	elapsedTime := time.Since(startTime)
	return elapsedTime
}
