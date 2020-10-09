package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/builder/format"
	"github.com/dapr-templates/daprme/pkg/model"
)

// ForPubSub collects pubsub info
func ForPubSub() (*model.Pubsub, error) {
	ps := &model.Pubsub{}
	fmt.Println("What type of pub/sub component:")
	for i, o := range model.PubsubComponentTypes() {
		fmt.Printf(fmt.Sprintf(" [%2d]: %s\n", i, o))
	}
	fmt.Println()

	var selectType string
	reader := bufio.NewReader(os.Stdin)
	for {
		i := readInt(reader, 0)
		if i < 0 || i >= len(model.PubsubComponentTypes()) {
			fmt.Println(outOfRangeMessage)
			return ForPubSub()
		}

		selectType = model.PubsubComponentTypes()[i]
		break
	}

	// comp and topic name
	ps.ComponentType = fmt.Sprintf("pubsub.%s", selectType)
	ps.ComponentName = ForString("Component name: ", fmt.Sprintf("%s-pubsub", format.CodeSafeString(selectType)))
	ps.TopicName = ForString("Topic name: ", "messages")

	return ps, nil
}
