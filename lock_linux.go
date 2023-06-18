package lock

import (
	"os/exec"
	"strings"
)

func isScreenLocked() bool {
	switch os.Getenv("XDG_CURRENT_DESKTOP") {
	case "Unity":
		return check(
			"gdbus call -e -d com.canonical.Unity -o /com/canonical/Unity/Session -m com.canonical.Unity.Session.IsLocked",
			"true",
		)
	case "KDE":
	case "XFCE":
	case "LXQt":
	case "MATE":
	case "Cinnamon":
	case "LXDE":
	case "Deepin":
	default: // default to gnome
		return check(
			"gdbus call -e -d org.gnome.SessionManager -o /org/gnome/SessionManager/Presence -m org.freedesktop.DBus.Properties.Get org.gnome.SessionManager.Presence status",
			"3>",
		)
	}
}

func check(cmdStr, contains string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cmd := strings.Split(cmdStr, " ")

	// This will output a list of tasks currently running
	output, _ := exec.CommandContext(ctx, cmd[0], cmd[1:]...).Output()

	// This will check if LogonUI.exe is in the list
	return strings.Contains(string(output), contains)
}
