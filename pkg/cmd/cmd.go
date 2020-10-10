package cmd

import (
	"os/exec"

	"github.com/pkg/errors"
)

// Exec executers the provided command
func Exec(appDir, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir = appDir
	if err := c.Run(); err != nil {
		return errors.Wrapf(err, "error executing command:\n%s", c.String())
	}
	return nil
}
