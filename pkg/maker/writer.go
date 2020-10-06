package maker

import (
	"os"
	tmpl "text/template"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

const (
	tmplPath = "template/app.tmpl"
)

// Make creates app in specified path
func Make(app *model.App, out string) error {
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

	t, err := tmpl.ParseFiles(tmplPath)
	if err != nil {
		return errors.Wrapf(err, "error parsing template: %s", tmplPath)
	}

	err = t.Execute(f, app)
	if err != nil {
		return errors.Wrapf(err, "error executing template: %s", tmplPath)
	}

	return nil
}
