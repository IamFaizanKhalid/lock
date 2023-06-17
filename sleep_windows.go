package lock

import (
	"time"
)

func listen(ch chan<- struct{}) {
	defer close(ch)

	// Create a ticker that ticks every second.
	ticker := time.NewTicker(1 * time.Second)

	lastTime := time.Now()

	// Loop forever.
	for {
		// Wait for the next tick.
		newTime := <-ticker.C

		// Check if the current time is more than 5 seconds after the previous time.
		if newTime.Sub(lastTime).Seconds() > 5 {
			// If so, the system has woken up.
			ch <- struct{}{}
		}

		// Set the previous time to the current time.
		lastTime = newTime
	}
}
