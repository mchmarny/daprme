package golang

import (
	"context"
	"fmt"
	"path"

	"github.com/dapr-templates/daprme/pkg/cmd"
	"github.com/dapr-templates/daprme/pkg/model"
)

const (
	httpPortDefault = 8080
	grpcPortDefault = 50505
)

// Golang represents language specific functionality for go
type Golang struct{}

// GetProjectConfig describes the project artifacts
func (g *Golang) GetProjectConfig() *model.Project {
	return &model.Project{
		Main:     "main.go",
		PortGRPC: grpcPortDefault,
		PortHTTP: httpPortDefault,
		Templates: map[string]string{
			"docker.tmpl": "Dockerfile",
			"ignore.tmpl": ".gitignore",
			"main.tmpl":   "main.go",
			"make.tmpl":   "Makefile",
		},
	}
}

// InitializeProject initializes project
func (g *Golang) InitializeProject(ctx context.Context, dir string, app *model.App) error {
	// init the modules
	appDir := path.Join(dir, app.Meta.Name)
	m := fmt.Sprintf("github.com/%s/%s", app.Meta.Owner, app.Meta.Name)
	if err := cmd.Exec(appDir, "go", "mod", "init", m); err != nil {
		return err
	}

	// ensure goimports
	if err := cmd.Exec(appDir, "go", "get", "golang.org/x/tools/cmd/goimports"); err != nil {
		return err
	}

	// remove unused imports and format the code
	if err := cmd.Exec(appDir, "goimports", "-w", "main.go"); err != nil {
		return err
	}

	// tidy after the format
	if err := cmd.Exec(appDir, "go", "mod", "tidy"); err != nil {
		return err
	}

	return nil
}
