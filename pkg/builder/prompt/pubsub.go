package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/builder/format"
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
		i := readInt(reader, 0)
		if i < 0 || i >= len(model.PubsubComponentTypes()) {
			return nil, errors.Errorf("input out of range: %d", i)
		}

		ps.ComponentType = model.PubsubComponentTypes()[i]
		break
	}

	// comp and topic name
	ps.ComponentName = ForString("Component name: ", fmt.Sprintf("%s-pubsub", format.CodeSafeString(ps.ComponentType)))
	ps.TopicName = ForString("Topic name: ", "messages")

	return ps, nil
}
