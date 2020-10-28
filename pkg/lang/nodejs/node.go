package nodejs

import (
	"context"
	"path"

	"github.com/dapr-templates/daprme/pkg/cmd"
	"github.com/dapr-templates/daprme/pkg/model"
)

const (
	httpPortDefault = 3000
	grpcPortDefault = 50505
)

// NodeJs represents language specific functionality for node
type NodeJs struct{}

// GetProjectConfig describes the project artifacts
func (n *NodeJs) GetProjectConfig() *model.Project {
	return &model.Project{
		Main:     "main.go",
		PortGRPC: grpcPortDefault,
		PortHTTP: httpPortDefault,
		Templates: map[string]string{
			"docker.tmpl":  "Dockerfile",
			"ignore.tmpl":  ".gitignore",
			"main.tmpl":    "app.js",
			"make.tmpl":    "Makefile",
			"package.tmpl": "package.json",
		},
	}
}

// InitializeProject initializes project
func (n *NodeJs) InitializeProject(ctx context.Context, dir string, app *model.App) error {
	// init the modules
	appDir := path.Join(dir, app.Meta.Name)

	// npp
	if err := cmd.Exec(appDir, "npm", "install"); err != nil {
		return err
	}

	return nil
}
