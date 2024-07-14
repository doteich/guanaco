package utils

import (
	"os/exec"
)

func setSysProcAttr(cmd *exec.Cmd) {
	// No special attributes needed for non-Windows systems
}
