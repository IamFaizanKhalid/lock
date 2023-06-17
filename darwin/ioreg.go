package darwin

import (
	"os/exec"

	"howett.net/plist"
)

func Check() bool {
	cmd, _ := exec.Command("ioreg", "-n", "Root", "-d1", "-a").Output()

	out := &output{}
	plist.Unmarshal(cmd, out)

	return out.IsScreenLocked()
}

type output struct {
	IOConsoleLocked bool
	IOConsoleUsers  []map[string]interface{}
}

func (n output) IsScreenLocked() bool {
	if len(n.IOConsoleUsers) == 0 {
		return false
	}

	_, ok := n.IOConsoleUsers[0]["CGSSessionScreenIsLocked"]
	return ok
}
