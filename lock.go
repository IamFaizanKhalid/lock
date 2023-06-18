package lock

import "time"

func routine(ch chan<- Event) {
	defer close(ch)

	// Create a ticker that ticks every second.
	ticker := time.NewTicker(1 * time.Second)

	lastTime := time.Now()
	wasLocked := false

	// Loop forever.
	for {
		// Wait for the next tick.
		newTime := <-ticker.C

		// Check the time difference between last time and current time.
		timeMissing := newTime.Sub(lastTime).Seconds() > 5

		// Check if the session is locked in different OS
		isLocked := isScreenLocked()

		// If there is a difference, that means the device was on sleep?
		if timeMissing {
			ch <- Event{
				Time:   lastTime, // because after this, the device was not responding
				Locked: true,
			}
		} else if !wasLocked && isLocked {
			// Device was not locked when checked last time, now it is locked.
			ch <- Event{
				Time:   newTime,
				Locked: true,
			}
		} else if wasLocked && !isLocked {
			// Device was locked when checked last time, now it is unlocked.
			ch <- Event{
				Time:   newTime,
				Locked: false,
			}
		}

		// Updating values for the next iteration.
		wasLocked = isLocked
		lastTime = newTime
	}
}
