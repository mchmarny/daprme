package builder

import (
	"github.com/dapr-templates/daprme/pkg/builder/prompt"
	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

// Start starts the wizard
func Start() (app *model.App, err error) {
	app = &model.App{}

	// name
	appName, err := prompt.ForString("App name: ", "my-app")
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	app.Name = appName

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

	// pubsub
	hasPubsub, err := prompt.ForBool("App subscribes to topic?")
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	if hasPubsub {
		list := make([]*model.Pubsub, 0)
		for {
			comp, err := prompt.ForPubSub()
			if err != nil {
				return nil, errors.Errorf("error getting pub/sub components: %v", err)
			}
			list = append(list, comp)
			more, err := prompt.ForBool("Add another PubSub component?")
			if err != nil {
				return nil, errors.Errorf("error parsing answer: %v", err)
			}
			if !more {
				break
			}
		}
		if len(list) > 0 {
			app.Pubsubs = list
		}
	}

	return
}
