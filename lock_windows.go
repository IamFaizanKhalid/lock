package lock

import (
	"context"
	"os/exec"
	"strings"
	"time"
)

func isScreenLocked() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// This will output a list of tasks currently running
	output, _ := exec.CommandContext(ctx, "tasklist").Output()

	// This will check if LogonUI.exe is in the list
	return strings.Contains(string(output), "LogonUI")
}
