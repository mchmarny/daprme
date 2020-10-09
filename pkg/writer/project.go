package writer

import (
	"fmt"
	"os"
	"path"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

const (
	outputDirPerm   = os.FileMode(int(0755))
	templateDir     = "template"
	templateCompDir = "component"
)

// Make creates project
func Make(app *model.App) error {
	if app == nil {
		return errors.Errorf("App instance required.")
	}

	// create directories
	outDir := path.Join(".", app.Name)
	if err := createDir(outDir); err != nil {
		return errors.Wrapf(err, "Error creating project dir: %s.", outDir)
	}

	configDir := path.Join(outDir, "config")
	if err := createDir(configDir); err != nil {
		return errors.Wrapf(err, "Error creating config dir: %s.", configDir)
	}

	// run templates
	makefilePath := path.Join(outDir, "Makefile")
	templateMakePath := path.Join(templateDir, "make.tmpl")
	if err := execTemplate(app, makefilePath, templateMakePath); err != nil {
		return errors.Wrap(err, "Error creating makefile.")
	}

	mainPath := path.Join(outDir, "main.go")
	templateMainPath := path.Join(templateDir, "main.tmpl")
	if err := execTemplate(app, mainPath, templateMainPath); err != nil {
		return errors.Wrap(err, "Error creating main.go.")
	}

	dockerPath := path.Join(outDir, "Dockerfile")
	templateDockerPath := path.Join(templateDir, "docker.tmpl")
	if err := execTemplate(app, dockerPath, templateDockerPath); err != nil {
		return errors.Wrap(err, "Error creating dockerfile.")
	}

	ignorePath := path.Join(outDir, ".gitignore")
	templateIgnorePath := path.Join(templateDir, "ignore.tmpl")
	if err := execTemplate(app, ignorePath, templateIgnorePath); err != nil {
		return errors.Wrap(err, "Error creating ignorefile.")
	}

	// components
	for _, c := range app.Bindings {
		if err := addComponent(&c.Component, configDir); err != nil {
			return errors.Wrap(err, "Error creating binding component.")
		}
	}

	// pubsub
	for _, c := range app.Pubsubs {
		if err := addComponent(&c.Component, configDir); err != nil {
			return errors.Wrap(err, "Error creating pubsub component.")
		}
	}

	// client components
	for _, c := range app.Components {
		if err := addComponent(c, configDir); err != nil {
			return errors.Wrap(err, "Error creating client component.")
		}
	}

	return nil
}

func addComponent(c interface{}, dir string) error {
	comp, ok := c.(model.Componentable)
	if !ok {
		return errors.Errorf("Note a Componentable: %T", c)
	}

	tmplPath := path.Join(templateDir, templateCompDir, fmt.Sprintf("%s.tmpl", comp.GetType()))
	outPath := path.Join(dir, fmt.Sprintf("%s.yaml", comp.GetName()))
	if err := execTemplate(comp, outPath, tmplPath); err != nil {
		return errors.Wrap(err, "Error creating ignorefile.")
	}

	return nil
}

func createDir(path string) error {
	if !pathExists(path) {
		if err := os.MkdirAll(path, outputDirPerm); err != nil {
			return errors.Wrapf(err, "Error creating config dir: %s.", path)
		}
	}
	return nil
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
