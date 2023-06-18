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

	// This will output registry with session information
	output, _ := exec.CommandContext(ctx, "ioreg", "-n", "Root", "-d1").Output()

	// This will check if output["IOConsoleUsers"][0]["CGSSessionScreenIsLocked"] exists
	return strings.Contains(string(output), "CGSSessionScreenIsLocked")
}
