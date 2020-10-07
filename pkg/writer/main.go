package writer

import (
	"os"
	tmpl "text/template"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

const (
	tmplMainPath = "template/main.tmpl"
)

// MakeMain creates app main in specified path
func MakeMain(app *model.App, out string) error {
	if app == nil {
		return errors.New("invalid app input")
	}

	if out == "" {
		return errors.New("output path required")
	}

	f, err := os.Create(out)
	if err != nil {
		return errors.Wrapf(err, "error creating output file: %s", out)
	}
	defer f.Close()

	t, err := tmpl.ParseFiles(tmplMainPath)
	if err != nil {
		return errors.Wrapf(err, "error parsing template: %s", tmplMainPath)
	}

	err = t.Execute(f, app)
	if err != nil {
		return errors.Wrapf(err, "error executing template: %s", tmplMainPath)
	}

	return nil
}
