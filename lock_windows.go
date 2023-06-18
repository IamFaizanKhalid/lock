package lock

import (
	"os/exec"
	"strings"
)

func isScreenLocked() bool {
	// This will output a list of tasks currently running
	cmd, _ := exec.Command("tasklist").Output()

	// This will check if LogonUI.exe is in the list
	return strings.Contains(string(cmd), "LogonUI")
}
