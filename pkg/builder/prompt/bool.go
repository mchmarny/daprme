package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ForBool prompts for bool answer
func ForBool(question string) bool {
	question = fmt.Sprintf("%s [y] Yes, [n] No", question)
	fmt.Printf(question + "\n> ")

	reader := bufio.NewReader(os.Stdin)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, try again.")
			return ForBool(question)
		}
		answer = strings.TrimSuffix(answer, "\n")
		return strings.ToUpper(answer) == "Y"
	}
}
