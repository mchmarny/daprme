package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/dapr-templates/daprme/pkg/lang"
	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
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
	if err := lang.Make(ctx, app, "test", testTargetDir); err != nil {
		t.Logf("Error making project: %v", err)
		t.FailNow()
	}
}
