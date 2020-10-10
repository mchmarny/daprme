package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/dapr-templates/daprme/pkg/project"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func getTestApp(t string) (app *model.App, err error) {
	testFile := fmt.Sprintf("test-data/%s.yaml", t)
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

func TestIntegrations(t *testing.T) {
	app, err := getTestApp("grpc")
	if err != nil {
		t.FailNow()
	}
	testIntegration(t, app)

	app, err = getTestApp("http")
	if err != nil {
		t.FailNow()
	}
	testIntegration(t, app)

	app, err = getTestApp("cli")
	if err != nil {
		t.FailNow()
	}
	testIntegration(t, app)
}

func testIntegration(t *testing.T, app *model.App) {
	testTargetDir := "./test"
	ctx := context.Background()
	if err := project.Make(ctx, app, testTargetDir); err != nil {
		t.Logf("Error making project: %v", err)
		t.FailNow()
	}

	if err := project.Initialize(ctx, testTargetDir, "test", app); err != nil {
		t.Logf("Error initializing project: %v", err)
		t.FailNow()
	}
}

func TestAppMarshaling(t *testing.T) {
	app, err := getTestApp("http")
	if err != nil {
		t.FailNow()
	}
	testMarshaling(t, app)
	app, err = getTestApp("grpc")
	if err != nil {
		t.FailNow()
	}
	testMarshaling(t, app)
	app, err = getTestApp("cli")
	if err != nil {
		t.FailNow()
	}
	testMarshaling(t, app)
}

func testMarshaling(t *testing.T, app *model.App) {
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
	assert.Equal(t, app.Meta, app2.Meta)
}
