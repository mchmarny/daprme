package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/builder/format"
	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

// ForComponents collects client component info
func ForComponents(list []string, comp string) ([]*model.Component, error) {
	out := make([]*model.Component, 0)
	for i, o := range list {
		fmt.Printf(fmt.Sprintf(" [%2d]: %s\n", i, o))
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		i := readInt(reader, 0)
		if i < 0 || i >= len(list) {
			return nil, errors.Errorf("input out of range: %d", i)
		}

		c := &model.Component{
			ComponentName: fmt.Sprintf("%s-%s", format.CodeSafeString(list[i]), comp),
			ComponentType: list[i],
		}

		out = append(out, c)

		if !ForBool("Add more?") {
			break
		}
	}

	return out, nil
}
