package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

// ForComponents collects client component info
func ForComponents(list []string) ([]*model.Component, error) {
	out := make([]*model.Component, 0)
	for i, o := range list {
		fmt.Printf(fmt.Sprintf(" [%2d]: %s\n", i, o))
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		i := readInt(reader)
		if i < 0 || i >= len(list) {
			return nil, errors.Errorf("input out of range: %d", i)
		}

		c := &model.Component{
			ComponentName: model.ToCodeSafeString(list[i]),
			ComponentType: list[i],
		}

		out = append(out, c)

		more, err := ForBool("Add more?")
		if err != nil {
			return nil, errors.Errorf("unable to read input: %v", err)
		}
		if !more {
			break
		}
	}

	return out, nil
}
