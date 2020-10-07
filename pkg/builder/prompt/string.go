package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ForString prompts for string answer
func ForString(question, fallback string) (string, error) {
	if fallback != "" {
		question = fmt.Sprintf("%s [%s]", question, fallback)
	}
	fmt.Printf(question + "\n> ")

	reader := bufio.NewReader(os.Stdin)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		answer = strings.TrimSuffix(answer, "\n")
		if answer == "" {
			answer = fallback
		}

		return answer, nil
	}
}
