package main

import (
	"io/ioutil"
	"testing"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/dapr-templates/daprme/pkg/writer"
)

func TestMarshaling(t *testing.T) {
	testFile := "test-data/app.yaml"
	b, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Logf("Error reading sample: %s", testFile)
		t.FailNow()
	}

	app, err := model.Unmarshal(b)
	if err != nil {
		t.Logf("Error parsing content: %s", b)
		t.FailNow()
	}

	b2, err := app.Marshal()
	if err != nil {
		t.Logf("Error marshaling app: %v", err)
		t.FailNow()
	}

	t.Logf("\n%s", b2)
}

func TestIntegration(t *testing.T) {
	testFile := "test-data/app.yaml"
	b, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Logf("Error reading sample: %s", testFile)
		t.FailNow()
	}

	app, err := model.Unmarshal(b)
	if err != nil {
		t.Logf("Error parsing content: %s", b)
		t.FailNow()
	}

	if err := writer.Make(app); err != nil {
		t.Logf("Error making project: %v", err)
		t.FailNow()
	}
}
