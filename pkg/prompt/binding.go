package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/model"
)

// ForBinding collects binding info
func ForBinding() *model.Component {
	b := &model.Component{}
	fmt.Println("What type of binding component:")
	list := model.InputBindingComponentTypes()
	for i, o := range list {
		fmt.Printf(" [%2d]: %s\n", i+1, o)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	i := readInt(reader, 0)
	if i < 1 || i > len(list) {
		fmt.Println(outOfRangeMessage)
		return ForBinding()
	}

	// comp name
	selected := list[i-1]
	b.Type = fmt.Sprintf("bindings.%s", selected)
	b.Name = ForString("Component name: ", fmt.Sprintf("%s-binding", codeSafeString(selected)))

	return b
}
