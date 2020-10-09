package main

import (
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/project"
	"github.com/dapr-templates/daprme/pkg/prompt"
)

var (
	// Version will be overritten during build
	Version = "v0.0.1-default"
)

func main() {
	prompt.Content(fmt.Sprintf("Starting daprme wizard (%s)", Version))

	// collect
	app, err := prompt.Start()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(-1)
	}

	// review
	prompt.Header("Review")
	prompt.Content(app)

	// create
	if prompt.ForBool("Create project?") {
		if err := project.Make(app); err != nil {
			fmt.Printf("Error creating project: %v", err)
			os.Exit(-1)
		}
	}

	prompt.Content(fmt.Sprintf("Project was created in: %s\n", app.Meta.Name))

	// init
	if prompt.ForBool("Initialize project?") {
		usr := prompt.ForString("GitHub org or username?", "me")
		if err := project.Initialize(usr, app.Meta.Name); err != nil {
			fmt.Printf("Error initializing project: %v", err)
			os.Exit(-1)
		}
	}

	prompt.Header("Done, Happy Dapring")
	os.Exit(0)
}
