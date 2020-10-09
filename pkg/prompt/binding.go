package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/format"
	"github.com/dapr-templates/daprme/pkg/model"
)

const (
	outOfRangeMessage = "Selection out of range, please try again."
)

// ForBinding collects binding info
func ForBinding() (*model.Component, error) {
	b := &model.Component{}
	fmt.Println("What type of binding component:")
	for i, o := range model.InputBindingComponentTypes() {
		fmt.Printf(fmt.Sprintf(" [%2d]: %s\n", i, o))
	}
	fmt.Println()

	var selectType string
	reader := bufio.NewReader(os.Stdin)
	for {
		i := readInt(reader, 0)
		if i < 0 || i >= len(model.InputBindingComponentTypes()) {
			fmt.Println(outOfRangeMessage)
			return ForBinding()
		}
		selectType = model.InputBindingComponentTypes()[i]
		break
	}

	// comp name
	b.Type = fmt.Sprintf("bindings.%s", selectType)
	b.Name = ForString("Component name: ", fmt.Sprintf("%s-binding", format.CodeSafeString(selectType)))

	return b, nil
}
