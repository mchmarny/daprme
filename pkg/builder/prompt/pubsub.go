package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

// ForPubSub collects pubsub info
func ForPubSub() (*model.Pubsub, error) {
	ps := &model.Pubsub{}
	fmt.Println("What type of PubSub component:")
	for i, o := range model.PubsubComponentTypes() {
		fmt.Printf(fmt.Sprintf(" [%2d]: %s\n", i, o))
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		i := readInt(reader)
		if i < 0 || i >= len(model.PubsubComponentTypes()) {
			return nil, errors.Errorf("input out of range: %d", i)
		}

		ps.ComponentType = model.PubsubComponentTypes()[i]
		break
	}

	// comp name
	compName, err := ForString("Component name: ", fmt.Sprintf("%s-pubsub", model.ToCodeSafeString(ps.ComponentType)))
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	ps.ComponentName = compName

	// topic name
	topicName, err := ForString("Topic name: ", "messages")
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	ps.TopicName = topicName

	return ps, nil
}
