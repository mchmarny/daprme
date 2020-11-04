package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/model"
)

// ForPubSub collects pubsub info
func ForPubSub() *model.PubSub {
	ps := &model.PubSub{}
	fmt.Println("What type of pub/sub component:")
	list := model.PubsubComponentTypes()
	for i, o := range list {
		fmt.Printf(" [%2d]: %s\n", i+1, o)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	i := readInt(reader, 0)
	if i < 1 || i > len(list) {
		fmt.Println(outOfRangeMessage)
		return ForPubSub()
	}

	// comp and topic name
	selected := list[i-1]
	ps.Type = fmt.Sprintf("pubsub.%s", selected)
	ps.Name = ForString("Component name: ", fmt.Sprintf("%s-pubsub", codeSafeString(selected)))
	ps.Topic = ForString("Topic name: ", "messages")

	return ps
}
