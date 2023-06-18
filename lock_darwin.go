package lock

import (
	"os/exec"
	"strings"
)

func isScreenLocked() bool {
	cmd, _ := exec.Command("ioreg", "-n", "Root", "-d1", "-a").Output()

	return strings.Contains(string(cmd), "CGSSessionScreenIsLocked")
}
