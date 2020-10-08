package builder

import (
	"github.com/dapr-templates/daprme/pkg/builder/print"
	"github.com/dapr-templates/daprme/pkg/builder/prompt"
	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

const (
	morePrompt = "Add another?"
)

// Start starts the wizard
func Start() (app *model.App, err error) {
	app = &model.App{}

	print.Header("Application")

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

	print.Header("Pub/Sub")

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

	print.Header("Binding")

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

	print.Header("Service Invocation")

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

	print.Header("Secrets")

	// secret
	app.Components = make([]*model.Component, 0)
	if prompt.ForBool("Use secrets?") {
		secretComp, err := prompt.ForComponents(model.SecretComponentTypes(), "secret")
		if err != nil {
			return nil, errors.Errorf("Error parsing answer: %v.", err)
		}
		app.Components = append(app.Components, secretComp...)
	}

	print.Header("Dapr Client")

	// client
	app.UsesClient = prompt.ForBool("Uses Dapr client?")
	if app.UsesClient {

		// state
		if prompt.ForBool("Add state components?") {
			stateComp, err := prompt.ForComponents(model.StateComponentTypes(), "store")
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, stateComp...)
		}

		// pubsub
		if prompt.ForBool("Add pub/sub components?") {
			pubsubComp, err := prompt.ForComponents(model.PubsubComponentTypes(), "pubsub")
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, pubsubComp...)
		}

		// binding
		if prompt.ForBool("Add output binding components?") {
			outBindComp, err := prompt.ForComponents(model.OutputBindingComponentTypes(), "binding")
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, outBindComp...)
		}

	}

	return
}
