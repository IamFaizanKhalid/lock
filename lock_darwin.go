package lock

import (
	"os/exec"
	"strings"
)

func isScreenLocked() bool {
	// This will output a plist with session information
	output, _ := exec.Command("ioreg", "-n", "Root", "-d1", "-a").Output()

	// This will check if output["IOConsoleUsers"][0]["CGSSessionScreenIsLocked"] exists
	return strings.Contains(string(output), "CGSSessionScreenIsLocked")
}
