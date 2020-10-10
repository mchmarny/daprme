package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ForInt prompts for int answer
func ForInt(question string, fallback int) int {
	question = fmt.Sprintf("%s [%d]", question, fallback)
	fmt.Printf(question + promptPrefix)
	reader := bufio.NewReader(os.Stdin)
	return readInt(reader, fallback)
}

func readInt(reader *bufio.Reader, fallback int) int {
	answer, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(formatErrorMessage)
		return readInt(reader, fallback)
	}
	answer = strings.TrimSuffix(answer, "\n")
	if len(answer) < 1 {
		return fallback
	}

	i, err := strconv.Atoi(answer)
	if err != nil {
		fmt.Printf(formatErrorMessage)
		return readInt(reader, fallback)
	}
	return i
}
