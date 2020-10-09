package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/model"
)

// ForPubSub collects pubsub info
func ForPubSub() (*model.PubSub, error) {
	ps := &model.PubSub{}
	fmt.Println("What type of pub/sub component:")
	list := model.PubsubComponentTypes()
	for i, o := range list {
		fmt.Printf(" [%2d]: %s\n", i, o)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	i := readInt(reader, 0)
	if i < 0 || i >= len(list) {
		fmt.Println(outOfRangeMessage)
		return ForPubSub()
	}

	// comp and topic name
	ps.Type = fmt.Sprintf("pubsub.%s", list[i])
	ps.Name = ForString("Component name: ", fmt.Sprintf("%s-pubsub", codeSafeString(list[i])))
	ps.Topic = ForString("Topic name: ", "messages")

	return ps, nil
}
