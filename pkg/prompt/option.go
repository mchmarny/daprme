package prompt

import (
	"bufio"
	"fmt"
	"os"
)

// ForOption prompts for options
func ForOption(question string, opts ...string) (string, error) {
	fmt.Printf(question + "\n> ")
	for i, o := range opts {
		fmt.Printf(fmt.Sprintf(" [%d] %s ", i, o))
	}
	fmt.Println()
	defaultOpt := opts[0]

	reader := bufio.NewReader(os.Stdin)
	for {
		i := readInt(reader, 0)
		if i < 0 || i >= len(opts) {
			fmt.Printf("Input out of range '%d', defaulting to %s.\n", i, defaultOpt)
			return ForOption(question, opts...)
		}
		return opts[i], nil
	}
}
