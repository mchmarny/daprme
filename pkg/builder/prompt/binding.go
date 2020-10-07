package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

// ForBinding collects binding info
func ForBinding() (*model.Binding, error) {
	b := &model.Binding{}
	fmt.Println("What type of PubSub component:")
	for i, o := range model.InputBindingComponentTypes() {
		fmt.Printf(fmt.Sprintf(" [%2d]: %s\n", i, o))
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		i := readInt(reader)
		if i < 0 || i >= len(model.InputBindingComponentTypes()) {
			return nil, errors.Errorf("input out of range: %d", i)
		}
		b.ComponentType = model.InputBindingComponentTypes()[i]
		break
	}

	// comp name
	compName, err := ForString("Component name: ", fmt.Sprintf("%s-pubsub", model.ToCodeSafeString(b.ComponentType)))
	if err != nil {
		return nil, errors.Errorf("unable to read input: %v", err)
	}
	b.ComponentName = compName

	return b, nil
}
