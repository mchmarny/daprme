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
		fmt.Printf(" [%2d]: %s\n", i, o)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		i := readInt(reader, 0)
		if i < 0 || i >= len(list) {
			fmt.Println(outOfRangeMessage)
			return ForComponents(list, suffix, comp)
		}

		c := &model.Component{
			Name: fmt.Sprintf("%s-%s", codeSafeString(list[i]), suffix),
			Type: fmt.Sprintf("%s.%s", comp, list[i]),
		}

		out = append(out, c)
		if !ForBool("Add more?") {
			break
		}
	}
	return out
}
