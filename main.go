package main

import (
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/builder"
	"github.com/dapr-templates/daprme/pkg/builder/prompt"
	"github.com/dapr-templates/daprme/pkg/writer"
)

var (
	// Version will be overritten during build
	Version = "v0.0.1-default"
)

func main() {
	app, err := builder.Start()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(-1)
	}

	fmt.Println()
	fmt.Println(app.String())
	fmt.Println()

	ok, err := prompt.ForBool("Create Dapr application?")
	if err != nil {
		fmt.Printf("Invalid input: %v", err)
		os.Exit(-1)
	}

	if ok {
		if err := writer.MakeMain(app, "./test/main.go"); err != nil {
			fmt.Printf("Error writting main: %v", err)
			os.Exit(-1)
		}
	}

	fmt.Println("DONE")
	os.Exit(0)
}
