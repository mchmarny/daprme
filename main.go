package main

import (
	"context"
	"os"

	"github.com/dapr-templates/daprme/pkg/project"
	"github.com/dapr-templates/daprme/pkg/prompt"
)

var (
	// Version will be overritten during build
	Version = "v0.0.1-default"

	daprVersion = "v0.11.2"

	targetDir = "." // TODO: make optional runtime flag
)

func main() {
	prompt.Print("\u2732 daprme (%s, Dapr: %s)", Version, daprVersion)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// collect
	app, err := prompt.Start(ctx)
	if err != nil {
		prompt.Print("Error: %v", err)
		os.Exit(-1)
	}

	// review
	prompt.Header("Review")
	prompt.Content(app)

	// create
	if prompt.ForBool("Create project?") {
		if err := project.Make(ctx, app, targetDir); err != nil {
			prompt.Print("Error creating project: %v", err)
			os.Exit(-1)
		}
	}

	prompt.Print("➜ Project was created in: %s", app.Meta.Name)

	// init
	if prompt.ForBool("Initialize project?") {
		usr := prompt.ForString("GitHub org or username?", "me")
		if err := project.Initialize(ctx, targetDir, usr, app); err != nil {
			prompt.Print("Error initializing project: %v", err)
			os.Exit(-1)
		}
	}

	prompt.Header("✓ Done, Happy Dapring")
	os.Exit(0)
}
