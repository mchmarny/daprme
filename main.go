package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

var (
	// Version will be overritten during build
	Version = "v0.0.1-default"
)

func main() {
	if err := wizrd(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(-1)
	}
	fmt.Println()
	fmt.Println("DONE")
	os.Exit(0)
}

func wizrd() error {
	app := &model.App{}

	// name
	appName, err := promptString("App name: ", "my-app")
	if err != nil {
		return errors.Errorf("unable to read input: %v", err)
	}
	app.Name = appName

	// protocol
	protocol, err := promptOption("App protocol: ", "HTTP", "gRPC")
	if err != nil {
		return errors.Errorf("unable to read input: %v", err)
	}
	var defaultPort int
	switch protocol {
	case "HTTP":
		app.Protocol = model.HTTP
		defaultPort = 8080
	case "gRPC":
		app.Protocol = model.GRPC
		defaultPort = 50050
	default:
		return errors.Errorf("invalid protocol input: %s", protocol)
	}

	// port
	appPort, err := promptInt("App protocol port: ", defaultPort)
	if err != nil {
		return errors.Errorf("unable to read input: %v", err)
	}
	app.Port = appPort

	// pubsub
	hasPubsub, err := promptBool("App subscribes to topic?")
	if err != nil {
		return errors.Errorf("unable to read input: %v", err)
	}
	if hasPubsub {
		list := make([]*model.Pubsub, 0)
		for {
			comp, err := promptPubSub()
			if err != nil {
				return errors.Errorf("error getting pub/sub components: %v", err)
			}
			list = append(list, comp)
			more, err := promptBool("Add another PubSub component?")
			if err != nil {
				return errors.Errorf("error parsing answer: %v", err)
			}
			if !more {
				break
			}
		}
		if len(list) > 0 {
			app.Pubsubs = list
		}
	}

	// print
	print(app)

	return nil
}

func print(app *model.App) {
	fmt.Printf(`
Ready to create your app: 

  Name:     %s
  Protocol: %v
  Port:     %d
	
`, app.Name, app.Protocol, app.Port)
}

func promptPubSub() (*model.Pubsub, error) {
	ps := &model.Pubsub{}
	fmt.Println("What type of PubSub component:")
	for i, o := range model.PubsubComponentTypes() {
		fmt.Printf(fmt.Sprintf(" [%2d]: %s\n", i, o))
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		answer = strings.TrimSuffix(answer, "\n")

		i, err := strconv.Atoi(answer)
		if err != nil {
			return nil, errors.Wrap(err, "input not a number")
		}

		if i < 0 || i > len(model.PubsubComponentTypes()) {
			return nil, errors.Errorf("input out of range: %d", i)
		}

		ps.ComponentType = model.PubsubComponentTypes()[i]
		break
	}

	// comp name
	compName, err := promptString("Component name: ", fmt.Sprintf("%s-pubsub", ps.ComponentType))
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	ps.ComponentName = compName

	// topic name
	topicName, err := promptString("Topic name: ", "messages")
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	ps.TopicName = topicName

	return ps, nil
}

func promptBool(question string) (bool, error) {
	question = fmt.Sprintf("%s [Y] Yes, [N] No", question)
	fmt.Printf(question + "\n> ")

	reader := bufio.NewReader(os.Stdin)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}

		answer = strings.TrimSuffix(answer, "\n")

		return strings.ToUpper(answer) == "Y", nil
	}
}

func promptInt(question string, fallback int) (int, error) {
	question = fmt.Sprintf("%s [%d]", question, fallback)
	fmt.Printf(question + "\n> ")

	reader := bufio.NewReader(os.Stdin)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}

		answer = strings.TrimSuffix(answer, "\n")
		if answer == "" {
			return fallback, nil
		}

		i, err := strconv.Atoi(answer)
		if err != nil {
			return 0, errors.Wrap(err, "input not a number")
		}

		return i, nil
	}
}

func promptOption(question string, opts ...string) (string, error) {
	fmt.Printf(question + "\n> ")
	for i, o := range opts {
		fmt.Printf(fmt.Sprintf(" [%d]: %s ", i, o))
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			return "", errors.Wrap(err, "error reading input")
		}

		answer = strings.TrimSuffix(answer, "\n")

		i, err := strconv.Atoi(answer)
		if err != nil {
			return "", errors.Wrap(err, "input not a number")
		}

		if i < 0 || i > len(opts) {
			return "", errors.Errorf("input out of range: %d", i)
		}

		return opts[i], nil
	}
}

func promptString(question, fallback string) (string, error) {
	if fallback != "" {
		question = fmt.Sprintf("%s [%s]", question, fallback)
	}
	fmt.Printf(question + "\n> ")

	reader := bufio.NewReader(os.Stdin)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		answer = strings.TrimSuffix(answer, "\n")

		if answer == "" {
			answer = fallback
		}

		return answer, nil
	}
}
