//go:build linux
// +build linux

package utils

import (
	"os/exec"
)

func setSysProcAttr(cmd *exec.Cmd) {
	// No special attributes needed for non-Windows systems
}
