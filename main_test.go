package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/dapr-templates/daprme/pkg/lang"
	"github.com/dapr-templates/daprme/pkg/model"
)

func getTestApp(t string) (app *model.App, err error) {
	return readManifest(fmt.Sprintf("test-data/%s.yaml", t))
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
	if err := lang.Make(ctx, app, testTargetDir); err != nil {
		t.Logf("Error making project: %v", err)
		t.FailNow()
	}
}
