package lock

import (
	"os/exec"
	"strings"
)

func isScreenLocked() bool {
	cmd, _ := exec.Command("tasklist").Output()

	return strings.Contains(string(cmd), "LogonUI")
}
