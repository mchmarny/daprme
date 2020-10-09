package cmd

import (
	"fmt"
	"os/exec"
	"path"

	"github.com/pkg/errors"
)

// InitProject initializes go project
func InitProject(usr, app string) error {
	if err := initModule(usr, app); err != nil {
		return err
	}
	return tidyModule(app)
}

func initModule(usr, app string) error {

	m := fmt.Sprintf("github.com/%s/%s", usr, app)
	c := exec.Command("go", "mod", "init", m)
	c.Dir = path.Join(".", app)

	if err := c.Run(); err != nil {
		return errors.Wrapf(err, "error executing command:\n%s", c.String())
	}

	return nil
}

func tidyModule(app string) error {
	c := exec.Command("go", "mod", "tidy")
	c.Dir = path.Join(".", app)

	if err := c.Run(); err != nil {
		return errors.Wrapf(err, "error executing command:\n%s", c.String())
	}

	return nil
}
