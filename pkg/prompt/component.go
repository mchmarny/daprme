package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/model"
)

// ForComponents collects client component info
func ForComponents(list []string, suffix, comp string) []*model.Component {
	out := make([]*model.Component, 0)
	for i, o := range list {
		fmt.Printf(" [%2d]: %s\n", i+1, o)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		i := readInt(reader, 0)
		if i < 1 || i > len(list) {
			fmt.Println(outOfRangeMessage)
			return ForComponents(list, suffix, comp)
		}

		selected := list[i-1]
		c := &model.Component{
			Name: fmt.Sprintf("%s-%s", codeSafeString(selected), suffix),
			Type: fmt.Sprintf("%s.%s", comp, selected),
		}

		out = append(out, c)
		if !ForBool("Add more?") {
			break
		}
	}
	return out
}
