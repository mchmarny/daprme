package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ForInt prompts for int answer
func ForInt(question string, fallback int) (int, error) {
	question = fmt.Sprintf("%s [%d]", question, fallback)
	fmt.Printf(question + "\n> ")
	reader := bufio.NewReader(os.Stdin)
	for {
		return readInt(reader), nil
	}
}

func readInt(reader *bufio.Reader) int {
	answer, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input, try again:")
		return readInt(reader)
	}
	answer = strings.TrimSuffix(answer, "\n")
	i, err := strconv.Atoi(answer)
	if err != nil {
		fmt.Printf("Input not a number: %s.\n", answer)
		return readInt(reader)
	}
	return i
}
