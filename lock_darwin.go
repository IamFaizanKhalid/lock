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

	wasLocked := false

	// Loop forever.
	for {
		// Wait for the next tick.
		<-ticker.C

		isLocked := checkCGSSession()

		if !wasLocked && isLocked {
			ch <- struct{}{}
		}

		wasLocked = isLocked
	}
}

func checkCGSSession() bool {
	cmd, _ := exec.Command("ioreg", "-n", "Root", "-d1", "-a").Output()

	return strings.Contains(string(cmd), "CGSSessionScreenIsLocked")
}
