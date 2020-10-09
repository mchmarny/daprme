package main

import (
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/builder"
	"github.com/dapr-templates/daprme/pkg/cmd"
	"github.com/dapr-templates/daprme/pkg/print"
	"github.com/dapr-templates/daprme/pkg/prompt"
	"github.com/dapr-templates/daprme/pkg/writer"
)

var (
	// Version will be overritten during build
	Version = "v0.0.1-default"
)

func main() {
	print.Content(fmt.Sprintf("Starting daprme wizard (%s)", Version))

	// collect
	app, err := builder.Start()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(-1)
	}

	// review
	print.Header("Review")
	print.Content(app)

	// create
	if prompt.ForBool("Create project?") {
		if err := writer.Make(app); err != nil {
			fmt.Printf("Error creating project: %v", err)
			os.Exit(-1)
		}
	}

	print.Header("Done")
	print.Content(fmt.Sprintf("Project was created in: %s\n", app.Meta.Name))

	// init
	if prompt.ForBool("Initialize go module? (go mod init ...)") {
		usr := prompt.ForString("GitHub org or username?", "me")
		if err := cmd.InitProject(usr, app.Meta.Name); err != nil {
			fmt.Printf("Error initializing project: %v", err)
			os.Exit(-1)
		}
	}

	os.Exit(0)
}
