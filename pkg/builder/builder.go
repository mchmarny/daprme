package builder

import (
	"fmt"

	"github.com/dapr-templates/daprme/pkg/builder/prompt"
	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

const (
	sectionLength = 80
	morePrompt    = "Add another?"
)

func printSection(title string) {
	s := fmt.Sprintf("= %s ", title)
	for i := len(s); i < sectionLength; i++ {
		s = s + "="
	}

	fmt.Println()
	fmt.Println(s)
	fmt.Println()
}

// Start starts the wizard
func Start() (app *model.App, err error) {
	app = &model.App{}

	printSection("Application")

	// name
	app.Name = prompt.ForString("Name: ", "my-app")

	// protocol
	protocol, err := prompt.ForOption("App protocol: ", model.HTTPProtocol, model.GRPCProtocol)
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	var defaultPort int
	switch protocol {
	case model.HTTPProtocol:
		app.Protocol = model.HTTPProtocol
		defaultPort = 8080
	case model.GRPCProtocol:
		app.Protocol = model.GRPCProtocol
		defaultPort = 50050
	default:
		return nil, errors.Errorf("invalid protocol input: %s", protocol)
	}

	// port
	appPort, err := prompt.ForInt("App protocol port: ", defaultPort)
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	app.Port = appPort

	printSection("Pub/Sub")

	// pubsub
	if prompt.ForBool("Subscribe to topic?") {
		list := make([]*model.Pubsub, 0)
		for {
			comp, err := prompt.ForPubSub()
			if err != nil {
				return nil, errors.Errorf("Error getting pub/sub components: %v.", err)
			}
			list = append(list, comp)
			if !prompt.ForBool(morePrompt) {
				break
			}
		}
		if len(list) > 0 {
			app.Pubsubs = list
		}
	}

	printSection("Binding")

	// binding
	if prompt.ForBool("Use input binding?") {
		list := make([]*model.Binding, 0)
		for {
			comp, err := prompt.ForBinding()
			if err != nil {
				return nil, errors.Errorf("Error getting binding components: %v.", err)
			}
			list = append(list, comp)
			if !prompt.ForBool(morePrompt) {
				break
			}
		}
		if len(list) > 0 {
			app.Bindings = list
		}
	}

	printSection("Service Invocation")

	// service
	if prompt.ForBool("Enable service invocation?") {
		list := make([]*model.Service, 0)
		for {
			comp, err := prompt.ForService()
			if err != nil {
				return nil, errors.Errorf("Error getting service: %v.", err)
			}
			list = append(list, comp)

			if !prompt.ForBool(morePrompt) {
				break
			}
		}
		if len(list) > 0 {
			app.Services = list
		}
	}

	printSection("Secrets")

	// secret
	app.Components = make([]*model.Component, 0)
	if prompt.ForBool("Use secrets?") {
		secretComp, err := prompt.ForComponents(model.SecretComponentTypes())
		if err != nil {
			return nil, errors.Errorf("Error parsing answer: %v.", err)
		}
		app.Components = append(app.Components, secretComp...)
	}

	printSection("Dapr Client")

	// client
	app.UsesClient = prompt.ForBool("Uses Dapr client?")
	if app.UsesClient {

		// state
		if prompt.ForBool("Add State client components?") {
			stateComp, err := prompt.ForComponents(model.StateComponentTypes())
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, stateComp...)
		}

		// pubsub
		if prompt.ForBool("Add PubSub client components?") {
			pubsubComp, err := prompt.ForComponents(model.PubsubComponentTypes())
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, pubsubComp...)
		}

		// binding
		if prompt.ForBool("Add Output Binding client components?") {
			outBindComp, err := prompt.ForComponents(model.OutputBindingComponentTypes())
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, outBindComp...)
		}

	}

	return
}
