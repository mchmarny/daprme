package builder

import (
	"github.com/dapr-templates/daprme/pkg/builder/print"
	"github.com/dapr-templates/daprme/pkg/builder/prompt"
	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

const (
	morePrompt     = "Add another?"
	appNameDefault = "my-app"
)

// Start starts the wizard
func Start() (app *model.App, err error) {
	app = &model.App{}

	print.Header("Application")

	// name
	app.Meta.Name = prompt.ForString("Name: ", appNameDefault)

	// protocol
	protocol, err := prompt.ForOption("App protocol: ", model.HTTPProtocol, model.GRPCProtocol)
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	var defaultPort int
	switch protocol {
	case model.HTTPProtocol:
		app.Meta.Protocol = model.HTTPProtocol
		defaultPort = 8080
	case model.GRPCProtocol:
		app.Meta.Protocol = model.GRPCProtocol
		defaultPort = 50050
	default:
		return nil, errors.Errorf("invalid protocol input: %s", protocol)
	}

	// port
	appPort, err := prompt.ForInt("App protocol port: ", defaultPort)
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	app.Meta.Port = appPort

	print.Header("Pub/Sub")

	// pubsub
	if prompt.ForBool("Subscribe to topic?") {
		list := make([]*model.PubSub, 0)
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
			app.PubSubs = list
		}
	}

	print.Header("Binding")

	// binding
	if prompt.ForBool("Use input binding?") {
		list := make([]*model.Component, 0)
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
		secretComp, err := prompt.ForComponents(model.SecretComponentTypes(), "secret", "secretstores")
		if err != nil {
			return nil, errors.Errorf("Error parsing answer: %v.", err)
		}
		app.Components = append(app.Components, secretComp...)
	}

	print.Header("Dapr Client")

	// client
	app.Meta.UsesClient = prompt.ForBool("Uses Dapr client?")
	if app.Meta.UsesClient {

		// state
		if prompt.ForBool("Add state components?") {
			stateComp, err := prompt.ForComponents(model.StateComponentTypes(), "store", "state")
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, stateComp...)
		}

		// pubsub
		if prompt.ForBool("Add pub/sub components?") {
			pubsubComp, err := prompt.ForComponents(model.PubsubComponentTypes(), "pubsub", "pubsub")
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, pubsubComp...)
		}

		// binding
		if prompt.ForBool("Add output binding components?") {
			outBindComp, err := prompt.ForComponents(model.OutputBindingComponentTypes(), "binding", "bindings")
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, outBindComp...)
		}

	}

	return
}
