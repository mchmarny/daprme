package writer

import (
	"os"
	"path"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

const (
	outputDirPerm      = os.FileMode(int(0755))
	templateMainPath   = "template/main.tmpl"
	templateMakePath   = "template/make.tmpl"
	templateDockerPath = "template/docker.tmpl"
	templateIgnorePath = "template/ignore.tmpl"
)

// Make creates project
func Make(app *model.App) error {
	if app == nil {
		return errors.Errorf("App instance required.")
	}

	outDir := path.Join(".", app.Name)
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outDir, outputDirPerm); err != nil {
			return errors.Wrapf(err, "Error creating project dir: %s.", outDir)
		}
	}

	configDir := path.Join(outDir, "config")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, outputDirPerm); err != nil {
			return errors.Wrapf(err, "Error creating project config dir: %s.", configDir)
		}
	}

	makefilePath := path.Join(outDir, "Makefile")
	if err := execTemplate(app, makefilePath, templateMakePath); err != nil {
		return errors.Wrap(err, "Error creating makefile.")
	}

	mainPath := path.Join(outDir, "main.go")
	if err := execTemplate(app, mainPath, templateMainPath); err != nil {
		return errors.Wrap(err, "Error creating main.go.")
	}

	dockerPath := path.Join(outDir, "Dockerfile")
	if err := execTemplate(app, dockerPath, templateDockerPath); err != nil {
		return errors.Wrap(err, "Error creating dockerfile.")
	}

	ignorePath := path.Join(outDir, ".gitignore")
	if err := execTemplate(app, ignorePath, templateIgnorePath); err != nil {
		return errors.Wrap(err, "Error creating ignorefile.")
	}

	return nil
}
