package writer

import (
	"os"
	tmpl "text/template"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

func execTemplate(app *model.App, out, temp string) error {
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

	t, err := tmpl.ParseFiles(temp)
	if err != nil {
		return errors.Wrapf(err, "error parsing template: %s", temp)
	}

	err = t.Execute(f, app)
	if err != nil {
		return errors.Wrapf(err, "error executing template: %s", temp)
	}

	return nil
}
