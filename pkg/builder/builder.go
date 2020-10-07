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

	// binding
	hasBinding, err := prompt.ForBool("App will use input binding?")
	if err != nil {
		return nil, errors.Errorf("Unable to read input: %v.", err)
	}
	if hasBinding {
		list := make([]*model.Binding, 0)
		for {
			comp, err := prompt.ForBinding()
			if err != nil {
				return nil, errors.Errorf("Error getting binding components: %v.", err)
			}
			list = append(list, comp)
			more, err := prompt.ForBool("Add another Binding component?")
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			if !more {
				break
			}
		}
		if len(list) > 0 {
			app.Bindings = list
		}
	}

	// service
	hasService, err := prompt.ForBool("App will be invoked by another service?")
	if err != nil {
		return nil, errors.Errorf("Unable to read input: %v.", err)
	}
	if hasService {
		list := make([]*model.Service, 0)
		for {
			comp, err := prompt.ForService()
			if err != nil {
				return nil, errors.Errorf("Error getting service: %v.", err)
			}
			list = append(list, comp)
			more, err := prompt.ForBool("Add another Service?")
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			if !more {
				break
			}
		}
		if len(list) > 0 {
			app.Services = list
		}
	}

	// secret
	app.Components = make([]*model.Component, 0)
	addSecretComp, err := prompt.ForBool("Add Secret components?")
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	if addSecretComp {
		secretComp, err := prompt.ForComponents(model.SecretComponentTypes())
		if err != nil {
			return nil, errors.Errorf("Error parsing answer: %v.", err)
		}
		app.Components = append(app.Components, secretComp...)
	}

	// client
	usesClient, err := prompt.ForBool("App uses Dapr client for API calls?")
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	if usesClient {
		app.UsesClient = usesClient

		// state
		addStateComp, err := prompt.ForBool("Add State client components?")
		if err != nil {
			return nil, errors.Errorf("unable to read input: %v", err)
		}
		if addStateComp {
			stateComp, err := prompt.ForComponents(model.StateComponentTypes())
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, stateComp...)
		}

		// pubsub
		addPubSubComp, err := prompt.ForBool("Add PubSub client components?")
		if err != nil {
			return nil, errors.Errorf("unable to read input: %v", err)
		}
		if addPubSubComp {
			pubsubComp, err := prompt.ForComponents(model.PubsubComponentTypes())
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, pubsubComp...)
		}

		// binding
		addOutBindComp, err := prompt.ForBool("Add Output Binding client components?")
		if err != nil {
			return nil, errors.Errorf("unable to read input: %v", err)
		}
		if addOutBindComp {
			outBindComp, err := prompt.ForComponents(model.OutputBindingComponentTypes())
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, outBindComp...)
		}

	}

	return
}
