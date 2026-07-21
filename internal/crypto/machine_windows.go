//go:build windows

package crypto

import (
	"golang.org/x/sys/windows/registry"
)

func machineID() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE)
	if err != nil {
		return "win-unknown"
	}
	defer k.Close()
	id, _, err := k.GetStringValue("MachineGuid")
	if err != nil || id == "" {
		return "win-unknown"
	}
	return id
}
