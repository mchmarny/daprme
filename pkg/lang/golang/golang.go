package golang

import (
	"context"
	"fmt"
	"os/exec"
	"path"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

// Golang represents language specific functionality for go
type Golang struct{}

// GetProjectConfig describes the project artifacts
func (g *Golang) GetProjectConfig() *model.Project {
	return &model.Project{
		Main: "main.go",
	}
}

// InitializeProject initializes project
func (g *Golang) InitializeProject(ctx context.Context, dir, usr, app string) error {
	// init the modules
	appDir := path.Join(dir, app)
	m := fmt.Sprintf("github.com/%s/%s", usr, app)
	if err := execCmd(appDir, "go", "mod", "init", m); err != nil {
		return err
	}

	// remove unused imports and format the code
	if err := execCmd(appDir, "goimports", "-w", "main.go"); err != nil {
		return err
	}

	// tidy after the format
	if err := execCmd(appDir, "go", "mod", "tidy"); err != nil {
		return err
	}

	return nil
}

func execCmd(appDir, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir = appDir
	if err := c.Run(); err != nil {
		return errors.Wrapf(err, "error executing command:\n%s", c.String())
	}
	return nil
}
