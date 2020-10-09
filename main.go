package main

import (
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/builder"
	"github.com/dapr-templates/daprme/pkg/builder/print"
	"github.com/dapr-templates/daprme/pkg/builder/prompt"
	"github.com/dapr-templates/daprme/pkg/writer"
)

var (
	// Version will be overritten during build
	Version = "v0.0.1-default"
)

func main() {
	print.Content(fmt.Sprintf("Starting daprme wizard (%s)", Version))

	app, err := builder.Start()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(-1)
	}

	print.Header("Review")
	print.Content(app)

	if prompt.ForBool("Create project?") {
		if err := writer.Make(app); err != nil {
			fmt.Printf("Error creating project: %v", err)
			os.Exit(-1)
		}
	}

	print.Header("Done")
	print.Content(fmt.Sprintf("Your project was created in %s directory.", app.Meta.Name))

	os.Exit(0)
}
