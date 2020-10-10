package prompt

import (
	"bufio"
	"fmt"
	"os"
)

// ForOption prompts for options
func ForOption(question string, opts ...string) string {
	fmt.Printf(question + promptPrefix)
	for i, o := range opts {
		fmt.Printf(" [%d] %s ", i, o)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	i := readInt(reader, 0)
	if i < 0 || i >= len(opts) {
		fmt.Printf(outOfRangeMessage)
		return ForOption(question, opts...)
	}
	return opts[i]
}
