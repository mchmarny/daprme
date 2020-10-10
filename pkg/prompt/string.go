package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ForString prompts for string answer
func ForString(question, fallback string) string {
	if fallback != "" {
		question = fmt.Sprintf("%s [%s]", question, fallback)
	}
	fmt.Printf(question + promptPrefix)

	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(formatErrorMessage)
		return ForString(question, fallback)
	}

	answer = strings.TrimSuffix(answer, "\n")
	if len(answer) < 1 {
		answer = fallback
	}

	return answer
}
