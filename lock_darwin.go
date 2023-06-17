package lock

import (
	"github.com/IamFaizanKhalid/lock/darwin"
	"time"
)

func listen(ch chan<- struct{}) {
	defer close(ch)

	// Create a ticker that ticks every second.
	ticker := time.NewTicker(1 * time.Second)

	wasLocked := false

	// Loop forever.
	for {
		// Wait for the next tick.
		<-ticker.C

		isLocked := darwin.Check()

		if !wasLocked && isLocked {
			ch <- struct{}{}
		}

		wasLocked = isLocked
	}
}
