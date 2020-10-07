package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ForBool prompts for bool answer
func ForBool(question string) (bool, error) {
	question = fmt.Sprintf("%s [Y] Yes, [N] No", question)
	fmt.Printf(question + "\n> ")

	reader := bufio.NewReader(os.Stdin)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}
		answer = strings.TrimSuffix(answer, "\n")
		return strings.ToUpper(answer) == "Y", nil
	}
}
