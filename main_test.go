package main

import (
	"io/ioutil"
	"testing"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/dapr-templates/daprme/pkg/writer"
)

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
