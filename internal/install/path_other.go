//go:build !windows

package install

import "fmt"

func ensureWindowsPath(dir string) error {
	return fmt.Errorf("ensureWindowsPath called on non-windows build")
}
