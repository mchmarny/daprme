package lang

import (
	"context"
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
func Make(ctx context.Context, app *model.App, usr, dir string) error {
	if app == nil {
		return errors.Errorf("app instance required")
	}
	if usr == "" {
		return errors.Errorf("user required")
	}

	// lang
	langProvider, err := MakeConfigurable(app.Meta.Lang)
	if err != nil {
		return errors.Wrap(err, "Error initializing language")
	}

	// create directories
	outDir := path.Join(dir, app.Meta.Name)
	if err := createDir(outDir); err != nil {
		return errors.Wrapf(err, "Error creating project dir: %s.", outDir)
	}

	configDir := path.Join(outDir, "config")
	if err := createDir(configDir); err != nil {
		return errors.Wrapf(err, "Error creating config dir: %s.", configDir)
	}

	projectConfig := langProvider.GetProjectConfig()
	if projectConfig == nil {
		return errors.Wrap(err, "Unable to create project config provider.")
	}

	// run templates
	for k, v := range projectConfig.Templates {
		makefilePath := path.Join(outDir, v)
		templateMakePath := path.Join(templateDir, app.Meta.Lang, k)
		if err := execTemplate(app, makefilePath, templateMakePath); err != nil {
			return errors.Wrapf(err, "Error executing template: %s", templateMakePath)
		}
	}

	// components
	for _, c := range app.Bindings {
		if err := addComponent(c, configDir); err != nil {
			return errors.Wrap(err, "Error creating binding component.")
		}
	}

	// pubsub
	for _, c := range app.PubSubs {
		if err := addComponent(c, configDir); err != nil {
			return errors.Wrap(err, "Error creating pubsub component.")
		}
	}

	// client components
	for _, c := range app.Components {
		if err := addComponent(c, configDir); err != nil {
			return errors.Wrap(err, "Error creating client component.")
		}
	}

	// init
	if err := langProvider.InitializeProject(ctx, dir, usr, app.Meta.Name); err != nil {
		return errors.Wrap(err, "Error initializing project")
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
