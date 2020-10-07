package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ForOption prompts for options
func ForOption(question string, opts ...string) (string, error) {
	fmt.Printf(question + "\n> ")
	for i, o := range opts {
		fmt.Printf(fmt.Sprintf(" [%d] %s ", i, o))
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			return "", errors.Wrap(err, "error reading input")
		}

		answer = strings.TrimSuffix(answer, "\n")

		i, err := strconv.Atoi(answer)
		if err != nil {
			return "", errors.Wrap(err, "input not a number")
		}

		if i < 0 || i > len(opts) {
			return "", errors.Errorf("input out of range: %d", i)
		}

		return opts[i], nil
	}
}
