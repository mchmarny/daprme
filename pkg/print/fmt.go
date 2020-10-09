package print

import (
	"fmt"
)

const (
	sectionLength = 80
)

// Content prints the provided object
func Content(v interface{}) {
	fmt.Println()
	fmt.Println(v)
	fmt.Println()
}

// Header prints provided title as header
func Header(title string) {
	s := fmt.Sprintf("=== %s ", title)
	for i := len(s); i < sectionLength; i++ {
		s = s + "="
	}

	fmt.Println()
	fmt.Println(s)
	fmt.Println()
}
