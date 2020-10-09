package main

import (
	"io/ioutil"
	"testing"

	"github.com/dapr-templates/daprme/pkg/cmd"
	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/dapr-templates/daprme/pkg/writer"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const (
	testFile = "test-data/app.yaml"
)

func getTestApp() (app *model.App, err error) {
	b, err := ioutil.ReadFile(testFile)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading test file: %s", testFile)
	}
	a, err := model.Unmarshal(b)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing test file content: %s", testFile)
	}
	return a, nil
}

func TestIntegration(t *testing.T) {
	app, err := getTestApp()
	if err != nil {
		t.FailNow()
	}

	if err := writer.Make(app); err != nil {
		t.Logf("Error making project: %v", err)
		t.FailNow()
	}

	if err := cmd.InitProject("test", app.Meta.Name); err != nil {
		t.Logf("Error initializing project: %v", err)
		t.FailNow()
	}
}

func TestMarshaling(t *testing.T) {
	app, err := getTestApp()
	if err != nil {
		t.FailNow()
	}

	b2, err := app.Marshal()
	if err != nil {
		t.Logf("Error marshaling app: %v", err)
		t.FailNow()
	}

	t.Logf("\n%s", b2)

	app2, err := model.Unmarshal(b2)
	if err != nil {
		t.FailNow()
	}

	assert.Equal(t, app, app2)
}
