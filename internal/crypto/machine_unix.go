//go:build !windows

package crypto

import (
	"os"
	"strings"
)

func machineID() string {
	data, err := os.ReadFile("/etc/machine-id")
	if err == nil {
		id := strings.TrimSpace(string(data))
		if id != "" {
			return id
		}
	}
	data, err = os.ReadFile("/var/lib/dbus/machine-id")
	if err == nil {
		id := strings.TrimSpace(string(data))
		if id != "" {
			return id
		}
	}
	return "unix-unknown"
}
