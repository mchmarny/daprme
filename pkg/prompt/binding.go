package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/model"
)

const (
	outOfRangeMessage = "Selection out of range, please try again."
)

// ForBinding collects binding info
func ForBinding() (*model.Component, error) {
	b := &model.Component{}
	fmt.Println("What type of binding component:")
	list := model.InputBindingComponentTypes()
	for i, o := range list {
		fmt.Printf(" [%2d]: %s\n", i, o)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	i := readInt(reader, 0)
	if i < 0 || i >= len(list) {
		fmt.Println(outOfRangeMessage)
		return ForBinding()
	}

	// comp name
	b.Type = fmt.Sprintf("bindings.%s", list[i])
	b.Name = ForString("Component name: ", fmt.Sprintf("%s-binding", codeSafeString(list[i])))

	return b, nil
}
