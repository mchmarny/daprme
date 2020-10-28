package main

import (
	"context"
	"flag"
	"io/ioutil"
	"os"

	"github.com/dapr-templates/daprme/pkg/lang"
	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/dapr-templates/daprme/pkg/prompt"
	"github.com/pkg/errors"
)

var (
	// Version will be overritten during build
	Version = "v0.0.1-default"

	daprVersion = "v0.11.3"
	targetDir   = ""
	sourceFile  = ""
)

func main() {
	flag.StringVar(&sourceFile, "file", "", "app manifest file path")
	flag.StringVar(&targetDir, "out", ".", "output directory")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if sourceFile != "" {
		app, err := readManifest(sourceFile)
		if err != nil {
			prompt.Print("Error: %v", err)
			os.Exit(-1)
		}
		if err := lang.Make(ctx, app, targetDir); err != nil {
			prompt.Print("Error creating project: %v", err)
			os.Exit(-1)
		}
		prompt.Print("➜ Project was created in: %s", targetDir)
		os.Exit(0)
	}

	// collect
	prompt.Print("\u2732 daprme (%s, Dapr: %s)", Version, daprVersion)
	app, err := prompt.Start(ctx)
	if err != nil {
		prompt.Print("Error: %v", err)
		os.Exit(-1)
	}

	// git
	prompt.Header("GitHub")
	app.Meta.Owner = prompt.ForString("Organization or username?", "me")

	// review
	prompt.Header("Review")
	prompt.Content(app)

	// create
	if prompt.ForBool("Create project?") {
		if err := lang.Make(ctx, app, targetDir); err != nil {
			prompt.Print("Error creating project: %v", err)
			os.Exit(-1)
		}
	}

	prompt.Print("➜ Project was created in: %s", targetDir)
	prompt.Header("✓ Done, Happy Dapring")
	os.Exit(0)
}

func readManifest(path string) (app *model.App, err error) {
	if path == "" {
		return nil, errors.New("empty manifest file path")
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading test file: %s", path)
	}
	a, err := model.Unmarshal(b)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing test file content: %s", path)
	}
	return a, nil
}
