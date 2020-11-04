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
		fmt.Printf(" [%d] %s ", i+1, o)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	i := readInt(reader, 0)
	if i < 1 || i > len(opts) {
		fmt.Println(outOfRangeMessage)
		return ForOption(question, opts...)
	}
	return opts[i-1]
}
