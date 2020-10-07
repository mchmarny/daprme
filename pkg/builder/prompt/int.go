package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ForInt prompts for int answer
func ForInt(question string, fallback int) (int, error) {
	question = fmt.Sprintf("%s [%d]", question, fallback)
	fmt.Printf(question + "\n> ")

	reader := bufio.NewReader(os.Stdin)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}

		answer = strings.TrimSuffix(answer, "\n")
		if answer == "" {
			return fallback, nil
		}

		i, err := strconv.Atoi(answer)
		if err != nil {
			return 0, errors.Wrap(err, "input not a number")
		}

		return i, nil
	}
}
