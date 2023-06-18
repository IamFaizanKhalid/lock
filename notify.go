package lock

import "time"

func Notify() <-chan struct{} {
	ch := make(chan struct{})

	go routine(ch)

	return ch
}

func routine(ch chan<- struct{}) {
	defer close(ch)

	// Create a ticker that ticks every second.
	ticker := time.NewTicker(1 * time.Second)

	lastTime := time.Now()
	wasLocked := false

	for {
		newTime := <-ticker.C

		// FIXME: first condition will work after system wakes from sleep

		isLocked := newTime.Sub(lastTime).Seconds() > 5 || isScreenLocked()

		if !wasLocked && isLocked {
			ch <- struct{}{}
		}

		wasLocked = isLocked
		lastTime = newTime
	}
}
