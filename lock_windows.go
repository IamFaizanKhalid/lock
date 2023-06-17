package lock

import (
	"os/exec"
	"strings"
	"time"
)

func listen(ch chan<- struct{}) {
	defer close(ch)

	// Create a ticker that ticks every second.
	ticker := time.NewTicker(1 * time.Second)

	lastTime := time.Now()
	wasLocked := false

	for {
		newTime := <-ticker.C

		// FIXME: first condition will work after system wakes from sleep

		isLocked := newTime.Sub(lastTime).Seconds() > 5 || isLogonUIRunnning()

		if !wasLocked && isLocked {
			ch <- struct{}{}
		}

		wasLocked = isLocked
		lastTime = newTime
	}
}

func isLogonUIRunnning() bool {
	cmd, _ := exec.Command("tasklist").Output()

	return strings.Contains(string(cmd), "LogonUI")
}
