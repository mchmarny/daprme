package project

import (
	"context"
	"fmt"
	"os/exec"
	"path"

	"github.com/pkg/errors"
)

// Initialize initializes go project
func Initialize(ctx context.Context, usr, app string) error {
	if err := initModule(ctx, usr, app); err != nil {
		return err
	}
	return tidyModule(ctx, app)
}

func initModule(ctx context.Context, usr, app string) error {
	m := fmt.Sprintf("github.com/%s/%s", usr, app)
	c := exec.Command("go", "mod", "init", m)
	c.Dir = path.Join(".", app)

	if err := c.Run(); err != nil {
		return errors.Wrapf(err, "error executing command:\n%s", c.String())
	}

	return nil
}

func tidyModule(ctx context.Context, app string) error {
	c := exec.Command("go", "mod", "tidy")
	c.Dir = path.Join(".", app)

	if err := c.Run(); err != nil {
		return errors.Wrapf(err, "error executing command:\n%s", c.String())
	}

	return nil
}
